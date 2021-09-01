package kernel

import (
	"github.com/RobyFerro/go-web-framework/cli"
	"github.com/RobyFerro/go-web-framework/register"
)

var (
	// Commands will export all registered commands
	// The following map of interfaces expose all available method that can be used by Go-Web CLI tool.
	// The map index determines the command that you've to run to for use the relative method.
	// Example: './goweb migration:up' will run '&command.MigrationUp{}' command.
	Commands = register.CommandRegister{
		List: map[string]interface{}{
			"database:seed":      &cli.Seeder{},
			"show:commands":      &cli.ShowCommands{},
			"cmd:create":         &cli.CmdCreate{},
			"controller:create":  &cli.ControllerCreate{},
			"generate:key":       &cli.GenerateKey{},
			"middleware:create":  &cli.MiddlewareCreate{},
			"migration:create":   &cli.MigrationCreate{},
			"migration:rollback": &cli.MigrateRollback{},
			"migration:up":       &cli.MigrationUp{},
			"model:create":       &cli.ModelCreate{},
			// Here is where you've to register your custom controller
		},
	}
	SingletonServices = register.ServiceRegister{
		List: []interface{}{
			RetrieveAppConf,
			CreateSessionStore,
		},
	}
	Services = register.ServiceRegister{
		List: []interface{}{},
	}
	CommandServices = register.ServiceRegister{
		List: []interface{}{},
	}
	Models = register.ModelRegister{
		List: []interface{}{},
	}
	Controllers = register.ControllerRegister{
		List: []interface{}{},
	}
	Middlewares = register.MiddlewareRegister{
		List: []interface{}{},
	}
)
