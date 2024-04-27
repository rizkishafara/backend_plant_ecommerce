package config

type DatabaseLocalType struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         string
}

type DatabaseDevType struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         string
}

type DatabaseProdType struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         string
}

const MaxRowsTable = 30

var MaxRowList = []int{10, 20, 30, 50, 100}

// For Dynamic DB
type ConfigDynamicType struct {
	DatabaseLocal DatabaseLocalType
	DatabaseDev   DatabaseDevType
	DatabaseProd  DatabaseProdType
}

var ConfigDynamic ConfigDynamicType

func LoadConfigDynamic() {
	ConfigDynamic = ConfigDynamicType{
		DatabaseLocal: DatabaseLocalType{
			Host:         "localhost",
			Username:     "root",
			Password:     "",
			DatabaseName: "grelegant",
			Port:         "8080",
		},
		DatabaseDev: DatabaseDevType{
			Host:         "103.171.85.168", //deploy dev host
			Username:     "admin",
			Password:     "root",
			DatabaseName: "devgre",
			Port:         "8080",
		},
		DatabaseProd: DatabaseProdType{
			Host:         "103.171.85.168", //deploy prod host
			Username:     "admin",
			Password:     "root",
			DatabaseName: "devgre",
			Port:         "8080",
		},
	}
}
