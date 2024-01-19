#

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

本地到ap好像没做

好，action有大病。
要大改出一个能和所有fedi通信的。不会改
先改出个和ns站通着再说。
不然没法测ap侧。

还有一个Auth。做了能做前端。
前端似乎可以做了。做吧。

### **cases**

几乎大更了，全要重来。

#### antenna
- [ ] 注册 - 查看 localuser 表
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
