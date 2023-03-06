# fedi-antenna

心血来潮的能够加入 Fediverse 的单人实例

## params

| variable      |                                        |
| ------------- | -------------------------------------- |
| user          | 就是用户名                             |
| user.json     | activityStream，其他实例会来获取的东西 |
| privateKey    | 自己存好                               |
| authorization | 大概放在 cookie 里面校验用             |

## interface

| path                                            |                             |
| ----------------------------------------------- | --------------------------- |
| /users/:user                                    | user.json                   |
| /users/:user/inbox                              | inbox，没啥好讲的           |
| /users/:user/outbox                             | 接受 post，校验使用 cookie? |
| /users/:user/followers                          |                             |
| /users/:user/following                          |                             |
| /users/:user/collection/featured                |                             |
| /inbox                                          | shared inbox                |
| .well-known/webfinger?resource=acct::user@:host | webfinger                   |

## 访问别人

是不是还没写了。

```json
{
  "@context": [
    "https://www.w3.org/ns/activitystreams",
    "https://w3id.org/security/v1",
    {
      "manuallyApprovesFollowers": "as:manuallyApprovesFollowers",
      "sensitive": "as:sensitive",
      "Hashtag": "as:Hashtag",
      "quoteUrl": "as:quoteUrl",
      "toot": "http://joinmastodon.org/ns#",
      "Emoji": "toot:Emoji",
      "featured": "toot:featured",
      "discoverable": "toot:discoverable",
      "schema": "http://schema.org#",
      "PropertyValue": "schema:PropertyValue",
      "value": "schema:value",
      "misskey": "https://misskey-hub.net/ns#",
      "_misskey_content": "misskey:_misskey_content",
      "_misskey_quote": "misskey:_misskey_quote",
      "_misskey_reaction": "misskey:_misskey_reaction",
      "_misskey_votes": "misskey:_misskey_votes",
      "isCat": "misskey:isCat",
      "vcard": "http://www.w3.org/2006/vcard/ns#"
    }
  ],
  "id": "https://p1.a9z.dev/follows/9a3qtdtypj/9bzteda35r",
  "type": "Follow",
  "actor": "https://p1.a9z.dev/users/9a3qtdtypj",
  "object": "https://v.meromeromeiro.top/users/meromero"
}
```

```json
{
  "@context": "https://www.w3.org/ns/activitystreams",
  "id": "https://mstdn.jp/9a7ddd0b-4e52-4486-98f0-11b8b51295fe",
  "type": "Follow",
  "actor": "https://mstdn.jp/users/meromero",
  "object": "https://v.meromeromeiro.top/users/test"
}
```

```json
{
  "@context": "https://www.w3.org/ns/activitystreams",
  "id": "https://mstdn.jp/users/seirenchinatu#delete",
  "type": "Delete",
  "actor": "https://mstdn.jp/users/seirenchinatu",
  "to": ["https://www.w3.org/ns/activitystreams#Public"],
  "object": "https://mstdn.jp/users/seirenchinatu",
  "signature": {
    "type": "RsaSignature2017",
    "creator": "https://mstdn.jp/users/seirenchinatu#main-key",
    "created": "2023-03-06T03:38:30Z",
    "signatureValue": "fT6t+pP025kpohxJoygcN/yyMtPYqvfB11hj1XqBgmFeGgfw1xhRFPPRENUa+mCQMlgqFWykDrCD647yrxpavuEmVNA397xs90u5JsXW9GLqJ3r7vbiVIwsZE60cWO8LyFxjzKUwpsOzGcn4oYQhg19c15C/vEsTG1xmArNtaGcnmhJBJw4RPHJek3uPZ9QqoaKzA1TDia4qK6WT5th9TnYGYD4VnoqitTOOPoP9nkwHkOMAEZNSGZWxSdUv3YobQixOL1+8r5G7v4g769kC5CKT1v0lB0wBaahxkO5dEmx+6kifUQ61HjTyqkvxeNRQLTTpcnlP4jottUc4p7SC6w=="
  }
}
```

# NEXT

处理 httpsig
自动回 fo
处理信息推送。

# MORE

注销

# 坑

full path form Ctx

```go
c.Request().URI().String()
```

app.Use is always first used.
