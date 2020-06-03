package redis

import (
	"testing"
)

func TestNewRedis(t *testing.T) {
	cli := NewRedis(&Options{})

	// client list:
	ret, err := cli.ClientList().Result()
	t.Log("client list:", ret, err)

	// set:
	cli.Set("hello", "world", 0)
	// get:
	ret, err = cli.Get("hello").Result()
	t.Log("hello:", ret)

	// ping:
	ret, err = cli.Ping().Result()
	t.Log("ping:", ret, err)

	// options:
	t.Logf("redis options: %+v", cli.Options())
}

func TestNewClient(t *testing.T) {
	cli := NewClient(&Options{})

	// client list:
	ret, err := cli.V1().ClientList().Result()
	t.Log("client list:", ret, err)

	// ping:
	ret, err = cli.V1().Ping().Result()
	t.Log("ping:", ret, err)

}
