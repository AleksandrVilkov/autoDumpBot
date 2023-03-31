package postgreSQL

import "psa_dump_bot/bot"

type Storage struct {
	psql PostgreSQL
}

func NewStorage() *Storage {
	return &Storage{}
}
func (s *Storage) GetConcerns() []bot.Concern {
	return nil
}
func (s *Storage) GetBrands(concern string) []bot.Brand {
	return nil
}
func (s *Storage) GetModels(brand string) []bot.Model {
	return nil
}
func (s *Storage) GetEngines(model string, brand string) []bot.Engine {
	return nil
}
func (s *Storage) GetBoltPatterns(model string, brand string) []bot.BoltPattern {
	return nil
}
func (s *Storage) SaveUser() bool {
	return false
}
func (s *Storage) UpdateUser() bool {
	return false
}
