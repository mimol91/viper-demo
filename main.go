package main

import (
	"fmt"

	"demo-viper/conf"
)

func main() {
	conf.Init()
	fmt.Print(conf.R)

	select {}

}
