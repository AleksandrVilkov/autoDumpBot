package bot

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"psa_dump_bot/bot/model"
)

func getMsgOnlyTextForChannel(e *model.Environment, test string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(e.Config.ValidateData.ChannelID, test)
}
func getMsgWithPhotoForChannel(e *model.Environment, photo []byte) tgbotapi.PhotoConfig {

	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photo,
	}
	return tgbotapi.NewPhotoUpload(e.Config.ValidateData.ChannelID, photoFileBytes)
}

func getMsgWithPhotoForChannelWithCaption(e *model.Environment, photo []byte, text string) tgbotapi.PhotoConfig {

	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photo,
	}

	msg := tgbotapi.NewPhotoUpload(e.Config.ValidateData.ChannelID, photoFileBytes)
	msg.Caption = text
	return msg
}

func getMsgWithSomePhotosForChannelWithCaption(e *model.Environment, msgs []tgbotapi.Message) tgbotapi.MediaGroupConfig {

	var listMediaVideoInput []interface{}

	for i := 0; i < len(msgs); i++ {
		msg := msgs[i]
		photos := *msg.Photo
		photo := photos[0]
		listMediaVideoInput = append(listMediaVideoInput, tgbotapi.NewInputMediaPhoto(photo.FileID))
	}

	return tgbotapi.NewMediaGroup(e.Config.ValidateData.ChannelID, listMediaVideoInput)
}
