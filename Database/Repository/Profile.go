package Repository

import (
	Table "Bot/Database/Repository/Table"
	"log"

	"gorm.io/gorm"
)

type Profile struct {
	database *gorm.DB
}

func NewProfileRepository(_database *gorm.DB) Profile {
	return Profile{
		database: _database,
	}
}

func (r Profile) Init() Profile {
	r.database.AutoMigrate(&Table.Profile{})
	return r
}

func (r Profile) CreateProfile(tp Table.Profile) error {
	err := r.database.Create(&tp).Error
	return err
}

func (r Profile) GetProfile(id int64) (Table.Profile, error) {
	tp := Table.Profile{Id: id}
	err := r.database.Where(tp).Take(&tp).Error
	if err != nil {
		log.Fatal(err)
		return tp, err
	}
	return tp, nil
}

func (r Profile) UpdateProfile(tp Table.Profile) error {
	err := r.database.Save(&tp).Error
	return err
}
