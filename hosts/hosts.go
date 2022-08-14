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

type M map[string]string

type HostUnique struct {
	Hide      bool
	IP        string
	Hostname  string
	GroupName M
}

func (m *M) toSlice() []string {
	s := make([]string, 0, len(*m))
	for _, v := range *m {
		s = append(s, v)
	}
	return s
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
		hide := strings.Contains(res[1], "#")
		for _, hostname := range strings.Split(res[3], " ") {
			hostname = strings.TrimSpace(hostname)
			if hostname == "" {
				continue
			}
			if !strings.Contains(res[2], ":") && !strings.Contains(res[2], ".") {
				continue
			}
			var groupNames = map[string]string{}
			for _, groupName := range strings.Split(res[5], "#") {
				groupName = strings.TrimSpace(groupName)
				if groupName == "" {
					continue
				}
				groupNames[groupName] = groupName
			}
			if len(groupNames) == 0 {
				groupNames = map[string]string{"uncategorized": "uncategorized"}
			}
			for _, groupName := range groupNames {
				if _, ok := f.Hosts[groupName]; !ok {
					f.Hosts[groupName] = map[string]*Host{}
				}
				f.Hosts[groupName][hostname] = &Host{
					Hide:      hide,
					IP:        res[2],
					Hostname:  hostname,
					GroupName: groupName,
				}
				log.Println(Host{
					Hide:      hide,
					IP:        res[2],
					Hostname:  hostname,
					GroupName: groupName,
				})
			}
		}
	}
}

func (f *MyHosts) Pretty() {
	f.HostsText = ""
	var allHostsUnique = map[string]HostUnique{}
	for groupName, hosts := range f.Hosts {
		for hostname, host := range hosts {
			var key string
			if host.Hide {
				key = fmt.Sprintf("%s_#", hostname)
			} else {
				key = fmt.Sprintf("%s_", hostname)
			}
			if _, ok := allHostsUnique[key]; !ok {
				allHostsUnique[key] = HostUnique{
					Hide:      host.Hide,
					IP:        host.IP,
					Hostname:  hostname,
					GroupName: M{groupName: groupName},
				}
			} else {
				allHostsUnique[key].GroupName[groupName] = groupName
			}
		}
	}
	for _, row := range allHostsUnique {
		groupName := strings.Join(row.GroupName.toSlice(), " # ")
		if row.Hide {
			f.HostsText += fmt.Sprintf("# %s %s # %s\n", row.IP, row.Hostname, groupName)
		} else {
			f.HostsText += fmt.Sprintf("%s %s # %s\n", row.IP, row.Hostname, groupName)
		}
	}
}

func (f *MyHosts) Add(groupName string, ip string, hostname string) {
	if _, ok := f.Hosts[groupName]; !ok {
		f.Hosts[groupName] = map[string]*Host{}
	}
	f.Hosts[groupName][hostname] = &Host{
		Hide:      false,
		IP:        ip,
		Hostname:  hostname,
		GroupName: groupName,
	}
}
