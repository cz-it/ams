package main

import (
	"encoding/json"
	"github.com/cz-it/ams"
	"github.com/cz-it/ams/utils"
	"os"
)

//ConfigJSONWrapper ConfigJsonWrapper
type ConfigJSONWrapper struct {
	ListenAddr string
	DBAddr     string
	DBName     string
	DBUser     string
	DBPasswd   string
	Plugins    []string
}

//LoadConfig load configure
func LoadConfig(filePath string) (err error) {
	fp, err := os.Open(filePath)
	if err != nil {
		utils.Logger.Error("Open Config file %s error: %s", filePath, err.Error())
		return
	}
	defer fp.Close()

	var config ConfigJSONWrapper
	decoder := json.NewDecoder(fp)
	if err = decoder.Decode(&config); err != nil {
		utils.Logger.Error("Decode Config Error:%s", err.Error())
		return
	}
	utils.Logger.Debug("Load Config file %s Success", filePath)

	ams.Config.ListenAddr = config.ListenAddr
	ams.Config.DBAddr = config.DBAddr
	ams.Config.DBName = config.DBName
	ams.Config.DBUser = config.DBUser
	ams.Config.DBPasswd = config.DBPasswd
	ams.Config.Plugins = config.Plugins
	err = nil
	return
}
