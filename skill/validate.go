package skill

import (
	"path/filepath"
	"regexp"
	"strings"
)

// nameRe 校验 name：小写字母、数字、连字符，不含连续连字符
var nameRe = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

// Problem 表示一个验证问题
type Problem struct {
	Field   string
	Message string
}

// Validate 验证 Skill 是否符合 Agent Skills 规范
// dir 为 skill 所在目录，用于校验 name 与目录名一致
func Validate(s *Skill, dir string) []Problem {
	var problems []Problem

	// 1. name 校验
	if s.Name == "" {
		problems = append(problems, Problem{Field: "name", Message: "name is required"})
	} else {
		if len(s.Name) > MaxNameLen {
			problems = append(problems, Problem{Field: "name", Message: "name must be at most 64 characters"})
		}
		if strings.HasPrefix(s.Name, "-") || strings.HasSuffix(s.Name, "-") {
			problems = append(problems, Problem{Field: "name", Message: "name must not start or end with hyphen"})
		}
		if strings.Contains(s.Name, "--") {
			problems = append(problems, Problem{Field: "name", Message: "name must not contain consecutive hyphens"})
		}
		if !nameRe.MatchString(s.Name) {
			problems = append(problems, Problem{Field: "name", Message: "name may only contain lowercase letters, numbers, and hyphens"})
		}
		if dir != "" {
			base := filepath.Base(filepath.Clean(dir))
			if base != s.Name {
				problems = append(problems, Problem{Field: "name", Message: "name must match parent directory name (got " + s.Name + ", dir is " + base + ")"})
			}
		}
	}

	// 2. description 校验
	if s.Description == "" {
		problems = append(problems, Problem{Field: "description", Message: "description is required"})
	} else if len(s.Description) > MaxDescLen {
		problems = append(problems, Problem{Field: "description", Message: "description must be at most 1024 characters"})
	}

	// 3. compatibility 可选，若提供则长度限制
	if s.Compatibility != "" && len(s.Compatibility) > MaxCompatLen {
		problems = append(problems, Problem{Field: "compatibility", Message: "compatibility must be at most 500 characters if provided"})
	}

	return problems
}

// ValidateLoad 加载并验证 skill，返回 skill 和验证问题
func ValidateLoad(dir string) (*Skill, []Problem, error) {
	s, err := Load(dir)
	if err != nil {
		return nil, nil, err
	}
	problems := Validate(s, dir)
	return s, problems, nil
}
