package app

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tidwall/redcon"
)

func RedisServerHandler(redisConnection *RedisClient, authClientes *AuthClient) func(redcon.Conn, redcon.Command) {
	debug := false
	if os.Getenv("DEBUG") == "true" {
		debug = true
	}
	return func(conn redcon.Conn, cmd redcon.Command) {
		if debug {
			fmt.Println(string(cmd.Raw))
		}
		switch strings.ToLower(string(cmd.Args[0])) {
		default:
			conn.WriteError("ERR unknown command '" + string(cmd.Args[0]) + "'")
		case "auth":
			if len(cmd.Args) != 2 {
				conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
				return
			}
			authClientes.SetClient(strings.Split(conn.RemoteAddr(), ":")[0], string(cmd.Args[1]))
			/// Tratar senha errada
			conn.WriteString("OK")
		case "ping":
			conn.WriteString("PONG")
		case "select":
			conn.WriteString("OK")
		case "quit":
			conn.WriteString("OK")
			conn.Close()
		case "set":
			if len(cmd.Args) != 3 {
				conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
				return
			}
			// Tratar problema na conexão
			redisConnection.Write(string(cmd.Args[1]), string(cmd.Args[2]), 0)
			conn.WriteString("OK")

		case "setex":
			if len(cmd.Args) != 4 {
				fmt.Println("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command" + string(len(cmd.Args)))
				conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
				return
			}
			// Tratar problema na conexão
			expiration, _ := strconv.Atoi(string(cmd.Args[2]))
			redisConnection.Write(string(cmd.Args[1]), string(cmd.Args[3]), expiration*1000)
			conn.WriteString("OK")
		case "get":
			if len(cmd.Args) != 2 {
				conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
				return
			}

			val, ok := redisConnection.Read(string(cmd.Args[1]))

			if !ok {
				conn.WriteNull()
			} else {
				conn.WriteBulk([]byte(val))
			}
		case "del":
			if len(cmd.Args) != 2 {
				conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
				return
			}
			ok := redisConnection.Delete(string(cmd.Args[1]))

			if !ok {
				conn.WriteInt(0)
			} else {
				conn.WriteInt(1)
			}
		}
	}
}
