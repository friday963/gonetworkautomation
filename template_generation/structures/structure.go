package structures


type Config struct {
	Interfaces    Interfaces    `yaml:"interfaces"`
	Dnsdomain     string        `yaml:"dnsdomain"`
	Snmp          Snmp          `yaml:"snmp"`
	PrivAuthUsers []string      `yaml:"priv_auth_users"`
	Vlans         []Vlans       `yaml:"vlans"`
	Vrfs          []Vrfs        `yaml:"vrfs"`
	PrefixLists   []PrefixLists `yaml:"prefix_lists"`
	RouteMap      []RouteMap    `yaml:"route_map"`
}
type Logical struct {
	Name string `yaml:"name"`
	IP   string `yaml:"ip"`
}
type Physical struct {
	Name string `yaml:"name"`
	IP   string `yaml:"ip"`
}
type Interfaces struct {
	Logical  []Logical  `yaml:"logical"`
	Physical []Physical `yaml:"physical"`
}
type Snmp struct {
	Users    []string `yaml:"users"`
	Location string   `yaml:"location"`
}
type Vlans struct {
	Name   string `yaml:"name"`
	Number int    `yaml:"number"`
}
type Vrfs struct {
	Name string `yaml:"name"`
}
type PrefixLists struct {
	Name       string `yaml:"name"`
	Seq        int    `yaml:"seq"`
	PermitDeny string `yaml:"permit_deny"`
	Prefix     string `yaml:"prefix"`
	Mask       int    `yaml:"mask"`
	LeGe       string `yaml:"le_ge"`
	LeGeMask   int    `yaml:"le_ge_mask"`
}
type RouteMap struct {
	Name       string `yaml:"name"`
	PrefixList string `yaml:"prefix_list"`
}