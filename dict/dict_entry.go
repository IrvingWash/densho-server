package dict

type DictEntry struct {
	Id          int
	Kanji       string
	Kana        string
	Translation string
}

type DictEntryPayload struct {
	Kanji       string
	Kana        string
	Translation string
}
