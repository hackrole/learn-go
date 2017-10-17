package main

import (
	"database/sql"
	"io/util"
	"log"

	"github.com/RichardKnop/go-fixtures"
	"github.com/urfave/cli"

	_ "github.com/lib/pq"
)

var (
	cliApp *cli.App
)


func init(){
	cliApp = cli.NewApp()
	cliApp.Name = "your-project"
	cliApp.Usage = "Project's usage"
	cliApp.Author = "your name"
	cliApp.Email = "your@email"
	cliApp.Version = "0.0.0"
}

func main() {
	db, err := sql.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil{
		log.Fatal(err)
	}
	defer db.Close()

	cliApp.Commands = []cli.Command{
		{
			Name: "loaddata",
			Usage: "load data from fixture",
			Action: func(c *cli.Context) error{
				data, err := ioutil.ReadFile(c.Args().First())
				if err != nil {
					return err
				}

				if err := fixtures.Load(data, db, "postgres"); err != nil{
					return err
				}
			},
		},
		{
			Name: "runserver",
			Usage: "run web server",
			Action: func(c *cli.Context) error{
				// run you webserver here
				return nil
			}
		}
	}
	cliApp.Run(os.Args)
}
