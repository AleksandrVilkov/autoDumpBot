package postgreSQL

import (
	"encoding/json"
	"log"
	model2 "psa_dump_bot/model"
	"strconv"
	"time"
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

	CAR_TABLE_NAME      = "CAR "
	USER_TABLE_NAME     = "CLIENT "
	REGION_TABLE_NAME   = "REGION "
	TEMPDATA_TABLE_NAME = "TEMPDATA "

	TEAMPLEATE_TIME = "2006-01-02T15:04:05Z07:00"
)

type Storage struct {
	psql PostgreSQL
}

func NewStorage() *Storage {
	return &Storage{}
}
func (s *Storage) CheckUser(login string) bool {
	query := SELECT + "login " + FROM + USER_TABLE_NAME + WHERE + "login = '" + login + "'"
	resultSearch := s.psql.GetRows(query)
	return resultSearch.Next()
}

func (s *Storage) SaveTempData(token string, c *model2.CallBack) bool {
	query := INSERT_INTO + TEMPDATA_TABLE_NAME +
		"(" + "token" + ", " +
		"createdDate" + ", " +
		"callback" + ") " + VALUES + "(" +
		" '" + token + "', " +
		" '" + time.Now().Format(TEAMPLEATE_TIME) + "', " +
		" '" + c.ToString() + "');"
	_, err := s.psql.SendQuery(query)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
func (s *Storage) FindTempDataByToken(token string) *model2.CallBack {
	query := SELECT + "callback " + FROM + TEMPDATA_TABLE_NAME + WHERE + "token = '" + token + "'"
	resultSearch := s.psql.GetRows(query)

	count := 0
	var stringCallback string
	for resultSearch.Next() {
		if count != 0 {
			log.Println("More than one callback found for the same token")
			return &model2.CallBack{}
		}
		err := resultSearch.Scan(&stringCallback)
		if err != nil {
			log.Println(err)

			return &model2.CallBack{}
		}
		count++
	}
	var result model2.CallBack
	err := json.Unmarshal([]byte(stringCallback), &result)
	if err != nil {
		log.Println(err)
		return &model2.CallBack{}
	}

	return &result
}
func (s *Storage) GetConcerns() []model2.Concern {
	query := SELECT + DISTINCT + "concern " + FROM + CAR_TABLE_NAME
	resultSearch := s.psql.GetRows(query)

	var result []model2.Concern
	for resultSearch.Next() {
		var concern string
		err := resultSearch.Scan(&concern)

		if err != nil {
			log.Println("Error scan concern in func GetConcerns()")
			return result
		}

		result = append(result, model2.Concern{Concern: concern})
	}
	return result
}
func (s *Storage) GetBrands(concern string) []model2.Brand {
	query := SELECT + DISTINCT + "brand " + FROM + CAR_TABLE_NAME + WHERE + "concern = '" + concern + "'"
	resultSearch := s.psql.GetRows(query)

	var result []model2.Brand

	for resultSearch.Next() {
		var brand string
		err := resultSearch.Scan(&brand)

		if err != nil {
			log.Println("Error scan brand in func GetBrands()")
			return result
		}
		result = append(result, model2.Brand{Brand: brand})
	}
	return result
}

func (s *Storage) GetModels(brand string) []model2.Model {
	query := SELECT + DISTINCT + "model " + FROM + CAR_TABLE_NAME + WHERE + "brand = '" + brand + "'"
	resultSearch := s.psql.GetRows(query)

	var result []model2.Model
	for resultSearch.Next() {
		var modelName string
		err := resultSearch.Scan(&modelName)
		if err != nil {
			log.Println("Error scan models in func GetModels()")
			return result
		}

		result = append(result, model2.Model{Model: modelName})
	}

	return result
}
func (s *Storage) GetEngines(carModel string, brand string) []model2.Engine {
	query := SELECT + DISTINCT + "engine " + FROM + CAR_TABLE_NAME + WHERE + "model ='" + carModel + "' " + AND + "brand ='" +
		brand + "'"

	resultSearch := s.psql.GetRows(query)
	var result []model2.Engine

	for resultSearch.Next() {
		var engineName string
		err := resultSearch.Scan(&engineName)
		if err != nil {
			log.Println("Error scan engine in func GetEngines()")
			return result
		}
		result = append(result, model2.Engine{EngineName: engineName})
	}
	return result
}
func (s *Storage) GetBoltPatterns(carModel string, brand string) []model2.BoltPattern {
	query := SELECT + DISTINCT + "boltPattern " + FROM + CAR_TABLE_NAME + WHERE + "model ='" + carModel + "' " + AND + "brand ='" +
		brand + "'"

	resultSearch := s.psql.GetRows(query)
	var result []model2.BoltPattern

	for resultSearch.Next() {
		var size string
		err := resultSearch.Scan(&size)
		if err != nil {
			log.Println("Error scan boltPattern in func GetBoltPatterns()")
			return result
		}
		result = append(result, model2.BoltPattern{
			BoltPatternSize: size,
		})
	}
	return result
}
func (s *Storage) SaveUser(u *model2.User) bool {
	carID := s.findCarId(u.UserCar.Concern, u.UserCar.Brand, u.UserCar.Model, u.UserCar.Engine)
	if carID == 0 {
		log.Println("failed to get car id from input")
		return false
	}
	userQueryIncerst := INSERT_INTO + USER_TABLE_NAME +
		"(" + "createddate" + ", " +
		"role" + ", " +
		"login" + ", " +
		"regionid" + " ," +
		"carid" + ") " + VALUES + "(" +
		" '" + u.CreateDate.Format(TEAMPLEATE_TIME) + "', " +
		" '" + string(u.Role) + "', " +
		" '" + u.Login + "', " +
		" '" + strconv.Itoa(u.Region.Id) + "', " +
		" '" + strconv.Itoa(carID) + "');"

	_, err := s.psql.SendQuery(userQueryIncerst)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (s *Storage) GetAllRegions() []model2.Region {
	query := SELECT + "*" + FROM + REGION_TABLE_NAME
	resultSearch := s.psql.GetRows(query)
	var result []model2.Region
	for resultSearch.Next() {
		var id int
		var name string
		err := resultSearch.Scan(&id, &name)
		if err != nil {
			log.Println("Error scan regions in func GetAllRegions()")
			return result
		}
		result = append(result, model2.Region{
			Id:         id,
			RegionName: name,
		})
	}
	return result
}
func (s *Storage) findCarId(concern model2.Concern, brand model2.Brand, model model2.Model, engine model2.Engine) int {
	//TODO подумать как сделать лучше

	query := SELECT + "*" + FROM + CAR_TABLE_NAME + WHERE + "model ='" + model.Model + "'" + AND + "concern ='" + concern.Concern + "' " + AND + "brand ='" + brand.Brand + "' " + AND +
		"engine ='" + engine.EngineName + "'"
	resultSearch := s.psql.GetRows(query)
	var id int
	iterator := 0
	for resultSearch.Next() {
		if iterator == 0 {
			iterator++
			var c string
			var b string
			var m string
			var e string
			var bp string
			var yearFrom string
			var yearTo string
			var class string
			err := resultSearch.Scan(&id, &c, &b, &m, &e, &bp, &yearFrom, &yearTo, &class)
			if err != nil {
				log.Println("Error find car ID in dataBase")
				return 0
			}
		} else {
			log.Println("More than one result found!")
			return 0
		}
	}

	return id
}
func (s *Storage) UpdateUser() bool {
	//TODO
	return false
}
