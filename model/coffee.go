package model

import (
	"encoding/json"
)

type Beverage struct {
	ID          uint    `gorm:"primaryKey"`
	Temperature string  `gorm:"size:100"`
	DrinkName   string  `gorm:"size:255;not null"`
	Price       float32 `gorm:"not null"`
	Calories    float32 `gorm:"not null"`
	ImageUrl    string  `gorm:"size:255"`
	Recipe      string  `gorm:"type:json"` // Store JSON as a string in the database
}

// Method to convert Recipe string into a map
func (b *Beverage) UnmarshalRecipe() (map[string]int, error) {
	var recipe map[string]int
	err := json.Unmarshal([]byte(b.Recipe), &recipe)
	if err != nil {
		return nil, err
	}
	return recipe, nil
}

// Method to convert a map into a Recipe string
func (b *Beverage) MarshalRecipe(recipe map[string]int) error {
	data, err := json.Marshal(recipe)
	if err != nil {
		return err
	}
	b.Recipe = string(data)
	return nil
}