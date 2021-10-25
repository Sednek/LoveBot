package main

import (
	botRepo "TelegramBot/pkg/bot"
	"TelegramBot/pkg/repo"
	"TelegramBot/pkg/service"
	"database/sql"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const (
	dbData = "user=sednek password=1 dbname=telegrambot sslmode=disable"
	botToken = "2069523748:AAHzvQaLpuUPx49i0hyepjBlQKx8vDctH78"
	ipApiKey = "f5bf67bb138398b5a6f96559d6932963"
)

func main(){
	//DB Connect
	db, err := sql.Open("postgres", dbData)
	if err != nil {
		log.Println(err)
	}
	log.Println("successful connect to database")
	rep, err := repo.NewRepo(db)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(rep)



	//Service initialize
	srv, err := service.NewService(ipApiKey)

	fmt.Println(srv)



	bot, err := botRepo.Connect(botToken)
	if err != nil {
		log.Println(err)
		return
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Println(err)
	}

	br, err := botRepo.NewBotRepo(bot, srv, updates)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(br)

	err = br.LoveBot()
	if err != nil {
		log.Println(err)
	}
}