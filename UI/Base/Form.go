package Base

import (
	Database "Bot/Database"
	Interface "Bot/Interface"
)

// Реализация формы с моделью
type FormConstructor struct {
	DB          Database.Database
	Name        string
	Model       Interface.IModel
	IsHasParent bool
}

func (lc *FormConstructor) New() Interface.IPage {
	f := new(Form)

	f.DB = lc.DB
	f.name = lc.Name
	f.model = lc.Model
	f.isHasParent = lc.IsHasParent

	if f.model.Init() {
		f.CreateKeyBoard()
	} else {
		f.isNeedToParent = true
	}

	return f
}

type Form struct {
	DB             Database.Database
	name           string
	board          Interface.IKeyboard
	isHasParent    bool
	isNeedToParent bool

	model Interface.IModel
}

func (p *Form) CreateKeyBoard() {
	kb := Keyboard{}
	kbr := KeyRow{}

	modelItems := p.model.GetButtons()
	for index, item := range modelItems {
		kbr.Keys = append(kbr.Keys, &Key{Name: item.GetTemplate(), Data: item.GetData()})
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

func (p *Form) GetName() string {
	return p.name
}

func (p *Form) OnProcessingMessage(text string) {
	p.model.OnProcessingMessage(text)
}

func (p *Form) OnProcessingKey(keyData string) {
	switch keyData {
	case OnBackToParent:
		{
			p.isNeedToParent = true
		}
	default:
		{
			p.model.OnProcessingKey(keyData)
		}
	}
}

func (p *Form) OnGetNextPage() Interface.IConstructor {
	return p.model.GetNextPage()
}

func (p *Form) OnBackToParent() bool {
	if p.isNeedToParent {
		p.model.Finish()
		return true
	}
	return p.model.IsNeedToParent()
}

func (p *Form) GetMessageText() Interface.ITemplate {
	return p.model.GetTemplate()
}

func (p *Form) GetKeyboard() Interface.IKeyboard {
	return p.board
}
