package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var gBot *tgbotapi.BotAPI
var gToken string
var gChatId int64

// var COLORS map[string]string
var COLORS = map[string]string{
	"красный":       "#FF0000",
	"малиновый":     "#DC143C",
	"розовый":       "#FFC0CB",
	"фуксия":        "#FF00FF",
	"коралловый":    "#FF7F50",
	"оранжевый":     "#FFA500",
	"жёлтый":        "#FFFF00",
	"желтый":        "#FFFF00",
	"золотый":       "#FFD700",
	"персиковый":    "#FFDAB9",
	"лавандовый":    "#E6E6FA",
	"фиолетовый":    "#EE82EE",
	"пурпурный":     "#800080",
	"индиго":        "#4B0082",
	"бежевый":       "#F5F5DC",
	"лайм":          "#00FF00",
	"зеленый":       "#008000",
	"зелёный":       "#008000",
	"оливковый":     "#808000",
	"голубой":       "#00FFFF",
	"аквамариновый": "#7FFFD4",
	"бирюзовый":     "#40E0D0",
	"синий":         "#0000FF",
	"белый":         "#FFFFFF",
	"серый":         "#808080",
	"черный":        "#000000",
	"чёрный":        "#000000",
	"серебряный":    "#C0C0C0",
	"бордовый":      "#800000",
}

func init() {

	os.Setenv(TOKEN_NAME_IN_OS, "6412549975:AAH49NSEAt0H933mFtfeMM5ls8rLQKeWPG4")
	gToken = os.Getenv(TOKEN_NAME_IN_OS)

	if gToken = os.Getenv(TOKEN_NAME_IN_OS); gToken == "" {
		panic(fmt.Errorf(`failed to load environment variable "%s"`, TOKEN_NAME_IN_OS))
	}

	var err error
	if gBot, err = tgbotapi.NewBotAPI(gToken); err != nil {
		log.Panic(err)
	}
	gBot.Debug = true
}

func isStartMessage(update *tgbotapi.Update) bool {
	return update.Message != nil && update.Message.Text == "/start"
}

func PrintSystemMessage(message string) {
	gBot.Send(tgbotapi.NewMessage(gChatId, message))
}

func printIntro(update *tgbotapi.Update) {
	PrintSystemMessage("Привет!")
	PrintSystemMessage("Этот бот поможет тебе получить шестнадцатиричный код цвета по его названию")
}

func giveCode() {

}

func main() {
	log.Printf("Authorized on account %s", gBot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := gBot.GetUpdatesChan(updateConfig)

	for update := range updates {

		colorName := strings.ToLower(update.Message.Text)
		colorCode, found := COLORS[colorName]

		if found {
			PrintSystemMessage(colorCode)
		} else if isStartMessage(&update) { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			gChatId = update.Message.Chat.ID
			printIntro(&update)

			//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			//msg.ReplyToMessageID = update.Message.MessageID

			//gBot.Send(msg)
		} else {
			PrintSystemMessage("К сожалению, такого цвета не нашлось!")
		}
	}
}
