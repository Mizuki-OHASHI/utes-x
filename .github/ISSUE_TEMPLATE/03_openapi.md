---
name: OpenAPI
about: 'Step3 OpenAPI'
title: '[Step 3] OpenAPI'
labels: curriculum
assignees: ''

---

## 概要

- OpenAPI は API の仕様を YAML/JSON で定義するフォーマットです。
- こちらのファイルを見て、実装すべきエンドポイントの確認をします。
- Open API の仕様に基づいて、エンドポイントに使用する構造体やメソッドを自動生成します。

## 手順

### 1. OpenAPI で仕様を確認する

OpenAPI の仕様は、`api/controller/openapi.yaml` に定義されています。

[こちらのコミット](https://github.com/Mizuki-OHASHI/utes-x/commit/e1db0be6f234c3c545ce8c3026215cda434f719f) を見て下さい。

<details>
<summary>
  OpenAPI の仕様
</summary>

`open-api/open-api.yaml`

```yaml
/posts/{post_id}/like:
  post:
    tags:
      - posts
    summary: Like a post
    parameters:
      - in: path
        name: post_id
        required: true
        schema:
          type: string
    requestBody:
      required: true
      content:
        application/json:
          schema:
            type: object
            properties:
              user_id:
                type: string
                format: ulid
            required:
              - user_id
    responses:
      '200':
        description: Like added
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Post'
```

</details>

\
こちらの YAML ファイルを見ると、`POST /posts/{post_id}/like` というエンドポイントが定義されていることがわかります。`{post_id}` はパスパラメータで、任意の値 (POST ID) が入ることを意味しています。

さらに、リクエストのボディでは
```json
{
  "user_id": "string"
}
```
のような JSON 形式で `user_id` を送信することが求められています。

成功した場合のレスポンスは、`200 OK` で、`Post` オブジェクトが返されることがわかります。`POST` オブジェクトは同ファイル内の別の場所 (`components/schemas/Post`) で定義されています。

なお、OpenAPI の仕様は YAML 形式で記述されていますが、こちらを見やすく表示できるサービスもあります。ぜひ活用してみてください。

例えば VS Code の拡張機能である [OpenAPI (Swagger) Editor](https://marketplace.visualstudio.com/items?itemName=42Crunch.vscode-openapi) を使うと、YAML ファイルを開くだけで、視覚的に確認できます。

<image src="../../assets/swagger-ui-example.png" alt="Swagger UI Example" width="600">

### 2. OpenAPI から構造体とメソッドを自動生成する

OpenAPI の仕様に基づいて、Go の構造体やメソッドを自動生成します。
以下のコマンドを実行して、`api/controller/openapi.yaml` から Go のコードを生成します。

```bash
make gen-open-api
```

make コマンドは `Makefile` に定義されたタスクを実行します。`gen-open-api` の実行内容は、`oapi-codegen` ツールを使って OpenAPI の仕様から Go のコードを生成するものです。

```makefile
# Makefile
gen-open-api:
	cd controller && \
	oapi-codegen -generate types,gin,spec,skip-prune -package controller -o ./schema.gen.go ./../../open-api/open-api.yaml
```

- [ ] `make gen-open-api` を実行して、OpenAPI からコードを生成し、`api/controller/schema.gen.go` ファイルが編集されていることを確認してください。
