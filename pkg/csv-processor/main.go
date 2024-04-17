package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type FindClubResp struct {
	Data struct {
		ClubID    int  `json:"club_id"`
		IsActive  bool `json:"is_active"`
		CreatedAt struct {
			Time  time.Time `json:"Time"`
			Valid bool      `json:"Valid"`
		} `json:"created_at"`
		DeletedAt struct {
			Time  time.Time `json:"Time"`
			Valid bool      `json:"Valid"`
		} `json:"deleted_at"`
		UpdatedAt struct {
			Time  time.Time `json:"Time"`
			Valid bool      `json:"Valid"`
		} `json:"updated_at"`
		Code         string `json:"code"`
		Club         string `json:"club"`
		Address      string `json:"address"`
		Observations string `json:"observations"`
		City         string `json:"city"`
		Country      string `json:"country"`
	} `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type PlayerDump struct {
	Params Params `json:"params"`
	Source Source `json:"source"`
}

type Params struct {
	Columns        []Column `json:"columns"`
	Width          string   `json:"width"`
	AutoHeight     string   `json:"autoheight"`
	RowsHeight     string   `json:"rowsheight"`
	ColumnsResize  string   `json:"columnsresize"`
	ColumnsReorder string   `json:"columnsreorder"`
	SelectionMode  string   `json:"selectionmode"`
	Sortable       string   `json:"sortable"`
	Editable       string   `json:"editable"`
}

type Source struct {
	ID         string      `json:"id"`
	UpdateRow  string      `json:"updaterow"`
	DataType   string      `json:"datatype"`
	DataFields []DataField `json:"datafields"`
	LocalData  []LocalData `json:"localdata"`
}

type Column struct {
	DefaultValue string `json:"defaultValue"`
	DataField    string `json:"dataField"`
	Width        string `json:"width"`
	Editable     string `json:"editable"`
	Text         string `json:"text"`
}

type DataField struct {
	Id        string `json:"id"`
	UpdateRow string `json:"updaterow"`
}

type LocalData struct {
	IdPlayer    string `json:"idJugador"`
	IdClub      string `json:"idClub"`
	FullName    string `json:"nomLlarg"`
	Sex         string `json:"sexe"`
	EloStd      string `json:"elo"`
	EloRapid    string `json:"eloR"`
	Title       string `json:"titol"`
	AgeCategory string `json:"descripcio"`
	ClubName    string `json:"nomClub"`
}

type Config struct {
	FromLocalPath   string
	ToDatabaseTable string
	DbDsn           string
}

type CSVClub struct {
	Code         string
	Club         string
	Address      string
	Observations string
	City         string
}

type CSVUser struct {
	Code             int
	Name             string
	Sex              string
	EloStd           string
	EloRapid         string
	Title            string
	ChessAgeCategory string
	ClubID           int
	ClubUserCode     string
	Country          string
}

type UserRequestBody struct {
	Code                int    `json:"code"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	Sex                 string `json:"sex"`
	Title               string `json:"title"`
	ChessAgeCategory    string `json:"chess_age_category"`
	ClubID              int    `json:"club_id"`
	ClubUserCode        string `json:"club_user_code"`
	Country             string `json:"country"`
	EloFideStandard     int    `json:"elo_fide_standard"`
	EloFideRapid        int    `json:"elo_fide_rapid"`
	EloNationalStandard int    `json:"elo_national_standard"`
	EloNationalRapid    int    `json:"elo_national_rapid"`
	EloRegionalStandard int    `json:"elo_regional_standard"`
	EloRegionalRapid    int    `json:"elo_regional_rapid"`
	Email               string `json:"email"`
}

type ClubRequestBody struct {
	Code         string `json:"code"`
	Club         string `json:"club"`
	Address      string `json:"address"`
	Observations string `json:"observations"`
	City         string `json:"city"`
}

func main() {
	var cfg Config
	flag.StringVar(&cfg.DbDsn, "db-dsn", os.Getenv("CHESS_DB_DSN"), "the destination table in the database")
	flag.StringVar(&cfg.FromLocalPath, "path", "", "the path to the CSV file")
	flag.StringVar(&cfg.ToDatabaseTable, "dest", "", "the destination table in the database")
	flag.Parse()
	fmt.Println("Migrating from", cfg.FromLocalPath, "to", cfg.ToDatabaseTable, "in", cfg.DbDsn)
	switch cfg.ToDatabaseTable {
	// for club creation
	case "clubs":
		if err := Migrate(cfg.FromLocalPath); err != nil {
			fmt.Println("Error migrating:", err)
			os.Exit(1)
		}

	// for user creation
	case "users":
		if err := migrateJSONPlayers(cfg.FromLocalPath); err != nil {
			fmt.Println("Error migrating:", err)
			os.Exit(1)
		}
	}
}

func migrateJSONPlayers(path string) error {
	var pd PlayerDump
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&pd)
	if err != nil {
		log.Fatal(err)
	}
	for i, ld := range pd.Source.LocalData {
		name := strings.Split(ld.FullName, ",")
		email := fmt.Sprintf("%d@gmail.com", i)
		fname := name[1]
		lname := name[0]
		playerCode, _ := strconv.Atoi(ld.IdPlayer)
		club := CSVClub{Club: ld.ClubName}
		ClubCode, _ := fetchClubID(club)
		elostd, err := strconv.Atoi(ld.EloStd)
		if err != nil {
			log.Fatal(err)
		}
		eloRapid, err := strconv.Atoi(ld.EloRapid)
		if err != nil {
			log.Fatal(err)
		}

		user := UserRequestBody{
			Code:                playerCode,
			FirstName:           fname,
			LastName:            lname,
			Sex:                 ld.Sex,
			EloRegionalStandard: elostd,
			EloRegionalRapid:    eloRapid,
			Title:               ld.Title,
			ChessAgeCategory:    catTranslate(ld.AgeCategory),
			ClubUserCode:        fmt.Sprintf("%d-%d", ClubCode, playerCode),
			Country:             "Spain",
			ClubID:              ClubCode,
			Email:               email,
		}
		if err := postUser(user); err != nil {
			fmt.Println("Error: ", user)
			log.Fatal(err)
		}
	}
	fmt.Println("Migration completed successfully")
	return nil
}

func migrateUsers(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	csvUsers := processUsers(data)

	for _, csvUser := range csvUsers {
		splitName := strings.Split(csvUser.Name, ",")
		lName := splitName[0]
		fName := strings.TrimSpace(splitName[1])
		word := strings.Split(csvUser.ChessAgeCategory, "-")
		translatedCategory := catTranslate(word[0])
		if len(word) > 1 {
			translatedCategory = translatedCategory + "-" + word[1]
		}
		cuid := fmt.Sprintf("%d-%d", csvUser.ClubID, csvUser.Code)
		user := UserRequestBody{
			Code:      csvUser.Code,
			FirstName: fName,
			LastName:  lName,
			Sex:       csvUser.Sex,
			// EloStd:           csvUser.EloStd,
			// EloRapid:         csvUser.EloRapid,
			Title:            csvUser.Title,
			ChessAgeCategory: translatedCategory,
			ClubID:           csvUser.ClubID,
			ClubUserCode:     cuid,
			Country:          "Spain",
		}

		if err := postUser(user); err != nil {
			fmt.Println("Error: ", err)
			log.Fatal(err)
		}
	}

	return nil
}

func processUsers(data [][]string) []CSVUser {
	var csvUser []CSVUser
	for i, line := range data {
		if i > 0 { // omit header line
			var rec CSVUser
			fmt.Println("Line: ", line)
			for j, field := range line {
				if j == 0 {
					val, err := strconv.Atoi(field)
					if err != nil {
						log.Fatal(err)
					}
					rec.Code = val
				} else if j == 1 {
					rec.Name = field
				} else if j == 2 {
					rec.Sex = field
				} else if j == 3 {
					rec.EloStd = field
				} else if j == 4 {
					rec.EloRapid = field
				} else if j == 7 {
					rec.Title = field
				} else if j == 8 {
					rec.ChessAgeCategory = field
				} else if j == 9 {
					clubCode, _ := fetchClubID(CSVClub{Club: field})
					fmt.Println("ClubCode: ", clubCode)
					rec.ClubID = clubCode
				}
			}
			csvUser = append(csvUser, rec)
		}
	}
	return csvUser
}

func Migrate(path string) error {
	file, err := os.Open(path)
	if err != nil {
		path, _ := os.Getwd()
		fmt.Println("wd: ", path)
		log.Fatal(err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	csvclubs := upload(data)

	for _, csvclub := range csvclubs {
		club := ClubRequestBody{
			Code:         csvclub.Code,
			Club:         csvclub.Club,
			Address:      csvclub.Address,
			Observations: csvclub.Observations,
			City:         csvclub.City,
		}
		if err := postClub(club); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

func postClub(club ClubRequestBody) error {
	c, err := json.Marshal(club)
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(c)
	resp, err := http.Post("http://127.0.0.1:8080/v1/club/create", "application/json", buf)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	fmt.Println("response status:", resp.Status)
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println("response :", bodyString)

	return nil
}

func postUser(user UserRequestBody) error {
	u, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(u)
	fmt.Println("USER BEFORE SENDING: ", buf)
	resp, err := http.Post("http://127.0.0.1:8080/v1/user/create", "application/json", buf)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	fmt.Println("response status FROM CREATING USER:", resp.Status)
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println("response :", bodyString)

	return nil
}

func upload(data [][]string) []CSVClub {
	var csvclub []CSVClub
	for i, line := range data {
		if i > 0 { // omit header line
			var rec CSVClub
			for j, field := range line {
				if j == 0 {
					rec.Code = field
				} else if j == 1 {
					rec.Club = field
				} else if j == 2 {
					rec.Address = field
				} else if j == 3 {
					rec.Observations = field
				} else if j == 4 {
					rec.City = field
				}
			}
			csvclub = append(csvclub, rec)
		}
	}
	return csvclub
}
func catTranslate(cat string) string {
	splitWrd := strings.Split(cat, "-")
	if len(splitWrd) > 1 {
		return catTranslate(splitWrd[0]) + "-" + splitWrd[1]
	}
	switch cat {
	case "Sènior":
		return "Senior"
	case "Veterà":
		return "Veteran"
	case "Sub":
		return "Sub"
	default:
		return cat
	}
}

func fetchClubID(club CSVClub) (int, error) {
	var fcr FindClubResp
	resp, err := http.Get("http://localhost:8080/v1/club/by-name/" + club.Club)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bodyBytes, &fcr)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("ClubID: ", fcr.Data)
	bodyString := string(bodyBytes)
	// fmt.Println("club :", club.Club)
	fmt.Println("response FROM GETTING CLUB DATA :", bodyString)
	return fcr.Data.ClubID, nil
}
