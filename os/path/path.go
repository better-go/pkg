package path

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/better-go/pkg/log"
)

var (
	defaultProjectName = "repo-name" // 项目根文件夹名称
)

// 获得的是运行时路径, 非项目路径
func RuntimePath() (dir string, err error) {
	// os.Executable() // 运行时, 非正确项目路径
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

// 当前路径:
func CurrentDir() (dir string, err error) {
	return os.Getwd() // 项目路径
}

// 项目根目录:
func ProjectRoot(projectName string) string {
	if projectName == "" {
		projectName = defaultProjectName
	}

	cur, err := CurrentDir()
	if err != nil {
		return ""
	}

	// 路径切分:
	paths := strings.Split(cur, projectName)
	log.Infof("cur=%v, paths: %+v", cur, paths)
	return paths[0] + projectName
}

// 配置文件路径:
func ConfigPath(rootName string, dir ...string) string {
	root := ProjectRoot(rootName)

	// file:
	fp := path.Join(dir...)
	return path.Join(root, fp)
}
