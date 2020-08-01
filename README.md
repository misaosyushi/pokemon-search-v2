# pokemon-search-v2

ポケモン検索BOT version2のリポジトリです

## ローカル環境構築

- Goの依存パッケージをインストール

```
go mod download
```

- Elasticsearchの起動

```
docker-compose up -d
```

- テストデータ投入

```
curl -H "Content-Type: application/json" -X POST "localhost:9200/pokemon/_bulk?pretty&refresh" --data-binary "@pokemon.json"
```

- データが入ったか確認

```
curl -X GET "localhost:9200/pokemon/_search?pretty" -H 'Content-Type: application/json' -d'
{
  "query": {
     "match" : { "name" : "サルノリ" }
  }
}
'
```

## APIコンテナイメージのビルド, 起動

- イメージのビルド

```
docker build -t pokemon-search src/
```

- コンテナ起動

```
docker run -p 8081:8081 -it pokemon-search
```
