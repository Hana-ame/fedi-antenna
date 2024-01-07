# 

### **stack**
~~很明显undo是需要把所有的东西都存一遍的。~~
mastodon做follow，unfollow，block，unblock，like(+undo)，announce(+undo)的接口
+ 打开mastodon操作一下然后copy上去
记得看一眼resp的返回是什么样子的。自己的返回也做一下。
+ 照着保存下来的写，gist上有一份的。
在mastodon加入了Account。给webfinger和联系什么的加上键，查找方式
ap侧主动的功能好像很完全了，去测ap侧被动的部分，在那之前去把mastodon写好
core的逻辑可能需要适当修改一下。
本质上是action需要cache，接到core上。
local的查询在dao上。
功能做完。
把utils挪一下，使得各个模块能尽可能自主运作。

### **cases**

#### antenna
- [x] 注册 - 查看 localuser 表
#### mastodon
- [x] 发嘟 - 查看 local notes 表
- [x] 发嘟 - fedi
  - [ ] 多个服务器。
- [ ] follow / accept / reject - 查看 local notifies 表
- [ ] like / undo - 查看 local notifies 表
- [ ] announce / undo - 查看 local notifies 表
#### webfinger
- [x] 查询user id
#### activitypub
- [x] 查询user
- [x] follow / undo - 查看 local notifies 表
- [ ] follow / accept / reject - 查看 local notifies 表
- [x] block / undo - 查看 local notifies 表
- [x] like / undo - 查看 local notifies 表
- [x] announce / undo - 查看 local notifies 表
- [ ] create note - 查看 local notifies 表
#### actions



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
git submodule add -b back/webfinger -f git@github.com:Hana-ame/fedi-antenna.git webfinger --depth 1
git submodule add -b back/core -f git@github.com:Hana-ame/fedi-antenna.git core --depth 1 
git submodule add -b back/activitypub -f git@github.com:Hana-ame/fedi-antenna.git activitypub --depth 1
git submodule add -b back/antenna -f git@github.com:Hana-ame/fedi-antenna.git antenna --depth 1 
git submodule add -b back/mastodon -f git@github.com:Hana-ame/fedi-antenna.git mastodon --depth 1 
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
