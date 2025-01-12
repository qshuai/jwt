cookie
===

这里只是简单的验证cookie数据可以完全由服务端来控制，这里也验证了`Cookie with Session id`可以完全通过调整客户端的数据，进而间接管理客户端认证的能力。这看起来是很傻的验证，不过也是知识体系的一部分。



### 验证步骤：

1. 运行服务端程序：

   ```shell
   $ go run main.go
   ```

2. 模拟登录接口，服务端通过后将认证信息返回给客户端：

   浏览器打开：http://127.0.0.1:8080/setcookie； 打开浏览器调试工具，查看该域下的Cookie数据，以下是我这边的展示：

   ![image.png](https://upload-images.jianshu.io/upload_images/4694144-2de9b0c3af618147.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

3. 模拟服务端控制客户端Cookie数据的情况：

   浏览器打开：http://127.0.0.1:8080/override_cookie； 你会发现客户端的Cookie完全按照服务端的数据进行了覆盖。

   ![image.png](https://upload-images.jianshu.io/upload_images/4694144-39ff56853c3af879.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)



### 结论：

在浏览器环境下，服务端可以完全控制Cookie中的数据，那么如果包含Session id的话，也同样可以进行控制。而JWT一旦发布，就很难进行吊销，只能等待其过期（否者条件额外的控制逻辑）。
