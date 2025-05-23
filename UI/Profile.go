package UI

import (
	Database "Bot/Database"
	Interface "Bot/Interface"
	Base "Bot/UI/Base"
	Controller "Bot/UI/Controller"
)

func CreateProfileFormConstructor(userId int64, db Database.Database) Interface.IConstructor {
	return &Base.FormConstructor{
		Name:        "profilePageForm",
		Controller:  &(Controller.Profile{Database: db, ID: userId}),
		IsHasParent: true,
	}
}
