# yas80 - Yet Another Assembler for Z80 and R800

# UNDER CONSTRUCTION 

## 主な特長

- マルチパスアセンブラ
  - 1st stage - 未定義シンボルの解決規定回数評価
  - 2nd stage - 生成コードの安定度評価
- 行書式
  - マルチステートメント
  - 継続行
  - 連続する文字列リテラルの結合
- スコープ管理
  - PROC
  - マクロ
- 関数
  - 複数行
  - 単一行のシンタックスシュガー
  - 値として参照可能
  - クロージャ
- シンボル
  - 定数
  - 変数（再代入可能）
  - シンボル結合演算子（##）
- ORG による生成コードの配置指定
- マクロ制御 EXITM
  - 後置 IF シンタクスシュガー
  - MACRO, REPT
- CHARAMAP
  - JSON ファイルの他、直接 JSON 文字列で生成
  - SETMAP による定義済み CHARMAP の更新
  - 未定義の文字に対する挙動選択
    - エラー（デフォルト）
    - 特定の文字へ置き換え
    - 元の文字コードのまま
- 配列リテラル
  - 可変長マクロ引数
  - REPT + 配列リテラル
- Z80
  - [Z80 Family CPU User Manual](https://www.zilog.com/docs/z80/z80cpu_um.pdf) 
  - IXH, IXL, IYH, IYL 非公開命令
- R800 
  - MUL 命令（mulub, muluw）
  - T States を出力（リストファイル）

## マルチパスアセンブラ

## スコープ

## ORG による生成コードの配置指定

## パフォーマンス

###  計測対象アセンブラ

| アセンブラ| バージョン | ファイルサイズ |
|  -- | -- | --: |
| yas80 | 0.1.0 (prototype) |    4,757,504  |
| z80as | 0.12 | 172,032 |
| tools80 | Release 6.48 (Ver. 6.6.66) | (jar) 186,306 |
| z80asm(z88dk) |Z80 Macro Assembler 22110-51889e5300-20231220 | 32,393,835 |
| ailz80asm |  1.0.31.0 | 68,322,858 |

### 計測対象ソース

| ソースファイル | 内容 |
| -- | -- | 
| [testdata/big.asm](testdata/big.asm) | Z80 公開命令を含むファイルを 45 回 include |
| [testdata/label.asm](label/big.asm) | big.asm を展開し、全命令行（32,295行）にラベルを付けたもの |

### 計測結果

#### [testdata/big.asm](testdata/big.asm)  100 回アセンブル

|アセンブラ | 総実行時間（秒）| 平均実行時間 |
| -- | --: | --: |
| yas80        |   4.379 | 0.044 |
| z80as        |   1.579 | 0.016 | 
| z80asm(z88dk)|  15.072 | 0.151 |
| tools80      |  40.184 | 0.402 |
| ailz80asm    | 104.192 | 1.042 |

#### [testdata/label.asm](testdata/label.asm)  10 回アセンブル

|アセンブラ | 総実行時間（秒）| 平均実行時間 |
| -- | --: | --: |
| yas80        |   0.644 | 0.064 |
| z80as        |   0.258 | 0.026 | 
| z80asm(z88dk)|   2.721 | 0.272 |
| tools80      |  24.633 | 0.246 |
| ailz80asm    |  62.493 | 6.249 |


## ライセンス

[MIT](LICENSE)