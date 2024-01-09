package main

import (
	"database/sql"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"xsis-code-test/routes"
)

func main() {
	var (
		db  *gorm.DB
		err error
	)
	if os.Getenv("POSTGRES_HOST") != "localhost" {
		dsn := fmt.Sprintf("postgresql://%s:%s@%s/%s",
			os.Getenv("POSTGRES_USR"),
			os.Getenv("POSTGRES_PWD"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_DB"),
		)
		sqlDB, err := sql.Open("pgx", dsn)
		if err != nil {
			log.Panic(err)
		}
		db, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{})
	} else {
		uri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_USR"),
			os.Getenv("POSTGRES_PWD"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_PRT"),
			os.Getenv("POSTGRES_SSL_MODE"),
			os.Getenv("POSTGRES_TIMEZONE"))
		db, err = gorm.Open(postgres.Open(uri), &gorm.Config{})
	}
	if err != nil {
		log.Panic("Cannot Connect to DB")
	}
	//db.AutoMigrate(model.Movie{})
	srv := routes.AppRoutes(db)
	log.Println("Listening Application on Port ", os.Getenv("APP_PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), srv); err != nil {
		log.Panic("App cannot start")
	}
}
