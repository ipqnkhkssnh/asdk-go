package skill

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	data := []byte(`---
name: test-skill
description: A test skill for parsing.
---

# Body

Content here.`)

	s, err := Parse(data, "/path/to/test-skill/SKILL.md")
	if err != nil {
		t.Fatal(err)
	}
	if s.Name != "test-skill" {
		t.Errorf("Name = %q, want test-skill", s.Name)
	}
	if s.Description != "A test skill for parsing." {
		t.Errorf("Description = %q", s.Description)
	}
	if s.Body != "# Body\n\nContent here." {
		t.Errorf("Body = %q", s.Body)
	}
	if s.Location != "/path/to/test-skill/SKILL.md" {
		t.Errorf("Location = %q", s.Location)
	}
}

func TestLoad(t *testing.T) {
	dir := filepath.Join("..", "testdata", "valid-skill")
	s, err := Load(dir)
	if err != nil {
		t.Fatal(err)
	}
	if s.Name != "valid-skill" {
		t.Errorf("Name = %q, want valid-skill", s.Name)
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name     string
		skill    *Skill
		dir      string
		wantAny  bool
	}{
		{
			name: "valid",
			skill: &Skill{Name: "valid-skill", Description: "A valid skill."},
			dir:  "valid-skill",
			wantAny: false,
		},
		{
			name:     "empty name",
			skill:    &Skill{Name: "", Description: "Desc"},
			dir:      "",
			wantAny:  true,
		},
		{
			name:     "name too long",
			skill:    &Skill{Name: strings.Repeat("a", 65), Description: "Desc"},
			dir:      "",
			wantAny:  true,
		},
		{
			name:     "name starts with hyphen",
			skill:    &Skill{Name: "-bad", Description: "Desc"},
			dir:      "",
			wantAny:  true,
		},
		{
			name:     "name has consecutive hyphens",
			skill:    &Skill{Name: "bad--name", Description: "Desc"},
			dir:      "",
			wantAny:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			problems := Validate(tt.skill, tt.dir)
			hasProblems := len(problems) > 0
			if hasProblems != tt.wantAny {
				t.Errorf("Validate() problems = %v, wantAny = %v", problems, tt.wantAny)
			}
		})
	}
}

func TestToPrompt(t *testing.T) {
	props := []Properties{
		{Name: "skill-a", Description: "First skill", Location: "/path/a/SKILL.md"},
		{Name: "skill-b", Description: "Second skill", Location: "/path/b/SKILL.md"},
	}
	out := ToPrompt(props)
	if out == "" {
		t.Fatal("ToPrompt returned empty")
	}
	if !strings.Contains(out, "<available_skills>") {
		t.Error("expected <available_skills>")
	}
	if !strings.Contains(out, "skill-a") || !strings.Contains(out, "skill-b") {
		t.Error("expected skill names in output")
	}
}
