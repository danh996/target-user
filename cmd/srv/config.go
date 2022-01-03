package main

import (
	"fmt"

	"github.com/urfave/cli"
)

type Config struct {
	HTTP        ConnAddress
	GRPC        ConnAddress
	ServiceName string

	Mongo    Mongo
	TokenKey string

	// for authentication

	// parameter
	Version string
	Region  string
}

type ConnAddress struct {
	Host string
	Port string
}

type Mongo struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func migrateFlags(action func(ctx *cli.Context) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		// for _, name := range ctx.FlagNames() {
		// 	ctx.GlobalSet(name, ctx.String(name))
		// }
		return action(ctx)
	}
}

func (c *ConnAddress) String() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
