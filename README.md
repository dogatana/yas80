# yas80


## はじめに

Go で書いた Z80/R800 マクロアセンブラです。

Go 言語を使い始めて『Go 言語でつくるインタプリタ』を読んでいるうちに何か作りたくなり、
それならば Z80 アセンブラを書いてみようと思い立ち、
その昔齧ったことのある yacc も利用できることもわかったので、
goyacc を使って書いてみました。

yas80 は Yet Another Assembler for Z80 and R800 の略で、
yacc の "ya" と同じです。(Yet Another xx を使いたかっただけ)

Z80 用アセンブラは既にいくつか利用可能ですが、
これまで使用している中で個人的に欲しかった仕様を盛り込んでいます。

## 動作環境

Go のクロスコンパイル機能を利用し次表の環境で利用可能です。
ただしこちらで確認しているのは windows - amd64 のみです。

| OS (goos)| CPU (garch)  |
| --       | --           |
| windows  | amd64        |
| linux    | amd64        |
| darwin   | amd64, arm64 |

## インストール

[Releases](https://github.com/dogatana/yad80/releases) からダウンロード、解凍してください。

## ソースからのビルド

```
git clone https://github.com/dogatana/yas80.git
cd yas80
go build
```

## 主な特長

- 寡黙（Rule of Silence に従い成功時は何も表示なし）
- マルチパスアセンブラ
- グローバル/PROC/マクロ スコープ
- 匿名ラベル(PROC内)
- ORG による生成コードの配置指定
- bin/mzt/t88 形式出力
- マルチステートメント
- 行継続
- 連続する文字列リテラルの結合
- 変数（再代入可能シンボル）
- シンボル結合演算子（##）
- 複数行定義の関数
- クロージャ
- EXITM マクロ制御構文
- CHARAMP 生成で直接 JSON 文字列を指定
- CHARMAP 適用時に未定義文字の扱いを選択可能（エラー、特定文字、元の文字）
- SETMAP による CHARMAP 定義内容更新
- 配列リテラル
- R800 命令（乗算命令+Z80非公開命令の一部）

## マニュアル

[別サイト](https://dogatana.github.io/yas80-docs/)に置いています。

## ライセンス

[MIT](LICENSE)

## アセンブル性能

実用性の指標となるアセンブル性能を計測しました。
合わせて容易に入手できる他のアセンブラと比較しています。


###  計測対象アセンブラ

| アセンブラ                            | バージョン                   | サイズ     | 備考            |  
|  --                                   | --                           | --:        | --              | 
| yas80.exe                             | 0.5.0                        |  3,710,464 |                 | 
| z80as.exe<br>(紅茶羊羹氏)             | 0.12                         |    172,032 |                 | 
| z80asm.exe<br>(z88dk)                 | 2.4                          | 39,447,620 | MSVCRT.dll 依存 |
| tools80.jar<br>(HAL 8999氏)           | Release 6.48<br>(Ver. 6.6.66)|    186,306 | 要 java.exe     |
| AILZ80ASM.exe<br>(Mitsuhito Ishino氏) | 1.0.31.0                     | 68,322,858 | 複数 dll 依存（おそらくは埋め込み） | 

### 計測対象ソース

疑似命令（ディレクティブ）はアセンブラ毎の仕様差が大きく共通ソースとすることが困難なため、
Z80 公開命令の範囲でアセンブル時間の比較を行っています。

公開命令でもアセンブラによっては対応していないものもあるため、
それを除いた最大公約数的なソースになっています。

またアセンブラのシンボル管理処理に負荷をかけるという観点で、
各行にラベルを付加したものも使用しています。

| ソースファイル                     | 行数   | 内容                         |
| --                                 | --:    |  --                          |
| [min.asm](performance/min.asm)     |    708 | Z80 公開命令のサブセット     |
| [max.asm](performance/max.asm)     | 32,340 | 生成コードが 64KB を超えない範囲で min.asm の内容を繰り返したもの |
| [label.asm](performance/label.asm) | 32,340 | max.asm の各行にラベルを付加 |

### 計測方法

- Python スクリプトで対象ソースファイルを所定回数アセンブルし、各回の実行時間を計測
- 実行時間の合計を総実行時間とし、総実行時間を回数で割った平均実行時間を求める

### 計測結果

#### min.asm 10 回

|アセンブラ | 平均実行時間（秒）|
| --        |       --:         |
| yas80     |             0.020 |
| z80as     |             0.005 |
| z80asm    |             0.016 |
| tools80   |             0.106 |
| AILZ80ASM |             0.524 |

![min](https://raw.githubusercontent.com/dogatana/yas80/main/performance/min.asm.svg)

#### max.asm 10 回

|アセンブラ | 平均実行時間（秒）|
| --        |       --:         |
| yas80     |             0.041 |
| z80as     |             0.015 |
| z80asm    |             0.143 |
| tools80   |             0.307 |
| AILZ80ASM |             0.942 |

![max](https://raw.githubusercontent.com/dogatana/yas80/main/performance/max.asm.svg)

#### label.asm 10 回

|アセンブラ | 平均実行時間（秒）|
| --        |       --:         |
| yas80     |             0.082 |
| z80as     |             0.026 |
| z80asm    |             0.279 |
| tools80   |             2.471 |
| AILZ80ASM |             6.246 |

![label](https://raw.githubusercontent.com/dogatana/yas80/main/performance/label.asm.svg)

#### ZMA v1.0.19 （参考）

ZMA v1.0.19 (t.hara氏)だと z80asm と tools80 の間にあたる 0.868 秒でした。<br>
ZMA は
- 書式が一般の Z80 アセンブラと異なる
- 対応 Z80 公開命令が更に絞り込まれている（模様）

ため専用ソース（書式変更、エラーとなる命令行コメント化）での計測となっています。

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
- AILZ80ASM  https://github.com/AILight/AILZ80ASM
- zma https://github.com/hra1129/zma
