package word

type Word struct {
	Id      int64  `json:"id"      sql:"id"`
	English string `json:"english" sql:"english"`
	Chinese string `json:"chinese" sql:"chinese"`
}

type English string
type Chinese string

type WordService interface {
	SearchWord(s interface{}) (*Word, error)
}
