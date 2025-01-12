package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // Needed for connect mysql
	"hcc/piano/lib/config"
	"hcc/piano/lib/logger"
	"strconv"
)

// Db : Pointer of mysql connection
var Db *sql.DB

// Init : Initialize mysql connection
func Init() error {
	var err error
	Db, err = sql.Open("mysql",
		config.Mysql.ID+":"+config.Mysql.Password+"@tcp("+
			config.Mysql.Address+":"+strconv.Itoa(int(config.Mysql.Port))+")/"+
			config.Mysql.Database+"?parseTime=true")
	if err != nil {
		logger.Logger.Println(err)
		return err
	}

	err = Db.Ping()
	if err != nil {
		logger.Logger.Println(err)
		return err
	}

	logger.Logger.Println("Connected to MySQL database")

	return nil
}

// End : Close mysql connection
func End() {
	if Db != nil {
		_ = Db.Close()
	}
}
