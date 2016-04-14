package main

import (
	"flag"
	"fmt"
	"github.com/cz-it/ams"
	"github.com/cz-it/ams/utils"
	"github.com/cz-it/golangutils/daemon"
)

func bootDBAgent() error {
	return nil
}

func loadPlugins() error {
	return nil
}

func serveAPI() error {
	return nil
}

func bootServer() {
	var err error
	err = bootDBAgent()
	if err != nil {
		utils.Logger.Error("Boot DB Agent Error %s", err.Error())
		return
	}

	err = loadPlugins()
	if err != nil {
		utils.Logger.Error("Load Plugins Error %s", err.Error())
		return
	}

	err = serveAPI()
	if err != nil {
		utils.Logger.Error("Serve API Error %s", err.Error())
		return
	}
}

func main() {
	if Flag.Version {
		fmt.Println("Cur Version:%s", ams.Version())
		return
	}

	if Flag.Config == "" {
		flag.Usage()
		return
	}

	if err := LoadConfig(Flag.Config); err != nil {
		println("Loading Config Error")
		return
	}
	fmt.Println(ams.Config)

	if Flag.Daemon {
		daemon.Boot("/tmp/magline.lock", "/tmp/magline.pid", func() {
			bootServer()
		})
	} else {
		bootServer()
	}

	println("[Testing]:End")
}
