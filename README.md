# fedi-antenna


## modules

utils 一些其他工具
webfinger webfinger相关工具
mycurl 与远程访问相关的内容
db 数据库

activitypub objects相关
actions 动作

### core/actions

所有和http相关的，不管内部东西？

好像还是不行，提供了所有需要的内容然后完成

core 接入点

## trivial 

- main.go > verify() 修改一下其中调用的函数名，规范
- webfinger 返回的时候注意大小写
- 更加不要紧的、
  - 其实不是很清楚rsa2017是不是要每时每刻更新的。失败了似乎是要更新的吧。忘了。

## notes

http的载体

url， endpoint


ap的载体

id actor等


url对应了资源

# dev

在 main.go 按下 F5 进行调试

## db

### sqlite

测试的时候使用navicat读一下内容

## script

跑一个访问3000端口的py，访问不到直接砍进程试试。
现在会出毛病吗，不会出就不写这个了，好麻烦

# log

从下往上读的，越靠上越近

## TODO

- **错误处理** 这只是个大主题
  - fetch隔壁user失败之后的问题
  - 。。。

inbox还没写
**db完全没写，连参考代码都没有的。**

- utils.PkPem 的方式太傻了，要不要改到数据库里面。
- 修改远程

其实是不是很简单，
签名各种还要分一下是哪个用户这样的。


## 08/30

- ~~user的返回~~
- ap的发送
  - 要弄好签名，签名直接从用户名取得pem文件这样吧。和用户系统分开来，在core的内容
- ap的接受

~~看test的时间戳都知道你有多摸~~

db写了，好像没啥问题，crud见test
只有单个，要多个么，好像不用。
考虑一下user的notification和hometimeline怎么处理

做到一半发现之后core会处理不好object
~~先做db~~

db考虑人和人之间relationship做一张表？
那也是之后的事情了

~~做一下db然后今天结束了吧~~

✔ 验证inbox的签名是可用的。
但是每次fetch user，这个之后修改吧

之前应该做了follow，批准follow
还有note的送信这些。

接下来是做Follow UnFollow之类的
数据库完全没考虑。之后做

user做好返回了，现在可以供测试
里面的东西一直没做，到时候需要把api再研究一下

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