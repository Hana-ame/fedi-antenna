# 再来一次
- mastodon 只和mastodon状态有关
  - dao 修改数据库的动作
  - handler 框架无关的处理
  - entities
- activitypub 
  - dao 修改数据库的动作
  - s2s 是tasks管理
  - handler 框架无关的处理
- webfinger
  - **todo**

先只做post吧。再也不想对着代码整了，前端做出来好不好。

似乎除了user其他ap都不需要保存的样子。

传到台式机的wsl里面  
跑go和nextjs  
记更新一下，go版本不对。  
然后在台式机wsl里面打开debug进行调试。  
如果要和nextjs一起的话  
记得装一下nginx  

todo
- [ ] 订阅relay，查询日本人的项目，想不起名字了，里面有example
   记得保存自己订阅了什么的结构体。
- [ ] 接受ap上的note，和user。
   看了一眼，如果是没遇到过的user的话是需要fetch一下的。
- [ ] 发送note
   主要是s2s的部分，差不多要重构
   记得留task结构体
   obj可能重做。
   队列以及如何执行。