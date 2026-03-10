package skill

import "errors"

var (
	ErrInvalidFrontmatter = errors.New("skill: invalid frontmatter, expected YAML between --- delimiters")
	ErrSkillNotFound      = errors.New("skill: SKILL.md not found in directory")
)
