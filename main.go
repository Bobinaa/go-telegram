package main

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	botToken := "" // YOUR_BOT_TOKEN

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Авторизован как @%s", bot.Self.UserName)

	// Получаем текущую дату и время
	currentTime := time.Now()

	// Вычисляем время до следующего запланированного утра
	nextMorning := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()+1, 6, 30, 0, 0, time.Local)
	duration := nextMorning.Sub(currentTime)

	// Ожидаем время до следующего утра
	time.Sleep(duration)

	for {
		// Отправляем сообщение другу
		msg := tgbotapi.NewMessage(166085806, "Доброе утро") //  YOUR_FRIEND_CHAT_ID

		_, err := bot.Send(msg)
		if err != nil {
			log.Println(err)
		}

		// Ждем до следующего утра
		time.Sleep(24 * time.Hour)
	}
}
