# asdk-go

Anthropic Agent Skills 标准的 Golang SDK。

参考规范: [agentskills.io/specification](https://agentskills.io/specification)

## 功能

- **Load** - 从目录加载 Skill（解析 SKILL.md）
- **Validate** - 验证 Skill 是否符合规范
- **ReadProperties** - 仅读取元数据（name、description、location），适用于预加载
- **ToPrompt** - 生成 `<available_skills>` XML 供 agent system prompt 使用

## 安装

```bash
go get github.com/ipqnkhkssnh/asdk-go/skill
```

## 使用

### Go API

```go
package main

import (
    "fmt"
    "path/filepath"

    "github.com/ipqnkhkssnh/asdk-go/skill"
)

func main() {
    // 加载并验证 skill
    s, problems, err := skill.ValidateLoad("./my-skill")
    if err != nil {
        panic(err)
    }
    if len(problems) > 0 {
        for _, p := range problems {
            fmt.Printf("%s: %s\n", p.Field, p.Message)
        }
        return
    }

    fmt.Printf("Skill: %s - %s\n", s.Name, s.Description)

    // 读取多个 skill 的属性并生成 prompt
    dirs := []string{"./skill-a", "./skill-b"}
    var props []skill.Properties
    for _, dir := range dirs {
        p, err := skill.ReadProperties(dir)
        if err != nil {
            continue
        }
        props = append(props, p)
    }
    prompt := skill.ToPrompt(props)
    fmt.Println(prompt)
}
```

### CLI

```bash
# 验证 skill
asdk validate path/to/skill

# 读取 skill 属性（输出 JSON）
asdk read-properties path/to/skill

# 生成 <available_skills> XML
asdk to-prompt path/to/skill-a path/to/skill-b
```

## SKILL.md 格式

```yaml
---
name: my-skill
description: 描述该 skill 的功能以及何时使用它
license: Apache-2.0
metadata:
  author: example-org
  version: "1.0"
---

# 技能正文

Markdown 格式的指令内容...
```

## License

Apache 2.0
