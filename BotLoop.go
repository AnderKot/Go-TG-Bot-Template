package main

import (
	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func BotLoop(bot *tgBotAPI.BotAPI, stopFlag chan<- struct{}) {
	//Устанавливаем время обновления
	u := tgBotAPI.NewUpdate(0)
	u.Timeout = 30

	//Получаем обновления от бота
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// Обработка сообщения
		go func(bot *tgBotAPI.BotAPI, update tgBotAPI.Update) {
			ID := GetChatID(update)
			if ID != 0 {
				stack := userRuns[ID]
				if stack.Action != nil {
					stack.Update = &update
					userRuns[ID] = userRuns[ID].Action(stack)
				} else {
					if update.Message != nil {
						userRuns[ID] = LoginMenu(CallStack{
							ChatID:  ID,
							Bot:     bot,
							Update:  &update,
							IsPrint: true,
						})
					}
				}
			}
		}(bot, update)
	}
	close(stopFlag)
}

var chopFile = tgBotAPI.FilePath("images/Chop.png")

func Chop(stack CallStack) CallStack {
	// Send "Work in progress"
	photo := tgBotAPI.NewPhoto(stack.ChatID, chopFile)
	photo.ReplyMarkup = tgBotAPI.NewRemoveKeyboard(true)
	_, _ = stack.Bot.Send(photo)
	// return on parent Run
	return ReturnOnParent(stack)
}

func GetChatID(update tgBotAPI.Update) int64 {
	if update.Message != nil {
		return update.Message.Chat.ID
	}

	if update.CallbackQuery != nil {
		return update.CallbackQuery.Message.Chat.ID
	}

	return -1
}

func ReturnOnParent(stack CallStack) CallStack {
	if stack.Parent != nil {
		stack.Parent.IsPrint = true
		stack.Parent.Update = nil
		return stack.Parent.Action(*stack.Parent)
	}
	return RunTemplate(CallStack{
		IsPrint: true,
		ChatID:  stack.ChatID,
		Bot:     stack.Bot,
	})
}
