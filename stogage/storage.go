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
	AND         = "AND "

	CAR_TABLE_NAME    = "CAR "
	USER_TABLE_NAME   = "CLIENT "
	REGION_TABLE_NAME = "REGION "

	TEAMPLEATE_TIME = "2006-01-02"
)

type Storage struct {
	psql PostgreSQL
}

func NewStorage() *Storage {
	return &Storage{}
}
func (s *Storage) GetConcerns() []bot.Concern {
	query := SELECT + DISTINCT + "concern " + FROM + CAR_TABLE_NAME
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
	query := SELECT + DISTINCT + "brand " + FROM + CAR_TABLE_NAME + WHERE + "concern = '" + concern + "'"
	resultSearch := s.psql.GetRows(query)

	var result []bot.Brand

	for resultSearch.Next() {
		var brand string
		err := resultSearch.Scan(&brand)

		if err != nil {
			log.Println("Error scan brand in func GetBrands()")
			return result
		}
		result = append(result, bot.Brand{Brand: brand})
	}
	return result
}

func (s *Storage) GetModels(brand string) []bot.Model {
	query := SELECT + DISTINCT + "model " + FROM + CAR_TABLE_NAME + WHERE + "brand = '" + brand + "'"
	resultSearch := s.psql.GetRows(query)

	var result []bot.Model
	for resultSearch.Next() {
		var modelName string
		err := resultSearch.Scan(&modelName)
		if err != nil {
			log.Println("Error scan models in func GetModels()")
			return result
		}

		result = append(result, bot.Model{Model: modelName})
	}

	return result
}
func (s *Storage) GetEngines(model string, brand string) []bot.Engine {
	query := SELECT + DISTINCT + "engine " + FROM + CAR_TABLE_NAME + WHERE + "model ='" + model + "' " + AND + "brand ='" +
		brand + "'"

	resultSearch := s.psql.GetRows(query)
	var result []bot.Engine

	for resultSearch.Next() {
		var engineName string
		err := resultSearch.Scan(&engineName)
		if err != nil {
			log.Println("Error scan engine in func GetEngines()")
			return result
		}
		result = append(result, bot.Engine{EngineName: engineName})
	}
	return result
}
func (s *Storage) GetBoltPatterns(model string, brand string) []bot.BoltPattern {
	query := SELECT + DISTINCT + "boltPattern " + FROM + CAR_TABLE_NAME + WHERE + "model ='" + model + "' " + AND + "brand ='" +
		brand + "'"

	resultSearch := s.psql.GetRows(query)
	var result []bot.BoltPattern

	for resultSearch.Next() {
		var size string
		err := resultSearch.Scan(&size)
		if err != nil {
			log.Println("Error scan boltPattern in func GetBoltPatterns()")
			return result
		}
		result = append(result, bot.BoltPattern{
			BoltPatternSize: size,
		})
	}
	return result
}
func (s *Storage) SaveUser(u *bot.User) bool {

	//userQueryIncerst := INSERT_INTO + USER_TABLE_NAME +
	//	"(" + "createddate" + ", " +
	//	"role" + ", " +
	//	"login" + ", " +
	//	"lastname" + " ," +
	//	"regionid" + " ," +
	//	"carid" + ") " + VALUES + "(" +
	//	" '" + u.CreateDate.Format(TEAMPLEATE_TIME) + "', " +
	//	" '" + string(u.Role) + "', " +
	//	" '" + u.Login + "', " +
	//	" '" + strconv.Itoa(u.Id) + "', " +
	//	" '" + strconv.Itoa(u.Region.Id) +
	//	" '" + strconv.Itoa(u.UserCar.Id) + "')"
	//
	//result, err := s.psql.SendQuery(userQueryIncerst)
	//TODO

	return false
}

func (s *Storage) GetAllRegions() []bot.Region {
	query := SELECT + "*" + FROM + REGION_TABLE_NAME
	resultSearch := s.psql.GetRows(query)
	var result []bot.Region
	for resultSearch.Next() {
		var id string
		var name string
		err := resultSearch.Scan(&id, &name)
		if err != nil {
			log.Println("Error scan regions in func GetAllRegions()")
			return result
		}
		result = append(result, bot.Region{
			RegionName: name,
		})
	}
	return result
}
func (s *Storage) findCarId(concern bot.Concern, brand bot.Brand, model bot.Model, engine bot.Engine) int {
	query := SELECT + "*" + FROM + CAR_TABLE_NAME + WHERE + "concern ='" + concern.Concern + "' " + AND + "brand ='" + brand.Brand + "' " + AND +
		"engine ='" + engine.EngineName + "';"
	resultSearch := s.psql.GetRows(query)
	var carid int
	err := resultSearch.Scan(&carid)
	if err != nil {
		log.Println("Error find car ID in dataBase")
		return carid
	}
	return carid
}
func (s *Storage) UpdateUser() bool {
	//TODO
	return false
}
