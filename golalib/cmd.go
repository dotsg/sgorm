package golalib

import (
	"os"

	"github.com/mitchellh/cli"
	"github.com/olachat/gola/drivers"
	"github.com/spf13/viper"
)

type cmd struct {
}

// GenCommandFactory wrap genCmd to be used in cli
func GenCommandFactory() (cli.Command, error) {
	return genCmd, nil
}

var genCmd *cmd

func (*cmd) Help() string {
	return "generate orm stubs from default toml config"
}

func (*cmd) Run(args []string) int {
	viper.SetConfigName("gola")
	wd, err := os.Getwd()
	if err != nil {
		wd = "."
	}

	configPaths := []string{wd}
	for _, p := range configPaths {
		viper.AddConfigPath(p)
	}
	viper.ReadInConfig()
	viper.AutomaticEnv()

	var driverName string
	var config drivers.Config
	code := -1

	driverName = "mysql"
	flag := true
	config = viper.GetStringMap(driverName)
	if len(config) != 0 {
		code += RunMySql(config)
		flag = false
	}

	driverName = "sqlite"
	config = viper.GetStringMap(driverName)
	if len(config) != 0 {
		code += RunSqlite(config)
		flag = false
	}

	if flag {
		println("Can't find db config")
	}
	return code
}

func (*cmd) Synopsis() string {
	return "generate orm stubs from default toml config"
}
