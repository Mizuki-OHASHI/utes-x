---
name: 環境構築
about: 'Step1 環境構築'
title: '[Step 1] 環境構築'
labels: curriculum
assignees: ''

---

## 生成AIの利用について

開発において生成AIを活用することは非常に有効です。特に、コードの自動生成やドキュメント作成など、繰り返し行う作業の効率化に大きなメリットがあります。筆者自身も普段から GitHub Copilot などの恩恵を受けています。

ただし、生成AIを利用する際にはいくつか重要な点に注意が必要です。最終的なコードの責任は、あくまでそれを採用した本人にあります。AIが提案するコードは常に正しいとは限らず、セキュリティ上の脆弱性やバグ、不適切な実装が含まれている可能性もあります。生成AIが書いたコードを理解しないまま使用することは非常に危険であり、無責任な行為であることを十分に自覚してください。

本研修においても、GitHub Copilot などによる補完 (関数名の候補提示や定型文の補助など) やデバッグなどの利用は問題ありませんが、一つの関数や処理全体を生成させて、**理解せずにそのまま採用することは控えてください**。理解を伴わずにコードを取り込むことで、学習の機会を失うだけでなく、後々のトラブルの原因にもなりかねません。

- [ ] 上記の注意点を理解した。

## 概要

一般的には、そのレポジトリで必要な環境設定を行う手順はルートの `README.md` に記載せれていることが多いです。本レポジトリでも同様です。

## 手順

### 1. パッケージのインストール

`/db` に移動します。
`README.md` の通り、以下のパッケージをインストールしてください。

- [ ] Docker
- [ ] Docker Compose
- [ ] Go (>= 1.23)
- [ ] make
- [ ] golang-migrate
- [ ] sqlboiler
- [ ] mockgen
- [ ] oapi-codegen

#### Note

- 研修では上の3つ (Docker, Docker Compose, Go) は完了しているはずです。
- Mac OS の場合、可能なものについては `homebrew` を使うと簡単にインストールできます。
- Mac OS の場合、`make` コマンドは標準でサポートされています。

### 2. 動作確認

`/api` に移動します。
`README.md` の通り、以下のコマンドを実行して、API サーバーが起動することを確認してください。

```bash
go run main.go
```

以下のようなレポジトリが出力されれば成功です。
- [ ] サーバーが起動することを確認した。

```bash
$ go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /posts                    --> utes-x-api/controller.(*ServerInterfaceWrapper).PostPosts-fm (4 handlers)
[GIN-debug] GET    /posts/:post_id           --> utes-x-api/controller.(*ServerInterfaceWrapper).GetPostsPostId-fm (4 handlers)
[GIN-debug] POST   /replies                  --> utes-x-api/controller.(*ServerInterfaceWrapper).PostReplies-fm (4 handlers)
[GIN-debug] GET    /users                    --> utes-x-api/controller.(*ServerInterfaceWrapper).GetUsers-fm (4 handlers)
[GIN-debug] POST   /users                    --> utes-x-api/controller.(*ServerInterfaceWrapper).PostUsers-fm (4 handlers)
[GIN-debug] GET    /users/:user_id/posts     --> utes-x-api/controller.(*ServerInterfaceWrapper).GetUsersUserIdPosts-fm (4 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8888
```
