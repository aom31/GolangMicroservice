package main

import (
	"flag"
	"fmt"
)

var (
	//flag . ประเภทค่า ( ชื่อคีย์ , ค่าเริ่มต้น , คำอธิบาย )
	port = flag.String("port", "8080", "port to be used")
	ip   = flag.String("ip", "localhost", "ip to be used")
)

func main() {
	flag.Parse()

	fmt.Print("Starting API")
}
