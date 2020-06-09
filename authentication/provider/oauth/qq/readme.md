


# QQ API 接入:


## QQ 登录接入: 

- 个人应用注册: https://connect.qq.com/
- [接入流程](https://zhuanlan.zhihu.com/p/35651406)
- [QQ 扫码登录接入](https://wiki.connect.qq.com/__trashed-2)
    - [ web vs app 接入方式](https://wiki.connect.qq.com/oauth2-0%E7%AE%80%E4%BB%8B)
- [【QQ登录】网站接入](https://wiki.open.qq.com/wiki/%E3%80%90QQ%E7%99%BB%E5%BD%95%E3%80%91%E7%BD%91%E7%AB%99%E6%8E%A5%E5%85%A5)
- [OAuth2.0简介](https://wiki.open.qq.com/wiki/mobile/OAuth2.0%E7%AE%80%E4%BB%8B)
- [使用Authorization_Code获取Access_Token](https://wiki.connect.qq.com/%E4%BD%BF%E7%94%A8authorization_code%E8%8E%B7%E5%8F%96access_token)
- [接口错误码](https://wiki.connect.qq.com/%E5%85%AC%E5%85%B1%E8%BF%94%E5%9B%9E%E7%A0%81%E8%AF%B4%E6%98%8E)
- [FAQ](https://wiki.open.qq.com/wiki/%E3%80%90QQ%E7%99%BB%E5%BD%95%E3%80%91FAQ)



## 参数说明: 


- `Access Token`: 是应用在调用OpenAPI访问和修改`用户数据`时必须传入的参数
- `openid`: 是此网站上唯一对应用户身份的标识，网站可将此ID进行存储便于用户下次登录时辨识其身份，或将其与用户在网站上的原有账号进行绑定


### ref: 

- https://www.eyesmoons.com/article/8
- https://learnku.com/articles/36658
- https://github.com/leo0o/simpleoauth
    - qq/wechat/weibo
- https://github.com/go-oauth2/oauth2
    - server lib
- https://github.com/golang/oauth2
    - client sdk




