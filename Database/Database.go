package Database

import (
	Repository "Bot/Database/Repository"

	"gorm.io/gorm"
)

type Database struct {
	Template Repository.Translation
	Profile  Repository.Profile
}

func InitDatabase(dialector gorm.Dialector) Database {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {

		panic("failed to connect database")
	}

	return Database{
		Template: Repository.NewTemplateRepository(db).Init(),
		Profile:  Repository.NewProfileRepository(db).Init(),
	}
}
