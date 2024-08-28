# 環境構築

## git
### 初回
```
// クローン(http)
git clone {URL}
//.env作成
touch .env
//.envに以下の内容を記述
DB_DATABASE=mazegame
DB_USER=shin
DB_PASSWORD=mazegame
DB_ROOT_PASSWORD=root
```
### mainブランチからブランチを切るとき
```
//ローカルのmainブランチを最新の状態にする
git pull origin main
//ローカルに新しいブランチを作成
git branch {ブランチ名}
//作成したブランチに移動
git switch {ブランチ名}
//ローカルにブランチを追加
git push origin {ブランチ名}
```
### リモートにpushする手順
```
//ステージングする
git add .
//commitする
git commit -m "{コミットメッセージ}"
// pushする
git push origin {ブランチ名}
```
### 間違えてpushしたコミットを消したいとき
```
//打ち消したいコミットのIDを確認
git log
//ローカルのcommitを打ち消す
git revert {コミットID}
//リモートのcommitを消すためにpushする
git push origin {ブランチ名}
```
## docker
```
//ビルド
docker compose build

//コンテナ立ち上げ
docker compose up

//DBコンテナ立ち上げが上手くいかなかった場合
docker compose up gabaithon-09-db -d
docker compose up
```