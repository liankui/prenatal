
# user service
get userinfo
[TOC]

## 整体说明
1.	字符串都为utf8格式;
1.	HTTP Headers:
	1.	Content-Type设置为：application/json
1.	DataTime格式参考RFC3339标准

## 错误处理
错误的具体信息将在error字段中返回。

### 错误码示例
```json
{
    "code": "400",
    "message": "Param Error"
}
```


### 状态码列表
| 状态码 | 说明 |
|---|---|
| 200 | 返回正常 |
| 400 | 参数错误 |
| 401 | 无access<br> key或key无效 |
| 500 | 服务器内部错误 |


## get user info

### 请求路径
```http
GET /user/v1/users/{id}
```


### 请求参数

#### Query 参数
| 参数名 | 参数类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `name` | `string` |  | 否 |  | get user info by id |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  | the user id |
| `name` | `string` |  | N |  | the user name |
| `password` | `string` |  | N |  |
| `email` | `string` |  | N |  |
| `phoneNumber` | `string` |  | N |  | 电话号码 |
| `industry` | `string` |  | N |  | 所在行业 |
| `company` | `string` |  | N |  | 公司名称 |
| `job` | `string` |  | N |  | 职业 |
| `role` | `string` |  | N |  | 职业角色 |
| `birthday` | `string` |  | N |  | 生日 |
| `education` | `string` |  | N |  | 学历 1-大专及以下 2-本科 3-硕士 4-博士及以上 |
| `major` | `string` |  | N |  | 专业 |
| `openid` | `string` |  | N |  | 用户在当前小程序的唯一标识 |
| `unionid` | `string` |  | N |  | 微信开放平台帐号下的唯一标识 |
| `nickname` | `string` |  | N |  | 用户昵称 |
| `avatarUrl` | `string` |  | N |  | 用户头像图片的URL |
| `gender` | `integer` | `Int64` | N |  | 用户性别 0-未知 1-男性 2-女性 |
| `country` | `string` |  | N |  | 用户所在国家 |
| `province` | `string` |  | N |  | 用户所在省份 |
| `city` | `string` |  | N |  | 用户所在城市 |
| `language` | `string` |  | N |  | 显示用户信息的语言(default: zh_CN) |
| `desc` | `string` |  | N |  | 声明获取用户个人信息后的用途 |
| `source` | `string` |  | N |  | 用户来源 |
| `eventId` | `integer` | `Int64` | N |  | 事件Id |
| `inviterId` | `integer` | `Int64` | N |  | 邀请者Id |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |
| `deleteTime` | `string` | `Timestamp` | N |  |  |
