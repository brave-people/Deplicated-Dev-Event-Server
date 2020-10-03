package responses

type Event struct {
	ID          uint   `gorm:"primary_key;auto_increment"  json:"id"`
	Title       string `gorm:"not null"                 json:"title"`
	Description string `gorm:"size:255;not null"    json:"nickname"`
	ImgURL      string `json:"imgURL"`
	Contact     string `json:"contact"`
	Year        uint   `json:"year"`
	Month       uint   `json:"month"`
	Day         uint   `json:"day"`
	Time        uint   `json:"time"`
	Free        uint   `json:"free"`
	Tages       []Tag  `gorm:"many2many:event_tags;" json:"tages"`
}
