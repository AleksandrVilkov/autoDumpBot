package model

type TempData struct {
	UserId           string
	Action           Action
	CarData          TempCarData
	SaleData         TempSaleData
	SearchData       TempSearchData
	SubscriptionData TempSubscriptionData
	Region           Region
}

type TempCarData struct {
	Concern     Concern
	CarBrand    Brand
	CarModel    Model
	CarEngine   Engine
	BoltPattern BoltPattern
}

type TempSearchData struct {
	Concern     Concern
	CarBrand    Brand
	CarModel    Model
	CarEngine   Engine
	BoltPattern BoltPattern
}

type TempSaleData struct {
	announcementText string
}

type TempSubscriptionData struct {
}
