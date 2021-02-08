

## [推荐查看工具](https://www.iminho.me/)

## 总览:
- [Oauth]
- [Waiting to write...]

--------------------

#### 简要描述：

- [登录]

#### 请求URL:

- http://localhost:8080/oauth2/api/v1/oauth.login

#### 请求方式：

- post

#### 请求参数:

- ` LoginReq ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`username` | 否|string|用户名   |
|`password` | 否|string|用户密码(md5加密)   |
|`timestamp` | 否|int64|当前时间戳(秒,差异系统时间10秒)   |
|`token` | 否|string|验证token md5(username &amp;#43; password &amp;#43; timestamp)   |


#### 请求示例:
```
{
     "password": "",
     "timestamp": 0,
     "token": "",
     "username": ""
}
```

#### 返回参数说明:

- ` LoginResp ` : 授权返回值

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`accessToken` | 否|string|访问令牌   |
|`refreshToken` | 否|string|刷新令牌   |
|`accessExpireTime` | 否|int64|访问令牌过期时间   |
|`refreshExpireTime` | 否|int64|刷新令牌过期时间   |


#### 返回示例:
	
```
{
     "accessExpireTime": 0,
     "accessToken": "",
     "refreshExpireTime": 0,
     "refreshToken": ""
}
```

#### 备注:

- 登录

--------------------

#### 简要描述：

- [刷新token]

#### 请求URL:

- http://localhost:8080/oauth2/api/v1/oauth.refresh_token

#### 请求方式：

- post

#### 请求参数:

- ` RefreshTokenReq ` : 刷新token

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`token` | 否|string|刷新令牌   |


#### 请求示例:
```
{
     "token": ""
}
```

#### 返回参数说明:

- ` RefreshTokenResp ` : 刷新token 返回值

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`accessToken` | 否|string|访问令牌   |
|`refreshToken` | 否|string|刷新令牌   |
|`accessExpireTime` | 否|int64|访问令牌过期时间   |
|`refreshExpireTime` | 否|int64|刷新令牌过期时间   |
|`userInfo` | 否|string|用户信息   |


#### 返回示例:
	
```
{
     "accessExpireTime": 0,
     "accessToken": "",
     "refreshExpireTime": 0,
     "refreshToken": "",
     "userInfo": ""
}
```

#### 备注:

- 刷新token

--------------------

#### 简要描述：

- [应用授权]

#### 请求URL:

- http://localhost:8080/oauth2/api/v1/oauth.authorize

#### 请求方式：

- post

#### 请求参数:

- ` AuthorizeReq ` : 授权请求接口

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`appId` | 否|string|应用的唯一标识   |
|`appKey` | 否|string|公匙   |
|`tokenType` | 否|string|令牌类型   |
|`timestamp` | 否|int64|当前时间戳(秒,差异系统时间10秒)   |
|`token` | 否|string|验证token md5(appid &amp;#43; appKey &amp;#43; timestamp &amp;#43; appSecret)   |


#### 请求示例:
```
{
     "appId": "",
     "appKey": "",
     "timestamp": 0,
     "token": "",
     "tokenType": ""
}
```

#### 返回参数说明:

- ` AuthorizeResp ` : 授权返回值

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`accessToken` | 否|string|访问令牌   |
|`refreshToken` | 否|string|刷新令牌   |
|`accessExpireTime` | 否|int64|访问令牌过期时间   |
|`refreshExpireTime` | 否|int64|刷新令牌过期时间   |


#### 返回示例:
	
```
{
     "accessExpireTime": 0,
     "accessToken": "",
     "refreshExpireTime": 0,
     "refreshToken": ""
}
```

#### 备注:

- 应用授权

--------------------

#### 简要描述：

- [校验token，并获取详细信息]

#### 请求URL:

- http://localhost:8080/oauth2/api/v1/oauth.check_token

#### 请求方式：

- post

#### 请求参数:

- ` CheckTokenReq ` : 校验token req请求

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`token` | 否|string|token   |


#### 请求示例:
```
{
     "token": ""
}
```

#### 返回参数说明:

- ` CheckTokenResp ` : 校验token 返回值

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`accessToken` | 否|string|访问令牌   |
|`userInfo` | 否|string|用户信息   |
|`accessExpireTime` | 否|int64|访问令牌过期时间   |


#### 返回示例:
	
```
{
     "accessExpireTime": 0,
     "accessToken": "",
     "userInfo": ""
}
```

#### 备注:

- 校验token，并获取详细信息
	

--------------------
--------------------

#### 自定义类型:


