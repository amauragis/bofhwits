package bofhwitsbot

import (
	// "fmt"

	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func sqliteOpen(bot *BofhwitsBot) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", bot.Configs.Sqlite.File)
	return db, err
}

func mysqlOpen(bot *BofhwitsBot) (*sql.DB, error) {

	return nil, nil
}

func (bot *BofhwitsBot) sqliteInit() error {

	if bot.Configs.Sqlite.File == "" {
		return errors.New("No valid sqlite file specified")
	}

	// TODO: Validate we have a good sqlite file?

	bot.dbOpen = sqliteOpen
	dbCon, oerr := sqliteOpen(bot)

	sqlStatement := `
	CREATE TABLE IF NOT EXISTS bofhwits_posts (post_id INTEGER PRIMARY KEY, user TEXT, post TEXT, requestor TEXT)
	`

	return nil
}

// TODO: split into sqlite/mysql/none itit functions
func (bot *BofhwitsBot) mysqlInit() error {
	// TODO: make this actually work!
	sqlcon, oerr := sql.Open("mysql", bot.Configs.Mysql.User+":"+bot.Configs.Mysql.Pass+"@tcp("+bot.Configs.Mysql.Host+":3306)/"+bot.Configs.Mysql.DB)
	if oerr != nil {
		bot.Log.Printf("DB Failure: %v\n", oerr)
	}
	defer sqlcon.Close()

	_, execErr :=
		sqlcon.Exec("CREATE TABLE IF NOT EXISTS bofhwits_posts ( post_id int PRIMARY KEY AUTO_INCREMENT, user VARCHAR(50), post VARCHAR(1000), requestor VARCHAR(50), ts TIMESTAMP);")

	if execErr != nil {
		bot.Log.Printf("Failed to init database: %v\n", execErr)
	}

	return nil
}
