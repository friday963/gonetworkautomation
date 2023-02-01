package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/scrapli/scrapligo/driver/options"
	"github.com/scrapli/scrapligo/platform"
	cfg "github.com/scrapli/scrapligocfg"
	"gopkg.in/yaml.v3"
)

type Maintenance struct {
	PreMaintenance  []string `yaml:"pre_maintenance"`
	Maintenance     []string `yaml:"maintenance"`
	PostMaintenance []string `yaml:"post_maintenance"`
}

func checker(err error)  {
	if err != nil {
		fmt.Println(err)
	}
}

func main()  {

	username := flag.String("username", "anexampleusername", "SSH username")
	password := flag.String("password", "anexmaplepass", "SSH password")
	flag.Parse()
	data, err := os.ReadFile("inputs.yml")
	checker(err)
	var maintenance Maintenance
	err = yaml.Unmarshal(data, &maintenance)
	checker(err)
	p,err := platform.NewPlatform(
		"arista_eos",
		"clab-ceostopo-ceos1",
		options.WithAuthNoStrictKey(),
		options.WithAuthUsername(*username),
		options.WithAuthPassword(*password),
	)
	checker(err)
	d,err := p.GetNetworkDriver()
	checker(err)
	err = d.Open()
	defer d.Close()
	fmt.Println("Starting pre-maintenance")
	for _,x := range maintenance.PreMaintenance {
		resp, err := d.SendCommand(x)
		fmt.Println(resp.Result + "\n\n")
		checker(err)
	}
	conf, err := cfg.NewCfg(d, "arista_eos")
	err = conf.Prepare()
	checker(err)
	_, err = conf.LoadConfigFromFile(maintenance.Maintenance[0], false)
	// fmt.Println(resp.Result)
	for _,x := range maintenance.PostMaintenance{
		resp, err := d.SendCommand(x)
		fmt.Println(resp.Result + "\n\n")
		checker(err)
	}

}
