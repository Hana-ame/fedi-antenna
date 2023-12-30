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