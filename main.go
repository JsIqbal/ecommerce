package main

import (
	_ "github.com/jsiqbal/ecommerce/docs"

	"github.com/jsiqbal/ecommerce/cmd"
)

// @title Ecommerce Assessment by IQBAL HOSSAIN
// @version 1.0
// @description This is the Assessment Ecomerce server. You can Follow Iqbal Hossain at https://github.com/JsIqbal

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:5000
// @BasePath /
// @query.collection.format multi
func main() {
	cmd.Execute()
}
