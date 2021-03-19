package mysql

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql" // Needed for connect mysql
	errors "innogrid.com/hcloud-classic/hcc_errors"

	"hcc/piano/lib/config"
	"hcc/piano/lib/logger"
)

// Db : Pointer of mysql connection
var (
	ctx context.Context
	Db  *sql.DB
)

// Type alias
type Rows = sql.Rows
type Result = sql.Result

// Prepare : Connect to mysql and prepare pointer of mysql connection
func Prepare() (func(), *errors.HccError) {
	var err error
	Db, err = sql.Open("mysql",
		config.Mysql.ID+":"+config.Mysql.Password+"@tcp("+
			config.Mysql.Address+":"+strconv.Itoa(int(config.Mysql.Port))+")/"+
			config.Mysql.Database+"?parseTime=true")
	if err != nil {
		return nil, errors.NewHccError(errors.ViolinNoVNCInternalInitFail, "mysql open")
	}

	timeTicker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	cancel := func() { done <- true; timeTicker.Stop() }
	go func() {
		for true {
			select {
			case <-done:
				return
			case <-timeTicker.C:
				err = Db.Ping()
				if err != nil {
					logger.Logger.Println(
						errors.NewHccError(errors.ViolinNoVNCInternalConnectionFail,
							"mysql connection lost, retry...").Error())
				}
				break
			}
		}
		if err != nil {

		}
	}()

	return cancel, nil
}