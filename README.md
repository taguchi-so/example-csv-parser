Windows のExcelが作成したCSVファイルをUnmarshalした。

```
encoding/csv
Text: テキスト
Textlrln: 数値
Numeric: URL
URL: セル内改行テキスト
Text: ああああああああああああ
Textlrln: 123456789
Numeric: https://example.com/
URL: あいうえお
かきくけこ
さしすせそ
Text: いいいいいいいい
Textlrln: 123456789
Numeric: http://example.com/
URL: あいうえお
かきくけこ
さしすせそ

github.com/gocarina/gocsv
Text: ああああああああああああ
Textlrln: あいうえお
かきくけこ
さしすせそ
Numeric: 123456789
URL: https://example.com/
Text: いいいいいいいい
Textlrln: あいうえお
かきくけこ
さしすせそ
Numeric: 123456789
URL: http://example.com/


```
