module github.com/better-go/pkg

go 1.13

require (
	github.com/Shopify/sarama v1.19.0
	github.com/ambelovsky/go-structs v1.1.0 // indirect
	github.com/ambelovsky/gosf v0.0.0-20201109201340-237aea4d6109
	github.com/ambelovsky/gosf-socketio v0.0.0-20201109193639-add9d32f8b19 // indirect
	github.com/beinan/fastid v0.0.0-20190107221622-c03a08f42c37
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/codahale/hdrhistogram v0.0.0-00010101000000-000000000000 // indirect
	github.com/fatih/structs v1.1.0
	github.com/gin-gonic/gin v1.7.4
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/go-redis/redis/v8 v8.11.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gocarina/gocsv v0.0.0-20200330101823-46266ca37bd3
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/go-querystring v1.0.0
	github.com/google/uuid v1.1.1
	github.com/gorilla/websocket v1.4.2
	github.com/jinzhu/gorm v1.9.12
	github.com/jinzhu/now v1.1.2
	github.com/jordan-wright/email v0.0.0-20200602115436-fd8a7622303e
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/machinebox/graphql v0.2.2
	github.com/matryer/is v1.4.0 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/micro/go-micro/v2 v2.8.0
	github.com/microcosm-cc/bluemonday v1.0.15
	github.com/opentracing/opentracing-go v1.1.1-0.20190913142402-a7454ce5950e
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/openzipkin/zipkin-go v0.2.1
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/rs/xid v1.2.1
	github.com/shurcooL/graphql v0.0.0-20200928012149-18c5c3165e3a
	github.com/streadway/amqp v1.0.0
	github.com/tal-tech/go-zero v1.1.5
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	github.com/urfave/cli/v2 v2.3.0
	github.com/ylywyn/jpush-api-go-client v0.0.0-20190906031852-8c4466c6e369
	go.uber.org/zap v1.15.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/net v0.0.0-20210917221730-978cfadd31cf // indirect
	golang.org/x/sys v0.0.0-20210921065528-437939a70204 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gorm.io/driver/mysql v0.2.0
	gorm.io/gorm v1.21.6
)

replace (
	github.com/codahale/hdrhistogram => github.com/HdrHistogram/hdrhistogram-go v0.9.0
	// github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 latest
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.2.0
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
