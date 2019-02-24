package pensamientos

import (
	"github.com/spf13/cobra"
	"system"
)

var Cmd = cobra.Command{
	Use: "pienso"
	Short: "daemon to serve web requests"
}

// Application holds global variables to be used in each request
type Application struct {
	Configuration          *Configuration
	// Template               *template.Template
	DB              	*mgo.Session
}

func newApplication(templatePath string) *Application {
	c := &system.Configuration{
		PublicPath:    "public",
		TemplatePath:  templatePath,
		DatabaseURI:   mongodb.URI(Cmd),
		IsDevelopment: flagDevelopment,
	}
	return &Application{Configuration: c}
}

func init() {
	Cmd.run = command
}