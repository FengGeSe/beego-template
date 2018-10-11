# beego-template
beego框架开发api服务器的项目模板

# 文件目录
```
conf/
  app.conf   # 应用的配置

controllers/
    base.go    # 定义了BaseController，对beego.Controller进行了扩展
    error.go   # 定义了ErrorController, 对 500、404 等错误进行了处理
    test.go    # 定义了TestController, 暴露出/test/v1接口用于测试

models/
    api.go    # 定义了接口的数据结构，响应结构体等
    pagination.go    # 定义了分页相关的数据结构

routers/
    router.go    # 路由

tests/
    default_test.go  # 单元测试

main.go  # 程序入口
```


# 测试接口
ps: domainname和appname在conf/app.conf里面定义
http://{ip}:{port}/{domainname}/{appname}/test/v1
例如：
1. say hello
http://localhost:8080/api.fenggese.com/beego-template/test/v1
```
{
    code: "ok",
    message: "success",
    data: {
        test: "hello, my appname is: beego-template"
    }
}
```
2. 服务端报错: http://localhost:8080/api.fenggese.com/beego-template/test/v1/400
```
{
    code: "test:Error500",
    message: "服务器出错",
    data: null
}
```
3. 客户端报错: http://localhost:8080/api.fenggese.com/beego-template/test/v1/500
```
{
    code: "test:Error400",
    message: "客户端请求错误",
    data: null
}
```
