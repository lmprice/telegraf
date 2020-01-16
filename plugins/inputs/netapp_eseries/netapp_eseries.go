package netapp_eseries

import (
	"fmt"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	//"io"
	"net/url"
	//"io/ioutil"
	"net/http"
	//"time"
)

type NetappESeriesPlugin struct {
	Address string
	Username string
	Password string
	systemFilter string
	//HostMappings bool
}

func (counters *NetappESeriesPlugin) Description() string {
	return "abc"
}

func (counters *NetappESeriesPlugin) Gather(acc telegraf.Accumulator) error {
	//start := time.Now()
	//sema := make(chan struct{})
	//tr := &http.Transport{IdleConnTimeout: 10 * time.Second,}

	//client := http.Client{Transport: tr}
	http.NewRequest("GET", counters.Address, nil )
	target, err := url.Parse(counters.Address)
	host := target.Host
	if err == nil {
		fmt.Println(target)
	}
	return nil
}

var sampleConfig = `
  ## Address of the Web Services instance you want to collect statistics from.
  # This can be a 
  address = "http://this-is-my-server:8080"
  
  ## The username to authenticate to the web services instance
  username = "admin"
  ## The password to authenticate to the web services instance
  password = ""

`

func (*NetappESeriesPlugin) SampleConfig() string {
	return sampleConfig
}

func init() {
	inputs.Add("netapp_eseries", func() telegraf.Input {
		return &NetappESeriesPlugin{}
	})
}

func main (){
	x := NetappESeriesPlugin{}
	x.Gather(nil )
}