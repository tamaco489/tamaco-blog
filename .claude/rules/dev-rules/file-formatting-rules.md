# ファイル フォーマット ルール

## 必須ルール

### 1. ファイル末尾の改行
**すべてのテキストファイルは最終行に改行文字で終わること**

### 2. 行末の空白禁止
**行の末尾に不要な空白文字（スペース、タブ）を残さないこと**

### 3. インデント規則
- **Go**: タブ文字
- **JavaScript/TypeScript/JSON**: 2スペース
- **SQL**: 4スペース
- **YAML**: 2スペース（タブ禁止）
- **Terraform**: 2スペース
- **Makefile**: タブ文字（仕様要件）

### 4. エンコーディング・行終端
- **文字エンコーディング**: UTF-8
- **行終端文字**: LF（Unix形式）

## 適用対象ファイル

**対象**: `.go`, `.js/.ts`, `.sql`, `.yaml/.yml`, `.json`, `.md`, `.sh`, `.tf`, `Makefile`, `.env*`
**除外**: バイナリファイル、自動生成ファイル、サードパーティライブラリ

## Claude Code 作業時の注意

**Write/Edit ツール使用時は必ず最終行に改行を含める**:

```javascript
// ✅ 正しい書き方
content: "最後の行\n"
content: `複数行の
ファイル内容
最終行\n`

// ❌ 間違った書き方
content: "最後の行"  // 改行なし
content: `複数行の
ファイル内容
最終行`  // 改行なし
```

## 検証・修正コマンド

**ファイル作成・更新後は必ず以下のコマンドを実行すること**:

```bash
# 最終行改行チェック
tail -c1 filename | od -An -tx1  # 0a なら改行あり

# 行末空白チェック
grep "[[:space:]]$" filename

# 最終行改行追加（必要時）
[ -n "$(tail -c1 filename)" ] && echo >> filename

# 行末空白削除（必要時）
sed -i 's/[[:space:]]*$//' filename
```

**Claude Code での実行例**:
```bash
# ファイル作成後に必ず実行
tail -c1 /path/to/created/file.ext | od -An -tx1
grep "[[:space:]]$" /path/to/created/file.ext
```

## 例外ケース

**Markdownでの行末2スペース**: 改行を意味するため意図的な場合は許可

