package main

import (
	"changeme/hosts"
	"context"
	"fmt"
	"log"
)

// App struct
type App struct {
	ctx     context.Context
	myHosts *hosts.MyHosts
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		myHosts: hosts.NewMyHosts(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	err := a.myHosts.Read()
	if err != nil {
		log.Println(err.Error())
	}
	a.myHosts.Split()
	a.myHosts.Pretty()
	a.myHosts.Print()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetHostList
func (a *App) GetHostList() map[string]map[string]*hosts.Host {
	return a.myHosts.Hosts
}

// GetHosts
func (a *App) GetHosts() string {
	return a.myHosts.HostsText
}

// SaveAllHosts
func (a *App) SaveAllHosts(hostsText string) string {
	a.myHosts.HostsText = hostsText
	a.myHosts.Split()
	a.myHosts.Pretty()
	err := a.myHosts.Write()
	if err != nil {
		return err.Error()
	}
	return ""
}
