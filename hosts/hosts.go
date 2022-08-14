package hosts

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type Host struct {
	Show      bool   `json:"show"`
	IP        string `json:"ip"`
	Hostname  string `json:"hostname"`
	GroupName string `json:"group_name"`
}

type Group struct {
	GroupName string           `json:"group_name"`
	Show      bool             `json:"show"`
	List      map[string]*Host `json:"list"`
}

type Hosts map[string]Group

type MyHosts struct {
	Path      string
	HostsText string
	Hosts     Hosts
}

type M map[string]string

type HostUnique struct {
	Show      bool
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
		Hosts:     Hosts{},
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
	f.Hosts = Hosts{}
	hosts := strings.Split(f.HostsText, "\n")
	for _, host := range hosts {
		re := regexp.MustCompile(`^([# ]*)([0-9]{0,3}\.[0-9]{0,3}\.[0-9]{0,3}\.[0-9]{0,3}|[a-zA-Z0-9:]{2,})[ ]+([a-zA-Z0-9. \-]*)+([#]+[ ]*(.*?))?$`)
		res := re.FindStringSubmatch(host)
		if len(res) == 0 {
			continue
		}
		show := !strings.Contains(res[1], "#")
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
					f.Hosts[groupName] = Group{
						GroupName: groupName,
						Show:      true,
						List:      map[string]*Host{},
					}
				}
				if !show {
					groupInfo := f.Hosts[groupName]
					groupInfo.Show = false
					f.Hosts[groupName] = groupInfo
				}
				f.Hosts[groupName].List[hostname] = &Host{
					Show:      show,
					IP:        res[2],
					Hostname:  hostname,
					GroupName: groupName,
				}
				log.Println(Host{
					Show:      show,
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
	for _, hosts := range f.Hosts {
		for hostname, host := range hosts.List {
			var key string
			if host.Show {
				key = fmt.Sprintf("%s_", hostname)
			} else {
				key = fmt.Sprintf("%s_#", hostname)
			}
			if _, ok := allHostsUnique[key]; !ok {
				allHostsUnique[key] = HostUnique{
					Show:      host.Show,
					IP:        host.IP,
					Hostname:  hostname,
					GroupName: M{hosts.GroupName: hosts.GroupName},
				}
			} else {
				allHostsUnique[key].GroupName[hosts.GroupName] = hosts.GroupName
			}
		}
	}
	for _, row := range allHostsUnique {
		groupName := strings.Join(row.GroupName.toSlice(), " # ")
		if row.Show {
			f.HostsText += fmt.Sprintf("%s %s # %s\n", row.IP, row.Hostname, groupName)
		} else {
			f.HostsText += fmt.Sprintf("# %s %s # %s\n", row.IP, row.Hostname, groupName)
		}
	}
}

func (f *MyHosts) Add(groupName string, ip string, hostname string) {
	if _, ok := f.Hosts[groupName]; !ok {
		f.Hosts[groupName] = Group{
			GroupName: groupName,
			Show:      true,
			List:      map[string]*Host{},
		}
	}
	f.Hosts[groupName].List[hostname] = &Host{
		Show:      true,
		IP:        ip,
		Hostname:  hostname,
		GroupName: groupName,
	}
}

func (f *MyHosts) SwitchByGroupName(groupName string, show bool) {
	if _, ok := f.Hosts[groupName]; !ok {
		return
	}
	groupInfo := f.Hosts[groupName]
	groupInfo.Show = show
	for hostname := range groupInfo.List {
		groupInfo.List[hostname].Show = show
	}
	f.Hosts[groupName] = groupInfo
}

func (f *MyHosts) SwitchByHostname(groupName string, hostname string, show bool) {
	if _, ok := f.Hosts[groupName]; !ok {
		return
	}
	groupInfo := f.Hosts[groupName]
	if _, ok := groupInfo.List[hostname]; !ok {
		return
	}
	groupInfo.List[hostname].Show = show
	if len(groupInfo.List) == 1 {
		groupInfo.Show = show
	}
	f.Hosts[groupName] = groupInfo
}
