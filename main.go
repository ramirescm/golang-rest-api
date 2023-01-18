package main

import (
	"log"

	"github.com/ramirescm/golang-rest-api/controllers"
	"github.com/ramirescm/golang-rest-api/database"
	_ "github.com/ramirescm/golang-rest-api/docs" // o swag init nao adicionou esse linha ...
)

//	@title			Swagger Example API
//	@version		1.0
//	@termsOfService	http://swagger.io/terms/

// @contact.name 	API Support
// @contact.email 	my@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:9000
// @BasePath	/api
func main() {

	err := database.InitializeDatabase()
	if err != nil {
		log.Fatal(err)
	}
	controllers.InitializeRouter()
}
