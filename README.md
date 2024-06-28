# Getting Started

```shell:
$ cp .env.example .env
$ make init
$ make migrate
$ make start
```

1. repositoryでインターフェースを実装
2. 具体的なdbとのやりとりや実装をinfra層へ
3. usecaseでリポジトリで定義したものを呼び出す
4. 最終的にresolverで組み合わせて使う
5. graphQL使用せずに個別ルーティングの場合はhandlerを作成