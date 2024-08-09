package dict

import "errors"

type Dict struct {
	entries []DictEntry
}

func NewDict() Dict {
	return Dict{}
}

func (d *Dict) Entries() *[]DictEntry {
	return &d.entries
}

func (d *Dict) AddEntry(entry *DictEntryPayload) {
	d.entries = append(
		d.entries,
		DictEntry{
			Id:          len(d.entries),
			Kanji:       entry.Kanji,
			Kana:        entry.Kana,
			Translation: entry.Translation,
		},
	)
}

func (d *Dict) RemoveEntry(id int) {
	var filteredEntries []DictEntry

	for _, e := range d.entries {
		if e.Id != id {
			filteredEntries = append(filteredEntries, e)
		}
	}

	d.entries = filteredEntries
}

func (d *Dict) UpdateEntry(updatedEntry *DictEntry) error {
	var entryToUpdateId = -1

	for i, e := range d.entries {
		if e.Id == updatedEntry.Id {
			entryToUpdateId = i
		}
	}

	if entryToUpdateId == -1 {
		return errors.New("entry to update not found")
	}

	d.entries[entryToUpdateId] = *updatedEntry

	return nil
}
