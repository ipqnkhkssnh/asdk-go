package skill

// ReadProperties 仅读取 skill 的元数据（name、description、location），
// 不解析完整 body，适用于 agent 启动时预加载所有 skill 的概要信息
func ReadProperties(dir string) (Properties, error) {
	s, err := Load(dir)
	if err != nil {
		return Properties{}, err
	}
	return s.Properties(), nil
}
