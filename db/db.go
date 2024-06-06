package db

import (
	"fmt"
	"log"
	"tanaman/config"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

func ConnectDB() *xorm.Engine {
	config := config.LoadConfig(".")

	engine, err := xorm.NewEngine(config.Db, config.DbDsn)
	if err != nil {
		fmt.Println("connect postgres error", err)
		log.Fatal(err)
		return nil
	}
	engine.ShowSQL()
	err = engine.Ping()
	if err != nil {
		fmt.Println("ping postgres error", err)
		log.Fatal(err)
		return nil
	}
	fmt.Println("connect postgres success")
	return engine
}
