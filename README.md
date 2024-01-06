# back/activitipub

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
- [ ] accept
  - just set to accpeted
- [ ] reject
  - just delete
- [x] block
- [x] undo follow
- [x] create note

### actions

- [x] follow
  - [ ] pedding and accepted while accepted
- [ ] block

### endpoints

- [ ] followers
- [ ] following
- [ ] outbox
