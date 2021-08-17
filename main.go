package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "", "学号")
	password := flag.String("password", "", "密码")
	flag.Parse()
	if *username == "" || *password == "" {
		flag.Usage()
		return
	}

	*password = Encrypt(*password)
	jnuID := Login(*username, *password)
	firstID := List(jnuID)
	mainTable, secondTable := Review(firstID, jnuID)
	_, msg := Checkin(jnuID, mainTable, secondTable)
	fmt.Println(msg)
}
