# core


## models

- local_user
  - activitypub.user
- local_relation
  - activitypub.follow
  - activitypub.block
- local_notify
  - activitypub.announce // 转发也在这里
  - activitypub.like
- local_note
  - activitypub.note
  - mastodon.status


### LocalRelation

请求，接受
(null) -> follow, pendding -> follow, accepted
请求，拒绝 / unfollow
(null) -> follow, pendding -> (null)
**注：请在action request return ok之后再设置自己家的表。**
block
(null)/follow, pendding/follow, accepted -> block, blocked
unblock
block, blocked -> null
**注：请在action request return ok之后再设置自己家的表。**
查询relation时需要查询。


### LocalNote
- activitypub.note
- mastodon.status
都需要转换。
公用Field.
AP Field.
mastodon Field.
