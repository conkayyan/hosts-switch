package main

import (
	"context"
	"fmt"
	"hosts-switch/hosts"
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

// GetInUseHostsText .
func (a *App) GetInUseHostsText() string {
	return a.myHosts.InUseHostsText
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

// SaveAllInUseHosts .
func (a *App) SaveAllInUseHosts(hostsText string) string {
	a.myHosts.HostsText = a.myHosts.NoInUseHostsText + "\n" + hostsText
	a.myHosts.Split()
	a.myHosts.Pretty()

	fmt.Println(a.myHosts.HostsText)

	if err := a.myHosts.Write(); err != nil {
		return err.Error()
	}
	return ""
}

// SaveAllGroupHosts .
func (a *App) SaveAllGroupHosts(groupName, hostsText string) string {
	a.myHosts.SetGroup(groupName, hostsText)
	a.myHosts.Pretty()
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
	a.myHosts.Split()
	a.myHosts.Pretty()

	if err := a.myHosts.Write(); err != nil {
		return err.Error()
	}
	return ""
}

// SaveAddHostsText .
func (a *App) SaveAddHostsText(hostsText string) string {
	a.myHosts.HostsText = a.myHosts.HostsText + "\n" + hostsText
	a.myHosts.Split()
	a.myHosts.Pretty()

	fmt.Println(a.myHosts.HostsText)

	if err := a.myHosts.Write(); err != nil {
		return err.Error()
	}
	return ""
}

// DeleteHost .
func (a *App) DeleteHost(groupName string, hostNameID int) string {
	a.myHosts.Delete(groupName, hostNameID)
	a.myHosts.Pretty()
	a.myHosts.Split()
	a.myHosts.Pretty()

	if err := a.myHosts.Write(); err != nil {
		return err.Error()
	}
	return ""
}

// SetGroupName .
func (a *App) SetGroupName(oldGroupName, groupName string) string {
	if oldGroupName == groupName {
		return ""
	}
	a.myHosts.SetGroupNameByOldGroupName(oldGroupName, groupName)
	a.myHosts.Pretty()
	a.myHosts.Split()

	if err := a.myHosts.Write(); err != nil {
		return err.Error()
	}
	return ""
}

// SwitchByGroupName .
func (a *App) SwitchByGroupName(groupName string, show bool) string {
	a.myHosts.SwitchByGroupName(groupName, show)
	a.myHosts.Pretty()
	a.myHosts.Split()
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
