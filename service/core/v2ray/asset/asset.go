package asset

import (
	"github.com/golang/protobuf/proto"
	"github.com/muhammadmuzzammil1998/jsonc"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
	"sync"
	"time"
	v2router "v2ray.com/core/app/router"
	"v2ray.com/core/common/strmatcher"
	"v2rayA/common/files"
	"v2rayA/core/dnsPoison"
	"v2rayA/global"
)

var v2rayLocationAsset *string

func GetV2rayLocationAsset() (s string) {
	if v2rayLocationAsset != nil {
		return *v2rayLocationAsset
	}
	switch global.ServiceControlMode {
	case global.SystemctlMode, global.ServiceMode:
		p, _ := GetV2rayServiceFilePath()
		out, err := exec.Command("sh", "-c", "cat "+p+"|grep Environment=V2RAY_LOCATION_ASSET").CombinedOutput()
		if err != nil {
			break
		}
		s = strings.TrimSpace(string(out))
		s = s[len("Environment=V2RAY_LOCATION_ASSET="):]
	}
	var err error
	if s == "" {
		//by default, v2ray working directory
		s, err = GetV2rayWorkingDir()
	}
	if err != nil {
		//fine, guess one
		s = "/etc/v2ray"
	} else {
		//save the result if not by guess
		v2rayLocationAsset = &s
	}
	return
}

func GetV2rayWorkingDir() (string, error) {
	switch global.ServiceControlMode {
	case global.SystemctlMode, global.ServiceMode:
		//从systemd的启动参数里找
		p, _ := GetV2rayServiceFilePath()
		out, err := exec.Command("sh", "-c", "cat "+p+"|grep ExecStart=").CombinedOutput()
		if err != nil {
			return "", newError(string(out)).Base(err)
		}
		arr := strings.SplitN(strings.TrimSpace(string(out)), " ", 2)
		return path.Dir(strings.TrimPrefix(arr[0], "ExecStart=")), nil
	case global.UniversalMode:
		//从环境变量里找
		out, err := exec.Command("sh", "-c", "which v2ray").CombinedOutput()
		if err == nil {
			return path.Dir(strings.TrimSpace(string(out))), nil
		}
	}
	return "", newError("not found")
}

func GetV2ctlDir() (string, error) {
	d, err := GetV2rayWorkingDir()
	if err == nil {
		_, err := os.Stat(d + "/v2ctl")
		if err != nil {
			return "", err
		}
		return d, nil
	}
	out, err := exec.Command("sh", "-c", "which v2ctl").Output()
	if err != nil {
		err = newError(string(out)).Base(err)
		return "", err
	}
	return path.Dir(strings.TrimSpace(string(out))), nil
}

func IsGFWListExists() bool {
	_, err := os.Stat(GetV2rayLocationAsset() + "/LoyalsoldierSite.dat")
	if err != nil {
		return false
	}
	return true
}
func IsGeoipExists() bool {
	_, err := os.Stat(GetV2rayLocationAsset() + "/geoip.dat")
	if err != nil {
		return false
	}
	return true
}
func IsGeositeExists() bool {
	_, err := os.Stat(GetV2rayLocationAsset() + "/geosite.dat")
	if err != nil {
		return false
	}
	return true
}
func GetGFWListModTime() (time.Time, error) {
	return files.GetFileModTime(GetV2rayLocationAsset() + "/LoyalsoldierSite.dat")
}
func IsCustomExists() bool {
	_, err := os.Stat(GetV2rayLocationAsset() + "/custom.dat")
	if err != nil {
		return false
	}
	return true
}

func GetConfigBytes() (b []byte, err error) {
	b, err = ioutil.ReadFile(GetConfigPath())
	if err != nil {
		log.Println(err)
		return
	}
	b = jsonc.ToJSON(b)
	return
}

func GetConfigPath() (p string) {
	p = "/etc/v2ray/config.json"
	switch global.ServiceControlMode {
	case global.SystemctlMode, global.ServiceMode:
		//从systemd的启动参数里找
		pa, _ := GetV2rayServiceFilePath()
		out, e := exec.Command("sh", "-c", "cat "+pa+"|grep ExecStart=").CombinedOutput()
		if e != nil {
			return
		}
		pa = strings.TrimSpace(string(out))[len("ExecStart="):]
		indexConfigBegin := strings.Index(pa, "-config")
		if indexConfigBegin == -1 {
			return
		}
		indexConfigBegin += len("-config") + 1
		indexConfigEnd := strings.Index(pa[indexConfigBegin:], " ")
		if indexConfigEnd == -1 {
			indexConfigEnd = len(pa)
		} else {
			indexConfigEnd += indexConfigBegin
		}
		p = pa[indexConfigBegin:indexConfigEnd]
	case global.UniversalMode:
		p = GetV2rayLocationAsset() + "/config.json"
	default:
	}
	return
}

var whitelistCn struct {
	siteList *v2router.GeoSiteList
	ipList   *v2router.GeoIPList
	sync.Mutex
}

func GetWhitelistCn(externIps []*v2router.CIDR, externDomains []*v2router.Domain) (wlIps *v2router.GeoIPMatcher, wlDomains *strmatcher.MatcherGroup, err error) {
	whitelistCn.Lock()
	defer whitelistCn.Unlock()
	dir := GetV2rayLocationAsset()
	if whitelistCn.siteList == nil {
		var siteList v2router.GeoSiteList
		b, err := ioutil.ReadFile(path.Join(dir, "geosite.dat"))
		if err != nil {
			return nil, nil, newError("GetWhitelistCn").Base(err)
		}
		err = proto.Unmarshal(b, &siteList)
		if err != nil {
			return nil, nil, newError("GetWhitelistCn").Base(err)
		}
		whitelistCn.siteList = &siteList
	}
	wlDomains = new(strmatcher.MatcherGroup)
	domainMatcher := new(dnsPoison.DomainMatcherGroup)
	fullMatcher := new(dnsPoison.FullMatcherGroup)
	for _, e := range whitelistCn.siteList.Entry {
		if e.CountryCode == "CN" {
			dms := e.Domain
			dms = append(dms, externDomains...)
			for _, dm := range dms {
				switch dm.Type {
				case v2router.Domain_Domain:
					domainMatcher.Add(dm.Value)
				case v2router.Domain_Full:
					fullMatcher.Add(dm.Value)
				case v2router.Domain_Plain:
					wlDomains.Add(dnsPoison.SubstrMatcher(dm.Value))
				case v2router.Domain_Regex:
					r, err := regexp.Compile(dm.Value)
					if err != nil {
						break
					}
					wlDomains.Add(&dnsPoison.RegexMatcher{Pattern: r})
				}
			}
			break
		}
	}
	domainMatcher.Add("lan")
	wlDomains.Add(domainMatcher)
	wlDomains.Add(fullMatcher)

	if whitelistCn.ipList == nil {
		var ipList v2router.GeoIPList
		b, err := ioutil.ReadFile(path.Join(dir, "geoip.dat"))
		if err != nil {
			return nil, nil, newError("GetWhitelistCn").Base(err)
		}
		err = proto.Unmarshal(b, &ipList)
		if err != nil {
			return nil, nil, newError("GetWhitelistCn").Base(err)
		}
		whitelistCn.ipList = &ipList
	}
	wlIps = new(v2router.GeoIPMatcher)
	for _, e := range whitelistCn.ipList.Entry {
		if e.CountryCode == "CN" {
			ips := e.Cidr
			ips = append(ips, externIps...)
			_ = wlIps.Init(ips)
			break
		}
	}
	return
}
func GetV2rayServiceFilePath() (path string, err error) {
	var out []byte

	if global.ServiceControlMode == global.SystemctlMode {
		out, err = exec.Command("sh", "-c", "systemctl status v2ray|grep /v2ray.service").CombinedOutput()
		if err != nil {
			err = newError(strings.TrimSpace(string(out)))
			if !strings.Contains(string(out), "not be found") {
				path = `/usr/lib/systemd/system/v2ray.service`
				return
			}
		}
	} else if global.ServiceControlMode == global.ServiceMode {
		out, err = exec.Command("sh", "-c", "service v2ray status|grep /v2ray.service").CombinedOutput()
		if err != nil || strings.TrimSpace(string(out)) == "(Reason:" {
			if !strings.Contains(string(out), "not be found") {
				path = `/lib/systemd/system/v2ray.service`
				return
			}
			if err != nil {
				err = newError(strings.TrimSpace(string(out)))
			}
		}
	} else {
		err = newError("commands systemctl and service not found")
		return
	}
	if err != nil {
		return
	}
	sout := string(out)
	l := strings.Index(sout, "/")
	r := strings.Index(sout, "/v2ray.service")
	if l < 0 || r < 0 {
		err = newError("fail: getV2rayServiceFilePath")
		return
	}
	path = sout[l : r+len("/v2ray.service")]
	return
}
