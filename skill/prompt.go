package skill

import (
	"bytes"
	"strings"
)

// ToPrompt 生成 <available_skills> XML 块，用于 agent 的 system prompt
// 参考 skills-ref: 建议将生成的 XML 注入到 agent 的 system prompt 中
func ToPrompt(props []Properties) string {
	var buf bytes.Buffer
	buf.WriteString("<available_skills>\n")
	for _, p := range props {
		buf.WriteString("<skill>\n")
		buf.WriteString("<name>\n")
		buf.WriteString(escapeXML(p.Name))
		buf.WriteString("\n</name>\n")
		buf.WriteString("<description>\n")
		buf.WriteString(escapeXML(p.Description))
		buf.WriteString("\n</description>\n")
		buf.WriteString("<location>\n")
		buf.WriteString(escapeXML(p.Location))
		buf.WriteString("\n</location>\n")
		buf.WriteString("</skill>\n")
	}
	buf.WriteString("</available_skills>")
	return buf.String()
}

func escapeXML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	s = strings.ReplaceAll(s, "'", "&apos;")
	return s
}
