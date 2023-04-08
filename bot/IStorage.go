package bot

import "psa_dump_bot/model"

type Storage interface {
	GetConcerns() []model.Concern
	GetBrands(concern string) []model.Brand
	GetModels(brand string) []model.Model
	GetEngines(model string, brand string) []model.Engine
	GetBoltPatterns(model string, brand string) []model.BoltPattern
	GetAllRegions() []model.Region
	SaveUser(user *model.User) bool
	CheckUser(login string) bool
	UpdateUser() bool
}
