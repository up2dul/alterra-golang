package model

type Book struct {
	Id			uint		`gorm:"primaryKey" json:"id"`
	Name		string 	`json:"name"`
	Author	string 		`json:"author"`
	Year		string	`json:"year"`
}
