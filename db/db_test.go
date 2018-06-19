package db

import (
	"fmt"
	"ksong/entity"
	"testing"
)

func TestDb(t *testing.T) {

	//初始化数据库
	Init()

	var music []entity.Music
	err := DB.Select(&music, "select * from m_music where music_id=?", 5)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	if len(music) == 0 {
		fmt.Println("no select row")
	}

	fmt.Println("select succ:", music)

	musics := make([]entity.Music, 0)
	DB.Select(&musics, "SELECT * FROM m_music ORDER BY music_id ASC")
	fmt.Println(musics)
	musicInfo := musics[0]
	fmt.Println(musicInfo.MusicName)
	jason, john := musics[0], musics[1]
	fmt.Printf("%#v\n%#v", jason, john)

}
