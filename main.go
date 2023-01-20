package main

import (
	"embed"
	"fmt"
	//"log"
	"strconv"

	api "github.com/seekr-osint/seekr/api"
	webServer "github.com/seekr-osint/seekr/webServer"
)

// Web server content
//
//go:embed web/*
var content embed.FS

var people = make(api.DataBase)

func main() {
  //fmt.Println(api.PortScan("scannme.nmap.org"))
	//if OPENBROWSER {
		//api.Openbrowser("http://localhost:5050")
	//}
	//api.Emails("fragenwert@gmail.com")
	//mail := api.GithubInfoDeep("niteletsplay",true)
	//fmt.Println(mail)
	fmt.Println("Welcome to seekr a powerful OSINT tool able to scan the web for " + strconv.Itoa(len(api.DefaultServices)) + "")
	go api.ServeApi(people, ":8080", "data.json") // TODO config parsing stuff
	RunWebServer()
	fmt.Println(api.ServicesHandler(api.DefaultServices, "9glenda"))
}

// //export RunWebServer
func RunWebServer() {

	var config = webServer.WebServerConfig{
		Type:    webServer.SingleBinary,
		Content: content,
		Dir:     "./web",
		Ip:      ":5050",
	}
	webServer.ParseConfig(config)
}
