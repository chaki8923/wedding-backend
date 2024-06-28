# Getting Started

```shell:
$ cp .env.example .env
$ make init
$ make migrate
$ make start
```

1. route.goでリクエストごとの処理振り分け
2. handlerで簡単なバリデーションと目的の処理(登録、更新、取得)行い値やstatusを返す。handleではバリデーションを使用するだけで実装はしない
(バリデーションの詳細実装はusecaseやカスタムバリデーション内)
3. usecaseではrepositoryで定義されたアクションを使用してusecaseを叶える
4. repositoryではエンティティに関わる処理をまとめておく。ここに処理は書かない
5. dbとやり取りするのはinfra層でのみにする(gorm.db)では引数にインスタンスしか渡せない
