package configs

import (
	"prakerja11/models/cerpen"
	"prakerja11/models/user"
)

func initMigrate() {
	DB.AutoMigrate(&user.User{}, &cerpen.Cerpen{})
}
