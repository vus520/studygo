package main

import (
	"fmt"
	goconf "github.com/Unknwon/goconfig"
)

var (
	ConfigFile = "config.ini"
)

func main() {
	Conf, err := goconf.LoadConfigFile(ConfigFile)

	if err != nil {
		fmt.Println("can't parse config file:" + ConfigFile)
	}

	v, err := Conf.GetValue("Demo", "key2")

	if err != nil {
		fmt.Println("can't get config value {key2}:" + v)
	} else {
		fmt.Println("Demo.key2 = " + v)
	}
}