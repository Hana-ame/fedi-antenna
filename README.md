# fedi-antenna


## modules

utils 一些其他工具
webfinger webfinger相关工具
mycurl 与远程访问相关的内容
db 数据库

activitypub objects相关
actions 动作

core 接入点

## trivial 

- main.go > verify() 修改一下其中调用的函数名，规范
- webfinger 返回的时候注意大小写
- 更加不要紧的、
  - 其实不是很清楚rsa2017是不是要每时每刻更新的。失败了似乎是要更新的吧。忘了。

# dev

在 main.go 按下 F5 进行调试

# log

从下往上读的，越靠上越近

## TODO

- **错误处理** 这只是个大主题
  - fetch隔壁user失败之后的问题
  - 。。。

inbox还没写
db还没写

其实是不是很简单，
签名各种还要分一下是哪个用户这样的。


## 08/30

- user的返回
- ap的发送
  - 要弄好签名，签名直接从用户名取得pem文件这样吧。和用户系统分开来，在core的内容
- ap的接受


~~测试webfinger的时候没反应~~
哦，是因为没写user？
~~但是没看见访问~~ 难绷，弄一下好了。
那这个算解决了

inbox怎么写

~~做webfinger的server部分。~~ 好了
verify也要fetch远程user，先不进行缓存，缓存之后在写
几个函数的命名之后再改？

TestGetUserIdFromAcct 成功 （meromero@p1.a9z.dev）

webfinger fetch 成功 （meromero@p1.a9z.dev）