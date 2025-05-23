package UI

import (
	Database "Bot/Database"
	Interface "Bot/Interface"
	Base "Bot/UI/Base"
)

func CreateMainMenu(userId int64, db Database.Database) Interface.IPage {
	p := Base.MenuConstructor{
		Name:     "mainMenu",
		Template: &Base.Template{Code: "mainMenuPage", Text: "[Main menu]"},
		Items: []Base.MenuItem{
			{
				Name:        &Base.Template{Code: "tournament", Text: "tournament"},
				Constructor: &tournamentMenuConstructor,
			}, {
				Name:        &Base.Template{Code: "yourActivity", Text: "yourActivity"},
				Constructor: Base.CHOP,
			}, {
				Name:        &Base.Template{Code: "profile", Text: "profile"},
				Constructor: CreateProfileFormConstructor(userId, db),
			},
		},
		IsHasParent: false,
	}

	return p.New()
}
