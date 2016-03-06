package main

import (
	redis "github.com/docker/go-redis-server"
	"github.com/phpor/go-redis-proxy/handler"
	"fmt"
)

func main() {
	conf := redis.DefaultConfig()
	conf.Host("127.0.0.1")
	conf.Port(6389)
	conf.Handler(&handler.MyHandler{})
	server, err := redis.NewServer(conf)
	if err != nil {
		print(err)
		print("abcd")
		return
	}
	err = server.ListenAndServe()
	fmt.Print(err)
}
