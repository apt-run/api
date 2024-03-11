package database

import (
	"context"
	"fmt"
	"io"
	"log"
	"main/configs"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sudo-adduser-jordan/gcolor"
)

var connection *pgxpool.Pool

func ConnectToDatabase() {

	// "postgres://username:password@localhost:5432/database_name"
	DATABASE_URL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		configs.Env("DB_USER"),
		configs.Env("DB_PASSWORD"),
		// configs.Env("DB_HOST"),
		"localhost",
		configs.Env("DB_PORT"),
		configs.Env("DB_NAME"),
	)

	// Docker Compose
	// "postgres://username:password@database:5432/database_name"
	// DATABASE_URL := fmt.Sprintf(
	// 	"postgres://%s:%s@%s:%s/%s",
	// 	configs.Env("DB_USER"),
	// 	configs.Env("DB_PASSWORD"),
	// 	configs.Env("DB_HOST"),
	// 	configs.Env("DB_PORT"),
	// 	configs.Env("DB_NAME"),
	// )

	var err error
	connection, err = pgxpool.New(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Unable to connection to database: %v\n",
			err,
		)
		os.Exit(1)
	}
	fmt.Print("	-----> ")
	fmt.Println(gcolor.GreenText("Connected to database."))
}

func MigrateDatabase() {
	// CreateSourceTable()
	CreatePackageTable()
	// CreateMaintainerTable()
	// CreateMetricsTable()

	fmt.Print("	-----> ")
	fmt.Println(gcolor.GreenText("Database migrated."))
}

func UpdateDebianList() {
	name := "Debian"
	list := GET_PACKAGE_LIST()
	_, err := connection.Exec(context.Background(),
		UPSERT_SOURCES,
		name, list,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("	-----> ")
	fmt.Println(gcolor.YellowText("Debian packages list updated."))
}

func CreateSourceTable() {
	_, err := connection.Exec(context.Background(),
		CREATE_SOURCE_TABLE,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("	-----> ")
	fmt.Println(gcolor.YellowText("Source table created."))

	UpdateDebianList()
}

func CreatePackageTable() {
	_, err := connection.Exec(context.Background(),
		CREATE_PACKAGE_TABLE,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("	-----> ")
	fmt.Println(gcolor.YellowText("Package table created."))
}

func GET_PACKAGE_LIST() []byte {
	response, err := http.Get("https://sources.debian.org/api/list")
	if err != nil {
		log.Fatalln(err)
	}
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return data
}
