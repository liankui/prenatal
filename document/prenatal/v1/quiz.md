
# quiz服务
quiz服务相关的接口
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


## 新增answer

### 请求路径
```http
POST /v1/Answer
```


### 请求参数

#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `questionId` | `string` |  | N |  |
| `questionCategory` | `string` |  | N |  |
| `answer` | `string` |  | N |  |
| `correctAnswer` | `string` |  | N |  |
| `userId` | `string` |  | N |  |
| `userName` | `string` |  | N |  |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `questionId` | `string` |  | N |  |
| `questionCategory` | `string` |  | N |  |
| `answer` | `string` |  | N |  |
| `correctAnswer` | `string` |  | N |  |
| `userId` | `string` |  | N |  |
| `userName` | `string` |  | N |  |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |


## 新增question

### 请求路径
```http
POST /v1/questions
```


### 请求参数

#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `category` | `string` |  | N |  |
| `question` | `string` |  | N |  |
| `options` | `Array<prenatal.QuestionOption>` |  | N |  |
| `explanation` | `string` |  | N |  |
| `userId` | `string` |  | N |  |
| `userName` | `string` |  | N |  |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |


#### `prenatal.QuestionOption`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `questionId` | `string` |  | N |  |
| `questionCategory` | `string` |  | N |  |
| `option` | `string` |  | N |  | 选项 |
| `content` | `string` |  | N |  | 选项的描述 |
| `isCorrect` | `boolean` |  | N |  |
| `explanation` | `string` |  | N |  |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `category` | `string` |  | N |  |
| `question` | `string` |  | N |  |
| `options` | `Array<prenatal.QuestionOption>` |  | N |  |
| `explanation` | `string` |  | N |  |
| `userId` | `string` |  | N |  |
| `userName` | `string` |  | N |  |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |


#### `prenatal.QuestionOption`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `questionId` | `string` |  | N |  |
| `questionCategory` | `string` |  | N |  |
| `option` | `string` |  | N |  | 选项 |
| `content` | `string` |  | N |  | 选项的描述 |
| `isCorrect` | `boolean` |  | N |  |
| `explanation` | `string` |  | N |  |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |


## 获取question详情

### 请求路径
```http
GET /v1/questions/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `id` | `string` |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `category` | `string` |  | N |  |
| `question` | `string` |  | N |  |
| `options` | `Array<prenatal.QuestionOption>` |  | N |  |
| `explanation` | `string` |  | N |  |
| `userId` | `string` |  | N |  |
| `userName` | `string` |  | N |  |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |


#### `prenatal.QuestionOption`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `questionId` | `string` |  | N |  |
| `questionCategory` | `string` |  | N |  |
| `option` | `string` |  | N |  | 选项 |
| `content` | `string` |  | N |  | 选项的描述 |
| `isCorrect` | `boolean` |  | N |  |
| `explanation` | `string` |  | N |  |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |


## 更新question信息

### 请求路径
```http
PUT /v1/questions/{id}
```


### 请求参数

#### Body 请求对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `category` | `string` |  | N |  |
| `question` | `string` |  | N |  |
| `options` | `Array<prenatal.QuestionOption>` |  | N |  |
| `explanation` | `string` |  | N |  |
| `userId` | `string` |  | N |  |
| `userName` | `string` |  | N |  |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |


#### `prenatal.QuestionOption`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `questionId` | `string` |  | N |  |
| `questionCategory` | `string` |  | N |  |
| `option` | `string` |  | N |  | 选项 |
| `content` | `string` |  | N |  | 选项的描述 |
| `isCorrect` | `boolean` |  | N |  |
| `explanation` | `string` |  | N |  |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |


### 返回值

#### 返回对象
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `category` | `string` |  | N |  |
| `question` | `string` |  | N |  |
| `options` | `Array<prenatal.QuestionOption>` |  | N |  |
| `explanation` | `string` |  | N |  |
| `userId` | `string` |  | N |  |
| `userName` | `string` |  | N |  |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |


#### `prenatal.QuestionOption`
| field | type | format | required | default | description |
|---|---|---|---|---|---|
| `id` | `string` |  | N |  |
| `questionId` | `string` |  | N |  |
| `questionCategory` | `string` |  | N |  |
| `option` | `string` |  | N |  | 选项 |
| `content` | `string` |  | N |  | 选项的描述 |
| `isCorrect` | `boolean` |  | N |  |
| `explanation` | `string` |  | N |  |
| `createTime` | `string` | `Timestamp` | N |  |  |
| `updateTime` | `string` | `Timestamp` | N |  |  |


## 删除question

### 请求路径
```http
DELETE /v1/questions/{id}
```


### 请求参数

#### Path 参数
| 参数名 | 参数类型 | 格式类型 | 说明 |
|---|---|---|---|
| `id` | `string` |  |  |


### 返回值

#### 返回对象
对象为空