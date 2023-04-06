package bot

import (
	"time"
)

func CreateUserFromTemp(td TempData) *User {
	return &User{
		CreateDate: time.Now(),
		Role:       USER_ROLE,
		Login:      td.UserId,
		Region:     td.Region,
		UserCar: UserCar{
			CreateDate:  time.Time{},
			Concern:     td.CarData.Concern,
			Model:       td.CarData.CarModel,
			Engine:      td.CarData.CarEngine,
			BoltPattern: td.CarData.BoltPattern,
			Brand:       td.CarData.CarBrand,
		},
	}
}
