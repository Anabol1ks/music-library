package models

type Song struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Group       string `gorm:"column:song_group;index" json:"group"`
	Title       string `gorm:"index" json:"title"`
	ReleaseDate string `json:"release_date"`
	Text        string `gorm:"type:text" json:"text"`
	Link        string `json:"link"`
}
