/**
* Author: CZ cz.theng@gmail.com
 */

package dbagent

import (
	"fmt"
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

func _TestInsertPlatformUID(t *testing.T) {
	err := InsertPlatformUID(1, "weixin", "abcdef", uint64(64))
	if err != nil {
		t.Error("InsertPlatformUID error with %s", err.Error())
		return
	}
}

func TestBindUID(t *testing.T) {
	err := BindUID(1, "weixin", "abcdef", uint64(1064))
	if err != nil {
		t.Error("BindUID error with %s", err.Error())
		return
	}
}

func TestQueryUID(t *testing.T) {
	UID, err := QueryUID(1, "weixin", "abcdef")
	if err != nil {
		t.Error("QueyBID error with %s", err.Error())
		return
	}
	fmt.Println("uid is ", UID)
}
