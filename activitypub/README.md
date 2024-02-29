# back/activitipub


when delete person, what should do.

most of functions are here

## data flow
` inbound -> controller -> handler -> actions -> fetch -> outbound

` inbound -> controller -> handler -> core/dao
`                                     core

## models

users are saved as cache.
publickey are saved as cache.
image are saved as cache.

notes are saved @core/localnote
fub are saved @core/localrelation
fav/ret/mention are saved @core/localnotify

reject/undo means delete in db.

## activitypub controller

### inbox

- [x] follow
  - [x] accept
  - [x] reject
- [x] block
- [ ] announce
- [ ] like
- [ ] undo
  - [x] follow
  - [x] block
  - [ ] announce
  - [ ] like
- [x] create note
- [ ] delete
  - [ ] user
  - [ ] note?
- poll

### actions

- [x] follow
  - [x] undo
  - [ ] pedding and accepted while accepted
- [x] block
  - [x] undo
- [ ] like
  - [ ] undo
- [ ] announce
  - [ ] undo
- [x] create note
- [ ] delete
  - [ ] user
  - [ ] note?

### endpoints

- [ ] followers
- [ ] following
- [ ] outbox


### unknown


- [ ] @
  - [ ] 发送
  - [ ] 接受


## note

undo relation 时用了删除
notifycation 时用了delete标记
不一样的做法