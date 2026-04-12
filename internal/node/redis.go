package node

import (
	"context"
	"log"
	"strings"

	"github.com/tidwall/redcon"
)

func Listen(ctx context.Context) error {
	err := redcon.ListenAndServe("localhost:6379", func(conn redcon.Conn, cmd redcon.Command) {
		switch strings.ToLower(string(cmd.Args[0])) {
		case "command":
			conn.Close()
		case "ping":
			conn.WriteAny("PONG")
		default:
			log.Println("unknow command", string(cmd.Args[0]))
			conn.WriteString("unknow command")
		}
	}, func(conn redcon.Conn) bool {
		return true
	}, func(conn redcon.Conn, err error) {
	})
	return err
}
