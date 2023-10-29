# go-architecture-multiple-commands

## 概要
Goチームが紹介しているパッケージ構成のうちの一つである`multiple-commands`の実装例を紹介します([参考リンク](https://go.dev/doc/modules/layout))

## 実行方法
- 足し算の場合
```go
cd add_cmd
go run ./... 3 2
```
- 引き算の場合
```go
cd subtract_cmd
go run ./... 3 2
```
