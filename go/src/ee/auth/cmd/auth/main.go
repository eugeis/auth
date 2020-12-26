package main

import (
	appAuth "ee/auth/app"
	"github.com/go-ee/utils/eh/app"
	"github.com/go-ee/utils/eh/app/filestore"
	"github.com/go-ee/utils/eh/app/memory"
	"github.com/go-ee/utils/eh/app/mongo"
	"github.com/go-ee/utils/lg"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
)

func main() {
	const productName = "Auth"

	var name, serverAddress, mongoUrl, targetFile, folderEventStore string
	var debug, secure bool
	var serverPort int

	commonFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        "name",
			Aliases:     []string{"n"},
			Usage:       "name of the app, used for backend data",
			Value:       "Auth",
			Destination: &name,
		}, &cli.BoolFlag{
			Name:        "secure",
			Aliases:     []string{"s"},
			Usage:       "activate secure mode",
			Destination: &secure,
		}, &cli.StringFlag{
			Name:        "address",
			Aliases:     []string{"a"},
			Usage:       "server address",
			Value:       "",
			Destination: &serverAddress,
		}, &cli.IntFlag{
			Name:        "port",
			Usage:       "server port",
			Value:       7070,
			Destination: &serverPort,
		}, &cli.BoolFlag{
			Name:        "debug",
			Aliases:     []string{"d"},
			Destination: &debug,
			Usage:       "Enable debug log level",
		},
	}

	runner := cli.NewApp()
	runner.Usage = name
	runner.Version = "1.0"

	lg.LogrusTimeAsTimestampFormatter()

	runner.Before = func(c *cli.Context) (err error) {
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}
		logrus.Debugf("execute %v", c.Command.Name)
		return
	}

	runner.Commands = []*cli.Command{
		{
			Name:  "mongo",
			Usage: "Start server with MongoDB backend",
			Flags: append(commonFlags, &cli.StringFlag{
				Name:        "mongo",
				Aliases:     []string{"m"},
				Usage:       "url of the MongoDB instance",
				Value:       "localhost",
				Destination: &mongoUrl,
			}),
			Action: func(c *cli.Context) (err error) {
				Auth := appAuth.NewAuth(mongo.NewAppMongo(
					&app.AppInfo{
						AppName:     name,
						ProductName: productName,
					}, &app.ServerConfig{
						ServerAddress: serverAddress,
						ServerPort:    serverPort,
					}, secure, mongoUrl))
				err = Auth.Start()
				return
			},
		}, {
			Name:  "memory",
			Usage: "Start server with memory backend",
			Flags: commonFlags,
			Action: func(c *cli.Context) (err error) {
				Auth := appAuth.NewAuth(memory.NewAppMemory(&app.AppInfo{
					AppName:     name,
					ProductName: productName,
				}, &app.ServerConfig{
					ServerAddress: serverAddress,
					ServerPort:    serverPort,
				}, secure))
				err = Auth.Start()
				return
			},
		}, {
			Name:  "fileStore",
			Usage: "Start server with file based event store",
			Flags: append(commonFlags, &cli.StringFlag{
				Name:        "fileStore",
				Aliases:     []string{"f"},
				Usage:       "folder for events",
				Value:       "store",
				Destination: &folderEventStore,
			}),
			Action: func(c *cli.Context) (err error) {
				Auth := appAuth.NewAuth(
					filestore.NewAppFileStore(&app.AppInfo{
						AppName:     name,
						ProductName: productName,
					}, &app.ServerConfig{
						ServerAddress: serverAddress,
						ServerPort:    serverPort,
					}, secure, folderEventStore))
				err = Auth.Start()
				return
			},
		},
		{
			Name:  "markdown",
			Usage: "Generate markdown help file",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "target",
					Aliases:     []string{"t"},
					Usage:       "Markdown target file name to generate",
					Value:       "site.md",
					Destination: &targetFile,
				},
			},
			Action: func(c *cli.Context) (err error) {
				if markdown, err := runner.ToMarkdown(); err == nil {
					err = ioutil.WriteFile(targetFile, []byte(markdown), 0)
				} else {
					logrus.Infof("%v", err)
				}
				return
			},
		},
	}

	if err := runner.Run(os.Args); err != nil {
		logrus.Infof("%v", err)
	}
}
