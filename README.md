# fedi-antenna

心血来潮的能够加入Fediverse的单人实例

webfinger

activitypub

(moonchan-api)

mastodon

都是企图自治的

都是

main

controller

core(optional, 需要调用其他数据的时候) 

dao

目录seemslike

main.go
```

webfinger
  controller
  dao
activitypub
  controller
  dao
core
  dao
  helperfunction
```

有空把其他的文件都整合一下，现在不看。
## params

|variable||
|---|---|
|user| 就是用户名|
|user.json| activityStream，其他实例会来获取的东西|
|privateKey| 自己存好|
|authorization| 大概放在cookie里面校验用|

## interface
|path||
|---|---|
|/users/:user|user.json|
|/users/:user/inbox|inbox，没啥好讲的|
|/users/:user/outbox|接受post，校验使用cookie?|
|/users/:user/followers||
|/users/:user/following||
|/users/:user/collection/featured||
|/inbox|shared inbox|
|.well-known/webfinger?resource=acct::user@:host|webfinger|

## 访问别人

是不是还没写了。