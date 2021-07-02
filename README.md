

## [推荐查看工具](https://www.iminho.me/)

## 总览:
- [Oauth]
- [Waiting to write...]

--------------------

#### 简要描述：

- [更新/删除用户]

#### 请求URL:

- http://localhost:8080/oauth2/api/v1/oauth.update_user

#### 请求方式：

- post

#### 请求参数:

- ` UpdateUserReq ` : 更新用户

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`rootName` | 否|string|创建者   |
|`userName` | 否|string|用户名(要操作的用户名)   |
|`op` | 否|int32|操作(1:添加，-1:删除)   |
|`password` | 否|string|用户密码(md5码)(如果有修改,不为空)   |
|`userInfo` | 否|string|应用信息(一般base64 json串)   |


#### 请求示例:
```
{
     "op": 0,
     "password": "",
     "rootName": "",
     "userInfo": "",
     "userName": ""
}
```

#### 返回参数说明:

- ` Empty ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |


#### 返回示例:
	
```
{}
```

#### 备注:

- 更新/删除用户

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
|`userName` | 否|string|用户名   |
|`accessExpireTime` | 否|int64|访问令牌过期时间   |
|`appId` | 否|string|应用的唯一标识   |


#### 返回示例:
	
```
{
     "accessExpireTime": 0,
     "accessToken": "",
     "appId": "",
     "userName": ""
}
```

#### 备注:

- 校验token，并获取详细信息

--------------------

#### 简要描述：

- [创建oauth]

#### 请求URL:

- http://localhost:8080/oauth2/api/v1/oauth.create_oauth

#### 请求方式：

- post

#### 请求参数:

- ` CreateOauthReq ` : 创建oauth

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`username` | 否|string|用户名   |
|`oauthInfo` | 否|string|应用信息(一般base64 json串)   |


#### 请求示例:
```
{
     "oauthInfo": "",
     "username": ""
}
```

#### 返回参数说明:

- ` CreateOauthResp ` : 创建oauth

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`appId` | 否|string|应用id   |
|`appKey` | 否|string|公匙   |
|`appSecret` | 否|string|私匙   |


#### 返回示例:
	
```
{
     "appId": "",
     "appKey": "",
     "appSecret": ""
}
```

#### 备注:

- 创建oauth

--------------------

#### 简要描述：

- [获取应用信息]

#### 请求URL:

- http://localhost:8080/oauth2/api/v1/oauth.get_app_list

#### 请求方式：

- post

#### 请求参数:

- ` GetAppListReq ` : 获取应用

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`username` | 否|string|用户名   |
|`appIds` | 否|[]string|应用id(如果有，用户名无效)   |
|`pageNo` | 否|int32|页数   |
|`pageSize` | 否|int32|每页条数   |


#### 请求示例:
```
{
     "appIds": [
          ""
     ],
     "pageNo": 0,
     "pageSize": 0,
     "username": ""
}
```

#### 返回参数说明:

- ` GetAppListResp ` : 获取应用

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`list` | 否|[]`oauth2.AppInfo`|返回数据   |
|`total` | 否|int64|总条数   |


#### 返回示例:
	
```
{
     "list": [
          {
               "appId": "",
               "oauthInfo": ""
          }
     ],
     "total": 0
}
```

#### 备注:

- 获取应用信息

--------------------

#### 简要描述：

- []

#### 请求URL:

- http://localhost:8080/oauth2/api/v1/oauth.get_login_info

#### 请求方式：

- post

#### 请求参数:

- ` GetLoginInfoReq ` : 获取用户信息请求参数

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`accessToken` | 否|string|用户token (跟用户名二选一)   |


#### 请求示例:
```
{
     "accessToken": ""
}
```

#### 返回参数说明:

- ` GetLoginInfoResp ` : 获取用户信息返回参数

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`username` | 否|string|用户名   |
|`accessExpireTime` | 否|int64|访问令牌过期时间   |
|`userInfo` | 否|string|用户信息(一般base64 json串)   |


#### 返回示例:
	
```
{
     "accessExpireTime": 0,
     "userInfo": "",
     "username": ""
}
```

#### 备注:

- 

--------------------

#### 简要描述：

- [获取用户列表]

#### 请求URL:

- http://localhost:8080/oauth2/api/v1/oauth.get_users

#### 请求方式：

- post

#### 请求参数:

- ` GetUsersReq ` : 获取用户信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`rootName` | 否|string|创建者   |
|`userName` | 否|string|用户名(为空获取所有用户列表)   |


#### 请求示例:
```
{
     "rootName": "",
     "userName": ""
}
```

#### 返回参数说明:

- ` GetUsersResp ` : 获取用户信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`list` | 否|[]`oauth2.UserInfo`|用户信息   |


#### 返回示例:
	
```
{
     "list": [
          {
               "userInfo": "",
               "username": ""
          }
     ]
}
```

#### 备注:

- 获取用户列表

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
|`userName` | 否|string|用户名(没有可以不用填写)   |


#### 请求示例:
```
{
     "appId": "",
     "appKey": "",
     "timestamp": 0,
     "token": "",
     "tokenType": "",
     "userName": ""
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

- [创建用户]

#### 请求URL:

- http://localhost:8080/oauth2/api/v1/oauth.create_user

#### 请求方式：

- post

#### 请求参数:

- ` CreateUserReq ` : 创建用户

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`rootName` | 否|string|创建者   |
|`userName` | 否|string|用户名   |
|`password` | 否|string|用户密码(md5码)   |
|`userInfo` | 否|string|应用信息(一般base64 json串)   |
|`regIP` | 否|string|注册时间   |


#### 请求示例:
```
{
     "password": "",
     "regIP": "",
     "rootName": "",
     "userInfo": "",
     "userName": ""
}
```

#### 返回参数说明:

- ` Empty ` : 

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |


#### 返回示例:
	
```
{}
```

#### 备注:

- 创建用户

--------------------

#### 简要描述：

- [List 删除应用]

#### 请求URL:

- http://localhost:8080/oauth2/api/v1/oauth.delete_app

#### 请求方式：

- post

#### 请求参数:

- ` DeleteAppReq ` : 删除应用

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`appId` | 否|string|appid   |
|`username` | 否|string|用户名   |


#### 请求示例:
```
{
     "appId": "",
     "username": ""
}
```

#### 返回参数说明:

- ` DeleteAppResp ` : 删除应用

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`rowsAffected` | 否|int64|受影响的条数   |


#### 返回示例:
	
```
{
     "rowsAffected": 0
}
```

#### 备注:

- List 删除应用

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
--------------------

#### 自定义类型:

#### ` oauth2 `


- ` AppInfo ` : 应用信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`appId` | 否|string|appid   |
|`oauthInfo` | 否|string|应用信息(一般base64 json串)   |



- ` UserInfo ` : 用户信息

|参数名|是否必须|类型|说明|
|:----    |:---|:----- |-----   |
|`username` | 否|string|用户名   |
|`userInfo` | 否|string|应用信息(一般base64 json串)   |




