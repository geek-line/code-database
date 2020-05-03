# code-database

## デプロイ先
https://code-database.com/

## 作業環境 
Go(1.14 darwin/amd64)  
MySQL(8.0.19)

## 使用ライブラリ
 - joho/godotenv
 - gorilla/sessions
 - go-sql-driver/mysql

## 作業ルール  
 - 作業はISSUEを立ててからISSUE名の入ったブランチをdevelopから切って作業する  
 - 作業を始める前にdevelopをpullする  
 - 作業が完了したらdevelopにpull requestを送る  
 - margeが終わったら、作業用ブランチは削除する
 - productブランチはamazonLinux環境での動作確認用
 - masterブランチは本番用ブランチ

## ISSUEの命名ルール  
ISSUE-(番号) (作業名)

## URLのガイド
| URL | 表示される画面 |
| ------------- | ------------- |
| /knowledges | 全記事の一覧画面 |
| /search  | 記事のキーワード検索結果画面  |
| /knowledge/[記事のid]  | 記事の内容画面  |
| /tags/[タグのid]  | タグがついている記事の一覧画面  |
| /admin/login/ | アドミンへのログイン画面 |
| /admin/knowledges/ | 記事の一覧画面 |
| /admin/knowledges/[記事のid] | 記事の編集画面 |
| /admin/knowledges/new/ | 記事の新規作成画面 |
| /admin/tags/ | タグの編集画面 |
| /admin/eyecatches/ | アイキャッチの編集画面 |

## フォルダの説明
### goファイル
| PATH | フォルダの中身 |
| ------------- | ------------- |
| /config/ | .envファイルを用いた環境変数の初期化 |
| /handlers/ | リクエストに対するハンドラ(表示の役割も兼ねる) |
| /middleware/ | リクエストに対してhandlersに処理がわたる前の中間処理(認証系など)を行う |
| /models/ | データベースとのやりとりを行う |
| /routes/  | サイト内に存在するPATHを管理する |
| /structs/  | 構造体の宣言用 |
| server.go | サーバーの立ち上げとリクエストに対するハンドラの振り分け |
### 静的ファイル(html/css/js/その他) 
| PATH | フォルダの中身 |
| ------------- | ------------- |
| /static/css | cssファイルの作成 |
| /static/js | jsファイルの作成 |
| /static/public | 画像の作成 |
| /template/ | htmlファイルの作成(goファイルのtemplateパッケージで使用) |
