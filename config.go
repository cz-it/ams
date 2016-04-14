/**
* Author: CZ cz.theng@gmail.com
 */

package ams

import ()

type config struct {
	ListenAddr string
	DBAddr     string
	DBName     string
	DBUser     string
	DBPasswd   string
	Plugins    []string
}

//Config is a siglone of config
var Config config
