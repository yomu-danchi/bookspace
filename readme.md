

## インストール手順

```shell
docker-compose up --build
```

## Tips

### DockerのDB接続
DBのホスト名がlocalhostだとソケット通信を試みるので失敗する。```127.0.0.1```のIPでの指定が必要。
```shell
mysql -h 127.0.0.1 -u ユーザー名 -p パスワード DB名
```
