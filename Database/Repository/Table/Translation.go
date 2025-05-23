package Table

type Translation struct {
	Code     string `gorm:"primaryKey;" csv:"code"`
	LangCode string `gorm:"primaryKey;" csv:"lang_code"`
	Text     string `csv:"text"`
}
