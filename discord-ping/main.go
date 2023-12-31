package main

import (
	"discord-ping/bot"
	"discord-ping/config"
	"fmt"
)

func main(){
	err := config.ReadConfig()

	if err != nil{
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}