package Table

type Profile struct {
	Id             int64  `gorm:"primaryKey;autoIncrement:false"`
	NickName       string `gorm:"default:''"`
	SteamID        string `gorm:"default:''"`
	IsValidSteamID bool   `gorm:"default:false"`
	LanguageCode   string `gorm:"default:'en'"`
}
