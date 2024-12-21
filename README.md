# WEB-TUTORIAL
web应用

## 目录
 http.Handler接口
 http.HandlerFunc类型
 路由
 五个内置handler
 NotFoundHandler
 RedirectHandler
 StripPrefix
 TimeoutHandler
 FileServer

 #URL的通用格式
 scheme://[userinfo@]host/path[?query][#fragment]
 不可以斜杠开头的URL被解释为：
 scheme:opaque[?query][#fragmen]

 ## 模版

ParseFiles函数可以解析模板文件，并返回*template.Template类型的对象。