# Go コーディングルール

## 概要

このドキュメントは、tamaco-blog バックエンドの Go 言語コーディング規約を定義します。

## エラーメッセージ

### エラーメッセージは小文字で開始

エラーメッセージは小文字で始めることを必須とします。

```go
// ✅ 正しい例
return fmt.Errorf("failed to connect database: %w", err)
return errors.New("user not found")

// ❌ 間違った例
return fmt.Errorf("Failed to connect database: %w", err)
return errors.New("User not found")
```

**理由**: Go の標準ライブラリおよびコミュニティの慣習に従い、エラーメッセージは小文字で開始します。これにより、エラーチェーンが読みやすくなります。

## 制御構文

### else 文の使用を避ける

if 文を使用する場合は、極力 else 文を使わず、多少冗長であっても if 文で記述することを推奨します。

```go
// ✅ 推奨例 - else を使わない
func processUser(user *User) error {
    if user == nil {
        return errors.New("user is nil")
    }

    if user.ID == 0 {
        return errors.New("user ID is required")
    }

    if user.Name == "" {
        return errors.New("user name is required")
    }

    // 正常な処理を実行
    return user.Save()
}

// ❌ 非推奨例 - else を使用
func processUser(user *User) error {
    if user == nil {
        return errors.New("user is nil")
    } else {
        if user.ID == 0 {
            return errors.New("user ID is required")
        } else {
            if user.Name == "" {
                return errors.New("user name is required")
            } else {
                return user.Save()
            }
        }
    }
}
```

**理由**: else 文を避けることで、コードの可読性が向上し、ネストが浅くなります。また、エラーケースを早期に処理することで、メインロジックが明確になります。

### スライスの事前メモリ確保を推奨

for文を使用する場合は、できるだけappendを使わずに、最初にスライスの要素数分のメモリを確保し、インデックスを使用して値を設定することを推奨します。

```go
// ✅ 推奨例 - 事前にメモリを確保してインデックスアクセス
func convertToUppercase(strs []string) []string {
    result := make([]string, len(strs))
    for i, str := range strs {
        result[i] = strings.ToUpper(str)
    }
    return result
}

// ✅ 推奨例 - サイズが分からない場合はcapを指定
func filterEvenNumbers(numbers []int) []int {
    result := make([]int, 0, len(numbers)) // cap を指定
    for _, num := range numbers {
        if num%2 == 0 {
            result = append(result, num)
        }
    }
    return result
}

// ❌ 非推奨例 - append のみを使用（メモリの再割り当てが発生する可能性）
func convertToUppercase(strs []string) []string {
    var result []string
    for _, str := range strs {
        result = append(result, strings.ToUpper(str))
    }
    return result
}
```

**理由**: 事前にメモリを確保することで、スライスの拡張時の再割り当てを避けられ、パフォーマンスが向上します。特に大きなスライスを扱う場合に効果的です。

