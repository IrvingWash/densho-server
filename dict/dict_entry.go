package dict

type DictEntry struct {
	Id          int    `json:"id"`
	Kanji       string `json:"kanji"`
	Kana        string `json:"kana"`
	Translation string `json:"translation"`
}

type DictEntryPayload struct {
	Kanji       string `json:"kanji"`
	Kana        string `json:"kana"`
	Translation string `json:"translation"`
}
