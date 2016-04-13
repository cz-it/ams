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
	db, err := sql.Open("mysql", "thesha3_db_ams:225306ams@/mql.freehostia.com")
	defer db.Close()
	if err != nil {
		fmt.Printf("open with error %s \n", err.Error())
	}
	fmt.Println("Open success")
	db.Exec(`
		CREATE TABLE IF NOT EXISTS t_ams_use_id (
    		f_index int NOT NULL AUTO_INCREMENT,
    		f_user_id bigint unsigned NOT NULL,
    		f_create_time datetime NOT NULL,
    		PRIMARY KEY(f_user_id),
    		INDEX USING BTREE (f_user_id)
		) ENGINE = InnoDB; `)
}
