package dict

type DictEntry struct {
	Id          int    `json:"id" sql:"id"`
	Kanji       string `json:"kanji" sql:"kanji"`
	Kana        string `json:"kana" sql:"kana"`
	Translation string `json:"translation" sql:"translation"`
}

type DictEntryPayload struct {
	Kanji       string `json:"kanji"`
	Kana        string `json:"kana"`
	Translation string `json:"translation"`
}
