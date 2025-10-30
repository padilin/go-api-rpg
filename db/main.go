package charDB

import (
	"fmt"
	"log"
)

func CharDB() {
	dsn := "test.db"

	InitDB(dsn)
	db := OpenDB(dsn)
	char := Character{
		Name: "test",
	}
	fmt.Println(char.ID)
	initChar, err := SaveCharacter(char, db)
	if err != nil {
		log.Printf("Failed to save character: %v", err)
		return
	}
	fmt.Println("Test Character: ", initChar.Name, " #", initChar.ID)

	getChar, err := RetrieveCharacterById(1, db)
	if err != nil {
		log.Printf("Failed to get character %d: %v", 1, err)
		return
	}
	fmt.Println("Test Character3: ", getChar.Name, " #", getChar.ID)
	// ctx := context.Background()
	// db.Debug().Where("id = ?", 1).First(ctx)
}
