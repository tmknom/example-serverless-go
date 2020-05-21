# example-serverless-go

Serverless Frameworkのサンプルコード。

## Requirements

- [Serverless Framework](https://www.serverless.com/)

## ビルド

依存ライブラリのダウンロードとバイナリのビルドを実行します。

```shell
make build
```

## デプロイ

バイナリをビルドしてLambdaへデプロイします。

```shell
make deploy
```

## ローカル実行

ローカル環境で実行します。

```shell
make invoke-local
```

## リモート実行

デプロイ済みのLambdaを実行します。

```shell
make invoke-remote
```

## Makefile

```text
build                          Build the application
clean                          Clean the binary
deploy                         Deploy a Serverless service
deps                           Install dependencies
generate-event                 Generate event
help                           Show help
info                           Display information about the service
invoke-local                   Invoke function locally
invoke-remote                  Invoke function remotely
logs                           Output the logs of a deployed function
remove                         Remove Serverless service and all resources
```

## License

Apache 2 Licensed. See LICENSE for full details.
