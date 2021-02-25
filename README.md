## 外部依存コマンド
- pandoc

## DBセットアップ

`docker`,`docker-compose`は予めインストール。

### 起動
```
$ docker-compose up -d
```
- mongodbは`localhost:27017`、webクライアントは`localhost:8800`から参照可
- cliは`mongo`コンテナ内にログインすることでも利用できる（`docker-compose exec mongo bash`）
- ユーザー名:`root`、パスワード:`passwd`

### 停止
```
$ docker-compose down
```
