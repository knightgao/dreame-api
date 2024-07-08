# 使用 API 操控 & 扩展 Dreame API

例如，虽然 Dreame API 本身没有直接支持支付，但是你可以通过系统扩展的 API 来实现支付功能。

又或者你想自定义渠道管理策略，也可以通过 API 来实现渠道的禁用与启用。

## 请求格式与响应格式
Dreame API 使用 JSON 格式进行请求和响应。

对于响应体，一般格式如下：
```json
{
  "message": "请求信息",
  "success": true,
  "data": {}
}
```

## API 列表
> 当前 API 列表不全，请自行通过浏览器抓取前端请求

如果现有的 API 没有办法满足你的需求，欢迎提交 issue 讨论。

### 获取当前登录用户信息
**GET** `/api/user/self`

### 为给定用户充值额度
**POST** `/api/topup`
```json
{
  "user_id": 1,
  "quota": 100000,
  "remark": "充值 100000 额度"
}
```
