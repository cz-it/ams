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
			f_master_user_id bigint unsigned ,
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

func query(db *sql.DB, table string, col string, cond ...interface{}) (value interface{}, err error) {
	var sql string
	sql = fmt.Sprintf("SELECT %s FROM %s ", col, table)
	if nil == cond || len(cond) == 0 {
		err = db.QueryRow(sql).Scan(value)
		return
	}
	sql += " WHERE "
	c := ""
	ctag := ""
	for i, v := range cond {
		if i >= len(cond)/2 {
			break
		}
		c += ctag + v.(string) + "=?"
		ctag = " AND "
	}
	sql += c
	var v interface{}
	err = db.QueryRow(sql, cond[len(cond)/2:]...).Scan(&v)
	value = uint64(v.(int64))
	return
}

func update(db *sql.DB, table string, col string, value interface{}, cond ...interface{}) (err error) {
	var sql string
	sql = fmt.Sprintf("UPDATE %s  SET %s=? ", table, col)
	if cond == nil || len(cond) == 0 {
		stmt, err := db.Prepare(sql)
		if err != nil {
			utils.Logger.Error("prepare with error %s \n", err.Error())
			return err
		}
		_, err = stmt.Exec()
		return err
	}
	sql += " WHERE "
	c := ""
	ctag := ""
	for i, v := range cond {
		if i >= len(cond)/2 {
			break
		}
		c += ctag + v.(string) + "=?"
		ctag = " AND "
	}
	sql += c
	stmt, err := db.Prepare(sql)
	if err != nil {
		utils.Logger.Error("prepare with error %s \n", err.Error())
		return
	}
	execarg := []interface{}{value}
	execarg = append(execarg, cond[len(cond)/2:]...)
	_, err = stmt.Exec(execarg...)
	return
}
