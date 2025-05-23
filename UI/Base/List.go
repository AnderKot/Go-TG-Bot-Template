package Base

import (
	Interface "Bot/Interface"
	"strconv"
	"strings"
)

// Hard code template
const (
	OnNextListPage string = "onNextListPage"
	OnPrefListPage string = "onPrefListPage"
)

type ListItem struct {
	Name        string
	Description string
	Constructor Interface.IConstructor
}

type Lister interface {
	GetListItems(no int) []ListItem
	GetNextIsExists(no int) bool
}

type ListConstructor struct {
	Name                string
	NextPageConstructor Interface.IConstructor
	Lister              Lister
	Columns             int
}

func (lc *ListConstructor) New() Interface.IPage {
	l := new(List)

	l.name = lc.Name
	l.lister = lc.Lister
	l.columns = lc.Columns

	l.CreateKeyBoardAndMessage()

	return l
}

type List struct {
	name           string
	messageBuilder strings.Builder
	board          Interface.IKeyboard
	isNeedToParent bool

	lister       Lister
	items        []ListItem
	selectedItem Interface.IConstructor
	curItemsNo   int

	columns int
}

func (l *List) CreateKeyBoardAndMessage() {
	l.messageBuilder.Reset()
	//l.messageText = l.repository.GetTemplate(l.name)
	kb := Keyboard{Rows: make([]Interface.IKeyRow, 0)}
	kbr := KeyRow{make([]Interface.IKey, 0)}

	if l.lister.GetNextIsExists(l.curItemsNo) {
		l.items = l.lister.GetListItems(l.curItemsNo)
		for index, item := range l.items {
			l.items = append(l.items)
			strIndex := "[" + strconv.Itoa(index+1) + "]"
			_, _ = l.messageBuilder.WriteString("\n" + strIndex + ": " + item.Description)
			kbr.Keys = append(kbr.Keys, &Key{Name: Template{Text: strIndex}, Data: item.Name})
			if (index+1)%l.columns == 0 {
				kb.Rows = append(kb.Rows, &kbr)
				kbr = KeyRow{make([]Interface.IKey, 0)}
				_, _ = l.messageBuilder.WriteString("\n")
			}
		}
		kbr = KeyRow{
			Keys: []Interface.IKey{
				Key{
					Name: OnNextListPageTemplate,
					Data: OnNextListPage,
				}}}
		kb.Rows = append(kb.Rows, &kbr)
	}

	if l.curItemsNo > 0 {
		kbr.Keys = append(kbr.Keys, Key{
			Name: OnPrefListPageTemplate,
			Data: OnPrefListPage,
		})
	}

	kb.Rows = append(kb.Rows, &KeyRow{Keys: []Interface.IKey{
		Key{Name: OnPrefListPageTemplate, Data: OnBackToParent},
	}})

	l.board = &kb
}

// IPage >>
// Common
func (l *List) GetName() string {
	return l.name
}

// Input
func (l *List) OnProcessingMessage(text string) {
}

func (l *List) OnProcessingKey(keyData string) {
	switch keyData {
	case OnBackToParent:
		{
			l.isNeedToParent = true
		}
	case OnNextListPage:
		{
			l.curItemsNo++
			l.CreateKeyBoardAndMessage()
		}
	case OnPrefListPage:
		{
			if l.curItemsNo > 0 {
				l.curItemsNo--
				l.CreateKeyBoardAndMessage()
			}
		}
	default:
		{
			index, _ := strconv.Atoi(keyData)
			l.selectedItem = l.items[index].Constructor
		}
	}
}

// Navigation
func (l *List) OnGetNextPage() Interface.IConstructor {
	return l.selectedItem
}

func (l *List) OnBackToParent() bool {
	return l.isNeedToParent
}

// Print
func (l *List) GetMessageText() Interface.ITemplate {
	return &Template{Text: l.messageBuilder.String()}
}

func (l *List) GetKeyboard() Interface.IKeyboard {
	return l.board
}

// IPage <<
