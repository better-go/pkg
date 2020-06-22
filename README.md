# pkg
golang pkg, common utils

## related:

- https://github.com/better-go/cookiecutter-go
    - `go 微服务`: 项目创建脚手架.

## quickstart:

- install:

```bash 

go get -u -v github.com/better-go/pkg

```

- usage:

```golang 

import (
	"github.com/better-go/pkg/random"
)

func Example() {
	// snowflake id:
	userID := random.SnowFlakeID()
	fmt.Printf("userID:%v\n", userID)
}

```


## features:

- [x] `log`: 集成 `uber.log`, 开箱即用.
- [x] sync/async: `go func()` 优雅异步.
- [x] `retry`: 通用重试wrap: 支持 backoff
- [x] `orm` 集成: `gorm` 扩展
    - mysql 支持
    - 集成 `gorm v1, v2` 版本
- [x] `cache`: 集成 `redis`
- [x] `mq`: 集成 `kafka`
- [x] `net`:
    - [x] http: 集成 `gin` + `auth` API 路由鉴权
- [x] `config`:
    - [x] toml 格式配置文件读取
- [x] `random`:
    - [x] `uuid`: ID 生成器
    - [x] `snowflake`: twitter 雪花算法: 18位纯数字(有序+递增+唯一)
- [x] `text`:
    - [x] xss: 预防 xss 校验
- [x] `time`: 大量时间方法扩展
- [x] `crypto`: 密码加密/验证
- [x] `com`: 第三方企业 API SDK 接入
    - [x] jpush: 极光推送
    - [x] sendcloud: 搜狐短信服务


## wiki:

- https://github.com/better-go/pkg/wiki
- go 编程规范

## requirements:

- uber.log
