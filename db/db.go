package db

import (
	"fmt"
	"log"
	"tanaman/config"

	_ "github.com/go-sql-driver/mysql"

	"xorm.io/xorm"
)

func ConnectDB() *xorm.Engine {
	config := config.LoadConfig(".")

	engine, err := xorm.NewEngine("mysql", config.DbDsn)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	engine.ShowSQL()
	err = engine.Ping()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println("connect mysql success")
	return engine
}
