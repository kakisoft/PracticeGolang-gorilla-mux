gorilla/mux のコードを書いて。
HelloWordレベルでいい。

________________________

### 改善候補

1. RESTサンプルを実装に合わせる
2. `addHandler` のテストを追加する
3. `fmt.Sscanf` を `strconv.Atoi` に変更する  
   現状では `1abc` のような値を `1` として受理する可能性があります
4. `gofmt` を適用する
5. READMEに全APIと実行例を記載する
6. 本番利用を想定するなら、HTTPサーバーに読み書きのタイムアウトを設定する
7. Petstore APIを実装する予定がないなら `petstore.yaml` を削除、またはサンプル資料だと明記する


