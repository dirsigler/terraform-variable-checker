package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

type tfdetails struct {
	VarDefault interface{}
	VarPosFn string
	VarPosLn int
}

func (t *tfdetails) GetVars(tfDir string) (err error) {
	module, _ := tfconfig.LoadModule(tfDir)
	
	tfconstraints := module.Variables
	for _, v := range tfconstraints {
		t.VarDefault = v.Default
		t.VarPosFn = v.Pos.Filename
		t.VarPosLn = v.Pos.Line
	}
	return err
}


func main() {
	var dir string
	flag.StringVar(&dir, "path", ".", "Path to the root Terraform module.")
	
	flag.Parse()
	var t tfdetails
	err := t.GetVars(dir)
	if err != nil {
		log.Fatal(err)
	}
	
	if t.VarDefault == nil {
		fmt.Printf("Empty default for variable %v in File %v at Line %v", t.VarDefault, t.VarPosFn, t.VarPosLn)
	}
}