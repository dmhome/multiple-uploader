package entity

type Music struct {
	MusicId   int  `json:"music_id" db:"music_id"`
	MusicName  string  `json:"music_name" db:"music_name"`
	MusicKey  string  `json:"music_key" db:"music_key"`
	LyricTxtKey  string  `json:"lyric_txt_key" db:"lyric_txt_key"`
	LyricLrcKey  string  `json:"lyric_lrc_key" db:"lyric_lrc_key"`
	DownloadNum  int  `json:"download_num" db:"download_num"`
	MusicType  int  `json:"music_type" db:"music_type"`
	SingerName  string  `json:"singer_name" db:"singer_name"`
	UploaderUid  int  `json:"uploader_uid" db:"uploader_uid"`
	CreateTime  string  `json:"create_time" db:"create_time"`
	UpdateTime  string  `json:"update_time" db:"update_time"`
	IsDel  int  `json:"is_del" db:"is_del"`
}