package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yamamoto-febc/terraform-provider-gmailfilter/gmailfilter"
	"github.com/yamamoto-febc/terraform-provider-gmailfilter/tools/tfdocgen"
)

const (
	templateDir = "templates"
	exampleDir  = "examples/website"
)

func main() {
	if len(os.Args) < 2 || 2 < len(os.Args) {
		fmt.Println("Usage: gen-gmailfilter-docs <destination-directory>")
		os.Exit(1)
	}

	destination := os.Args[1]
	provider := tfdocgen.Provider{
		Name:              "gmailfilter",
		TerraformProvider: gmailfilter.Provider(),
		DisplayNameFunc: func(name string) string {
			d, ok := definitions[name]
			if !ok {
				return name
			}
			return d.displayName
		},
		CategoryNameFunc: func(name string) string {
			d, ok := definitions[name]
			if !ok {
				return ""
			}
			return d.category
		},
		CategoriesFunc: func() []string {
			return categories
		},
	}
	if err := provider.GenerateDocs(templateDir, exampleDir, destination); err != nil {
		log.Fatal(err)
	}
}

type definition struct {
	displayName string
	category    string
}

const (
	CategoryCommon = "Gmail Filter Settings"
)

var categories = []string{
	CategoryCommon,
}

var definitions = map[string]definition{
	"gmailfilter": {
		displayName: "Gmail Filter",
	},
	"gmailfilter_filter": {
		displayName: "Filter",
		category:    CategoryCommon,
	},
	"gmailfilter_label": {
		displayName: "Label",
		category:    CategoryCommon,
	},
}
