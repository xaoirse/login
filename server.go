package main

import (
	"flag"
	"log"

	"github.com/xaoirse/logbook/graph/model"
	_ "github.com/xaoirse/logbook/graph/plugin"
	"github.com/xaoirse/logbook/router"
)

const defaultPort = "4000"

func main() {

	port := flag.String("p", "4000", "Port")
	secret := flag.String("s", "", "Secret")
	flag.Parse()

	db := model.GetDb()
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalln("Error when closing db:", err)
		}
	}()

	router.Start(db, port, secret)
}
