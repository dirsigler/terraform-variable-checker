package main

import (
	"flag"
	"fmt"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

func main() {
	var dir string
	flag.StringVar(&dir, "path", ".", "Path to the root Terraform module.")

	flag.Parse()

	module, _ := tfconfig.LoadModule(dir)

	tfconstraints := module.Variables
	for i, v := range tfconstraints {

		d := &v.Default
		if *d == nil {
			fmt.Printf("The variable %v has no configured default Value!\n", i)
		}
	}
}