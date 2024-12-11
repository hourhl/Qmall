# Qmall

## 项目简介

本项目是2024年字节青训营11月的后端项目，使用微服务架构来实现电子商城。具体要求见[项目方案](./项目方案.md)

参考[Gmall](https://github.com/cloudwego/biz-demo/tree/main/gomall)完成

## 技术栈
\- Go - Hertz   -Kitex  -Consul   - OpenTelemetry   - Gorm   -cwgo   -Redis

## 项目详细信息
### 基本功能

1. 认证中心
   * 分发身份令牌
   * 校验身份令牌
2. 用户服务
   * 创建用户（注册）
   * 登录
3. 商品服务
   * 查询单个商品信息
   * 查询同类商品信息
   * 根据信息搜索商品
4. 购物车服务
   * 加入购物车
   * 清空购物车
   * 获取购物车信息
5. 订单服务
   * 创建订单
   * 获取订单
6. 结算
   * 订单结算
7. 支付
   * 支付
8. 链路追踪 - 集成OpenTelemetry


## 运行

* 获取源码
   ```shell
   git clone https://github.com/hourhl/Qmall.git
   ```
* 启动consul作为服务注册中心
   ```shell
   cd Qmall
   docker-compose up -d
   ```
  可以访问[consul的UI界面](http://127.0.0.1:8500/ui/dc1)来查看注册的服务
* 启动模块

   以user服务为例(其依赖于auth服务，故先启动auth服务)
   ```shell
   # 启动auth服务
   cd app/auth
   go mod tidy
   go run .
   # 启动user服务
   cd ../user
   go mod tidy
   docker-compose up -d
   go run .
   ```
  可以在[consul的UI界面](http://127.0.0.1:8500/ui/dc1/services)看到auth服务和user服务注册成功


## 完善 or 优化

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
4. 购物车服务
   * 添加购物车操作前的身份验证逻辑
5. 订单服务
   * 修改订单信息
   * 订单定时取消
5. 支付
   * 取消支付
   * 定时取消支付



## 参考资料

[cloudwego/gomall视频资料](https://space.bilibili.com/3494360534485730/channel/collectiondetail?sid=2632484)

[kitex文档](https://www.cloudwego.io/zh/docs/kitex/)

[golang-jwt文档](https://pkg.go.dev/github.com/golang-jwt/jwt#section-documentation)



