package bot

type TempUserData struct {
	User struct {
		Id string
	}
	Action struct {
		MainAction  string
		LastCommand string
	}
	CarData struct {
		Concern     Concern
		CarBrand    Brand
		CarModel    Model
		CarEngine   Engine
		BoltPattern string
	}
	SaleData struct {
	}
	SubscriptionData struct {
	}
}

type CallBack struct {
	Type CallbackType `json:"type"`
	Data string       `json:"data"`
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
}
