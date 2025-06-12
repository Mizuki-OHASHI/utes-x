---
name: 構成の理解
about: 'Step2 構成の理解'
title: '[Step 2] 構成の理解'
labels: curriculum
assignees: ''

---

## 概要

環境構築が完了したら、次はディレクトリ構成や技術スタックを理解しましょう。

## 手順

### 1. ディレクトリ構成の確認

バックエンドの本体は `/api` ディレクトリに配置されています。

```bash
api
├── controller/ # コントローラー (HTTP のリクエストを受け取り、レスポンスを返す)
├── dao/ # データアクセスオブジェクト (DAO) (データベースとのやり取りを行う)
│   └── mock/ # モック (テスト用のモック実装)
├── model/ # モデル (データの構造を定義)
├── sqlboiler/ # SQLBoiler による ORM の設定
│   ├── entity/ # 自動生成されたエンティティ
│   └── templates/ # SQLBoiler のテンプレート
└── usecase/ # ユースケース (ビジネスロジックを実装)
```

研修プログラムインターン編の [「ソースコードの改善」の章](https://utokyotechclub.gitbook.io/curriculum/intern/ssukdonogo) の構成に従っています。

- [ ] 研修プログラムインターン編の「ソースコードの改善」の章を確認した。

<img src="https://utokyotechclub.gitbook.io/~gitbook/image?url=https%3A%2F%2F686112546-files.gitbook.io%2F%7E%2Ffiles%2Fv0%2Fb%2Fgitbook-x-prod.appspot.com%2Fo%2Fspaces%252F15Gri4KP0qK959o0eDJi%252Fuploads%252FhhgWZL2BuS52gbCvcZQ3%252F%25E3%2582%25B9%25E3%2582%25AF%25E3%2583%25AA%25E3%2583%25BC%25E3%2583%25B3%25E3%2582%25B7%25E3%2583%25A7%25E3%2583%2583%25E3%2583%2588%25202023-03-24%252020.50.49.png%3Falt%3Dmedia%26token%3D6cfa7814-3317-41e5-99f6-c5f5bfdaa7d6&width=400&dpr=2&quality=100&sign=6981a6ef&sv=2" alt="ディレクトリ構成" width="600">


### 2. 依存性注入

<!-- TODO: 依存性注入についての説明と、実際に Go でどのように実装するのかを記述する -->

依存性注入は、ソフトウェアの設計パターンの一つで、オブジェクトの依存関係を外部から注入することで、コードの柔軟性とテスト容易性を向上させる手法です。
依存性注入を使用することで、コンポーネント間の結合度を下げ、テストやメンテナンスが容易になります。
依存性注入は、以下のようなメリットがあります。
- **テスト容易性**: モックやスタブを使用して、依存関係を簡単に置き換えることができるため、ユニットテストが容易になります。
- **柔軟性**: 依存関係を外部から注入することで、コンポーネントの実装を変更しても、他のコンポーネントに影響を与えにくくなります。
- **再利用性**: 依存関係を外部から注入することで、同じコンポーネントを異なるコンテキストで再利用しやすくなります。
- **可読性**: 依存関係が明示的になるため、コードの可読性が向上します。

依存性注入は、以下のような方法で実装されます。

```go
// 実装したいメソッドたちのインターフェース (入力と出力の定義)
type CustomInterface interface {
    Method1() string
    Method2() int
}

// 実装したいメソッドたちで使用する他のインターフェース
type customStruct struct {
    dependency OtherInterface
}

// 実装したいメソッドたちのコンストラクタ
func NewCustomStruct(dependency OtherInterface) *customStruct {
    return &CustomStruct{dependency: dependency}
}

// 実際のメソッドの実装
func (c *customStruct) Method1() string {
    return c.dependency.SomeMethod()
}

func (c *customStruct) Method2() int {
    return c.dependency.AnotherMethod()
}
```

この例では、`CustomInterface` が実装したいメソッドのインターフェースを定義し、`customStruct` がその実装を提供しています。このメソッド群で使用したい別のインターフェース (依存関係) はコンストラクタで注入され、`NewCustomStruct` 関数を通じて外部から提供されます。

```
CustomInterface <--依存-- OtherInterface
```

こうすることで、`OtherInterface` 部分をモックに置き換えて、対象となる `CustomInterface` のみを単独でテストすることが可能になります。

### 3. 技術スタックの確認

上記の構成において、各レイヤーで使用されている技術スタックを確認します。詳細な使い方は、次以降のステップで学んでいきます。

#### Controller

コントローラーは、HTTP リクエストを受け取り、レスポンスを返す役割を担います。

- `Gin`: Go 用の HTTP フレームワークで、ルーティングやミドルウェアの管理を行います。
- `oapi-codegen`: OpenAPI 仕様からコードを生成するツールで、API のエンドポイントを定義します。


#### DAO

データアクセスオブジェクト (DAO) は、データベースとのやり取りを行います。

- `sqlboiler`: Go 用の ORM (Object-Relational Mapping) ツールで、データベースのテーブルと Go の構造体をマッピングします。実際にあるデータベースに対して、それに対応する Go の構造体やメソッドを自動生成します。
- `mockgen`: Go のモック生成ツールで、DAO のインターフェースをモック化してテストを容易にします。

### 3. 本カリキュラムでやること

本カリキュラムでは、新たに投稿に対する LIKE 機能 (Create/Read) を実装します。
投稿に対する LIKE 機能を実装するために、以下のステップを進めます。

- [Step 3] OpenAPI
  - OpenAPI は API の仕様を YAML/JSON で定義するフォーマットです。
  - こちらのファイルを見て、実装すべきエンドポイントの確認をします。
  - Open API の仕様に基づいて、エンドポイントに使用する構造体やメソッドを自動生成します。
- [Step 4] DBマイグレーション
  - DBマイグレーションは、データベースのスキーマを変更するための手順です。
  - 新たに LIKE 機能を実装するために、データベースのテーブルを追加します。
  - マイグレーションには `golang-migrate` を使用します。
- [Step 5] SQLBoiler
  - 前のステップでマイグレーションしたデータベースをもとに、SQLBoiler の自動生成を実行します。
  - SQLBoiler を使って、データベースのテーブルと Go の構造体をマッピングします。
  - これにより、SQL クエリを直接書くことなく、Go のコードでデータベース操作を行うことができます。
  - 実装した LIKE 機能に必要な DAO のメソッドを追加します。
- [Step 6] ユースケース
  - ユースケースは、ビジネスロジックを実装するレイヤーです。
  - 投稿に対する LIKE 機能のビジネスロジックを実装します。
- [Step 7] コントローラー
  - コントローラーは、HTTP リクエストを受け取り、レスポンスを返す役割を担います。
  - ユースケースで実装したビジネスロジックを呼び出し、HTTP レスポンスを返すコントローラーのメソッドを実装します。
- [Step 8] 単体テスト (後日追加予定)
  -  ここではロジックが複雑になりがちなユースケースレイヤーの単体テストを実装します。

### 4. 準備

最後に、次のステップに進む前に、データベースにダミーデータを投入しておきましょう。

- [ ] `api/script/seed_loader.go` に移動して、`seed_loader.go` を実行します。
- [ ] 実行後に、`curl 'http://localhost:8888/users'` などの API を叩いてみて、データが正しく投入されていることを確認します。DB に接続してみるのも良いでしょう。
