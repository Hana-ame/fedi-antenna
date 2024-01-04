# 

**todo**
- 要干啥来着
- [x] pretty还没安装好，可能是死了
- [x] note的接受和发送
  - [x] 发送
  - [ ] 接受
  - [ ] 删除
  - [ ] 编辑
- [ ] FUB
  - [ ] 估计要改db
- [ ] 喜欢
  - [ ] 发送
  - [ ] 接受
- [ ] 转发
  - [ ] 发送
  - [ ] 接受
- [ ] @
  - [ ] 发送
  - [ ] 接受
- [ ] 注册的mastodon实现(这个是不是不在api里啊)

- [ ] 挪
  - [ ] publickey可以往core挪
  - [ ] user已经往core挪了



```sh
git submodule add -b back/webfinger -f git@github.com:Hana-ame/fedi-antenna.git webfinger --depth 1
git submodule add -b back/core -f git@github.com:Hana-ame/fedi-antenna.git core --depth 1 
git submodule add -b back/activitypub -f git@github.com:Hana-ame/fedi-antenna.git activitypub --depth 1
git submodule add -b back/antenna -f git@github.com:Hana-ame/fedi-antenna.git antenna --depth 1 
git submodule add -b back/mastodon -f git@github.com:Hana-ame/fedi-antenna.git mastodon --depth 1 
```

why it not checkouted to a branch. there must be something wrong


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
