package model

import "encoding/json"

type CallBack struct {
	Token      string
	UserId     string             `json:"userId"`
	Subsection CallbackSubsection `json:"subsection"`
	Action     Action             `json:"action"`
	UserData   struct {
		RegionName string `json:"name"`
		RegionId   int    `json:"id"`
	}
	CarData struct {
		Concern         string `json:"concern"`
		Brand           string `json:"brand"`
		Model           string `json:"model"`
		EngineName      string `json:"engineName"`
		BoltPatternSize string `json:"boltPattern"`
	} `json:"data"`
}

func (c *CallBack) ToString() string {
	str, _ := json.Marshal(c)
	return string(str)
}

type ActionCallBack struct {
	Action Action `json:"action"`
}
type SubsectionCallBack struct {
	Subsection CallbackSubsection `json:"subsection"`
	Data       string             `json:"data"`
}
