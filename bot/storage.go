package bot

type Storage interface {
	GetConcerns() []Concern
	GetBrands(concern string) []Brand
	GetModels(brand string) []Model
	GetEngines(model string, brand string) []Engine
	GetBoltPatterns(model string, brand string) []BoltPattern
	saveUser() bool
	updateUser() bool
}
