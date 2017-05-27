package main

import (
	"github.com/kensodev/jgflow"
	"github.com/segmentio/go-prompt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "Login to your Jira instance",
			Action: func(c *cli.Context) error {
				username := prompt.String("Type your username please")
				password := prompt.PasswordMasked("Please type your password. Don't worry, nothing will show on the screen")
				jiraDomain := prompt.String("What is the instance URL for Jira (no need for https://)")
				err, dir := jgflow.GetUserDir()

				if err != nil {
					return err
				}

				loginDetails := jgflow.LoginDetails{
					Username:   username,
					Password:   password,
					JiraDomain: jiraDomain,
				}

				loginService := &jgflow.LoginService{}
				loginService.Login(loginDetails, dir)

				return nil
			},
		},
	}

	app.Run(os.Args)
}
