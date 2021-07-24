package main

import (
	"net/http"
	"os"
	"log"
	
	"github.com/joho/godotenv"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joseph-sx/kh-bot/commands"
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
	webhookUrl := (MainUrl+bot.Token)
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(webhookUrl))
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

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			
			switch update.Message.Command() {
			case "help":
				msg.Text = "Available Commands \n /joke \n /pokemon name \n /sayhi  \n /status."
			case "sayhi":
				msg.Text = "Hi :)"
			case "status":
				msg.Text = "I'm ok."
			case "joke":
				joke := commands.Joke()
				msg.ParseMode = "html"
				msg.Text = joke
			case "withArgument":
				msg.Text = "You supplied the following argument: " + update.Message.CommandArguments()
			case "pokemon":
				arg := update.Message.CommandArguments() 
				poke := commands.Pokemon(arg)
				msg.ParseMode = "markdown"
				if arg != ""{
					log.Printf("==================== { POKEMON } ====================")
					log.Println(poke)
					msg.Text = poke
				}else{
					msg.Text = "No pokemon name provided to fetch data"
				}
			case "html":
				msg.ParseMode = "html"
				msg.Text = "This will be interpreted as HTML, click <a href=\"https://www.example.com\">here</a>"
			default:
				msg.Text = "I don't know that command"
			}
			typing := tgbotapi.NewChatAction(update.Message.Chat.ID, tgbotapi.ChatTyping);
			bot.Send(typing)
			bot.Send(msg)
		}
	}
}
