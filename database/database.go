package database

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"main/configs"
	"main/utils"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sudo-adduser-jordan/gcolor"
)

var connection *pgxpool.Pool

func ConnectToDatabase() {

	// "postgres://username:password@database:5432/database_name"
	DATABASE_URL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		configs.Env("DB_USER"),
		configs.Env("DB_PASSWORD"),
		configs.Env("DB_HOST"),
		configs.Env("DB_PORT"),
		configs.Env("DB_NAME"),
	)

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
	// CreatePackageTable()
	// CreateStatsTable()

	fmt.Print("	-----> ")
	fmt.Println(gcolor.GreenText("Database migrated."))
}

func CreateSourceTable() {
	_, err := connection.Exec(context.Background(),
		CREATE_SOURCE_TABLE,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("	-----> ")
	fmt.Println(gcolor.YellowText("Table sources created."))

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
	fmt.Println(gcolor.YellowText("Table package created."))
}

func CreateStatsTable() {
	DeleteStatsTable()
	_, err := connection.Exec(context.Background(),
		CREATE_STATS_TABLE,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("	-----> ")
	fmt.Println(gcolor.YellowText("Table stats created."))

	InsertDebianStats()
}

func UpdateDebianList() {
	name := "Debian"
	list := utils.GetList()

	_, err := connection.Exec(context.Background(),
		UPSERT_SOURCES,
		name, list,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("	-----> ")
	fmt.Println(gcolor.YellowText("Debian sources updated."))
}

// add progress and time to complete
func InsertDebianStats() {
	err := utils.DownloadFile("tmp/by_inst.txt", "https://popcon.debian.org/by_inst")
	if err != nil {
		fmt.Println(err.Error())
	}

	file, err := os.Open("tmp/by_inst.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)

	var rows []PackageStats
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		if strings.Contains(line, "---------------------------") {
			break
		}
		if line[0] != '#' &&
			len(line) > 0 {
			stats := PackageStats{}
			stats.Rank, err = strconv.Atoi(tokens[0])
			if err != nil {
				fmt.Println(err.Error())
			}
			stats.Name = tokens[1]
			stats.Inst, err = strconv.Atoi(tokens[2])
			if err != nil {
				fmt.Println(err.Error())
			}
			stats.Vote, err = strconv.Atoi(tokens[3])
			if err != nil {
				fmt.Println(err.Error())
			}
			stats.Old, err = strconv.Atoi(tokens[4])
			if err != nil {
				fmt.Println(err.Error())
			}
			stats.Recent, err = strconv.Atoi(tokens[5])
			if err != nil {
				fmt.Println(err.Error())
			}
			stats.NoFiles, err = strconv.Atoi(tokens[6])
			if err != nil {
				fmt.Println(err.Error())
			}
			str_builder := tokens[7]
			for _, token := range tokens[8:] {
				str_builder += " "
				str_builder += token
				if strings.Contains(token, ")") {
					break
				}
			}
			stats.Maintainer = str_builder[1 : len(str_builder)-1]
			rows = append(rows, stats)
		}
	}

	_, err = connection.CopyFrom(
		context.Background(),
		pgx.Identifier{"stats"},
		[]string{"rank", "name", "installs", "vote", "old", "recent", "nofiles", "maintainer"},
		pgx.CopyFromSlice(len(rows), func(i int) ([]any, error) {
			return []any{
				rows[i].Rank,
				rows[i].Name,
				rows[i].Inst,
				rows[i].Vote,
				rows[i].Old,
				rows[i].Recent,
				rows[i].NoFiles,
				rows[i].Maintainer,
			}, nil
		}),
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Print("	-----> ")
	fmt.Println(gcolor.YellowText("Debian stats created."))
}

func DeleteStatsTable() {
	_, err := connection.Exec(context.Background(),
		DROP_TABLE_STATS,
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	// process tmp/data.txt
	fmt.Print("	-----> ")
	fmt.Println(gcolor.RedText("Table stats dropped."))

}
