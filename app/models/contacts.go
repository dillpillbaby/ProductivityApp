package models

type Contact struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Name  string `json:"title"`
  Phone_Number string `json:"phoneNumber"`
}