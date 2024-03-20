package csv

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

type Config struct {
	FromLocalPath   string
	ToDatabaseTable string
	DbDsn           string
}

func Migrate() error {
	var cfg Config
	flag.StringVar(&cfg.DbDsn, "db-dsn", os.Getenv("CHESS_DB_DSN"), "the destination table in the database")
	flag.StringVar(&cfg.FromLocalPath, "path", "", "the path to the CSV file")
	flag.StringVar(&cfg.ToDatabaseTable, "dest", "", "the destination table in the database")
	flag.Parse()
	fmt.Println("Migrating from", cfg.FromLocalPath, "to", cfg.ToDatabaseTable, "in", cfg.DbDsn)

	file, err := os.Open(cfg.FromLocalPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
	return nil
}
