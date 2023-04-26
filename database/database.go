package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/minacio00/adopet/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbCredentials struct {
	client   string
	user     string
	password string
	port     string
	host     string
	database string
	ssl      string
}

var Db *gorm.DB

func (db *dbCredentials) formatStr() string {
	return "user=" + db.user + " host=" + db.host + " port=" + db.port + " password=" + db.password + " dbname=" + db.database + " sslmode=" + db.ssl
}

func Connectdb() {
	sql.Drivers()
	creds := dbCredentials{
		client: "postgresql", user: "postgres",
		password: "postgresql", port: "5432",
		host:/*"db"*/ "192.168.1.4",
		database: "go-adopet",
		ssl:      "disable",
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: creds.formatStr(),
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
	if err != nil {
		println(err)
		panic(err)
	}

	// println("Connection Opened to database")
	db.Logger = logger.Default.LogMode(logger.Info)
	db.AutoMigrate(&models.Abrigo{}, &models.Tutor{}, &models.Pet{})

	Db = db

}
