/**
* Author: CZ cz.theng@gmail.com
 */

package dbagent

import (
	"github.com/cz-it/ams"
	"testing"
	"time"
)

func initConfig() {
	ams.Config.DBAddr = "mysql.freehostia.com:3306"
	ams.Config.DBName = "thesha3_db_ams"
	ams.Config.DBUser = "thesha3_db_ams"
	ams.Config.DBPasswd = "225306ams"
}

func TestInit(t *testing.T) {
	initConfig()
	err := Init()
	if err != nil {
		t.Error("Init error with %s", err.Error())
		return
	}

	err = InsertUID(uint64(time.Now().Unix()))
	if err != nil {
		t.Error("Insert error with %s", err.Error())
	}
}
