
# 毕业项目
## 交易系统 - 微服务改造
- 微服务架构（BFF、Service、Admin、Job、Task 分模块）
    - 架构分层，前端-PHP-Go
    - PHP作为BFF层
    - Go进行业务模块分布式服务处理
    - 存MySQL前使用Kafak解耦，开发独立Kafka Consumer Job
- API 设计（包括 API 定义、错误码规范、Error 的使用）
    - 统一使用gRPC
    - 统一使用gRPC错误码
    - 目前使用Sentinel ErrorCode方式
    - 使用err.Wrap
    - 对goroutine进行封装，能够捕捉panic信息
- gRPC 的使用
    - 统一使用gRPC，放弃HTTP/1.1
    - 实现服务注册与发现，和负载均衡
    - 使用拦截器做请求响应的统一处理
    - 使用元数据进行超时的处理
- Go 项目工程化（项目结构、DI、代码分层、ORM 框架）
    - 项目结构优化，抛弃传统的MVC架构
    - 面向接口编程、多态、模板方法设计模式
    - 使用go-orm处理MySQL数据库使用
- 并发的使用（errgroup 的并行链路请求
    - 使用errgroup管理程序的启动和退出
- 微服务中间件的使用（ELK、Opentracing、Prometheus、Kafka）
    - 统一使用Kafka消息队列做异步处理、削峰填谷
    - 使用Prometheus自定义指标监控
    - 计划使用ELK进行日志收集与分析
- 缓存的使用优化（一致性处理、Pipeline 优化)
    - 交易系统追求时效性，先更新缓存，再落地数据库
    - 定时任务对所有订单做操作，需要进行批处理


# 毕业总结
- 受益匪浅，还需要回炉重造
