package models
type Color struct{
	ID uint `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	ColorCode string `json:"colorCode"`
}