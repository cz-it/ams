/**
* Author: CZ cz.theng@gmail.com
 */

package dbagent

import (
	"database/sql"
	"fmt"
	"github.com/cz-it/ams"
	"github.com/cz-it/ams/utils"
	//for mysql
	_ "github.com/go-sql-driver/mysql"
)

func connect() (db *sql.DB, err error) {
	var dbInfo string
	dbInfo = fmt.Sprintf("%s:%s@tcp(%s)/%s", ams.Config.DBUser, ams.Config.DBPasswd, ams.Config.DBAddr, ams.Config.DBName)
	db, err = sql.Open("mysql", dbInfo)
	if err != nil {
		utils.Logger.Error("open db  with error %s \n", err.Error())
		return
	}
	utils.Logger.Info("Open DB(%s) success", ams.Config.DBName)
	err = db.Ping()
	if err != nil {
		utils.Logger.Error("ping with error %s \n", err.Error())
		return
	}
	return
}

func createTables(db *sql.DB) (err error) {
	// t_user_id
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS t_ams_use_id (
    		f_user_id bigint unsigned NOT NULL,
    		f_create_time datetime NOT NULL,
    		PRIMARY KEY(f_user_id),
    		INDEX USING BTREE (f_user_id)
		) ENGINE = InnoDB; `)

	if err != nil {
		utils.Logger.Error("exec with error %s \n", err.Error())
		return
	}

	// t_open_platform
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS t_ams_open_id (
			f_platform tinyint unsigned NOT NULL,
			f_app_id char(128) ,
			f_open_id char(128) NOT NULL,
			f_user_id bigint unsigned NOT NULL,
			f_master_user_id bigint unsigned NOT NULL,
			PRIMARY KEY (f_platform, f_app_id, f_open_id),
			INDEX USING BTREE (f_platform, f_app_id, f_open_id)
		) ENGINE = InnoDB; `)

	if err != nil {
		utils.Logger.Error("exec with error %s \n", err.Error())
		return
	}
	return
}

func insert(db *sql.DB, table string, colval ...interface{}) (err error) {
	var sql string
	col := ""
	coltag := ""
	valtag := ""
	val := ""
	for i, v := range colval {
		if i < len(colval)/2 {
			col += coltag + v.(string)
			coltag = ", "
		} else {
			val += valtag + "?"
			valtag = ", "
		}
	}
	sql = fmt.Sprintf("INSERT %s (%s) VALUES (%s)", table, col, val)
	stmt, err := db.Prepare(sql)
	if err != nil {
		utils.Logger.Error("prepare with error %s \n", err.Error())
		return
	}
	_, err = stmt.Exec(colval[len(colval)/2:]...)
	if err != nil {
		utils.Logger.Error("exec with error %s \n", err.Error())
		return
	}
	return
}

func query(db *sql.DB, table string, col string, colval ...interface{}) (value interface{}, err error) {

	return
}

/*
	stmt, err := db.Prepare(`INSERT t_ams_use_id ( f_user_id, f_create_time) values (?,?)`)
	if err != nil {
		fmt.Printf("prepare with error %s \n", err.Error())
		return
	}
		res, err := stmt.Exec(22, "2014-04-03 12:23:23")
		if err != nil {
			fmt.Printf("exec with error %s \n", err.Error())
			return
		}
		id, err := res.LastInsertId()
		if err != nil {
			fmt.Printf("Last inser id with error %s \n", err.Error())
			return
		}
	fmt.Println(sql.Drivers())
	fmt.Println("stmt:", stmt)
}
*/
