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
