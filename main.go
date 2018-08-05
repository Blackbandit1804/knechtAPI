package main

import (
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		panic("INVALID ARGS: Please enter database DSN as argument!\n" +
			  "username:password@tcp(127.0.0.1)/knecht")
	}
	
	mysql, err := NewMySql(args[0])
	check(err)

	fmt.Println("Running...")

	_, err = CreateHttpServer(5665, mysql)
	check(err)

	// rows, _ := mysql.Query("SELECT * FROM userbots")
	// for rows.Next() {
	// 	var botid, ownerid, prefix, dummy string
	// 	rows.Scan(&botid, &ownerid, &prefix, &dummy, &dummy, &dummy)
	// 	fmt.Println(botid, ownerid, prefix)
	// }

	fmt.Println(mysql, err)
}