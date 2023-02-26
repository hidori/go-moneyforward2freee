# go-moneyforward2freee

MoneyForward 口座の「家計簿データの出力（CSV 形式）」を freee 口座の「明細アップロード（CSV 形式）」に変換します。

## ビルド

```bash
go build -o ./bin/moneyforward2freee ./main.go
```

## 使用方法

以下のコマンドラインを実行すると、標準出力に変換結果を出力します。

### コマンドライン
```sh
moneyforward2freee ./example.csv
```

### 変換結果
```csv
"Date","Amount","Description"
"2022/01/28",181,"AMAZON.CO.JP"
"2022/01/28",530,"AMAZON.CO.JP"
"2022/01/29",130,"APPLE COM BILL"
"2022/01/29",1326,"AMAZONDOWNLOADS"
"2022/01/29",663,"AMAZONDOWNLOADS"
```
