package test

import (
	"gin_web_gin/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

// 假设 global 包定义如下（根据实际项目调整）
var global struct {
	CONFIG any // 可以替换为具体结构体
}

// 替换 getConfigPath() 以便控制测试路径
func getConfigPath() string {
	return os.Getenv("TEST_CONFIG_PATH")
}

// 创建临时配置文件
func createTempConfigFile(t *testing.T, content string) string {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	err := os.WriteFile(path, []byte(content), 0644)
	require.NoError(t, err)
	return path
}

// TC01: 正常情况 - 配置文件有效
func TestViper_Success(t *testing.T) {
	configContent := `
	app:
  	name: test-app
	port: 8080
`
	path := createTempConfigFile(t, configContent)
	t.Setenv("TEST_CONFIG_PATH", path)

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("unexpected panic: %v", r)
		}
	}()

	v := core.Viper()
	assert.NotNil(t, v)
	assert.Equal(t, "test-app", v.GetString("app.name"))
}
