package pensamientos

import (
	sys "github.com/connorvanderhook/pienso/system/configuration"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
)

var Cmd = cobra.Command{
	Use:   "pienso",
	Short: "daemon to serve web requests",
}

var flagDevelopment = true

// Application holds global variables to be used in each request
type Application struct {
	Configuration *sys.Configuration
	DB            *mongo.Session
}

func newApplication(templatePath string) *Application {
	c := &sys.Configuration{
		PublicPath:   "public",
		TemplatePath: templatePath,
		// DatabaseURI:   mongo.DialWithTimeout("", 0),
		IsDevelopment: flagDevelopment,
	}
	return &Application{Configuration: c}
}

func init() {
	Cmd.Run = command
}

func command(cmd *cobra.Command, list []string) {

}
