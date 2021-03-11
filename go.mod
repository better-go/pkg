module github.com/better-go/pkg

go 1.13

require (
	github.com/Shopify/sarama v1.19.0
	github.com/ambelovsky/go-structs v1.1.0 // indirect
	github.com/ambelovsky/gosf v0.0.0-20201109201340-237aea4d6109
	github.com/ambelovsky/gosf-socketio v0.0.0-20201109193639-add9d32f8b19 // indirect
	github.com/beinan/fastid v0.0.0-20190107221622-c03a08f42c37
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/go-redis/redis/v8 v8.0.0-beta.4
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gocarina/gocsv v0.0.0-20200330101823-46266ca37bd3
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/uuid v1.1.1
	github.com/gorilla/websocket v1.4.2
	github.com/jinzhu/gorm v1.9.12
	github.com/jinzhu/now v1.1.1
	github.com/jordan-wright/email v0.0.0-20200602115436-fd8a7622303e
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/machinebox/graphql v0.2.2
	github.com/matryer/is v1.4.0 // indirect
	github.com/micro/go-micro/v2 v2.8.0
	github.com/microcosm-cc/bluemonday v1.0.2
	github.com/opentracing/opentracing-go v1.1.1-0.20190913142402-a7454ce5950e
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.4.5
	github.com/openzipkin/zipkin-go v0.2.1
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron/v3 v3.0.1
	github.com/rs/xid v1.2.1
	github.com/shurcooL/graphql v0.0.0-20200928012149-18c5c3165e3a
	github.com/streadway/amqp v1.0.0
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	github.com/ylywyn/jpush-api-go-client v0.0.0-20190906031852-8c4466c6e369
	go.uber.org/zap v1.15.0
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/tools v0.0.0-20200501005904-d351ea090f9b // indirect
	google.golang.org/protobuf v1.24.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gorm.io/driver/mysql v0.2.0
	gorm.io/gorm v0.2.7
)

replace (
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
