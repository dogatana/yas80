# yas80 - Yet Another Assembler for Z80 and R800

## はじめに

Go で書いた Z80/R800 マクロアセンブラです。

Go 言語を使い始め『Go 言語でつくるインタプリタ』を読んでいるうちに何か作りたくなり、
それならば Z80 アセンブラを書いてみようと思い立ったたものの、
手書きパーサにするどうか悩んでいたところ、
その昔齧ったことのある yacc も利用できることもわかったので、
go-yacc を使って書いてみました。

こんな経緯から yas80 の "ya" は yacc の "ya" と同じです。
(Yet Another xx を使いたかっただけとも)

Z80 用アセンブラは既にいくつか利用可能ですが、
これまで使用している中で個人的に欲しかった仕様を盛り込んでいます。
仕様はともかく、性能面では満足いくものになりました。

## 主な特長

- マルチパスアセンブラ
- グローバル/PROC/マクロ スコープ
- ORG による生成コードの配置指定
- bin/mzt/t88 形式出力
- マルチステートメント
- 行継続
- 連続する文字列リテラルの結合
- 変数（再代入可能シンボル）
- シンボル結合演算子（##）
- 複数行関数
- クロージャ
- EXITM マクロ制御構文
- CHARAMP 生成で直接 JSON 文字列を指定
- CHARMAP 適用時に未定義文字の扱いを選択可能（エラー、特定文字、元の文字）
- SETMAP による CHARMAP 定義内容更新
- 配列リテラル
- R800 命令（乗算命令、Z80非公開命令）

## 仕様説明

[yas80-docs](https://dogatana.github.io/yas80-docs/)

## ライセンス

[MIT](LICENSE)

## アセンブル性能

### 計測観点

アセンブル性能を計測するとともに、入手容易なアセンブラと比較したものです。
疑似命令（ディレクティブ）はアセンブラ毎の仕様差が大きくて共通ソースとすることが困難なため、
Z80 公開命令の範囲でアセンブル時間の比較を行ったものとなっています。

公開命令でもアセンブラによってはエラーになる場合もあるために、
それを除いた最大公約数的なソースになっています。

またアセンブラのシンボル管理処理に負荷をかける（かもしれない）という観点で、
各行にラベルを付加したものも使用しています。

結果は下にありますが、やはりCPUバイナリを直接生成するものが速く、
中でも z80as は非常に高速です。

###  計測対象アセンブラ

| アセンブラ    | バージョン                | サイズ     | 備考            |  
|  --           | --                        | --:        | --              | 
| yas80.exe     | 0.4.0 (prototype)         |  3,693,568 |                 | 
| z80as.exe     | 0.12                      |    172,032 |                 | 
| z80asm.exe    | 2.4                       | 39,447,620 | MSVCRT.dll 依存 |
| tools80.jar   | Release 6.48 (Ver. 6.6.66)|    186,306 | 要 java.exe     |
| ailz80asm.exe | 1.0.31.0                  | 68,322,858 | 複数 dll 依存（おそらく埋め込み） | 

### 計測対象ソース

| ソースファイル                     | 行数   | 内容                         |
| --                                 | --:    |  --                          |
| [min.asm](performance/min.asm)     |    708 | Z80 公開命令のサブセット     |
| [max.asm](performance/max.asm)     | 32,340 | 生成コードが 64KB を超えない範囲で min.asm の内容を繰り返したもの |
| [label.asm](performance/label.asm) | 32,340 | max.asm の各行にラベルを付加 |

### 計測方法

- Python スクリプトで対象ソースファイルを所定回数アセンブルし、各回の実行時間を計測
- 実行時間の合計を総実行時間とし、総実行時間を回数で割った平均実行時間を求める

### 計測結果

#### min.asm 100回

|アセンブラ    | 総実行時間（秒）| 平均実行時間（秒） |
| --           |       --:       |                --: |
| yas80        |           0.781 |              0.008 |
| z80as        |           0.428 |              0.004 | 
| z80asm(z88dk)|           1.024 |              0.010 |
| tools80      |          10.500 |              0.105 |
| ailz80asm    |          48.457 |              0.485 |

![min](https://github.com/dogatana/yas80/performance/min.svg)

#### max.asm 100回

|アセンブラ    | 総実行時間（秒）| 平均実行時間（秒） |
| --           |       --:       |                --: |
| yas80        |           4.120 |              0.041 |
| z80as        |           1.414 |              0.014 | 
| z80asm(z88dk)|          14.196 |              0.142 |
| tools80      |          30.844 |              0.308 |
| ailz80asm    |          98.436 |              0.984 |

![max](https://github.com/dogatana/yas80/performance/max.svg)

#### label.asm 100回

|アセンブラ    | 総実行時間（秒）| 平均実行時間（秒） |
| --           |       --:       |                --: |
| yas80        |           6.243 |              0.062 |
| z80as        |           2.550 |              0.025 | 
| z80asm(z88dk)|          27.537 |              0.275 |
| tools80      |         249.643 |              2.496 |
| ailz80asm    |         685.302 |              6.853 |

![label](https://github.com/dogatana/yas80/performance/label.svg)


## 参考

yas80 作成にあたり、次の書籍、サイトを参考にさせていただきました。

- Thorsten Ball 著、設樂 洋爾 訳. Go言語でつくるインタプリタ. O'Reilly Japan, 2018.<br>https://www.oreilly.co.jp/books/9784873118222/  
- Jon Bodner著、武舎 広幸 訳. 初めてのGo言語 第2版. O'Reilly Japan, 2025.<br>https://www.oreilly.co.jp/books/9784814401192/
- 近藤 嘉雪 著. yaccによるCコンパイラプログラミング. ソフトバンク, 1990
- [goyaccで構文解析を行う](https://qiita.com/k0kubun/items/1b641dfd186fe46feb65)
- [MZTファイルの仕様について覚書](https://mzakd.cool.coocan.jp/starthp/mzt.html) - [AKD's site](https://mzakd.cool.coocan.jp/main.html)
- [T88Format](https://quagma.sakura.ne.jp/manuke/t88format.html) - [Manuke Station](https://quagma.sakura.ne.jp/manuke/index_j.html)
- [DumpListEditor](https://bugfire2009.ojaru.jp/download.html) - [PC-8001を懐かしむページ](https://bugfire2009.ojaru.jp/index.html)
- z80as - [We Love MZ-700](http://www.maroon.dti.ne.jp/youkan/mz700/index.html)
- z80asm - [z88dk](https://z88dk.org/site/)
- z80as - [OUT of STANDARD [PC-8001]](http://upd780c1.g1.xrea.com/pc-8001/index.html#UTL) 
- ailz80asm  https://github.com/AILight/AILZ80ASM
