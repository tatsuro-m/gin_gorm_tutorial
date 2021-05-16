package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	C   *gorm.DB
	err error
)

func ConnectDB() {
	C, err = gorm.Open(postgres.Open(os.Getenv("MAIN_DSN")), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func Close() {
	db, err := C.DB()
	if err != nil {
		panic(err)
	}

	err = db.Close()
	if err != nil {
		panic(err)
	}
}
