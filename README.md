# 24

不知道脑子抽了什么疯反正mastodon改到单独的db里面了，有前端就可以上测试了。  
ap的部分接受和发送还能分开做，先做接收然后再做发送再补接收。服了。


还是有点不明要干啥。
先做mastodon兼容api，然后在这之上兼容antenna，再兼容ap。
mastodon要用的也就只有mastodon几个entities吧。
那些个写出来就ok了，存储量之类的优化以后搞。
头好痛头好痛。

- [x] antenna account
  - 在handler里面create了mastoodn的和local的
- [x] mastodon follow unfollow block
- [x] mastodon post
- [x] mastodon fav reblog
  
不对，编译不起来吧。
只能全部改完再看结果了。
可能会遇到顺序反了之类的，save和read都做了要test

code in webfinger is in bad quality, do not use it.

ap recv 的 model 完全是放在orderedmap里面的啊。
那么cache的话借用action的model。
在做完mastodon之后思考。


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

### core/webfinger
client
通过 x@x.x 找到 webfinger object
```go
func FetchWebfingerObj(acct string) (o *orderedmap.OrderedMap, err error) 
```
server
通过 name host 生成webfinger。大概要改了加功能。
```go
func CreateWebfingerObj(username, host string) (o *orderedmap.OrderedMap, err error)
```
接在`gin.go`里

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

## curl

```sh
curl -v https://fedi.moonchan.xyz/remoteusers?acct=meromero@p1.a9z.dev
# post user for test
curl -v -X POST \
-H 'Content-Type: application/json' \
-d @script/test.json \
https://fedi.moonchan.xyz/api/v1/users 
```

s2s测试
```sh
# webfinger
curl -v https://fedi.moonchan.xyz/.well-known/webfinger?resource=acct:nanaka@fedi.moonchan.xyz

# AS:Person
curl -v \
--header 'Accept: application/ld+json' \
--header 'Accept: application/activity+json' \
https://fedi.moonchan.xyz/users/Nanaka 

```

### TEST

从下往上读的，越靠上越近

## TODO


```txt
inbox 
  +
分流
  +-----+
  +    记录
  +
  +-------------+
stream处理   存入各表
```

stream处理就是stream处理
需要做很多的处理，所以还是存入各表然后轮询的方式做最开始的prototype

```txt
action
  +-----------+
 首次发送     记录/存入各表
  
```

接下来可能注重本地db的搭建。
话说单实例用sqlite也行么。。

就做喜欢转发回复的好了。

说不定设计反了。再说吧
misskey也是先是sns再是加入fedi的。
擦。

- **错误处理** 这只是个大主题
  - fetch隔壁user失败之后的问题
  - 。。。
- user
  - ~~只做了开始的部分~~
  - ~~后面还要做修改，db会进去~~
  - 有架子了
- inbox还没写
  - ~~记录~~
  - 要怎么处理
  - 根据不同的内容分表？
- **~~db完全没写，连参考代码都没有的。~~** 开始写了
  - ~~不会，没想过~~
- trivial
  - ~~utils.PkPem 的方式太傻了，要不要改到数据库里面。~~  改了
  - 修改远程 <- 这是在说啥
  - 人与人之间的关系独立建表？
- action
  - 动作发出去的object要往本地存一份
  - 至少有自己的object被访问的db
  - 还有供本地用户用的分表
  - 那么 db 还需要存这些东西

### activitypub api

先弄清楚所有的api
- [x] 发note
  - [x] 带@
    - [ ] 带多个@
  - [x] 带投票
    - [ ] 返回投票
  - [x] 带权限
    - 影响 to 和 cc
  - [x] 带warning
    - 试过了，是warnign
  - [x] 带图片
  - [x] 带sensitive图片
    - [ ] blurhash 怎么搓
- [x] 修改note
  - update
- [x] 删note
  - 下面的好像没影响
  - [ ] 带@
  - [ ] 带投票
  - [ ] 带权限
  - [ ] 修改过的会怎么样
  - [ ] 被转发过的
  - [ ] 被点赞过的
- [ ] 转发note
  - [ ] 取消
- [ ] 点赞note
  - [ ] 取消
- [ ] 投票note

- [x] follow
  - follow
  - [x] 批准follow
    - accept
  - [x] 不批准
    - reject
  - [x] unfollow
    - undo
- [x] block
  - block
- [x] 解除block
  - undo

note的id是独一无二的。
不知道改用时间戳还是内部的objectid
修改的时候要怎么弄呃。
难不成额外做个内部的uuid去link。不好弄吧。
加个key做往前记录的链接之类的。
history列表应该是本地实现吧。

return的时候要不要加上status啊，覆盖率？是不是进入了正确的分支？


先通过`/api/v1/accounts/lookup?acct=[Username]`查找id,直接返回obj
`https://mona.do/api/v1/accounts/[id]`这里读取obj里的id
local产生action是
`https://mona.do/api/v1/accounts/[id]/[action]`
action in `[block, unblock, follow, unfollow, mute, unmute]`
unfollow 有一个结构体
全都是POST

#  log

## 09/08

设计好db之后
再去做更多的设计

本地用户看到的东西

db怎么弄呃现在是

其实是不是很简单，
签名各种还要分一下是哪个用户这样的。


写到哪里了，要做本地的user返回到s2s上吧。

local访问remote user通了的

总之开始写s2s的user返回

post user
创建成功了吗
request from s2s
能被访问了吗 

1. webfinger notfound
2. post user for test
  - 两个list没有正确marshal，需要处理的
3. AS:Person
  - ~~大小写出了问题~~
  - 加type case 啥啥的好了，但是要删库重来的

win下面不分大小写可能还行，但是存储的时候要分大小写。
把privatekey pem塞到数据库里了

可以了，远程能访问了

- 创建user apiPostUser
- db保存
- webfinger 访问 user
- s2s访问user

成功了。

接下来做action？

## 09/07

~~加了user的本地api~~ done
- ~~也需要往core里面挪~~ done
- ~~还没测~~ done 

~~core缓存User~~ done
- ~~api也加一个~~ done

插入不同的acct（uniqe）之后报错
会返回error的。
TODO：但是不太懂sql会返回什么，gorm也只是一知半解。
咋办啊。

加了cache超时

remote访问user
昨晚remote访问user的互通就歇了

做了localuser
webfinger 的 server 做在哪里了

返回err还算对，webfinger.CreateWebfingerObj其实没测试。希望对的



## 08/31

server的db在另一个路径。

accept/reject的id

accpet/reject通过了。

根据to的部分分开？
权限也是根据这个分开的

再写本地用的db
大概是user和note和通知这三部分
follow 也和objects有关
~~block~~ 本地objects

objects大概要往详细里面改。
麻了

## 08/30

- ~~user的返回~~
- ~~ap的发送~~
  - ~~要弄好签名，签名直接从用户名取得pem文件这样吧。和用户系统分开来，在core的内容~~
- ~~ap的接受~~

~~看test的时间戳都知道你有多摸~~

db写了，好像没啥问题，crud见test
只有单个，要多个么，好像不用。
考虑一下user的notification和hometimeline怎么处理

做到一半发现之后core会处理不好object
~~先做db~~

db考虑人和人之间relationship做一张表？
那也是之后的事情了

~~做一下db然后今天结束了吧~~

✔ ~~验证inbox的签名是可用的。~~
但是每次fetch user，这个之后修改吧

✔ ~~之前应该做了follow，批准follow~~
还有note的送信这些。

✔ ~~接下来是做Follow UnFollow之类的~~
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
### https://moonchan.xyz/page/

mastodon entities流程
复制entities-script.js
复制document
cd mastodon/entities
node open_entity.mjs 类名
黏贴
保存退出

mastodon methods流程
复制method-script.js
复制document(注意复制到Response)
黏贴到本地controll
如果有model，复制到model
handler加入入口(等待之后写)
保存退出

### **stack**
怎么去做error的判定
读取mastodon status的时候通过core给的函数来吧。
所以可能是AP

绷不住了，又tm改一遍db，重写得了。
找不到的话去gist存一份。不然太蛋疼了。
db也要理。他妈的早知道画图。

**https://gorm.io/docs/transactions.html**
本地到ap好像没做

好，action有大病。
要大改出一个能和所有fedi通信的。不会改
先改出个和ns站通着再说。
不然没法测ap侧。

还有一个Auth。做了能做前端。
前端似乎可以做了。做吧。

### db

#### core
localuser
存一些邮箱或者privatekey之类的
localrelation
存一些fub，key是apID
localnotify
存一些announce和like


#### mastodon
user
存一些别的内容
status
存note用的


### **cases**

几乎大更了，全要重来。
~~卧槽我之前全都做通了的啊?~~

#### antenna
- [x] 注册 - 查看 localuser 表和 mastodonuser 表
  - [x] 能从activitypub看到
#### mastodon
- [ ] 发嘟 - 查看 local notes 表
- ~~[ ] 发嘟 - fedi (暂时没做)~~
  - [ ] 多个服务器。
- [ ] follow / accept / reject - 查看 local relation 表
- [ ] like / undo - 查看 local notifies 表
- [ ] announce / undo - 查看 local notifies 表
#### webfinger
- [ ] 查询user id
  - [ ] cache
#### activitypub
- [ ] 查询user
  - [ ] cache
- [ ] follow / undo - 查看 local notifies 表
- [ ] follow / accept / reject - 查看 local notifies 表
- [ ] block / undo - 查看 local notifies 表
- [ ] like / undo - 查看 local notifies 表
- [ ] announce / undo - 查看 local notifies 表
- [ ] create note - 查看 local notifies 表
~~#### actions~~



### **todo**
先把activitypub的action和controller搓出来
action先mock了。要加处理用的闭包线程的
activitypub/controller还没测试
- [x] 哦，controller改成orderedmap了还没写完了。
  - [x] 访问本地的时候没有去local user这个库。
  - [x] attribute to 没显示
  - [ ] 再跑一下createnote，然后可以抓fav和reblog了
    - [x] 流程为 create， f， accept， note， 
    - [ ] fav/ret
  - 虽然乱七八糟的但是还是正常能够fub
- [ ] gorm只update非空的键
写完之后测试：(我在写什么??)
- [ ] user是否通的。
- [ ] fub是否通的。

- [ ] 再把mastodon的api搓出来

### **points**

去submodule的readme看。

### **known issue**
- [x] 新建的user不设置icon的时候会取出image表中的任意一条。
  - 有修改未验证
  - 从postman里面看应该好了
- [x] undo报错没object
  - 是修改了getordefault的问题
- [ ] 直接取消关注会有一个reject但是object是空白的情况

- [ ] httpsig
  - [ ] 删除未知的user时会有一个null，不管可能也不要紧。
- [ ] action
- [ ] 当重复接受到follow等actions时
- [ ] 当发送action没有成功时/成功时怎么处理

### **done**
- [x] webfinger
  - [ ] also known as
- [ ] antenna
  - [x] register
- [ ] activitypub
- [ ] mastodon
  - [x] post notes

### console log

```sh
git submodule add -b back/actions -f git@github.com:Hana-ame/fedi-antenna.git actions --depth 1
git submodule add -b back/webfinger -f git@github.com:Hana-ame/fedi-antenna.git webfinger --depth 1
git submodule add -b back/core -f git@github.com:Hana-ame/fedi-antenna.git core --depth 1 
git submodule add -b back/activitypub -f git@github.com:Hana-ame/fedi-antenna.git activitypub --depth 1
git submodule add -b back/antenna -f git@github.com:Hana-ame/fedi-antenna.git antenna --depth 1 
git submodule add -b back/mastodon -f git@github.com:Hana-ame/fedi-antenna.git mastodon --depth 1 
git submodule add -b back/actions -f git@github.com:Hana-ame/fedi-antenna.git actions --depth 1


 git submodule init 
 git submodule update --init --recursive
```

```sh
git submodule foreach --recursive 'git stash save 整理db'
git submodule foreach --recursive 'git stash list'
```

why it not checkouted to a branch. there must be something wrong


## structure

not up to date.

main 
=>
`activitypub/controller`
`webfinger/controller`
进行校验的地方，保证给入的数据都是正常的
=>
`core/handler/antenna`
`core/handler/activitypub`
`core/handler/mastodon`
进行处理的地方
=>
`core/action`
涉及远端的处理
=>
`core`
=>
`activitypub/action`
`webfinger/action`
=>
`core/dao`
==============>
`core/utils`
`tools/*`
