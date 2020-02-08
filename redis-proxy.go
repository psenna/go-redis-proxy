package main

import (
	"log"
	"os"

	"github.com/psenna/go-redis-proxy/app"
	"github.com/tidwall/redcon"
)

var addr = ":6380"

var password = "12345"

func main() {
	redisConnection := app.NewRedisClient()

	authClients := app.GetAuthClient()

	go log.Printf("started server at %s", addr)
	err := redcon.ListenAndServe(addr,
		app.RedisServerHandler(redisConnection, authClients),
		func(conn redcon.Conn) bool {
			// use this function to accept or deny the connection.
			// log.Printf("accept: %s", conn.RemoteAddr())
			return true
		},
		func(conn redcon.Conn, err error) {
			// this is called when the connection has been closed
			// log.Printf("closed: %s, err: %v", conn.RemoteAddr(), err)
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}

func checkAuth(string) {

}

func initVariables() {
	if os.Getenv("ADDR") != "" {
		addr = os.Getenv("ADDR")
	}

	if os.Getenv("PASSWORD") != "" {
		password = os.Getenv("PASSWORD")
	}
}
