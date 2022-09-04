package redis

import (
	"github.com/gomodule/redigo/redis"
	rg "github.com/redislabs/redisgraph-go"
)

func Init() {
	conn, _ := redis.Dial("tcp", "127.0.0.1:6379")
	graph := rg.GraphNew("social", conn)
	graph.Delete()
}
