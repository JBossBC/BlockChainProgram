package dao

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"testing"
)

func TestInitDB(t *testing.T) {
	TestLoadXMlFile(t)
	_, err := sql.Open("mysql", getMysqlDSN())
	if err != nil {
		log.Panic(fmt.Errorf("database connection error:  %s", err.Error()).Error())
		return
	}
}

func TestLoadXMlFile(t *testing.T) {
	file, err := os.Open("./dbConfig.xml")
	if err != nil {
		panic(err.Error())
	}
	err = xml.NewEncoder(file).Encode(&DBConfig)
	if err != nil {
		panic(err.Error())
	}

}
