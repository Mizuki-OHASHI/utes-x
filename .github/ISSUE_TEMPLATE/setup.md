---
name: 環境構築
about: 'Step1 環境構築'
title: '[Step 1] 環境構築'
labels: curriculum
assignees: ''

---

## 概要

一般的には、そのレポジトリで必要な環境設定を行う手順はルートの `README.md` に記載せれていることが多いです。本レポジトリでも同様です。

## 手順

### 1. パッケージのインストール

`/db` に移動します。
`README.md` の通り、以下のパッケージをインストールしてください。

- [ ] Docker
- [ ] Docker Compose
- [ ] Go (>= 1.23)
- [ ] golang-migrate
- [ ] sqlboiler
- [ ] mockgen
- [ ] oapi-codegen

#### Note

- 研修では上の3つ (Docker, Docker Compose, Go) は完了しているはずです。
- Mac OS の場合、可能なものについては `homebrew` を使うと簡単にインストールできます。

### 2. 動作確認

`/api` に移動します。
`README.md` の通り、以下のコマンドを実行して、API サーバーが起動することを確認してください。

```bash
go run main.go
```
