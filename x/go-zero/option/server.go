package option

// main server 启动器:
type Server interface {
	Run(configFile string)
}
