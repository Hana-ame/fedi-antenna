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