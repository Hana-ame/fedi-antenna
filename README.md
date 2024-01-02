## actions

actions in submodule give a way directly to access the remote resources.
they are supposed to make a request, and not so reliable thus (should it leave a log? do it after the porject is done.)

## model

that provide the model of activitypub

though there is plenty of 


Createable
- note(Create)
Sendable
- Block(Unblock)
- Follow(Unfollow)
Acceptable
- Follow(Accpet)
Recectable
- Follow(Reject)
Deleteable
- ???(Delete)

Used
- User
  - PublicKey
  - Image
  - Tag
- Follow
- Block
- Note
  - Mention
  - Collection

## (root)

that provide a useful method to regist to the gin.Router
and it is the only one
