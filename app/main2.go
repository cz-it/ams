/**
* Author: CZ cz.theng@gmail.com
 */

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func _main() {
	db, err := sql.Open("mysql", "thesha3_db_ams:225306ams@tcp(mysql.freehostia.com:3306)/thesha3_db_ams")
	defer db.Close()
	if err != nil {
		fmt.Printf("open with error %s \n", err.Error())
		return
	}
	fmt.Println("Open success")
	err = db.Ping()
	if err != nil {
		fmt.Printf("ping with error %s \n", err.Error())
		return
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS t_ams_use_id (
    		f_user_id bigint unsigned NOT NULL,
    		f_create_time datetime NOT NULL,
    		PRIMARY KEY(f_user_id),
    		INDEX USING BTREE (f_user_id)
		) ENGINE = InnoDB; `)

	if err != nil {
		fmt.Printf("exec with error %s \n", err.Error())
		return
	}

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
	fmt.Printf(" id is %d", id)
}
