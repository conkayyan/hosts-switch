package hosts

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type Host struct {
	ID        int    `json:"id"`
	Show      bool   `json:"show"`
	IP        string `json:"ip"`
	Hostname  string `json:"hostname"`
	GroupName string `json:"group_name"`
}

type List map[int]*Host

type Group struct {
	GroupName string `json:"group_name"`
	Show      bool   `json:"show"`
	List      List   `json:"list"`
}

type ListByGroup map[string]Group

type MyHosts struct {
	Path        string      `json:"path"`
	HostsText   string      `json:"hosts_text"`
	TotalNum    int         `json:"total_num"`
	List        List        `json:"list"`
	ListByGroup ListByGroup `json:"list_by_group"`
}

type M map[string]string

func (m *M) toSlice() []string {
	s := make([]string, 0, len(*m))
	for _, v := range *m {
		s = append(s, v)
	}
	return s
}

func NewMyHosts() *MyHosts {
	return &MyHosts{
		Path:        getHostPath(),
		HostsText:   "",
		TotalNum:    0,
		List:        List{},
		ListByGroup: ListByGroup{},
	}
}

func (f *MyHosts) Read() error {
	iot, err := os.ReadFile(f.Path)
	if err != nil {
		return err
	}
	f.HostsText = string(iot)
	f.HostsText = strings.ReplaceAll(f.HostsText, "	", " ")
	return nil
}

func (f *MyHosts) Write() error {
	return os.WriteFile(f.Path, []byte(f.HostsText), 666)
}

func (f *MyHosts) Print() {
	log.Println(f.HostsText)
}

func (f *MyHosts) Split() {
	f.ListByGroup = ListByGroup{}
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
				if _, ok := f.ListByGroup[groupName]; !ok {
					f.ListByGroup[groupName] = Group{
						GroupName: groupName,
						Show:      true,
						List:      map[int]*Host{},
					}
				}
				if !show {
					groupInfo := f.ListByGroup[groupName]
					groupInfo.Switch(false)
					f.ListByGroup[groupName] = groupInfo
				}
				f.TotalNum++
				f.ListByGroup[groupName].List[f.TotalNum] = &Host{
					ID:        f.TotalNum,
					Show:      show,
					IP:        res[2],
					Hostname:  hostname,
					GroupName: groupName,
				}
				f.List[f.TotalNum] = &Host{
					ID:        f.TotalNum,
					Show:      show,
					IP:        res[2],
					Hostname:  hostname,
					GroupName: groupName,
				}
				log.Println(Host{
					ID:        f.TotalNum,
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
	for _, group := range f.ListByGroup {
		for _, row := range group.List {
			if row.Show {
				f.HostsText += fmt.Sprintf("%s %s # %s\n", row.IP, row.Hostname, group.GroupName)
			} else {
				f.HostsText += fmt.Sprintf("# %s %s # %s\n", row.IP, row.Hostname, group.GroupName)
			}
		}
	}
}

func (f *MyHosts) Add(groupName string, ip string, hostname string) {
	f.TotalNum++
	if _, ok := f.ListByGroup[groupName]; !ok {
		f.ListByGroup[groupName] = Group{
			GroupName: groupName,
			Show:      true,
			List:      map[int]*Host{},
		}
	}
	f.ListByGroup[groupName].List[f.TotalNum] = &Host{
		ID:        f.TotalNum,
		Show:      true,
		IP:        ip,
		Hostname:  hostname,
		GroupName: groupName,
	}
	f.List[f.TotalNum] = &Host{
		ID:        f.TotalNum,
		Show:      true,
		IP:        ip,
		Hostname:  hostname,
		GroupName: groupName,
	}
}

func (f *MyHosts) Delete(groupName string, hostNameID int) {
	if _, ok := f.ListByGroup[groupName]; !ok {
		f.ListByGroup[groupName] = Group{
			GroupName: groupName,
			Show:      true,
			List:      map[int]*Host{},
		}
	}
	delete(f.ListByGroup[groupName].List, hostNameID)
	if len(f.ListByGroup[groupName].List) == 0 {
		delete(f.ListByGroup, groupName)
	}
}

func (f *MyHosts) SwitchByGroupName(groupName string, show bool) {
	if _, ok := f.ListByGroup[groupName]; !ok {
		return
	}
	groupInfo := f.ListByGroup[groupName]
	groupInfo.Switch(show)
	for id := range groupInfo.List {
		groupInfo.List[id].Switch(show)
		f.List[id].Switch(show)

	}
	f.ListByGroup[groupName] = groupInfo
}

func (f *MyHosts) SwitchByHostNameId(groupName string, hostNameID int, show bool) {
	if _, ok := f.ListByGroup[groupName]; !ok {
		return
	}
	groupInfo := f.ListByGroup[groupName]
	if _, ok := groupInfo.List[hostNameID]; !ok {
		return
	}
	groupInfo.List[hostNameID].Switch(show)
	groupInfo.Switch(true)
	for _, row := range groupInfo.List {
		if !row.Show {
			groupInfo.Switch(false)
		}
	}
	f.ListByGroup[groupName] = groupInfo
	f.List[hostNameID].Switch(show)
}

func (f *MyHosts) SetGroupNameByHostNameId(groupName string, hostNameID int) {
	if _, ok := f.List[hostNameID]; !ok {
		return
	}

	oldGroupName := f.List[hostNameID].GroupName
	if oldGroupName == groupName {
		return
	}

	delete(f.ListByGroup[oldGroupName].List, hostNameID)
	if len(f.ListByGroup[oldGroupName].List) == 0 {
		delete(f.ListByGroup, oldGroupName)
	} else {
		oldGroupInfo := f.ListByGroup[oldGroupName]
		oldGroupInfo.Switch(true)
		for _, row := range oldGroupInfo.List {
			if !row.Show {
				oldGroupInfo.Switch(false)
			}
		}
	}

	if _, ok := f.ListByGroup[groupName]; !ok {
		f.ListByGroup[groupName] = Group{
			GroupName: groupName,
			Show:      true,
			List:      map[int]*Host{},
		}
	}

	f.List[hostNameID].SetGroupName(groupName)
	f.ListByGroup[groupName].List[hostNameID] = f.List[hostNameID]

	groupInfo := f.ListByGroup[groupName]
	groupInfo.Switch(true)
	for _, row := range groupInfo.List {
		if !row.Show {
			groupInfo.Switch(false)
		}
	}
}

func (f *MyHosts) GetAllGroupNames() []string {
	var groupNames []string
	for _, g := range f.ListByGroup {
		groupNames = append(groupNames, g.GroupName)
	}
	return groupNames
}

func (h *Host) Switch(show bool) {
	h.Show = show
}

func (h *Host) SetGroupName(groupName string) {
	h.GroupName = groupName
}

func (g *Group) Switch(show bool) {
	g.Show = show
}
