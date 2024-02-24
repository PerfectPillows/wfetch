package main

import (
	"fmt"

	"helper"
)

func main() {

	totalMemory, err := helper.GetTotalRAM()
	if err != nil {
		panic(err)
	}

	usedMemory, err := helper.GetUsedRAMAmount()
	if err != nil {
		panic(err)
	}

	userName, err := helper.GetUserName()
	if err != nil {
		panic(err)
	}

	manufacturer, model, err := helper.GetHostInfo()
	if err != nil {
		panic(err)
	}

	uptime, err := helper.GetUptimeInfo()
	if err != nil {
		panic(err)
	}

	line := ""
	for i := 0; i < len(userName); i++ {
		line += "-"
	}
	teal := "\033[38;5;6m"
	reset := "\033[0m"

	helper.ReadOSAsciiArt()
	fmt.Println()
	fmt.Println(teal + userName + reset)
	fmt.Println(line)
	fmt.Println(teal, "OS : ", reset, helper.GetOSVersion())
	fmt.Println(teal, "HOST : ", reset, manufacturer, model)
	fmt.Println(teal, "UPTIME : ", reset, uptime)
	fmt.Println(teal, "RAM USAGE : ", reset, usedMemory, "GB / ", totalMemory, "GB")
}
