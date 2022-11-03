package db

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type DbConfig struct {
	Service  string
	Host     string
	Port     string
	Username string
	Password string
}

var Dbc DbConfig = DbConfig{}
var dbpass string

func LoadDbConfig() {

	ConfigFile := "OraDBCOnfig.json"
	//...................................
	//Reading into struct type from a JSON file
	//...................................
	content, err := os.ReadFile(ConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &Dbc)
	if err != nil {
		log.Fatal(err)
	}

	//TestAES()
	if strings.HasPrefix(Dbc.Password, "##NLP##") {
		dbpass = Dbc.Password[7:]
		Dbc.Password = Encriptar(dbpass)

		//...................................
		//Writing struct type to a JSON file
		//...................................
		content, err = json.MarshalIndent(Dbc, " ", " ")
		if err != nil {
			fmt.Println(err)
		}
		err = os.WriteFile(ConfigFile, content, 0644)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		//dbpass = Encriptar(Dbc.Password)
		dbpass = Desencriptar(Dbc.Password)
	}

}

// ## Loading the CSV data
func LoadCSV(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Cannot open '%s': %s\n", path, err.Error())
	}
	defer f.Close()
	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Cannot read CSV data:", err.Error())
	}
	return rows
}
