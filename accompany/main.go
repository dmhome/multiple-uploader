package main

import (
	"bufio"
	"fmt"
	"io"
	"ksong/db"
	"ksong/utils"
	"os"
	"strings"
	"sync"
	"ksong/qiniu"
)

type MusicInfo struct {
	fileMusicPath, fileMusicKey, fileTxtPath, fileTxtKey, fileLrcPath, fileLrcKey , authorName, muiscName string
}

var wg sync.WaitGroup

func main() {
	//初始化数据库
	db.Init()
	defer db.DB.Close()

	chm := make(chan MusicInfo)

	for i:=0; i<10; i++ {
		go work(chm)
	}

	basePath := "../music/accompany/"

	fi, err := os.Open(basePath + "歌曲列表.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {

		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		linestr := string(a)
		musicInfo := strings.Split(linestr, " ")

		if len(musicInfo) != 4 {
			continue
		}

		logId := utils.GetLogId()

		fileMusicPath := basePath + musicInfo[0]
		fileMusicKey := logId + ".mp3"
		fileTxtPath := basePath + musicInfo[1]
		fileTxtKey := logId + ".txt"

		fileLrcPath := basePath + musicInfo[2]
		fileLrcKey := logId + ".lrc"

		m := MusicInfo{
			fileMusicPath:fileMusicPath,
			fileMusicKey:fileMusicKey,
			fileTxtPath:fileTxtPath,
			fileTxtKey:fileTxtKey,
			fileLrcPath:fileLrcPath,
			fileLrcKey:fileLrcKey,
			authorName: musicInfo[3],
			muiscName: strings.Replace(musicInfo[0], ".mp3", "", -1),
		}
		wg.Add(1)
		chm <- m
	}
	wg.Wait()
	fmt.Println("success")

}

func work(ch chan MusicInfo)  {

	for musicInfo := range ch {
		fmt.Println("开始上传")
		fmt.Printf("%#v \n", musicInfo)

		fileexist,_ := utils.FileExists(musicInfo.fileTxtPath)
		if !fileexist {
			fmt.Println(musicInfo.fileTxtPath + "文件不存在")
			wg.Done()
			continue
		}
		fileexist,_ = utils.FileExists(musicInfo.fileLrcPath)
		if !fileexist {
			fmt.Println(musicInfo.fileLrcPath + "文件不存在")
			wg.Done()
			continue
		}


		fileMusicFlag := qiniu.Upload(musicInfo.fileMusicPath, musicInfo.fileMusicKey)
		fileTxtFlag := qiniu.Upload(musicInfo.fileTxtPath, musicInfo.fileTxtKey)
		fileLrcFlag := qiniu.Upload(musicInfo.fileLrcPath, musicInfo.fileLrcKey)

		fmt.Printf("fileMusicPath:%s fileMusicKey:%s fileMusicFlag:%t \n", musicInfo.fileMusicPath, musicInfo.fileMusicKey, fileMusicFlag)
		fmt.Printf("fileTxtPath:%s fileTxtKey:%s fileTxtFlag:%t \n", musicInfo.fileTxtPath, musicInfo.fileTxtKey, fileTxtFlag)
		fmt.Printf("fileLrcPath:%s fileLrcKey:%s fileLrcFlag:%t \n", musicInfo.fileLrcPath, musicInfo.fileLrcKey, fileLrcFlag)

		sql := "insert into m_music(music_name, music_key, lyric_txt_key, lyric_lrc_key, music_type, singer_name, uploader_uid, create_time, update_time) values(?,?,?,?,?,?,?,?,?)"
		_, err := db.DB.Exec(sql, musicInfo.muiscName, musicInfo.fileMusicKey, musicInfo.fileTxtKey, musicInfo.fileLrcKey, 2, musicInfo.authorName, 10000, utils.GetcurDateTime(), utils.GetcurDateTime())
		if err != nil {
			fmt.Println(err)
			wg.Done()
		}

		fmt.Println("上传完成")
		wg.Done()
	}

}
