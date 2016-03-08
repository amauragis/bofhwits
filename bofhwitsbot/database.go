package bofhwitsbot

import (
	// "fmt"

	"database/sql"
	"errors"

	"github.com/amauragis/sanitize"

	// database driver tomfoolery
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func sqliteOpen(bot *BofhwitsBot) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", bot.Configs.Sqlite.File)
	return db, err
}

func mysqlOpen(bot *BofhwitsBot) (*sql.DB, error) {
	db, err := sql.Open("mysql", bot.Configs.Mysql.User+":"+bot.Configs.Mysql.Pass+"@tcp("+bot.Configs.Mysql.Host+":3306)/"+bot.Configs.Mysql.DB)
	return db, err
}

func (bot *BofhwitsBot) sqliteInit() error {

	if bot.Configs.Sqlite.File == "" {
		return errors.New("No valid sqlite file specified")
	}

	// TODO: Validate we have a good sqlite file?

	bot.dbOpen = sqliteOpen

	dbCon, err := bot.dbOpen(bot)
	if err != nil {
		bot.Log.Printf("err: %q\n", err)
		return err
	}

	defer dbCon.Close()

	sqlStatement := `
	CREATE TABLE IF NOT EXISTS bofhwits_posts (post_id INTEGER PRIMARY KEY, user TEXT, post TEXT, requestor TEXT, ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL);
	`
	_, err = dbCon.Exec(sqlStatement)
	if err != nil {
		bot.Log.Printf("%q: %s\n", err, sqlStatement)
		return err
	}

	bot.Log.Printf("sqlite inited successfully")

	return nil
}

func (bot *BofhwitsBot) mysqlInit() error {

	if bot.Configs.Mysql.User == "" ||
		bot.Configs.Mysql.Pass == "" ||
		bot.Configs.Mysql.Host == "" ||
		bot.Configs.Mysql.DB == "" {

		return errors.New("Invalid sqlite parameters!")
	}
	bot.dbOpen = mysqlOpen

	dbCon, err := bot.dbOpen(bot)
	if err != nil {
		bot.Log.Printf("DB Failure: %v\n", err)
	}
	defer dbCon.Close()

	sqlStatement := `
	CREATE TABLE IF NOT EXISTS bofhwits_butts (post_id INTEGER PRIMARY KEY AUTO_INCREMENT, user VARCHAR(50), post VARCHAR(1000), requestor VARCHAR(50), ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP);
	`

	_, err = dbCon.Exec(sqlStatement)
	if err != nil {
		bot.Log.Printf("%q: %s\n", err, sqlStatement)
		return err
	}

	return nil
}

func (bot *BofhwitsBot) postSQL(user string, msg string, requestor string) {
	msg = sanitize.HTML(msg)
	user = sanitize.HTML(user)
	requestor = sanitize.HTML(requestor)

	sqlcon, err := bot.dbOpen(bot)
	if err != nil {
		bot.con.Privmsg(bot.Configs.Channel, "Could not connect db for some reason...")
		bot.Log.Printf("DB Failure: %v\n", err)
	}
	defer sqlcon.Close()

	bot.Log.Printf("DB: INSERT INTO bofhwits_posts (user, post, requestor) VALUES (%s, %s, %s)", user, msg, requestor)

	stmt, err := sqlcon.Prepare("INSERT INTO bofhwits_posts (user, post, requestor) VALUES (?, ?, ?)")
	stmt.Exec(user, msg, requestor)

	if err != nil {
		bot.con.Privmsg(bot.Configs.Channel, "Could not db for some reason...")
		bot.Log.Printf("DB Failure: %v\n", err)
	}
}
