### 使い方

```sh
# 認証
aws-vault exec <profile>

# コマンド
# PutObject（アップロード用URL生成）
go run main.go -m put -b <bucket_name> -k <object_key>

# GetObject（ダウンロード用URL生成）
go run main.go -m get -b <bucket_name> -k <object_key>

# PostObject（フォームアップロード用URL生成）
go run main.go -m post -b <bucket_name> -k <object_key>

# DeleteObject（削除用URL生成）
go run main.go -m delete -b <bucket_name> -k <object_key>
```

### パラメータ

- `-m <method>`: 必須。URLの種類を指定します。選択肢: `get`（ダウンロード用）, `put`（アップロード用）, `post`（フォームアップロード用）, `delete`（削除用）
- `-b <bucket_name>`: 必須。S3バケット名
- `-k <object_key>`: 必須。S3オブジェクトキー
