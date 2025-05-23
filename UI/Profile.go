package UI

import (
	Database "Bot/Database"
	Interface "Bot/Interface"
	Base "Bot/UI/Base"
	Model "Bot/UI/Model"
)

func CreateProfileFormConstructor(userId int64, db Database.Database) Interface.IConstructor {
	return &Base.FormConstructor{
		Name:        "profilePageForm",
		Model:       &(Model.Profile{Database: db, ID: userId}),
		IsHasParent: true,
	}
}
