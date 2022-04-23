# code-database
[![CI](https://github.com/geek-line/code-database/actions/workflows/ci.yml/badge.svg)](https://github.com/geek-line/code-database/actions/workflows/ci.yml)
## デプロイ先
https://code-database.com/

## 作業環境 
Docker

## 使用ライブラリ
 - joho/godotenv
 - gorilla/sessions
 - go-sql-driver/mysql

## 本番ビルド/サーバ立ち上げ
フロントエンドのビルド
```
npm run build
```
バックエンドのビルド・サーバー立ち上げ
```
make start
```
(option)バックエンドのビルドのみ
```
make build
```

## 開発用ビルド/サーバー立ち上げ方法
開発用サーバー立ち上げ
```
make app
```

## 作業ルール  
 - 作業を始める前にdevelopをpullする  
 - 作業が完了したらdevelopにpull requestを送る  
 - margeが終わったら、作業用ブランチは削除する
 - productブランチはamazonLinux環境での動作確認用
 - masterブランチは本番用ブランチ

## URLのガイド
| URL | 表示される画面 |
| ------------- | ------------- |
| /knowledges | 全記事の一覧画面 |
| /search  | 記事のキーワード検索結果画面  |
| /knowledges/[記事のid]  | 記事の内容画面  |
| /tags/[タグのid]  | タグがついている記事の一覧画面  |
| /categories  | カテゴリ一覧画面  |
| /categories/[カテゴリのid]  | カテゴリ詳細画面  |
| /about  | aboutページ  |
| /privacy  | プライバシーポリシー  |
| /categories/[カテゴリのid]  | カテゴリ詳細画面  |
| /admin/login/ | アドミンへのログイン画面 |
| /admin/knowledges/ | 記事の一覧画面 |
| /admin/knowledges/[記事のid] | 記事の編集画面 |
| /admin/knowledges/new/ | 記事の新規作成画面 |
| /admin/tags/ | タグの編集画面 |
| /admin/eyecatches/ | アイキャッチの編集画面 |

## フォルダの説明
### バックエンド
| PATH | フォルダの中身 |
| ------------- | ------------- |
| /config/ | .envファイルを用いた環境変数の初期化 |
| /handlers/ | リクエストに対するハンドラ(表示の役割も兼ねる) |
| /middleware/ | リクエストに対してhandlersに処理がわたる前の中間処理(認証系など)を行う |
| /models/ | データベースとのやりとりを行う |
| /routes/  | サイト内に存在するPATHを管理する |
| /structs/  | 構造体の宣言用 |
| /development/ | 開発用に使う関数のパッケージ |
| server.go | サーバーの立ち上げとリクエストに対するハンドラの振り分け |
### フロントエンド

- typescript + webpackでhtml(goのテンプレート)/css/jsファイルを生成
- `src/`ディレクトリにファイルを格納
  - componentsディレクトリ: htmlを含むページごとのhtml/css/tsファイルの組を記述。(ディレクトリ名がhtmlのファイル名としてビルドされる)
  - helpersディレクトリ: componentsディレクトリにてimportされる共通化ファイル
- 画像は`public/`ディレクトリに格納

### ビルドされた後の静的ファイル(ファイルは原則編集禁止)
| PATH | フォルダの中身 |
| ------------- | ------------- |
| /dist/css | cssファイルの作成 |
| /dist/js | jsファイルの作成 |
| /dist/template/ | htmlファイルの作成(goファイルのtemplateパッケージで使用) |
