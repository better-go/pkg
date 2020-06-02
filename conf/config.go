package conf

import (
	path2 "path"

	"pkg/log"
	"pkg/os/path"
)

var (
	defaultAppFolder = "app"
	// 默认配置路径:
	defaultConfigDir = "configs/configs.toml"
)

func LoadConfig(dist interface{}, rootName string, appName string) error {
	filepath := path2.Join(defaultAppFolder, appName, defaultConfigDir)

	// file path:
	fp := path.ConfigPath(rootName, filepath)
	log.Infof("current config path: %v", fp)

	// parse:
	if err := TomlInit(fp); err != nil {
		log.Errorf("toml config parse error, fp=%v, err=%v", fp, err)
		return err
	}

	// scan:
	if err := TomlScan(&dist, ""); err != nil {
		log.Errorf("toml config scan error, fp=%v, err=%v", fp, err)
		return err
	}

	log.Debugf("config map: %+v", &dist)

	return nil
}
