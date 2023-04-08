package bot

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
	"math"
	"psa_dump_bot/model"
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

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetMD5HashFromCallBack(c *model.CallBack) string {
	hasher := md5.New()
	data, _ := json.Marshal(c)
	hasher.Write(data)
	return hex.EncodeToString(hasher.Sum(nil))
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
