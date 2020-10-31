package ipseach

import (
	"github.com/ipipdotnet/ipdb-go"
	"y7utils/src/ipseach/iptocity"
)

type ipObject struct {
	Ipdbobj   *ipdb.City
	Ip2region *iptocity.Iptocity
}

func NewIpObj(ip2regionpath string, ipdbpath string) *ipObject {
	var ipObj ipObject
	if ip2regionpath != "" {
		ipObj.Ip2region, _ = iptocity.New(ip2regionpath)
	}
	if ipdbpath != "" {
		ipObj.Ipdbobj, _ = ipdb.NewCity(ipdbpath)
	}
	return &ipObj
}
func (this *ipObject) finIp(ip string) (ret map[string]string) {
	if this.Ip2region != nil {
		ipinfo, err := this.Ip2region.MemorySearch(ip)
		if err != nil {
			panic(err)
		}
		ret["country"] = ipinfo.Country
		ret["province"] = ipinfo.Province
		ret["city"] = ipinfo.City
		ret["isp"] = ipinfo.ISP
		if ret["country"] != "" && ret["province"] != "" {
			return ret
		}
	}
	if this.Ipdbobj != nil {
		ipdbinfo, err := this.Ipdbobj.FindMap(ip, "CN")
		if err != nil {
			panic(err)
		}
		ret["country"] = ipdbinfo["country_name"]
		ret["province"] = ipdbinfo["region_name"]
		ret["city"] = ipdbinfo["city_name"]
		ret["isp"] = ipdbinfo["isp_domain"]
	}
	return
}
