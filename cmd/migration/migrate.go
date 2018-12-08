package main

import (
	"flag"
	"fmt"
	"github.com/fzerorubigd/balloon/cmd"
	"os"

	"github.com/fzerorubigd/balloon/pkg/cli"
	"github.com/fzerorubigd/balloon/pkg/initializer"
	"github.com/fzerorubigd/balloon/pkg/log"
	"github.com/fzerorubigd/balloon/pkg/migration"
	"github.com/fzerorubigd/balloon/pkg/postgres/model"
)

var (
	action = flag.String("action", "up", "up/down is supported, default is up")
	n      int
)

func main() {
	ctx := cli.Context()
	cmd.InitializeConfig()

	defer initializer.Initialize(ctx)()

	flag.Parse()
	var err error
	m := &model.Manager{}
	if *action == "up" {
		n, err = migration.Do(m, migration.Up, 0)
		fmt.Printf("\n\n%d migration is applied\n", n)
	} else if *action == "down" {
		n, err = migration.Do(m, migration.Down, 1)
		fmt.Printf("\n\n%d migration is applied\n", n)
	} else if *action == "down-all" {
		n, err = migration.Do(m, migration.Down, 0)
		fmt.Printf("\n\n%d migration is applied\n", n)
	} else if *action == "redo" {
		n, err = migration.Do(m, migration.Down, 1)
		if err == nil {
			n, err = migration.Do(m, migration.Up, 1)
		}
		fmt.Printf("\n\n%d migration is applied\n", n)

	} else if *action == "list" {
		migration.List(m, os.Stdout)
	}

	if err != nil {
		log.Fatal("Error on migration", log.Err(err))
	}
}
