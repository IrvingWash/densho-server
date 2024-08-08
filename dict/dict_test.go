package dict_test

import (
	"densho/dict"
	"testing"
)

func TestAddEntry(t *testing.T) {
	dictionary := dict.NewDict()

	dictionary.AddEntry(&dict.DictEntryPayload{
		Kanji:       "Hello",
		Kana:        "Haro",
		Translation: "Привет",
	})

	if len(*dictionary.Entries()) != 1 {
		t.Errorf("Expected 1 entry got %d", len(*dictionary.Entries()))
	}

	dictionary.AddEntry(&dict.DictEntryPayload{
		Kanji:       "Lollipop",
		Kana:        "Roripopu",
		Translation: "Леденец",
	})

	if (*dictionary.Entries())[1].Kanji != "Lollipop" {
		t.Errorf("Expected Lollipop, got %s", (*dictionary.Entries())[1].Kanji)
	}
}

func TestRemoveEntry(t *testing.T) {
	dictionary := dict.NewDict()

	dictionary.AddEntry(&dict.DictEntryPayload{
		Kanji:       "Hello",
		Kana:        "Haro",
		Translation: "Привет",
	})

	dictionary.AddEntry(&dict.DictEntryPayload{
		Kanji:       "Lollipop",
		Kana:        "Roripopu",
		Translation: "Леденец",
	})

	dictionary.RemoveEntry(0)

	if len(*dictionary.Entries()) != 1 {
		t.Errorf("Expected 1 but got %d", len(*dictionary.Entries()))
	}

	if (*dictionary.Entries())[0].Kanji != "Lollipop" {
		t.Errorf("Expected Lollipop but got %s", (*dictionary.Entries())[0].Kanji)
	}
}

func TestUpdateEntry(t *testing.T) {
	dictionary := dict.NewDict()

	dictionary.AddEntry(&dict.DictEntryPayload{
		Kanji:       "Hello",
		Kana:        "Haro",
		Translation: "Привет",
	})

	wrongEntry := (*dictionary.Entries())[0]

	wrongEntry.Kana = "Harro"

	dictionary.UpdateEntry(&wrongEntry)

	if len(*dictionary.Entries()) != 1 {
		t.Errorf("Expected 1 but got %d", len(*dictionary.Entries()))
	}

	if (*dictionary.Entries())[0].Kana != "Harro" {
		t.Errorf("Expected Harro but got %s", (*dictionary.Entries())[0].Kana)
	}
}
