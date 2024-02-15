package main

import (
	"fmt"
	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Run func(CallStack) CallStack
type CallStack struct {
	ChatID      int64
	Bot         *tgBotAPI.BotAPI
	Update      *tgBotAPI.Update
	Action      Run
	IsPrint     bool
	Parent 		*CallStack
	Data        string
}

// Данные на время жизни приложения
var userRuns = map[int64]CallStack{}
var userDatas = map[int64]*userData{}

func main() {
	fmt.Println("Cтарт Go")

	stopFlag := make(chan struct{})
	//Создаем бота
	bot, err := tgBotAPI.NewBotAPI("") //os.Getenv("TOKEN")
	if err != nil {
		panic(err)
	}

	go BotLoop(bot, stopFlag)
	go ConsoleLoop(bot, stopFlag)

	<-stopFlag
	fmt.Println("Cтоп Go")
}
