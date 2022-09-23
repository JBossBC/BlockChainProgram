package dao

import (
	"database/sql"
	"encoding/xml"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"strings"
	"time"
)

// var MongoClient *mongo.Client
type MysqlConfig struct {
	Username string `xml:"username"`
	Password string `xml:"password"`
	Host     string `xml:"host"`
	Port     int    `xml:"port"`
	DataBase string `xml:"database"`
}
type RedisConfig struct {
	Host     string `xml:"host"`
	Password string `xml:"password"`
	Port     int    `xml:"port"`
	Database string `xml:"database"`
}

type Config struct {
	Mysql *MysqlConfig `xml:"mysql"`
	Redis *RedisConfig `xml:"redis"`
}

var DBConfig *Config
var MysqlClient *sql.DB

func InitDB() {
	log.Println("Starting Init DB.... ")
	InitConfigFile()
	log.Println("Starting init mysql database...")
	db, err := sql.Open("mysql", getMysqlDSN())
	if err != nil {
		panic(fmt.Sprintf("init mysql error:%s", err.Error()))
	}
	db.SetConnMaxLifetime(10 * time.Second)
	db.SetMaxOpenConns(1000)
	db.SetMaxIdleConns(500)
	db.SetConnMaxIdleTime(5 * time.Second)
	log.Println("init mysql database success")
	MysqlClient = db
}

func getMysqlDSN() string {
	if DBConfig == nil {
		panic("DBconfig is nil pointer")
	}
	var MysqlDsn = strings.Builder{}
	MysqlDsn.WriteString(DBConfig.Mysql.Username)
	MysqlDsn.WriteString(":")
	MysqlDsn.WriteString(DBConfig.Mysql.Password)
	MysqlDsn.WriteString("@tcp(")
	MysqlDsn.WriteString(DBConfig.Mysql.Host)
	MysqlDsn.WriteString(":")
	MysqlDsn.WriteString(fmt.Sprintf("%d", DBConfig.Mysql.Port))
	MysqlDsn.WriteString(")/")
	MysqlDsn.WriteString(DBConfig.Mysql.DataBase)
	return MysqlDsn.String()
}

//func getRedisDSN() string {
//
//}

func InitConfigFile() {
	log.Println("Starting init database config file... ")
	file, err := os.OpenFile("./dao/dbConfig.xml", os.O_APPEND, 0644)
	defer file.Close()
	if err != nil {
		panic(fmt.Sprintf("database config init error :%s", err.Error()))
	}
	DBConfig = &Config{}
	err = xml.NewDecoder(file).Decode(DBConfig)
	if err != nil {
		panic(fmt.Sprintf("xml analy error: %s", err.Error()))
	}
	log.Println("init database config success")
}
