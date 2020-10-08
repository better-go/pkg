package micro

import (
	"github.com/better-go/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/web"
)

func RegisterMicroWithGin(opt ...web.Option, )  {
	svc := web.NewService(opt...)

	svc.Init()

	r := gin.Default()


	svc.Handle("/", r)

	// run:
	if err := svc.Run(); err !=nil{
		log.Errorf("server run error: %v", err)
	}






}

