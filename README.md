# yas80 - Yet Another Assembler for Z80 and R800

<div style="color:red;font-size:2em;font-weight:bold">UNDER CONSTRUCTION </div> 

## 主な特長

- マルチパスアセンブラ
  - 1st stage
    - 未定義シンボルの解決まで最大 256 回の評価
  - 2nd staget
    - 生成コードが安定するまで評価
- 行書式
  - マルチステートメント
  - 継続行
  - 連続する文字列リテラルの結合
- スコープ管理
  - .name スコープ内
  - @name マクロ内
- 関数
  - 複数行の定義
  - 1行 = 1式 の関数定義シンタックスシュガー
  - 値として利用可能
  - クロージャ
- シンボル
  - 定数
  - 変数（再代入可能）
  - シンボル結合演算子（##）
- ORG による生成コードの配置指定
- マクロ制御 EXITM
  - 後置 IF シンタクスシュガー（EXITM IF expr）
- CHARAMAP
  - JSON ファイルの他、直接 JSON 文字列で生成
  - SETMAP によるロード後の CHARMAP の修正
  - CHARAMP に未定義の文字に対する挙動の選択
    - エラー（デフォルト）
    - 特定の文字へ置き換え
    - 元の文字コードのまま
- 配列リテラルの利用
  - 可変長マクロ引数
  - REPT + 配列リテラルで配列要素を値としたマクロ展開
- Z80 命令
  - Z80 CPU Users Manual 掲載の命令 ( IN Flags, (C) を除く)
  - IXH, IXL, IYH, IYL 非公開命令
- R800 対応
  - MUL 命令（mulub, muluw）
  - T States を出力（リストファイル）

## マルチパスアセンブラ

## スコープ

## ORG による生成コードの配置指定

## パフォーマンス

###  計測対象

| Assembler | Version | File Size (byte) |
|  -- | -- | --: |
| yas80 | 0.1.0 (prototype) |    4,757,504  |
| z80asm(z88dk) |Z80 Macro Assembler 22110-51889e5300-20231220 | 32,393,835 |
| ailz80asm |  1.0.31.0 | 68,322,858 |

### 計測結果

<style>
.number {
  font-family: consolas, monotype; 
  text-align: right;
}
</style>
<table>
<thead>
  <tr>
    <th>Source</th>
    <th>Execution Count</th>
    <th>Asselmber</th>
    <th>Total Execution Time (sec)</th>
    <th>Avarege Execution Time (sec)</th>
  </tr>
</thead>
<tbody>
  <tr>
    <td rowspan="3"><a href="testdata/all.asm">testdata/all.asm</td>
    <td rowspan="3" style="text-align:center">100</td>
    <td>yas80</td>
    <td class="number">4.560</td>
    <td class="number">0.046</td>
  </tr>
  <tr>
    <td>z80asm(z88dk)</td>
    <td class="number">15.411</td>
    <td class="number">0.154</td>
  </tr>
    <td>ailz80asm</td>
    <td class="number">106.820</td>
    <td class="number">1.048</td>
  </tr>
  <tr>
    <td rowspan="3"><a href="testdata/label.asm">testdata/label.asm</td>
    <td rowspan="3" style="text-align:center">100</td>
    <td>yas80</td>
    <td class="number">6.414</td>
    <td class="number">0.064</td>
  </tr>
  <tr>
    <td>z80asm(z88dk)</td>
    <td class="number">27.126</td>
    <td class="number">0.271</td>
  </tr>
    <td>ailz80asm</td>
    <td class="number">680.028</td>
    <td class="number">6.800</td>
  </tr>

</tbody>
</table>

- all.asm は include を含み、展開後 32,295 行の命令文となる
- label.asm は all.asm を展開した 32,295 行全てにラベルを付与したもの
- 生成ファイルはどちらも 65,510 byte

## ライセンス

[MIT](LICENSE)