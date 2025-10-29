package charDB

import (
	"context"
	"errors"
	"log"

	"gorm.io/gorm"
)

// Currency represents a specific type of currency in the game
type Currency struct {
	gorm.Model
	Name        string  `gorm:"not null" json:"name"`
	Code        string  `gorm:"not null;uniqueIndex" json:"code"`
	Amount      float64 `gorm:"default:0" json:"amount"`
	MaxAmount   float64 `gorm:"default:999999" json:"max_amount"`
	Description string  `gorm:"description"`
	CharacterID uint
}

// Stats represents the base attributes of a character
type Stats struct {
	gorm.Model
	CharacterID  uint `gorm:"uniqueIndex" json:"character_id"`
	Strength     int  `gorm:"default:10" json:"strength"`
	Dexterity    int  `gorm:"default:10" json:"dexterity"`
	Constitution int  `gorm:"default:10" json:"constitution"`
	Intelligence int  `gorm:"default:10" json:"intelligence"`
	Wisdom       int  `gorm:"default:10" json:"wisdom"`
	Charisma     int  `gorm:"default:10" json:"charisma"`
}

// Class represents a character class with its abilities and attributes
type Class struct {
	gorm.Model
	Name        string      `gorm:"not null;uniqueIndex" json:"name"`
	Description string      `json:"description"`
	Abilities   []Ability   `gorm:"foreignKey:ClassID" json:"abilities"`
	Stats       string      `gorm:"type:json" json:"stats"`      // Base stats modifiers stored as JSON
	Equipment   string      `gorm:"type:json" json:"equipment"`  // Allowed equipment types stored as JSON
	Attributes  string      `gorm:"type:json" json:"attributes"` // Flexible attributes stored as JSON
	CharacterID  Character  `gorm:"foreignKey:ClassID" json:"characters,omitempty"`
}

// Ability represents a special ability or skill
type Ability struct {
	gorm.Model
	Name        string `gorm:"not null" json:"name"`
	Description string `json:"description"`
	Cost        int    `gorm:"default:0" json:"cost"`
	Cooldown    int    `gorm:"default:0" json:"cooldown"` // in seconds
	Effect      string `json:"effect"`
	ClassID     uint   `json:"class_id"`
}

// Character represents the main character entity
type Character struct {
	gorm.Model
	Name       string    `gorm:"not null" json:"name"`
	Level      int       `gorm:"default:1" json:"level"`
	Experience int64     `gorm:"default:0" json:"experience"`
	ClassID    uint      `json:"class_id"`
	Class      []Class    `gorm:"foreignKey:ClassID" json:"class,omitempty"`
	Stats      Stats     `gorm:"foreignKey:CharacterID" json:"stats"`
	Currencies []Currency   `gorm:"foreignKey:CharacterID" json:"currencies"`
	Inventory  []Item    `gorm:"foreignKey:CharacterID" json:"inventory"`
	Equipment  []Equipment `gorm:"foreignKey:CharacterID" json:"equipment"`
	Status     string    `gorm:"default:active" json:"status"`
	Attributes string    `gorm:"type:json" json:"attributes"`
}

// Item represents an item in the game
type Item struct {
	gorm.Model
	Name        string  `gorm:"not null" json:"name"`
	Description string  `json:"description"`
	Type        string  `gorm:"not null;index" json:"type"`
	Stats       string  `gorm:"type:json" json:"stats"` // Stats stored as JSON
	Value       float64 `gorm:"default:0" json:"value"`
	Stackable   bool    `gorm:"default:false" json:"stackable"`
	Quantity    int     `gorm:"default:1" json:"quantity"`
	CharacterID *uint   `json:"character_id,omitempty"` // For inventory items
	EquipmentID *uint   `json:"equipment_id,omitempty"` // For equipped items
}

// Equipment represents equipped items in different slots
type Equipment struct {
	gorm.Model
	CharacterID uint  `gorm:"uniqueIndex" json:"character_id"`
	Head        *Item `gorm:"foreignKey:EquipmentID" json:"head"`
	Body        *Item `gorm:"foreignKey:EquipmentID" json:"body"`
	Legs        *Item `gorm:"foreignKey:EquipmentID" json:"legs"`
	Feet        *Item `gorm:"foreignKey:EquipmentID" json:"feet"`
	MainHand    *Item `gorm:"foreignKey:EquipmentID" json:"main_hand"`
	OffHand     *Item `gorm:"foreignKey:EquipmentID" json:"off_hand"`
	Necklace    *Item `gorm:"foreignKey:EquipmentID" json:"necklace"`
	Ring1       *Item `gorm:"foreignKey:EquipmentID" json:"ring1"`
	Ring2       *Item `gorm:"foreignKey:EquipmentID" json:"ring2"`
}

func SaveCharacter(character Character, db *gorm.DB) Character {
	ctx := context.Background()
	result := gorm.WithResult()
	err := gorm.G[Character](db, result).Create(ctx, &character)
	if err != nil {
		log.Fatal("Failed to save character: ", err)
	}
	return character
}

func SaveCharacterBulk(characters []Character, db *gorm.DB) []Character {
	result := db.Create(&characters)
	if result.Error != nil {
		log.Fatal("Failed to save character array: ", result.Error)
	}
	return characters
}

func RetrieveCharacterById(ID uint, db *gorm.DB) Character {
	var character Character
	result := db.First(&character, ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Fatal("Failed to find character with id: ", ID)
	}
	return character
}
