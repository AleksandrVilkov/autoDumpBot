package bot

import (
	"psa_dump_bot/model"
)

type TempStorage interface {
	SaveTempData(token string, c *model.CallBack) bool
	FindTempDataByToken(token string) *model.CallBack
}
