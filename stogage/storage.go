package postgreSQL

import "psa_dump_bot/bot"

type Storage struct {
	psql PostgreSQL
}

func (s *Storage) GetConcerns() []bot.Concern {

}
func (s *Storage) GetBrands(concern string) []bot.Brand {

}
func (s *Storage) GetModels(brand string) []bot.Model {

}
func (s *Storage) GetEngines(model string, brand string) []bot.Engine {

}
func (s *Storage) GetBoltPatterns(model string, brand string) []bot.BoltPattern {

}
func (s *Storage) saveUser() bool {

}
func (s *Storage) updateUser() bool {

}
