package main

import (
	"io/ioutil"
	"fmt"
	"strings"
	"sync"
	"ksong/db"
	"ksong/utils"
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

	//开启10个协程
	for i:=0; i<10; i++ {
		go work(chm)
	}

	basePath := "../music/original/"

	files, _ := ioutil.ReadDir(basePath)
	for _, f := range files {
		//fmt.Println(f.Name())

		filename := f.Name()
		fileinfoss := strings.Split(filename, "-")
		if len(fileinfoss) < 2 {
			fmt.Println(filename)
			continue
		}

		authName :=  strings.TrimSpace(fileinfoss[0])
		musicName := strings.TrimSpace(fileinfoss[1])

		logId := utils.GetLogId()

		fileMusicPath := basePath + filename
		fileMusicKey := logId + ".mp3"

		fileTxtKey := ""

		fileLrcKey := ""

		m := MusicInfo{
			fileMusicPath:fileMusicPath,
			fileMusicKey:fileMusicKey,
			fileTxtPath:"",
			fileTxtKey:fileTxtKey,
			fileLrcPath:"",
			fileLrcKey:fileLrcKey,
			authorName: authName,
			muiscName: strings.Replace(musicName, ".mp3", "", -1),
		}
		//fmt.Println(m)
		wg.Add(1)
		chm <- m
	}


	wg.Wait()
	fmt.Println("success")
}


func work(ch chan MusicInfo)  {

	for musicInfo := range ch {
		fmt.Println("开始上传")
		fmt.Printf("%#v\n", musicInfo)
		fileMusicFlag := qiniu.Upload(musicInfo.fileMusicPath, musicInfo.fileMusicKey)
		fmt.Printf("fileMusicPath:%s fileMusicKey:%s fileMusicFlag:%t \n", musicInfo.fileMusicPath, musicInfo.fileMusicKey, fileMusicFlag)

		sql := "insert into m_music(music_name, music_key, lyric_txt_key, lyric_lrc_key, music_type, singer_name, uploader_uid, create_time, update_time) values(?,?,?,?,?,?,?,?,?)"
		_, err := db.DB.Exec(sql, musicInfo.muiscName, musicInfo.fileMusicKey, musicInfo.fileTxtKey, musicInfo.fileLrcKey, 1, musicInfo.authorName, 100000, utils.GetcurDateTime(), utils.GetcurDateTime())
		if err != nil {
			panic(err)
		}

		fmt.Println("上传完成")
		wg.Done()
	}

}
