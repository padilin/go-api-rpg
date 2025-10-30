package charDB

import (
	"os"
	"testing"
)

func TestSaveAndRetrieveCharacter(t *testing.T) {
	tmp, err := os.CreateTemp("", "testdb-*.db")
	if err != nil {
		t.Fatalf("failed to create temp db file: %v", err)
	}
	dbPath := tmp.Name()
	tmp.Close()
	defer os.Remove(dbPath)

	InitDB(dbPath)
	db := OpenDB(dbPath)

	c := Character{Name: "unittest"}
	saved, err := SaveCharacter(c, db)
	if err != nil {
		t.Fatalf("SaveCharacter failed: %v", err)
	}
	if saved.ID == 0 {
		t.Fatalf("expected saved character to have ID, got 0")
	}

	got, err := RetrieveCharacterById(saved.ID, db)
	if err != nil {
		t.Fatalf("RetrieveCharacterById failed: %v", err)
	}
	if got.Name != "unittest" {
		t.Fatalf("expected name 'unittest', got '%s'", got.Name)
	}

	got.Name = "updated"
	updated, err := SaveCharacter(got, db)
	if err != nil {
		t.Fatalf("SaveCharacter (update) failed: %v", err)
	}
	if updated.Name != "updated" {
		t.Fatalf("expected updated name 'updated', got '%s'", updated.Name)
	}
}

func TestSaveCharacterBulk(t *testing.T) {
	tmp, err := os.CreateTemp("", "testdb-*.db")
	if err != nil {
		t.Fatalf("failed to create temp db file: %v", err)
	}
	dbPath := tmp.Name()
	tmp.Close()
	defer os.Remove(dbPath)

	InitDB(dbPath)
	db := OpenDB(dbPath)

	chars := []Character{{Name: "bulk1"}, {Name: "bulk2"}}
	saved, err := SaveCharacterBulk(chars, db)
	if err != nil {
		t.Fatalf("SaveCharacterBulk failed: %v", err)
	}
	if len(saved) != 2 {
		t.Fatalf("expected 2 saved characters, got %d", len(saved))
	}
	for i := range saved {
		if saved[i].ID == 0 {
			t.Fatalf("expected saved[%d] to have ID, got 0", i)
		}
	}
}
