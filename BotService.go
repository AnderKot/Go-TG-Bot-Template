package Bot

import (
	Database "Bot/Database"
	Interface "Bot/Interface"
	Stack "Bot/Stack"
	"fmt"
	"sync"

	UI "Bot/UI"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewBotService(botAPI *tgBotAPI.BotAPI, database Database.Database) BotService {
	return BotService{
		BotAPI:           botAPI,
		UserMutexes:      map[int64]*sync.Mutex{},
		UserPagesStack:   map[int64]Stack.MyStack[Interface.IPage]{},
		UserChatMessages: map[int64]Message{},
		Database:         database,
	}
}

type Message struct {
	id       int
	text     string
	keyboard Interface.IKeyboard
}

type BotService struct {
	BotAPI           *tgBotAPI.BotAPI
	UserMutexes      map[int64]*sync.Mutex
	UserPagesStack   map[int64]Stack.MyStack[Interface.IPage]
	UserChatMessages map[int64]Message
	Database         Database.Database
}

func (bs *BotService) Start() {
	u := tgBotAPI.NewUpdate(0)
	u.Timeout = 20

	//Получаем обновления от бота
	updates := bs.BotAPI.GetUpdatesChan(u)

	for u := range updates {
		// Обработка сообщения
		go func(bot *tgBotAPI.BotAPI, update tgBotAPI.Update) {
			chatID := GetChatIDFromUpdate(update)
			langCode := GetLangCodeFromUpdate(update)

			SaveLock(bs.UserMutexes[chatID])

			stack := bs.UserPagesStack[chatID]
			exist, page := stack.DeStack()
			if !exist {
				p := UI.CreateMainMenu(chatID, bs.Database)
				page = &p
			}

			// Обработка сообщений
			if update.CallbackQuery != nil {
				(*page).OnProcessingKey(update.CallbackQuery.Data)
			}
			if update.Message != nil {
				(*page).OnProcessingMessage(update.Message.Text)
			}

			// Навигация по страницам
			nextPageConstructor := (*page).OnGetNextPage()
			if nextPageConstructor != nil {
				stack.AddStack(page)
				p := nextPageConstructor.New()
				page = &p
			}

			if (*page).OnBackToParent() {
				oldExist, oldPage := stack.DeStack()
				if oldExist {
					page = oldPage
				} else {
					p := UI.CreateMainMenu(chatID, bs.Database)
					page = &p
				}
			}

			// Ответ пользователю
			t := (*page).GetMessageText()
			newText := bs.Database.Template.GetTemplateText(t, langCode)
			args := t.GetArgs()
			if len(args) > 0 {
				newText = fmt.Sprintf(newText, args)
			}

			newKeyboard := (*page).GetKeyboard()

			oldMessage, isOldExist := bs.UserChatMessages[chatID]
			if isOldExist {
				if oldMessage.text != newText {
					isOk := bs.sendEdit(chatID, oldMessage.id, newText, bs.GenerateKeyboard(newKeyboard, langCode), false)
					if isOk {
						bs.UserChatMessages[chatID] = Message{
							id:       oldMessage.id,
							text:     newText,
							keyboard: newKeyboard,
						}
					} else {
						panic("send edit message error")
					}
				}
			} else {
				isOk, messageID := bs.sendNew(chatID, newText, bs.GenerateKeyboard(newKeyboard, langCode), false)
				if isOk {
					bs.UserChatMessages[chatID] = Message{
						id:       messageID,
						text:     newText,
						keyboard: newKeyboard,
					}
				} else {
					panic("send new message error")
				}
			}

			stack.AddStack(page)
			bs.UserPagesStack[chatID] = stack
			SaveUnlock(bs.UserMutexes[chatID])
		}(bs.BotAPI, u)
	}
}

func (bs *BotService) Final() {

}

func (bs *BotService) sendNew(chat int64, text string, keyboard *tgBotAPI.InlineKeyboardMarkup, isMarkdown bool) (bool, int) {
	NewMsgRequest := tgBotAPI.NewMessage(chat, text)
	if isMarkdown {
		NewMsgRequest.ParseMode = "Markdown"
	}
	if keyboard != nil {
		if keyboard.InlineKeyboard != nil {
			NewMsgRequest.ReplyMarkup = keyboard
		}
	}
	NewMsgRespons, err := bs.BotAPI.Send(NewMsgRequest)
	if err != nil {
		return false, 0
	}

	return true, NewMsgRespons.MessageID
}

func (bs *BotService) sendEdit(chat int64, oldMessageID int, text string, keyboard *tgBotAPI.InlineKeyboardMarkup, isMarkdown bool) bool {
	NewMsgRequest := tgBotAPI.NewEditMessageText(chat, oldMessageID, text)
	if isMarkdown {
		NewMsgRequest.ParseMode = "Markdown"
	}
	if keyboard != nil {
		if keyboard.InlineKeyboard != nil {
			NewMsgRequest.ReplyMarkup = keyboard
		}
	}
	_, err := bs.BotAPI.Request(NewMsgRequest)
	if err != nil {
		return err.(*tgBotAPI.Error).Message == "Bad Request: message is not modified: specified new message content and reply markup are exactly the same as a current content and reply markup of the message"
	}

	return true
}

func (bs *BotService) deleteMsg(chat int64, messageID int) {
	_, err := bs.BotAPI.Request(tgBotAPI.NewDeleteMessage(chat, messageID))
	if err != nil {
		return
	}
}

func (bs *BotService) GenerateKeyboard(kbd Interface.IKeyboard, langCode string) *tgBotAPI.InlineKeyboardMarkup {
	var botKeyboard tgBotAPI.InlineKeyboardMarkup
	rows := kbd.GetRows()
	for _, row := range rows {
		var botRow []tgBotAPI.InlineKeyboardButton
		keys := row.GetKeys()
		for _, key := range keys {
			botRow = append(botRow, tgBotAPI.NewInlineKeyboardButtonData(bs.Database.Template.GetTemplateText(key.GetTemplate(), langCode), key.GetData()))
		}
		if len(botRow) > 0 {
			botKeyboard.InlineKeyboard = append(botKeyboard.InlineKeyboard, botRow)
		}
	}
	if len(botKeyboard.InlineKeyboard) > 0 {
		return &botKeyboard
	}
	return nil
}
