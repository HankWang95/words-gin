package word

type Word struct {
	Id      int64  `json:"id"      sql:"id"`
	English string `json:"english" sql:"english"`
	Chinese string `json:"chinese" sql:"chinese"`
}

type WordService interface {
	SearchWordWithChinese(word string) (*Word, error)
	SearchWordWithEnglish(word string) (*Word, error)
}
