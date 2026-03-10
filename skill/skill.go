// Package skill 实现 Anthropic Agent Skills 标准的 Golang SDK
//
// 参考规范: https://agentskills.io/specification
package skill

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	SkillFileName = "SKILL.md"
	MaxNameLen    = 64
	MaxDescLen    = 1024
	MaxCompatLen  = 500
)

// Skill 表示一个 Agent Skill，符合 Anthropic Agent Skills 标准
type Skill struct {
	// 必需字段
	Name        string `yaml:"name"`
	Description string `yaml:"description"`

	// 可选字段
	License      string            `yaml:"license,omitempty"`
	Compatibility string           `yaml:"compatibility,omitempty"`
	Metadata     map[string]string `yaml:"metadata,omitempty"`
	AllowedTools string            `yaml:"allowed-tools,omitempty"`

	// 解析后填充
	Body     string // Markdown 正文
	Location string // SKILL.md 的完整路径
}

// Properties 用于轻量级元数据读取（仅 name、description、location）
type Properties struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
}

// Load 从指定目录加载 Skill
// dir 为 skill 目录路径（包含 SKILL.md 的目录）
func Load(dir string) (*Skill, error) {
	skillPath := filepath.Join(dir, SkillFileName)
	data, err := os.ReadFile(skillPath)
	if err != nil {
		return nil, err
	}

	absPath, err := filepath.Abs(skillPath)
	if err != nil {
		absPath = skillPath
	}

	return Parse(data, absPath)
}

// Parse 从 SKILL.md 内容解析 Skill
// location 为 SKILL.md 的路径，用于生成 prompt 等
func Parse(data []byte, location string) (*Skill, error) {
	content := string(data)
	parts := strings.SplitN(content, "---", 3)
	if len(parts) < 3 {
		return nil, ErrInvalidFrontmatter
	}

	// parts[0] 可能是空或 BOM，parts[1] 是 YAML，parts[2] 是 body（可能以 \n 开头）
	var s Skill
	if err := yaml.Unmarshal([]byte(strings.TrimSpace(parts[1])), &s); err != nil {
		return nil, err
	}

	s.Body = strings.TrimLeft(parts[2], "\n\r")
	s.Location = location
	return &s, nil
}

// Properties 返回技能的轻量级属性（用于 agent 启动时预加载）
func (s *Skill) Properties() Properties {
	return Properties{
		Name:        s.Name,
		Description: s.Description,
		Location:    s.Location,
	}
}
