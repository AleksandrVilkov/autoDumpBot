package model

type TempStorage interface {
	SaveTempData(token string, c *CallBack) bool
	FindTempDataByToken(token string) *CallBack
}
