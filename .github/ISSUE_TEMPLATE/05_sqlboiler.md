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
  sqlboiler.PostWhere.UserID.EQ("01JXCX2QSSEK1Q4M4HWYCAYM7V"),
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
  sqlboiler.PostWhere.ID.EQ("01JXCX2QSSEK1Q4M4HWYCAYM7V"),
  qm.Load(sqlboiler.PostRels.Replies),
).One(ctx, p.db)
```

撮ってきた DTO (Data Transfer Object) から、リレーションを取得するには、以下のようにします。

```go
// dao/posts.conv.go
repliesDto := postDto.R.GetReplies()
```

なお、JOIN される側のテーブルのカラムについての WHERE 条件を指定する場合には明示的に JOIN 句を指定する必要があるので注意して下さい。

### 2. Model の実装

- [ ] まずは、`model/posts.go` に LIKE 機能に関する構造体を定義します。

```go
// ...
type PostLike struct {
  ID model.ID
  PostID model.ID
  UserID model.ID
  CreatedAt time.Time
  UpdatedAt *time.Time
  User *User
}
// ...
```

- [ ] さらに、既存の `Post` 構造体に、`Likes` フォールドを追加します。

```go
// ...
type Post struct {
  // ...
  Likes []PostLike
}
// ...
```

### 3. DAO の実装

SQLBoiler を使って、DAO (Data Access Object) を実装します。DAO は、データベース操作を抽象化し、ビジネスロジックからデータベースの詳細を隠蔽します。

#### LIKE の作成

まずは、LIKE の作成の D を実装します。LIKE 機能は投稿に関係するものですから、今回は既存のファイル `dao/posts.go` に追加します。

- [ ] はじめに、DAO のインターフェースに LIKE 機能を追加します。

```go
type Post interface {
	GetMany(ctx context.Context, query GetManyQuery) ([]model.Post, error)
	Create(ctx context.Context, post model.Post) (*model.Post, error)
	CreateReply(ctx context.Context, replyTo model.ID, userID model.ID, reply model.Post) (*model.Reply, error)
	GetWithReplies(ctx context.Context, postID model.ID) (*model.PostWithReplies, error)
  CreateLike(ctx context.Context, like model.PostLike) (*model.PostLike, error) // <-- 追加
}
```

この状態では、インターフェースだけが定義され、中身はまだ実装されていません。そのため、静的解析ではエラーが発生していると思います。以下で、実装を進めていきます。

- [ ]  まず、`CreateLike` として、関数を追加します。

```go
// ...
func (p *Post) CreateLike(ctx context.Context, like model.PostLike) (*model.PostLike, error) {
  // LIKE を作成する処理を実装します。
  // SQLBoiler を使って、likes テーブルに新しいレコードを挿入します。
}
```

これを実装した時点で、静的解析のエラーは解消されるはずです。

次に、既存の `GetMany` および `GetWIthReplies` 関数に、LIKE の情報を取得する処理を追加します。
これにあたって、まずは変換関数を修正しましょう。

- [ ] `dao/posts.conv.go` について、以下の修正を行なってください。

```go
// 1. LIKE モデルへ変換する関数を追加する
func toPostLikeModel(likeDto sqlboiler.Like) (*model.PostLike, error) {
  // PostLike 構造体に変換する処理を実装します。
  // User フィールドについては `users.conv.go` の `toUserModel` 関数を使用してください。
}

// 2. LIKE モデルのスライスを変換する関数を追加する
func toPostLikeModels(likesDto sqlboiler.LikeSlice) ([]*model.PostLike, error) {
  // PostLike のスライスに変換する処理を実装します。
  // 各 LIKE DTO を toPostLikeModel 関数を使って変換してください。
}

// 3. toPostModel 関数を修正する
func toPostModel(postDto sqlboiler.Post) (*model.Post, error) {
  // PostLike を取得する処理を追加します。
}
```

変換関数を変更しただけだと、`postDto.R.GetLikes()` はロードされず、`nil` になってしまいます。適切に LIKE の情報を取得するために、Eager Loading を使いましょう。ここでは、Likes を取得するだけでなく、それに関連するユーザー情報も読み込む必要があることに注意してください。

- [ ] 変換関数が実装できたら、Eager Loading を使って、LIKE の情報を取得するように `GetMany` および `GetWithReplies` 関数を修正します。
