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

研修プログラムの [「ソースコードの改善」の章](https://utokyotechclub.gitbook.io/curriculum/intern/ssukdonogo) の構成に従っています。

<img src="../../assets/image.avif" alt="ディレクトリ構成" width="600">

### 2. 技術スタックの確認

上記の構成において、各レイヤーで使用されている技術スタックを確認します。詳細な使い方は、次以降のステップで学んでいきます。

#### Controller

コントローラーは、HTTP リクエストを受け取り、レスポンスを返す役割を担います。

- `Gin`: Go 用の HTTP フレームワークで、ルーティングやミドルウェアの管理を行います。
- `oapi-codegen`: OpenAPI 仕様からコードを生成するツールで、API のエンドポイントを定義します。


#### DAO

データアクセスオブジェクト (DAO) は、データベースとのやり取りを行います。

- `sqlboiler`: Go 用の ORM (Object-Relational Mapping) ツールで、データベースのテーブルと Go の構造体をマッピングします。実際にあるデータベースに対して、それに対応する Go の構造体やメソッドを自動生成します。
- `mockgen`: Go のモック生成ツールで、DAO のインターフェースをモック化してテストを容易にします。
