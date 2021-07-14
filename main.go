package main

import (
	"net/http"
	"os"
	"log"
	"github.com/joho/godotenv"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	err := godotenv.Load()
	tgToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	MainUrl := os.Getenv("MAIN_URL")


	bot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		log.Printf("$TELEGRAM_BOT_TOKEN Must Be Set")
		log.Printf(tgToken);
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(MainUrl+bot.Token))
	if err != nil {
		log.Fatal("error al setear webhook")
		log.Fatal(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal("error al obtener info de webhook")
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	updates := bot.ListenForWebhook("/" + bot.Token)
	// go http.ListenAndServeTLS("0.0.0.0:"+ os.Getenv("PORT"), "cert.pem", "key.pem", nil)
	go http.ListenAndServe("0.0.0.0:"+ os.Getenv("PORT"), nil)

	for update := range updates {
		log.Printf("==================== { UPDATE } ====================")
		// log.Printf("%+v\n", update)
		if update.Message == nil { // ignore any non-Message Updates
			log.Printf("El mensaje viene vacío")
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		typing := tgbotapi.NewChatAction(update.Message.Chat.ID, tgbotapi.ChatTyping);
		bot.Send(typing)
		bot.Send(msg)
	}
}