---
name: SQLBoiler
about: 'Step5 SQLBoiler'
title: '[Step 5] SQLBoiler'
labels: curriculum
assignees: ''

---

## 概要

SQLBoiler を使って、データベースのテーブルと Go の構造体をマッピングする方法を学びます。

適宜 `db/migrations` ディレクトリにある SQL クエリを参照したり、DBに接続してテーブル構造を確認したりしながら進めてください。

## 手順

### 1. SQLBoiler の理解

SQLBoiler は、Go 用の ORM (Object-Relational Mapping) ツールで、データベースのテーブルと Go の構造体をマッピングします。これにより、SQL クエリを直接書くことなく、Go のコードでデータベース操作を行うことができます。

例えば投稿から、投稿主の投稿を取得する場合、以下のような SQL クエリを実行します。

```sql
SELECT * FROM posts WHERE user_id = '01JXCX2QSSEK1Q4M4HWYCAYM7V';
```

これは、 SQLBoiler を使うと、以下のように記述できます。

```go
// dao/posts.go
posts, err := sqlboiler.Posts(
  sqlboiler.PostWhere.UserID.EQ(query.UserID.String()),
).All(ctx, p.db)
```

これだけではまだあまり便利さがわからないかもしれません。
SQLBoiler では、外部キーによって関連付けられたテーブルのデータも簡単に取得できます。

例えば、投稿に紐づく返信を取得する場合、SQL では以下のようになるかと思います。

```sql
SeLECT posts.*, replies.* FROM posts
JOIN replies ON posts.id = replies.post_id WHERE posts.id = '01JXCX2QSSEK1Q4M4HWYCAYM7V';
```

SQLBoiler では、明示的に JOIN しなくても、eager loading を使うことで、以下のように簡単に取得できます。

```go
// dao/posts.go
postDto, err := sqlboiler.Posts(
  sqlboiler.PostWhere.ID.EQ(postID.String()),
  qm.Load(sqlboiler.PostRels.Replies),
).One(ctx, p.db)
```

撮ってきた DTO (Data Transfer Object) から、リレーションを取得するには、以下のようにします。

```go
// dao/posts.conv.go
repliesDto := postDto.R.GetReplies()
```

なお、JOIN される側のテーブルのカラムについての WHERE 条件を指定する場合には明示的に JOIN 句を指定する必要があるので注意して下さい。

