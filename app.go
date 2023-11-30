package main

import (
	"changeme/hosts"
	"context"
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

// GetList .
func (a *App) GetList() hosts.List {
	return a.myHosts.List
}

// GetListByGroup .
func (a *App) GetListByGroup() hosts.ListByGroup {
	return a.myHosts.ListByGroup
}

// GetHostsText .
func (a *App) GetHostsText() string {
	return a.myHosts.HostsText
}

// SaveAllHosts .
func (a *App) SaveAllHosts(hostsText string) string {
	a.myHosts.HostsText = hostsText
	a.myHosts.Split()
	a.myHosts.Pretty()

	if err := a.myHosts.Write(); err != nil {
		return err.Error()
	}
	return ""
}

// AddHost .
func (a *App) AddHost(groupName string, ip string, hostname string) string {
	a.myHosts.Add(groupName, ip, hostname)
	a.myHosts.Pretty()

	if err := a.myHosts.Write(); err != nil {
		return err.Error()
	}
	return ""
}

// DeleteHost .
func (a *App) DeleteHost(groupName string, hostNameID int) string {
	a.myHosts.Delete(groupName, hostNameID)
	a.myHosts.Pretty()

	if err := a.myHosts.Write(); err != nil {
		return err.Error()
	}
	return ""
}

// SetGroupName .
func (a *App) SetGroupName(hostNameID int, groupName string) string {
	a.myHosts.SetGroupNameByHostNameId(groupName, hostNameID)
	a.myHosts.Pretty()

	if err := a.myHosts.Write(); err != nil {
		return err.Error()
	}
	return ""
}

// SwitchByGroupName .
func (a *App) SwitchByGroupName(groupName string, show bool) string {
	a.myHosts.SwitchByGroupName(groupName, show)
	a.myHosts.Pretty()

	if err := a.myHosts.Write(); err != nil {
		return err.Error()
	}
	return ""
}

// GetAllGroupNames .
func (a *App) GetAllGroupNames() []string {
	return a.myHosts.GetAllGroupNames()
}

// SwitchByHostnameId .
func (a *App) SwitchByHostnameId(groupName string, hostNameID int, show bool) string {
	a.myHosts.SwitchByHostNameId(groupName, hostNameID, show)
	a.myHosts.Pretty()

	if err := a.myHosts.Write(); err != nil {
		return err.Error()
	}
	return ""
}
