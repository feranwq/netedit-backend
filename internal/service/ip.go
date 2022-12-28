package service

import (
	"context"
	"fmt"
	"math/rand"
	"netedit/utility/utilexec"
	"os"
	"strings"
	"time"
)

type IIP interface {
	GetIP(ctx context.Context) (ip []string, err error)
	UpdateIP(ctx context.Context, oldIP, newIP string) (err error)
	AddIP(ctx context.Context, ip string) (err error)
	DeleteIP(ctx context.Context, ip string) (err error)
}

var ipService = ipImpl{}

func IP(deviceName string) IIP {
	ipService.deviceName = deviceName
	return &ipService
}

type ipImpl struct{
	deviceName string
}

func (s *ipImpl) GetIP(ctx context.Context) (ipList []string, err error) {
	cmd := "/bin/sh"
	arg := []string{"-c", fmt.Sprintf("grep IPADDR /etc/sysconfig/network-scripts/ifcfg-%s:* 2>/dev/null || :", s.deviceName)}
	out := utilexec.String(ctx, cmd, arg[:]...)
	ipAll := strings.Split(out, "\n")
	for _, i := range ipAll {
		if i != "" {
			ip := strings.Split(i, "=")
			ipList = append(ipList, ip[1])
		}
	}
	return
}

func (s *ipImpl) UpdateIP(ctx context.Context, oldIP, newIP string) (err error) {
	cmd := "/bin/sh"
	arg := []string{"-c", fmt.Sprintf("grep %s /etc/sysconfig/network-scripts/ifcfg-%s:* 2>/dev/null || :", oldIP, s.deviceName)}
	out := utilexec.String(ctx, cmd, arg[:]...)
	if out == "" {
		return fmt.Errorf("ip %s no exist", oldIP)
	}
	arg = []string{"-c", fmt.Sprintf("grep %s /etc/sysconfig/network-scripts/ifcfg-%s:* 2>/dev/null || :", newIP, s.deviceName)}
	out = utilexec.String(ctx, cmd, arg[:]...)
	if out != "" {
		return fmt.Errorf("ip %s already exist", newIP)
	}
	arg = []string{"-c", fmt.Sprintf("sed -i 's|%s|%s|g' /etc/sysconfig/network-scripts/ifcfg-%s:* || :", oldIP, newIP, s.deviceName)}
	utilexec.Cmd(ctx, cmd, arg[:]...)
	arg = []string{"-c", fmt.Sprintf("sed -i 's|%s|%s|g' /etc/firewalld/direct.xml || :", oldIP, newIP)}
	utilexec.Cmd(ctx, cmd, arg[:]...)
	arg = []string{"-c", "systemctl restart network && firewall-cmd --reload"}
	utilexec.Cmd(ctx, cmd, arg[:]...)
	return
}

func (s *ipImpl) AddIP(ctx context.Context, ip string) (err error) {
	cmd := "/bin/sh"
	arg := []string{"-c", fmt.Sprintf("grep %s /etc/sysconfig/network-scripts/ifcfg-%s:* 2>/dev/null || :", ip, s.deviceName)}
	out := utilexec.String(ctx, cmd, arg[:]...)
	if out != "" {
		return fmt.Errorf("ip %s already exist", ip)
	}
	rand.Seed(time.Now().UnixNano())
	cfgName := fmt.Sprintf("%s:%d",s.deviceName, rand.Intn(1000))
	arg = []string{"-c", fmt.Sprintf("ls /etc/sysconfig/network-scripts/ifcfg-%s || :", cfgName)}
	out = utilexec.String(ctx, cmd, arg[:]...)
	cfgContent := fmt.Sprintf(`
		cat << EOF >> /etc/sysconfig/network-scripts/ifcfg-%s
		DEVICE=%s
		IPADDR=%s
		PREFIX=24
		EOF`, cfgName, cfgName, ip)
	cfgContent = strings.Replace(cfgContent, "\t", "", -1)
	if strings.HasPrefix(out, "ls:") {
		arg = []string{"-c", cfgContent}
		utilexec.Cmd(ctx, cmd, arg[:]...)
		arg = []string{"-c", fmt.Sprintf(`grep -q %s /etc/firewalld/direct.xml || sed -i '3i   <passthrough ipv="ipv4">-t nat POSTROUTING -o %s -j MASQUERADE -s %s/24</passthrough>' /etc/firewalld/direct.xml`, ip, s.deviceName, ip)}
		utilexec.Cmd(ctx, cmd, arg[:]...)
		arg = []string{"-c", "systemctl restart network && firewall-cmd --reload"}
		utilexec.Cmd(ctx, cmd, arg[:]...)
	} else {
		s.AddIP(ctx, ip)
	}
	return
}

func (s *ipImpl) DeleteIP(ctx context.Context, ip string) (err error) {
	cmd := "/bin/sh"
	arg := []string{"-c", fmt.Sprintf("grep %s /etc/sysconfig/network-scripts/ifcfg-%s:* 2>/dev/null || :", ip, s.deviceName)}
	out := utilexec.String(ctx, cmd, arg[:]...)
	if out == "" {
		return fmt.Errorf("ip %s no exist", ip)
	}
	arg = []string{"-c", fmt.Sprintf("grep -H %s /etc/sysconfig/network-scripts/ifcfg-%s:* || :", ip, s.deviceName)}
	out = utilexec.String(ctx, cmd, arg[:]...)
	out = strings.Trim(out, "\n")
	if strings.HasSuffix(out, ip) {
		cfgName := out[:strings.LastIndex(out, ":")]
		err = os.Remove(cfgName)
	}
	arg = []string{"-c", fmt.Sprintf(`sed -i '/%s\/24/d' /etc/firewalld/direct.xml`, ip)}
	utilexec.Cmd(ctx, cmd, arg[:]...)
	arg = []string{"-c", "systemctl restart network && firewall-cmd --reload"}
	utilexec.Cmd(ctx, cmd, arg[:]...)
	return
}
