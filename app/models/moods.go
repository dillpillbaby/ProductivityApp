package models
import (
	"gorm.io/gorm"
	"errors"
)
type Mood struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Name  string `json:"name"`
  Description string `json:"description"`
  PositivityLevel int `json:"positivityLevel" gorm:"check:positivity_level >= 1 AND positivity_level <= 10"`
  ColorID int `json:"colorId"`
  Color Color `gorm:"foreignKey:ColorID"`
}

// BeforeSave is a GORM hook that runs before saving the model
func (m *Mood) BeforeSave(tx *gorm.DB) (err error) {
	if m.PositivityLevel < 1 || m.PositivityLevel > 10 {
		return errors.New("positivity level must be between 1 and 10")
	}
	return
}