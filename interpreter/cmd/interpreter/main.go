package main

import (
	"flag"
	"fmt"

	"github.com/project-kessel/starlark-unified-schema/internal/lang"
)

func main() {
	srcDirArg := flag.String("src", "schema", "The path to the directory containing Starlark schema (.star) files.")
	flag.Parse()

	loader := lang.NewLoader(*srcDirArg)
	processor := lang.NewProcessor(loader)

	src_files := flag.Args()
	if len(src_files) > 0 {
		for _, src_file := range src_files {
			err := processor.ProcessModule(src_file, nil)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		err := processor.ProcessAllModules(nil)
		if err != nil {
			fmt.Println(err)
		}
	}
}
