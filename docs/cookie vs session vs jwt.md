### Cookie vs Session:

![image.png](https://upload-images.jianshu.io/upload_images/4694144-52f07fd1dee33743.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

> Reference: 
>
> - https://www.ifb.me/zh/blog/backend/web-cun-chu-zhi-zhen

#### 区别：

- 存储位置： Cookie存储在客户端(浏览器)，而session存储在服务器端；
- 存储容量：Cookie一般限制单个domain下不能超过4KB（http://browsercookielimits.iain.guru/），而session的存储可以很大；
- 安全性：由于Cookie存储在客户端，可以被用户删除和修改；而session存储在服务端，客户端通过session ID与服务端进行交互，相比于Cookie更加安全些；
- 传输：Cookie的传输一般依赖与浏览器Cookie的行为，而在App和小程序环境下Cookie是不生效的，需要手动处理Cookie的传递；而Session的传递严格意义上来说可以不依赖Cookie，可以通过Header传递，也可以适配Bear Authorization的规范；
- 跨域：在跨域请求时，Cookie是不会自动传递的。





### Session vs JWT:

![Session vs JWT](https://tech.lucumt.info/img/security/session-vs-jwt.gif)

>  References: 
>
> - https://tech.lucumt.info/docs/security/session-vs-jwt/
> - https://www.youtube.com/watch?v=fyTxwIa-1U0

#### 区别：

- 存储依赖：在分布式系统中，多个服务需要同时写入和读取session，故需要session的单独存储；对该存储的操作会增加接口响应的延时；另外如果存储系统故障，则会导致整个Authorization系统的故障。
- 失效控制：服务端可以操作session数据来快速失效部分session，已达到避免潜在认证风险带来的问题；而JWT是使用计算来代替了存储，不能通过修改数据来完成认证的快速失效，一般只能依赖JWT的过期时间（可以通过Refresh Token机制来缩短JWT的过期时间，进而降低潜在风险）。