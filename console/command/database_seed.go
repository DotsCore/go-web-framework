package command

import (
	"fmt"
	"reflect"
	"strings"

	gwf "github.com/RobyFerro/go-web-framework"
	"github.com/jinzhu/gorm"
)

// Seeder will handle database seeding.
type Seeder struct {
	Signature   string
	Description string
}

// Register this command
func (c *Seeder) Register() {
	c.Signature = "database:seed <name>"
	c.Description = "Execute database seeder"
}

// Run this command
// Todo: Improve this method to run a single seeder
func (c *Seeder) Run(kernel *gwf.HttpKernel, args string, console map[string]interface{}) {
	err := kernel.Container.Invoke(func(db *gorm.DB) {
		models := kernel.Models

		if len(args) > 0 {
			extractSpecificModel(args, &models)
		}

		seed(models, db)
	})

	if err != nil {
		gwf.ProcessError(err)
	}
}

// Extract the specified models from model list
func extractSpecificModel(name string, models *[]interface{}) {
	var newModels []interface{}
	for _, m := range *models {
		modelName := reflect.TypeOf(m).Name()

		if strings.ToLower(name) == strings.ToLower(modelName) {
			newModels = append(newModels, m)
			break
		}
	}

	*models = newModels
}

// Parse model register and run every seed
func seed(models []interface{}, db *gorm.DB) {
	for _, m := range models {
		fmt.Printf("\nCreating items for model %s...\n", reflect.TypeOf(m).Name())
		v := reflect.ValueOf(m)
		method := v.MethodByName("Seed")
		method.Call([]reflect.Value{reflect.ValueOf(db)})

		fmt.Printf("Success!\n")
	}

	fmt.Println("Seeding complete!")
}
