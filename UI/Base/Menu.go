package Base

import (
	Interface "Bot/Interface"
	"strconv"
)

// Затычка
var CHOP = &MenuConstructor{
	Name:        "CHOP",
	Template:    &Template{Code: "CHOP", Text: "[CHOP]"},
	Items:       []MenuItem{},
	IsHasParent: true,
}

// Реализация страниц меню >>
type MenuConstructor struct {
	Name        string
	Template    Interface.ITemplate
	Items       []MenuItem
	IsHasParent bool
}

type Menu struct {
	name           string
	template       Interface.ITemplate
	board          Interface.IKeyboard
	isHasParent    bool
	isNeedToParent bool

	items        []MenuItem
	selectedItem Interface.IConstructor
}

type MenuItem struct {
	Name        Interface.ITemplate
	Constructor Interface.IConstructor
}

func (lc *MenuConstructor) New() Interface.IPage {
	l := new(Menu)

	l.name = lc.Name
	l.template = lc.Template
	l.items = lc.Items
	l.isHasParent = lc.IsHasParent

	l.CreateKeyBoard()

	return l
}

func (p *Menu) CreateKeyBoard() {
	kb := Keyboard{Rows: make([]Interface.IKeyRow, 0)}
	kbr := KeyRow{make([]Interface.IKey, 0)}

	for index, item := range p.items {
		kbr.Keys = append(kbr.Keys, &Key{Name: item.Name, Data: strconv.Itoa(index)})
		if (index+1)%3 == 0 {
			kb.Rows = append(kb.Rows, &kbr)
			kbr = KeyRow{make([]Interface.IKey, 0)}
		}
	}
	kb.Rows = append(kb.Rows, &kbr)

	if p.isHasParent {
		kb.Rows = append(kb.Rows, &KeyRow{Keys: []Interface.IKey{
			Key{Name: OnBackToParentTemplate, Data: OnBackToParent},
		}})
	}

	p.board = &kb
}

// IPage >>

// Common
func (p *Menu) GetName() string {
	return p.name
}

// Input
func (p *Menu) OnProcessingMessage(text string) {
}

func (p *Menu) OnProcessingKey(keyData string) {
	switch keyData {
	case OnBackToParent:
		{
			p.isNeedToParent = true
		}
	default:
		{
			index, _ := strconv.Atoi(keyData)
			p.selectedItem = p.items[index].Constructor
		}
	}
}

// Navigation
func (p *Menu) OnGetNextPage() Interface.IConstructor {
	return p.selectedItem
}

func (p *Menu) OnBackToParent() bool {
	return p.isNeedToParent
}

// Print
func (p *Menu) GetMessageText() Interface.ITemplate {
	return p.template
}

func (p *Menu) GetKeyboard() Interface.IKeyboard {
	return p.board
}

// IPage <<
