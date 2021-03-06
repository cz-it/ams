/**
* Author: CZ cz.theng@gmail.com
 */

package dbagent

import (
	"database/sql"
	"github.com/cz-it/ams/utils"
	"time"
)

//DBAgent is a Database proxy
type DBAgent struct {
	db *sql.DB
}

//agent is a siglone DBAgent
var agent DBAgent

//Init init a default DBAgent agnet object use config
func Init() error {
	return agent.init()
}

//InsertUID insert a UID to db
func InsertUID(UID uint64) error {
	return agent.insertUID(UID)
}

//QueryUID query a user's UID
func QueryUID(platform int, appID string, openID string) (UID uint64, err error) {
	return agent.queryUID(platform, appID, openID)
}

//InsertPlatformUID insert a usre's UID to db
func InsertPlatformUID(platform int, appID string, openID string, UID uint64) error {
	return agent.insertPlatformUID(platform, appID, openID, UID)
}

//BindUID bind a platform UID to master UID
func BindUID(platform int, appID string, openID string, UID uint64) error {
	return agent.bindUID(platform, appID, openID, UID)
}

func (ag *DBAgent) init() error {
	db, err := connect()
	if err != nil {
		return err
	}
	err = createTables(db)
	if err != nil {
		return err
	}
	ag.db = db
	return err
}

func (ag *DBAgent) insertUID(UID uint64) error {
	err := insert(ag.db, "t_ams_use_id", "f_user_id", "f_create_time", UID, time.Now().Format("2006-01-02 15:04:05"))
	return err
}

func (ag *DBAgent) queryUID(platform int, appID string, openID string) (UID uint64, err error) {
	rst, err := query(ag.db, "t_ams_open_id", "f_master_user_id", "f_platform", "f_app_id", "f_open_id", platform, appID, openID)
	if err != nil {
		return
	}
	if rst != nil {
		utils.Logger.Error("rst is null")
		UID = rst.(uint64)
	}
	return
}

func (ag *DBAgent) insertPlatformUID(platform int, appID string, openID string, UID uint64) error {
	err := insert(ag.db, "t_ams_open_id", "f_platform", "f_app_id", "f_open_id", "f_user_id", platform, appID, openID, UID)
	return err
}

func (ag *DBAgent) bindUID(platform int, appID string, openID string, UID uint64) error {
	err := update(ag.db, "t_ams_open_id", "f_master_user_id", UID, "f_platform", "f_app_id", "f_open_id", platform, appID, openID)
	return err
}
