package UI

import (
	Base "Bot/UI/Base"
)

var tournamentMenuConstructor = Base.MenuConstructor{
	Name:     "tournamentMenu",
	Template: &Base.Template{Code: "tournamentMenuPage", Text: "[Tournament menu]"},
	Items: []Base.MenuItem{
		{
			Name:        &Base.Template{Code: "create", Text: "create"},
			Constructor: Base.CHOP,
		}, {
			Name:        &Base.Template{Code: "yours", Text: "yours"},
			Constructor: Base.CHOP,
		}, {
			Name:        &Base.Template{Code: "matchmaking", Text: "matchmaking"},
			Constructor: Base.CHOP,
		}, {
			Name:        &Base.Template{Code: "running", Text: "running"},
			Constructor: Base.CHOP,
		}, {
			Name:        &Base.Template{Code: "ended", Text: "ended"},
			Constructor: Base.CHOP,
		},
	},
	IsHasParent: true,
}
