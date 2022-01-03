package main

import (
	"github.com/urfave/cli"
)

// NewApp creates an app with sane defaults.
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Action = cli.ShowAppHelp
	app.Name = "shop-user"
	app.Author = "Le Cong Danh"
	app.Email = "lcdanh996@gmail.com"
	app.Usage = "Shop Ecomerce"
	return app
}

var (

	// HTTPPortFlag ...
	HTTPHostFlag = cli.StringFlag{
		Name:   "http.host",
		Usage:  "HTTP Host Listen",
		EnvVar: "HTTP_HOST",
		Value:  "localhost",
	}

	// HTTPPortFlag ...
	HTTPPortFlag = cli.StringFlag{
		Name:   "http.port",
		Usage:  "HTTP Port Listen",
		EnvVar: "HTTP_PORT",
		Value:  "10010",
	}

	// GRPCPortFlag ...
	GRPCHostFlag = cli.StringFlag{
		Name:   "GRPC.host",
		Usage:  "GRPC Host Listen",
		EnvVar: "HTTP_HOST",
		Value:  "",
	}

	// GRPCPortFlag ...
	GRPCPortFlag = cli.StringFlag{
		Name:   "grpc.port",
		Usage:  "Grpc Port Listen",
		EnvVar: "GRPC_PORT",
		Value:  "8283",
	}

	// MongoHostFlag ...
	MongoHostFlag = cli.StringFlag{
		Name:   "mongo.host",
		Usage:  "Mongo host",
		EnvVar: "MONGO_HOST",
		Value:  "mongodb://localhost",
	}

	// MongoPortFlag ...
	MongoPortFlag = cli.StringFlag{
		Name:   "mongo.port",
		Usage:  "Mongo port",
		EnvVar: "MONGO_PORT",
		Value:  "27017",
	}

	// MongoDatabaseFlag ...
	MongoDatabaseFlag = cli.StringFlag{
		Name:   "mongo.database",
		Usage:  "Mongo database",
		EnvVar: "MONGO_DATABASE",
		Value:  "shop",
	}

	// MongoDatabaseUsernameFlag ...
	MongoDatabaseUsernameFlag = cli.StringFlag{
		Name:   "mongo.database.username",
		Usage:  "Mongo database username",
		EnvVar: "MONGO_DATABASE_USERNAME",
		Value:  "",
	}

	// MongoDatabasePasswordFlag ...
	MongoDatabasePasswordFlag = cli.StringFlag{
		Name:   "mongo.database.password",
		Usage:  "Mongo database.password",
		EnvVar: "MONGO_DATABASE_PASSWORD",
		Value:  "",
	}

	TokenKeyFlag = cli.StringFlag{
		Name:   "token.key",
		Usage:  "token key",
		EnvVar: "TOKEN_KEY",
		Value:  "ezexjSqGSzefFUFRxlxgMLNlPjrRWsiA",
	}
)

var (
	startCommand = cli.Command{
		Action:      Start,
		Name:        "start",
		Usage:       "Bootstrap and start worker server",
		ArgsUsage:   "<genesisPath>",
		Flags:       []cli.Flag{},
		Category:    "Crawler Worker",
		Description: `Used to start crawler worker, clone data from omada cloud`,
	}
)
