# pkg
golang pkg, common utils


## Features:

- [x] log: 集成 `uber.log`, 开箱即用.
- [x] sync/async: `go func()` 优雅异步.
- [x] retry: 通用重试wrap: 支持 backoff
- [x] orm 集成: `gorm` 扩展
    - mysql 支持
    - 集成 `gorm v1, v2` 版本
- [x] cache: 集成 `redis`
- [x] mq: 集成 `kafka`
- [x] net:
    - [x] http: 集成 `gin` + `auth` API 路由鉴权
- [x] config:
    - [x] toml 格式配置文件读取
- [x] text:
    - [x] xss: 预防 xss 校验
- [x] com: 第三方企业 API SDK 接入
    - [x] jpush: 极光推送
    - [x] sendcloud: 搜狐短信服务


## require:

- uber.log
