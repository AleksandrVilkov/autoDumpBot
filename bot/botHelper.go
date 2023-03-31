package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
	"math"
)

func CheckError(e error) {
	if e != nil {
		log.Print(e)
	}
}
func CheckFatalError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func CreateInlineKeyBoard(callback map[string]string, column int) tgbotapi.InlineKeyboardMarkup {
	var rowsCount = int(math.Ceil(float64(float64(len(callback)) / float64(column))))

	//TODO проверять кратность строк и колонок.
	//Если возможна пустая кнопка - подбирать допустимый размер. В дальнейшем убрать колонку, и определять автоматически
	rows := make([][]tgbotapi.InlineKeyboardButton, rowsCount)
	keys := make([]string, 0, len(callback))
	for k := range callback {
		keys = append(keys, k)
	}
	iter := 0
	for i := 0; i < rowsCount; i++ {
		row := make([]tgbotapi.InlineKeyboardButton, column)
		for k := 0; k < column; k++ {
			if iter < len(callback) {
				callbackData := callback[keys[iter]]
				row[k] = tgbotapi.InlineKeyboardButton{
					Text:                         keys[iter],
					URL:                          nil,
					CallbackData:                 &callbackData,
					SwitchInlineQuery:            nil,
					SwitchInlineQueryCurrentChat: nil,
					CallbackGame:                 nil,
					Pay:                          false,
				}
				iter++
			}
		}
		rows[i] = row
	}
	return tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}
}
