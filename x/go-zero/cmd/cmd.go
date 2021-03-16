package cmd

import (
	"os"
	"strings"

	"github.com/better-go/pkg/log"
	"github.com/better-go/pkg/x/go-zero/option"
	"github.com/urfave/cli/v2"
)

/*

go-zero:
	cmd 服务启动器:
		serverType:
			- inner: 对内 rpc server
			- outer: 对外 http api server
			- job: 对内 job server
			- admin: 对内 admin http api sever

*/

// run server:
func Runner(inner option.Server, outer option.Server, job option.Server, admin option.Server) {
	// parse cmd args:
	r := &cli.App{
		Name:    "Queue Service",
		Version: "v0.0.1",

		//
		// 参数解析:
		//
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "serverType",
				Aliases: []string{"r"},
				Usage:   "run server type: [inner/outer/job/admin]",
				EnvVars: []string{"SERVER_TYPE", "serverType"},
			},
			&cli.StringFlag{
				Name:    "configFile",
				Aliases: []string{"f"},
				Usage:   "config file path: [configs/configs.yaml]",
				EnvVars: []string{"CONFIG_FILE", "configFile"},
				//FilePath: "configs/configs.yaml", // 会自动解析文件内容
			},
		},

		//
		// 执行动作:
		//
		Action: func(ctx *cli.Context) error {
			// 服务类型:(小写转换)
			serverType := strings.ToLower(ctx.String("serverType"))
			// 配置文件路径:(可以改成读内容)
			configFile := ctx.String("configFile")

			log.Infof("run action: serverType=%v, configFile=%v", serverType, configFile)

			// dispatch:
			switch serverType {
			case "inner":
				inner.Run(configFile) // 对内: rpc sever
			case "outer":
				outer.Run(configFile) // 对外: http api server
			case "job":
				job.Run(configFile) // 对内: job server
			case "admin":
				admin.Run(configFile) // 对内: admin api server
			}
			return nil
		},

		Before:   nil,
		After:    nil,
		Commands: nil,
	}

	log.Infof("ready to start server: cmd.args=%+v", os.Args)

	// do run:
	if err := r.Run(os.Args); err != nil {
		log.Errorf("start server failed, error: %v", err)
	}
}
