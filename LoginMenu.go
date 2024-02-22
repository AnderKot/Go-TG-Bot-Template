package main

import (
	"fmt"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func LoginMenu(stack CallStack) CallStack {
	data := userDatas[stack.ChatID]
	stack.Action = LoginMenu
	if data == nil {
		if RequestData(stack) {
			return RunTemplate(CallStack{
				ChatID:  stack.ChatID,
				Bot:     stack.Bot,
				IsPrint: true,
			})
		}
		LanguageCode := "eng"

		if stack.Update.Message != nil {
			LanguageCode = stack.Update.Message.From.LanguageCode
		}

		if stack.Update.CallbackQuery != nil {
			LanguageCode = stack.Update.CallbackQuery.Message.From.LanguageCode
		}

		text := fmt.Sprintf(GetTemplate("LoginMenu", LanguageCode))
		msg := tgBotAPI.NewMessage(stack.ChatID, text)
		_, _ = stack.Bot.Request(msg)

		return stack
	}

	return RunTemplate(CallStack{
		ChatID:  stack.ChatID,
		Bot:     stack.Bot,
		IsPrint: true,
	})
}

func RequestData(stack CallStack) bool {
	if stack.Update != nil {
		if stack.Update.Message == nil {
			return false
		}

		if stack.Update.Message.Text == "🔑" {
			data := userData{
				userName:     stack.Update.Message.From.UserName,
				firstName:    stack.Update.Message.From.FirstName,
				lastName:     stack.Update.Message.From.LastName,
				languageСode: stack.Update.Message.From.LanguageCode,
			}
			userDatas[stack.ChatID] = &data
			return true
		}
	}
	return false
}
