### Schema Definition

```
Post = {
author: User
title: string,
content: []bytes,
tags: []string,
published: boolean,
created_at: date,
likes: []User.id
}
```

```
User = {
firstname: string,
lastname: string,
email: string,
password: string,
password_reset_token: string,
active: boolean,
}
```

```
Comment = {
author: User,
body: text,
upvotes: []User.id,
downvotes: []User.id
}
```
