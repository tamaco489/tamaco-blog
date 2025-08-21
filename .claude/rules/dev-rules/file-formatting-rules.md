# ファイル フォーマット ルール

## 概要

このドキュメントは、コードやドキュメントファイルを作成・編集する際の統一的なフォーマットルールを定義します。

## 必須ルール

### 1. ファイル末尾の改行（EOF newline）

**すべてのテキストファイル（コード、SQL、Markdown、設定ファイル等）は、最終行に改行文字で終わること**

```
✅ 正しい例:
line 1
line 2
line 3
[EOF newline]

❌ 間違った例:
line 1
line 2
line 3[EOF without newline]
```

**理由**:

- POSIX 標準の要件
- Git での差分表示が正常に動作
- 一部のツールやエディタでの互換性確保
- ファイル結合時の問題回避

### 2. 行末の不要な空白（Trailing whitespace）

**行の末尾に不要な空白文字（スペース、タブ）を残さないこと**

```
✅ 正しい例:
const message = "Hello World";[改行]
console.log(message);[改行]

❌ 間違った例:
const message = "Hello World";   [空白][改行]
console.log(message);[改行]
```

**理由**:

- コードの見た目の一貫性
- Git での不要な差分を避ける
- エディタでの表示問題回避
- ファイルサイズの無駄削減

### 3. 統一したインデント

**インデントは以下のルールに従うこと**

#### Go 言語

- **タブ文字** を使用
- インデントレベル: 1 タブ = 8 スペース相当

#### JavaScript/TypeScript/JSON

- **2 スペース** を使用
- タブ文字は使用しない

#### SQL

- **4 スペース** を使用
- タブ文字は使用しない

#### YAML

- **2 スペース** を使用
- タブ文字は絶対に使用しない（YAML 仕様）

#### Makefile

- **タブ文字** を使用（Makefile 仕様の要件）

### 4. 文字エンコーディング

**すべてのテキストファイルは UTF-8 エンコーディングを使用すること**

### 5. 行終端文字（Line endings）

**LF（Unix 形式）を使用し、CRLF（Windows 形式）は使用しないこと**

## 適用対象ファイル

### 必須適用

- `.go` - Go ソースコード
- `.js`, `.ts`, `.jsx`, `.tsx` - JavaScript/TypeScript
- `.sql` - SQL ファイル
- `.yaml`, `.yml` - YAML 設定ファイル
- `.json` - JSON 設定ファイル
- `.md` - Markdown ドキュメント
- `.sh` - シェルスクリプト
- `Makefile` - Makefile ファイル
- `.env*` - 環境変数ファイル

### 除外対象

- バイナリファイル（画像、動画、圧縮ファイル等）
- 自動生成ファイル（`*.gen.go`, `node_modules/`等）
- サードパーティライブラリファイル

## チェック・検証方法

### 1. エディタ設定

**VS Code**:

```json
{
  "files.trimTrailingWhitespace": true,
  "files.insertFinalNewline": true,
  "files.trimFinalNewlines": true
}
```

### 2. Git Pre-commit Hook

```bash
#!/bin/sh
# .git/hooks/pre-commit

# Check for trailing whitespace
if git diff --cached --check; then
    echo "✅ No trailing whitespace found"
else
    echo "❌ Trailing whitespace found. Please fix before committing."
    exit 1
fi

# Check for final newline
for file in $(git diff --cached --name-only); do
    if [ -f "$file" ] && [ "$(tail -c1 "$file" | wc -l)" -eq 0 ]; then
        echo "❌ File $file does not end with newline"
        exit 1
    fi
done
```

### 3. コマンドラインでの確認

```bash
# 行末の空白をチェック
grep -r "[[:space:]]$" . --include="*.go" --include="*.js" --include="*.sql"

# 最終行の改行をチェック
find . -name "*.go" -o -name "*.js" -o -name "*.sql" | xargs -I {} sh -c 'if [ "$(tail -c1 "{}" | wc -l)" -eq 0 ]; then echo "No final newline: {}"; fi'
```

## 自動修正

### 1. 行末空白の削除

```bash
# 特定のファイル
sed -i 's/[[:space:]]*$//' filename

# ディレクトリ内の全ファイル
find . -name "*.go" -exec sed -i 's/[[:space:]]*$//' {} \;
```

### 2. 最終行改行の追加

```bash
# 特定のファイル
echo >> filename

# より安全な方法
[ -n "$(tail -c1 filename)" ] && echo >> filename
```

## Claude Code 作業時の注意事項

**Claude Code を使用してファイルを作成・編集する際は、必ず以下を実行前に確認すること**:

1. **Write/Edit ツール使用時**:

   - ファイル内容の最終行に改行が含まれているか確認
   - 各行の末尾に不要な空白がないか確認

2. **MultiEdit ツール使用時**:

   - 各編集内容に不要な空白が含まれていないか確認
   - 最終的なファイル内容が正しい改行で終わるか確認

3. **ファイル作成後の検証**:
   - `cat -n` コマンドで内容と行番号を確認
   - 必要に応じて `hexdump -C` で文字コードレベルの確認

## 例外ケース

### Markdown ファイルでの行末スペース

Markdown では行末の 2 スペースが改行を意味するため、意図的な場合は許可する:

```markdown
これは一行目です。  
これは二行目です。
```

ただし、意図しない行末スペースは削除すること。

## まとめ

これらのルールを守ることで:

- コードベース全体の一貫性向上
- Git での差分管理の改善
- 開発ツール間の互換性確保
- 予期しないエラーの回避

**すべての新規作成・編集ファイルでこれらのルールを適用すること**。
