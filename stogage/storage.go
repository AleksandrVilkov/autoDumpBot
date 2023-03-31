package postgreSQL

import (
	"log"
	"psa_dump_bot/bot"
)

const (
	PARAMS_PATH = "/home/vilkov/GolandProjects/autoDumpBot/config/config.yaml"

	INSERT_INTO = "INSERT INTO "
	DELETE_FROM = "DELETE FROM"
	DISTINCT    = "DISTINCT "
	WHERE       = "WHERE "
	VALUES      = "VALUES "
	SELECT      = "SELECT "
	FROM        = "FROM "
	ALL         = "* "
)

type Storage struct {
	psql PostgreSQL
}

func NewStorage() *Storage {
	return &Storage{}
}
func (s *Storage) GetConcerns() []bot.Concern {
	query := SELECT + DISTINCT + "concern " + FROM + "CAR"
	resultSearch := s.psql.GetRows(query)

	var result []bot.Concern
	for resultSearch.Next() {
		var concern string
		err := resultSearch.Scan(&concern)

		if err != nil {
			log.Println("Error scan concern in func GetConcerns()")
			return result
		}

		result = append(result, bot.Concern{Concern: concern})
	}
	return result
}
func (s *Storage) GetBrands(concern string) []bot.Brand {
	//TODO
	return nil
}
func (s *Storage) GetModels(brand string) []bot.Model {
	//TODO
	return nil
}
func (s *Storage) GetEngines(model string, brand string) []bot.Engine {
	//TODO
	return nil
}
func (s *Storage) GetBoltPatterns(model string, brand string) []bot.BoltPattern {
	//TODO
	return nil
}
func (s *Storage) SaveUser() bool {
	//TODO
	return false
}
func (s *Storage) UpdateUser() bool {
	//TODO
	return false
}
