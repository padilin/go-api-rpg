package charDB

import (
	"context"
	"fmt"
)

func charDB() {
	dsn := "test.db"

	InitDB(dsn)
	db := OpenDB(dsn)
	char := Character{
		Name: "test",
	}
	initChar := saveCharacter(char, db)
	fmt.Println("Test Character: ", initChar.Name, " #", initChar.ID)

	getChar := retrieveCharacterById(1, db)
	fmt.Println("Test Character3: ", getChar.Name, " #", getChar.ID)
	ctx := context.Background()
	db.Debug().Where("id = ?", 2).First(ctx)
}
