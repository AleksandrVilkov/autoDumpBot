package model

import (
	"time"
)

type User struct {
	Id         int
	CreateDate time.Time
	Role       Role
	Login      string
	Region     Region
	UserCar    UserCar
}

type UserCar struct {
	Id          int
	CreateDate  time.Time
	Concern     Concern
	Model       Model
	Engine      Engine
	BoltPattern BoltPattern
	Brand       Brand
}

type ButtonData struct {
	Token string `json:"token"`
}

type Concern struct {
	Concern string `json:"concern"`
}

type Brand struct {
	Brand string `json:"brand"`
}

type Model struct {
	Model string `json:"model"`
}

type Engine struct {
	EngineName string `json:"engineName"`
}

type BoltPattern struct {
	BoltPatternSize string `json:"boltPattern"`
}

type Region struct {
	Id         int    `json:"id"`
	RegionName string `json:"name"`
}
