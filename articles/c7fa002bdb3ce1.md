---
title: "OpenAPIで空文字を弾く方法"
emoji: "🈳"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: [ "openapi" ]
published: true
---

# 結論

```yml
properties:
  name:
    type: string
    minLength: 1
```

# おまけ

requiredはキーが必須になるだけで、空文字が送られて来た時に弾くわけではない。

```yml
properties:
  name:
    type: string
    required: true
```