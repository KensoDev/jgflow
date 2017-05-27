package jgflow

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/go-ini/ini"
	"log"
	"os/user"
)

type LoginDetails struct {
	Username   string
	Password   string
	JiraDomain string
}

type LoginService struct {
}

func (s *LoginService) Login(loginDetails LoginDetails, dir string) bool {
	err := s.SaveLoginDetails(loginDetails, dir)

	if err != nil {
		log.Fatal(err)
	}

	err, loadedDetails := s.LoadLoginDetails(dir)
	if err != nil {
		log.Fatal(err)
	}

	jiraClient, err := jira.NewClient(nil, loadedDetails.JiraDomain)
	if err != nil {
		panic(err)
	}

	res, err := jiraClient.Authentication.AcquireSessionCookie(loadedDetails.Username, loadedDetails.Password)

	if err != nil || res == false {
		fmt.Printf("Result: %v\n", res)
		panic(err)
	}

	fmt.Println("Successfully logged in to Jira, congrats!")
	return true
}

func GetUserDir() (error, string) {
	usr, err := user.Current()
	if err != nil {
		return err, ""
	}
	return nil, usr.HomeDir
}

func getFileLocation(dir string) string {
	return fmt.Sprintf("%s/%s", dir, ".jira.ini")
}

func (s *LoginService) LoadLoginDetails(dir string) (error, LoginDetails) {
	cfg, err := ini.InsensitiveLoad(getFileLocation(dir))

	if err != nil {
		return err, LoginDetails{}
	}

	section, err := cfg.GetSection("")

	if err != nil {
		return err, LoginDetails{}
	}

	username, err := section.GetKey("username")
	if err != nil {
		return err, LoginDetails{}
	}
	password, err := section.GetKey("password")
	if err != nil {
		return err, LoginDetails{}
	}
	domain, err := section.GetKey("domain")
	if err != nil {
		return err, LoginDetails{}
	}
	loginDetails := LoginDetails{
		Username:   username.String(),
		Password:   password.String(),
		JiraDomain: domain.String(),
	}
	return nil, loginDetails
}

func (s *LoginService) SaveLoginDetails(loginDetails LoginDetails, dir string) error {
	cfg := ini.Empty()

	cfg.Section("").NewKey("username", loginDetails.Username)
	cfg.Section("").NewKey("password", loginDetails.Password)
	cfg.Section("").NewKey("domain", fmt.Sprintf("https://%s", loginDetails.JiraDomain))

	cfg.SaveTo(getFileLocation(dir))
	return nil
}
