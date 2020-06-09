package redis

import (
	"context"
	"testing"
	"time"
)

func TestNewRedis(t *testing.T) {
	cli := NewRedis(&Options{})
	ctx := context.Background()

	// client list:
	ret, err := cli.ClientList(ctx).Result()
	t.Log("client list:", ret, err)

	// set:
	cli.Set(ctx, "hello", "world", 0)
	// get:
	ret, err = cli.Get(ctx, "hello").Result()
	t.Log("hello:", ret)

	// ping:
	ret, err = cli.Ping(ctx).Result()
	t.Log("ping:", ret, err)

	// options:
	t.Logf("redis options: %+v", cli.Options())
}

func TestNewClient(t *testing.T) {
	cli := NewClient(&Options{})
	ctx := context.Background()

	// client list:
	ret, err := cli.V1().ClientList(ctx).Result()
	t.Log("client list:", ret, err)

	// ping:
	ret, err = cli.V1().Ping(ctx).Result()
	t.Log("ping:", ret, err)

}

func TestExpiration_Days(t *testing.T) {
	exp := NewExpiration()
	in := []uint32{
		0,
		1,
		3,
		5,
	}

	t.Log("duration test:")

	for _, item := range in {

		// wait:
		time.Sleep(exp.Seconds(item))
		t.Logf("wait [%v]s,  see me...", item)
	}
}
