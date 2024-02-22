package main

import (
	"database/sql"
	"fmt"
	_ "github.com/glebarez/go-sqlite"
)

var Version = 2

func GetConnection() *sql.DB {
	db, err := sql.Open("sqlite", "DataBase.db")
	if err != nil {
		panic(err)
	}
	return db
}

func InitDatabase() {
	fmt.Println("Start Init Database SQlite")
	db := GetConnection()
	defer db.Close()

	rows, err := db.Query(
		"SELECT cur FROM version",
	)
	if err != nil {
		CreateTables(db)
	} else {
		version := -1
		rows.Next()
		_ = rows.Scan(&version)
		if version != Version {
			CreateTables(db)
		}
	}
	fmt.Println("Done Init Database SQlite")
}

func CreateTables(db *sql.DB) {
	db.Exec("DROP TABLE IF EXISTS version; CREATE TABLE version(cur INTEGER PRIMARY KEY);")
	db.Exec("INSERT INTO version (cur) VALUES ($1);",
		Version,
	)

	db.Exec("DROP TABLE IF EXISTS templates; CREATE TABLE templates(name TEXT, code TEXT,template TEXT, PRIMARY KEY(name, code));")
	// Init Data
	db.Exec("INSERT INTO templates (name,code,template) VALUES" +
		" ('RunTemplate','ru','[Шаблон меню]\nШаблон текста\n%s')," +
		" ('RunTemplate','enu','[Template menu]\nTemplate text\n%s')," +
		" ('LoginMenu','ru','[Логин меню]\nПриветствуем !\n\nДля входа введите ключ от акаунта 🔑')," +
		" ('LoginMenu','enu','[[Login menu]\nWelcome!\n\nTo log in, enter your account key 🔑')," +
		" ('back','ru','⬅️ Назад')," +
		" ('back','enu','⬅️ Back'),",
	)
}

func GetTemplate(name string, code string) string {
	var template = "Template not specified !"
	db := GetConnection()
	defer db.Close()

	rows, _ := db.Query(
		"SELECT template from templates WHERE (name = $1 AND code = $2) OR (name = $3 AND code = 'enu')",
		name,
		code,
		name,
	)

	rows.Next()
	_ = rows.Scan(&template)

	return template
}
