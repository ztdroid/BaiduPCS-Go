package pcsweb

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DB = connection()
)

func connection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
	  panic("failed to connect database")
	}
	// defer db.Close()
	return db
}

