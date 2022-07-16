package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	//flag . ประเภทค่า ( ชื่อคีย์ , ค่าเริ่มต้น , คำอธิบาย )
	port = flag.String("port", "8080", "port to be used")
	ip   = flag.String("ip", "localhost", "ip to be used")
)

func main() {
	flag.Parse()                         // parse/sent data value from command line to be use in flag
	flags := models.NewFlags(*ip, *port) // use * for ref to address of value  if not use it show number of value instead

	fmt.Print("Starting API")

	logger := log.New(os.Stdout, "auth", 1)
	route := routers.NewRoute(logger, flags)
	engine := route.RegisterRoutes()

	url, _ := flags.GetApplicationUrl()
	engine.Run(*url)

}
