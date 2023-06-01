package models

type Blog struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Image  string `json:"image"`
	UserID int    `json:"user_id"`
	User   User   `json:"user";gorm:"foreignkey:UserID"`
}
