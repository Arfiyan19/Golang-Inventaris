package main
import (
	"inventory-system/database"
	"inventory-system/routes"
)
func main() {
	database.Connect()
	r := routes.SetupRouter()
	r.Run(":8080")
}
