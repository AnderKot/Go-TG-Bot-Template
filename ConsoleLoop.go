package main

import (
	"fmt"
	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func ConsoleLoop(bot *tgBotAPI.BotAPI, stopFlag chan<- struct{}) {
	var command string
	for command != "stop" {
		fmt.Scan(&command)

	Switch:
		switch command {
		case "STAM" +
			"":
			{
				var text string
				fmt.Scan(&text)
				if text != "" {
					for chatId, data := range userDatas {
						text = fmt.Sprintf(SelectTemplate("LoginMenu", data.languageСode))
						if text == "Template not specified !\n" {
							break Switch
						}
						msg := tgBotAPI.NewMessage(chatId, text)
						_, _ = bot.Request(msg)
					}
				}
			}
		default:
			fmt.Println("Команда не распозана")
			fmt.Println("Список команд:")
			fmt.Println("STAM - Отправисть всем пользователям сообщение")
		}
	}
	close(stopFlag)
}
