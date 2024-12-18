# Qmall

## 项目简介

本项目是2024年字节青训营11月的后端项目，使用微服务架构来实现电子商城。具体要求见[项目方案](./项目方案.md)

参考[Gmall](https://github.com/cloudwego/biz-demo/tree/main/gomall)完成

## 技术栈
\-Go -cwgo -Kitex -Gorm -MySQL -Redis -Consul -Prometheus -OpenTelemetry          

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
8. 缓存服务
   * 用户登录后生成的token存储在redis中
   * 商品查询后存储在redis中
9. 可观测性服务
   * Prometheus - 时间序列指标
   * OpenTelemetry - 链路追踪
   * Loki - 日志
   * Grafana - 提供可视化


## 运行

* 获取源码
   ```shell
   git clone https://github.com/hourhl/Qmall.git
   ```
* 启动项目的依赖，包括consul作为服务注册中心、Grafana作为可视化监控等
   ```shell
   cd Qmall
   docker-compose up -d
   ```
  

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

* 监控
  * 通过consul查看注册的服务 : http://127.0.0.1:8500/ui/dc1
  * 通过Grafana进行可视化监控 ：http://127.0.0.1:3000
  * 通过Jaeger进行可视化监控 : http://127.0.0.1:16686

   

## 参考资料

[cloudwego/gomall视频资料](https://space.bilibili.com/3494360534485730/channel/collectiondetail?sid=2632484)

[kitex文档](https://www.cloudwego.io/zh/docs/kitex/)

[golang-jwt文档](https://pkg.go.dev/github.com/golang-jwt/jwt#section-documentation)



