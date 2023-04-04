package bot

import tgbotapi "github.com/Syfaro/telegram-bot-api"

func getMsgOnlyTextForChannel(e *Environment, test string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(e.Config.ValidateData.ChannelID, test)
}
func getMsgWithPhotoForChannel(e *Environment, photo []byte) tgbotapi.PhotoConfig {

	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photo,
	}
	return tgbotapi.NewPhotoUpload(e.Config.ValidateData.ChannelID, photoFileBytes)
}

func getMsgWithPhotoForChannelWithCaption(e *Environment, photo []byte, text string) tgbotapi.PhotoConfig {

	photoFileBytes := tgbotapi.FileBytes{
		Name:  "picture",
		Bytes: photo,
	}

	msg := tgbotapi.NewPhotoUpload(e.Config.ValidateData.ChannelID, photoFileBytes)
	msg.Caption = text
	return msg
}

func getMsgWithSomePhotosForChannelWithCaption(e *Environment, photos [][]byte, text string) tgbotapi.MediaGroupConfig {

	//TODO not working
	photoFilesBytes := make([]interface{}, len(photos))
	for i := 0; i < len(photos); i++ {
		photoFilesBytes[i] = tgbotapi.PhotoConfig{
			BaseFile: tgbotapi.BaseFile{
				BaseChat:    tgbotapi.BaseChat{},
				File:        photos[i],
				FileID:      "",
				UseExisting: false,
				MimeType:    "",
				FileSize:    0,
			},
			Caption:   "",
			ParseMode: "",
		}
	}

	var msg = tgbotapi.NewMediaGroup(e.Config.ValidateData.ChannelID, photoFilesBytes)
	return msg
}
