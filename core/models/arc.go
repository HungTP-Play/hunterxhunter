package models

type Arc struct {
	Id          string      `json:"id" gorm:"primaryKey"`
	Name        string      `json:"name"`
	FromChapter int         `json:"from_chapter"`
	ToChapter   int         `json:"to_chapter"`
	FromEpisode int         `json:"from_episode"`
	ToEpisode   int         `json:"to_episode"`
	Characters  []Character `json:"characters" gorm:"type:jsonb"`
}
