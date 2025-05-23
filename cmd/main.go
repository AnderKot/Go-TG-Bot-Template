package main

import (
	Bot "Bot"
	Database "Bot/Database"
	"fmt"
	"log"
	"os"
	"time"

	tgBotAPI "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/ini.v1"
	"gorm.io/driver/sqlite"
)

func main() {
	f, err := os.OpenFile("log"+time.Now().Format("2006-01-02 15-04-05")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	config, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	runArgs := config.Section("runArgs")

	botAPI, err := tgBotAPI.NewBotAPI(runArgs.Key("BotKey").String())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	d := sqlite.Open("Database.db")
	database := Database.InitDatabase(d)
	defer database.Template.ExportTemplates()

	bot := Bot.NewBotService(botAPI, database)

	defer bot.Final()
	bot.Start()
}

//postgres.Open("host="+dbHostName+" port="+dbPort+" user="+dbLogin+" password="+dbPass+" dbname=service sslmode=disable")

/*
	dbHostName, exists := os.LookupEnv("DB_HOST_NAME")
	if !exists {
		if err := godotenv.Load(); err != nil {
			log.Print("No .env file found")
		}
		dbHostName, exists = os.LookupEnv("DB_HOST_NAME")
		if !exists {
			panic("DB_HOST_NAME")
		}
	}

	dbLogin, exists := os.LookupEnv("DB_USERS_USER")
	if !exists {
		panic("DB_USERS_USER")
	}

	dbPort, exists := os.LookupEnv("DB_USERS_PORT")
	if !exists {
		panic("DB_USERS_PORT")
	}

	dbPass, exists := os.LookupEnv("DB_PASSWORD")
	if !exists {
		panic("DB_PASSWORD")
	}
*/
