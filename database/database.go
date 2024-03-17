package database

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"main/configs"
	"net/http"
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
	CreateStatsTable()

	fmt.Print("	-----> ")
	fmt.Println(gcolor.GreenText("Database migrated."))
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

func InsertDebianStats() {
	url := "https://popcon.debian.org/by_inst"
	filepath := "tmp/by_inst.txt"

	if err := os.Mkdir("temp", 0777); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Print("	-----> ")
	fmt.Println(gcolor.YellowText("Directory created: tmp"))

	fmt.Print("	-----> ")
	fmt.Println(gcolor.YellowText("Downloading " + url + "..."))
	err := DownloadFile(filepath, url)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Print("	-----> ")
	fmt.Println(gcolor.GreenText("Finished downloading file :" + filepath))

	file, err := os.Open("tmp/by_inst")
	if err != nil {
		log.Fatalf(err.Error())
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
		log.Fatalf(err.Error())
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

	fmt.Print("	-----> ")
	fmt.Println(gcolor.RedText("Table stats dropped."))

}

func DownloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(gcolor.GreenText("File downloaded: " + filepath))
	return nil
}
