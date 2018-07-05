package form

type WordForm struct {
	Word string `form:"word"`
	Translation string `form:"translation"`
}

