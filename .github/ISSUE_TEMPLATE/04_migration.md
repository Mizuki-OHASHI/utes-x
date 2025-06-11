---
name:  DBマイグレーション
about: 'Step4 DBマイグレーション'
title: '[Step 4] DBマイグレーション'
labels: curriculum
assignees: ''

---

## 概要

- DBマイグレーションは、データベースのスキーマを変更するための手順です。
- 新たに LIKE 機能を実装するために、データベースのテーブルを追加します。
- マイグレーションには `golang-migrate` を使用します。

## 手順

複数の環境 (開発環境、ステージング環境、本番環境 など) で同じ SQL クエリを実行して、データベースのスキーマを変更するために、マイグレーションツールを使用します。

### 1. マイグレーションファイルの作成

`db/` ディレクトリに移動し、マイグレーションファイルを作成します。

マイグレーションファイルは、`db/migrations` ディレクトリに配置します。
まず、マイグレーションファイルを作成します。
ここでは `Makefile` にマイグレーションファイルを作成するためのターゲットが定義されています。

以下のように、環境変数 `NAME` としてマイグレーション名を指定して、`make migrate-create` コマンドを実行します。

```bash
NAME=add_likes_table make migrate-create
```

これにより、`db/migrations/` ディレクトリにマイグレーションファイルが作成されます。

- [ ] マイグレーションファイルを作成した。

### 2. マイグレーションファイルの編集

マイグレーションファイルを編集して、LIKE 機能に必要なテーブルを追加します。
先ほどのコマンドで

- `***_add_likes_table.up.sql`
- `***_add_likes_table.down.sql`

のようなファイルが作成されます。
`up.sql` ファイルには、テーブルを追加する SQL クエリを記述します。
`down.sql` ファイルには、テーブルを削除する SQL クエリを記述します。
`up.sql` を実行した後に、ダウングレードしたくなった場合は `down.sql` を実行することで、元の状態に戻すことができ流ようにクエリを記述します。

今回は新たに `posts` と `users` テーブルを結びつける `likes` テーブルを追加します。

前ステップで確認した OpenAPI の仕様を考慮すると、`likes` テーブルは以下のような構造になるでしょう。

| カラム名 | データ型 | 説明 | NULL制約 |
| --- | --- | --- | --- |
| id | VARCAHAR(32) | LIKE の ID (ULID 形式), プライマリーキー | NOT NULL |
| post_id | VARCHAR(32) | 投稿の ID (ULID 形式), 外部キー | NOT NULL |
| user_id | VARCHAR(32) | ユーザーの ID (ULID 形式), 外部キー | NOT NULL |
| created_at | TIMESTAMP | LIKE の作成日時 | NOT NULL |
| updated_at | TIMESTAMP | LIKE の更新日時 | NULL |

これの構成のテーブルを作成する SQL クエリを一旦考えてみてください。
下のトグルを開くと、SQL クエリの例が表示されます。

- [ ] `***.up.sql` ファイルにテーブルを追加する SQL クエリを記述してください。

<details>
<summary>SQL クエリの例</summary>

```sql
BEGIN;

CREATE TABLE likes (
    id VARCHAR(32) PRIMARY KEY,
    post_id VARCHAR(32) NOT NULL,
    user_id VARCHAR(32) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (post_id) REFERENCES posts(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

COMMIT;
```

</details>

続いて、`***.down.sql` ファイルには、テーブルを削除する SQL クエリを記述します。

- [ ] `***.down.sql` ファイルにテーブルを削除する SQL クエリを記述してください。

<details>
<summary>SQL クエリの例</summary>

```sql
BEGIN;

DROP TABLE IF EXISTS likes;

COMMIT;
```

</details>
