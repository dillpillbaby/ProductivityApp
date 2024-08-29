package models
import (
	"time"
)
type JournalEntry struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Description string `json:"description"`
  MoodID int `json:"moodId"`
  Mood Mood `gorm:"foreignKey:MoodID"`
  CreatedAt    time.Time      
  UpdatedAt    time.Time     
}