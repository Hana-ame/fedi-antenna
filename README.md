# 
```sh
git submodule add -b back/webfinger -f git@github.com:Hana-ame/fedi-antenna.git webfinger --depth 1
git submodule add -b back/core -f git@github.com:Hana-ame/fedi-antenna.git core --depth 1 
git submodule add -b back/activitypub -f git@github.com:Hana-ame/fedi-antenna.git activitypub --depth 1
```

why it not checkouted to a branch. there must be something wrong


main 
=>
`activitypub/controller`
`webfinger/controller`
=>
`core/handler`
`core/action`
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

## oauth2.0


