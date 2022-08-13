package hosts

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type Host struct {
	Hide      bool   `json:"hide"`
	IP        string `json:"ip"`
	Hostname  string `json:"hostname"`
	GroupName string `json:"group_name"`
}

type MyHosts struct {
	Path      string                      `json:"path"`
	HostsText string                      `json:"hosts_text"`
	Hosts     map[string]map[string]*Host `json:"hosts"`
}

func NewMyHosts() *MyHosts {
	return &MyHosts{
		Path:      getHostPath(),
		HostsText: "",
		Hosts:     map[string]map[string]*Host{},
	}
}

func (f *MyHosts) Read() error {
	iot, err := ioutil.ReadFile(f.Path)
	if err != nil {
		return err
	}
	f.HostsText = string(iot)
	f.HostsText = strings.ReplaceAll(f.HostsText, "	", " ")
	return nil
}

func (f *MyHosts) Write() error {
	return ioutil.WriteFile(f.Path, []byte(f.HostsText), 666)
}

func (f *MyHosts) Print() {
	log.Println(f.HostsText)
}

func (f *MyHosts) Split() {
	f.Hosts = map[string]map[string]*Host{}
	hosts := strings.Split(f.HostsText, "\n")
	for _, host := range hosts {
		re := regexp.MustCompile(`^([# ]*)([0-9]{0,3}\.[0-9]{0,3}\.[0-9]{0,3}\.[0-9]{0,3}|[a-zA-Z0-9:]{2,})[ ]+([a-zA-Z0-9. \-]*)+([#]+[ ]*(.*?))?$`)
		res := re.FindStringSubmatch(host)
		if len(res) == 0 {
			continue
		}
		if _, ok := f.Hosts[res[5]]; !ok {
			f.Hosts[res[5]] = map[string]*Host{}
		}
		hide := strings.Contains(res[1], "#")
		for _, hostname := range strings.Split(res[3], " ") {
			hostname = strings.TrimSpace(hostname)
			if hostname == "" {
				continue
			}
			if !strings.Contains(res[2], ":") && !strings.Contains(res[2], ".") {
				continue
			}
			f.Hosts[res[5]][hostname] = &Host{
				Hide:      hide,
				IP:        res[2],
				Hostname:  hostname,
				GroupName: res[5],
			}
			log.Println(Host{
				Hide:      hide,
				IP:        res[2],
				Hostname:  hostname,
				GroupName: res[5],
			})
		}
	}
}

func (f *MyHosts) Pretty() {
	f.HostsText = ""
	for groupName, hosts := range f.Hosts {
		for hostname, host := range hosts {
			if host.Hide {
				f.HostsText += fmt.Sprintf("# %s %s # %s\n", host.IP, hostname, groupName)
			} else {
				f.HostsText += fmt.Sprintf("%s %s # %s\n", host.IP, hostname, groupName)
			}
		}
	}
}
