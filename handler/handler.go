package handler

import "github.com/hoisie/redis"

type MyHandler struct {

}

var client redis.Client

func init() {
	client.Password = "test1234"
	client.Addr = "172.16.18.39:6379"
	client.Auth("test1234")
}
func (h *MyHandler) Set(key string, value []byte) error {
	return client.Set(key, value)
}
func (h *MyHandler) Get(key string) ([]byte, error) {
	return client.Get(key)
}
// 不知道怎么实现quit命令,这个框架实现的很不完整，handler拿不到连接句柄，也拿不到request
func (h *MyHandler) Quit() (error) {
	return nil
}
