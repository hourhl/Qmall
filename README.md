# Qmall

## 项目简介

本项目是2024年字节青训营11月的后端项目。

参考[b站cloudwego的电商项目](https://space.bilibili.com/3494360534485730/channel/collectiondetail?sid=2632484)完成

### 基本功能

1. 认证中心
   * 分发身份令牌
   * 校验身份令牌
2. 用户服务
   * 创建用户（注册）
   * 登录s
3. 商品服务
   * 查询商品信息（单个商品、批量商品）
4. 购物车服务
   * 创建购物车
   * 清空购物车
   * 获取购物车信息
5. 订单服务
   * 创建订单
6. 结算
   * 订单结算
7. 支付
   * 支付



## 技术栈

\- Go - Hertz   -Kitex  -Consul   - OpenTelemetry   - Gorm   -cwgo   -Redis



## 项目结构





## 完善 or 优化

### 前端

* 根据功能适配前端



### 优化服务

1. 认证中心
   * 续期身份令牌
2. 用户服务
   * 用户登出
   * 删除用户
   * 更新用户
3. 商品服务
   * 创建商品
   * 修改商品信息
   * 删除商品
4. 订单服务
   * 修改订单信息
   * 订单定时取消
5. 支付
   * 取消支付
   * 定时取消支付





## 参考资料

[cloudwego/gomall视频资料](https://space.bilibili.com/3494360534485730/channel/collectiondetail?sid=2632484)

[微服务架构下的统一身份认证与授权](https://mtide.net/%E5%BE%AE%E6%9C%8D%E5%8A%A1%E6%9E%B6%E6%9E%84%E4%B8%8B%E7%9A%84%E7%BB%9F%E4%B8%80%E8%BA%AB%E4%BB%BD%E8%AE%A4%E8%AF%81%E4%B8%8E%E6%8E%88%E6%9D%83.html)

