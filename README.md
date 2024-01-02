# core

this works as the middle layer for the project

## webfinger
```
webfinger
  core.IsAccountExist()
    dao.ReadAccount()
```
like that

### core/actions: handle local attemptions.
                                     => core                 -> 
local/controller => **core/actions** => activitypub/actions  -> core/dao
                                     => webfinger/actions    -> 

### core/dao: handle all data visits
core/actions -> core -> activitypub/actions => **core/dao**

### core/handler
activitypub/controller => **core/handler**
webfinger/controller   =>                    

### core

caution:

core 是十分接近底层的地方


temparory settings

actions and handlers  are accurately the same layer. just remain it here for less job.

core are much more close to datas, retrun a rather relaieble data interface 

handlers / actions provide the abstruct methods fot handling the data recived from controllers.


如果可以试着把core中的大部分文件夹当作动作名词。会怎么样。

handle 用作接受并且返回
action 用作