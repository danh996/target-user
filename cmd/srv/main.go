package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var (
	app    = NewApp()
	server = new(srv)
)

func init() {
	app.Action = cli.ShowAppHelp
	app.Commands = []cli.Command{
		startCommand,
	}

	app.Flags = []cli.Flag{
		HTTPHostFlag,
		HTTPPortFlag,
		GRPCHostFlag,
		GRPCPortFlag,
		MongoHostFlag,
		MongoPortFlag,
		MongoDatabaseFlag,
		MongoDatabaseUsernameFlag,
		MongoDatabasePasswordFlag,
		TokenKeyFlag,
	}
}

// Start ...
func Start(ctx *cli.Context) error {
	if err := server.loadConfig(ctx); err != nil {
		return err
	}

	if err := server.connectMongo(); err != nil {
		return err
	}

	// if err := server.loadRedis(); err != nil {
	// 	return err
	// }

	// // if err := server.loadPubSubs(); err != nil {
	// // return err
	// // }

	// if err := server.loadLogger(); err != nil {
	// 	return err
	// }

	if err := server.loadAuthenticator(); err != nil {
		return err
	}

	if err := server.loadRepository(); err != nil {
		return err
	}

	if err := server.loadDomain(); err != nil {
		return err
	}

	if err := server.loadDelivery(); err != nil {
		return err
	}

	if err := server.loadGRPCServer(); err != nil {
		return err
	}

	server.startGRPCServer()

	return nil
}

func main() {

	if err := app.Run(os.Args); err != nil {

		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
