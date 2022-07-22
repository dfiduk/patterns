/*
Copyright (c) 2021-2022 MIND Software LLC. All Rights Reserved.

This file is part of MIND Multicloud Mobility Software Suite.
For more information visit https://mindsw.io
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/mind-sw/mind-libs/executor"
	"github.com/mind-sw/mind-libs/winps"
	"github.com/mind-sw/orchestrator/libs/common"

	"github.com/sirupsen/logrus"
)

const BASEDIR = "c:/mind"
const LOCKFILE = BASEDIR + "/mind.lock"

var env []string
var execLib executor.Executor
var log *logrus.Logger
var errorChan chan error
var wg sync.WaitGroup
var errorCode int

func prepare(eConfig *executor.Config) string {

	config := flag.String("c", fmt.Sprintf("%s/mind.conf", BASEDIR), "MIND migration config")
	jobGUID := flag.String("g", "", "guid of execution")
	timeout := flag.Int("t", 600, "timeout")
	URI := flag.String("u", "", "callback URL")
	loglvl := flag.String("l", "info", "log leve: info, debug or...")
	flag.Parse()

	ll, err := logrus.ParseLevel(*loglvl)
	if err != nil {
		ll = logrus.DebugLevel
	}
	log.SetLevel(ll)

	eConfig.Basepath = BASEDIR
	eConfig.Lockfile = LOCKFILE
	eConfig.JobGUID = *jobGUID
	eConfig.Timeout = *timeout
	eConfig.URI = *URI
	eConfig.Logger = log

	return *config
}

func parseConfig(configPath string) (config common.MigrationConfig, err error) {

	if configPath != "" {
		cfd, err := os.Open(configPath)
		if err != nil {
			err = fmt.Errorf("cannot open config file %s: %s", configPath, err.Error())
			return config, err
		}
		defer cfd.Close()

		decoder := json.NewDecoder(cfd)
		err = decoder.Decode(&config)
		if err != nil {
			return config, err
		}

	}

	return config, nil
}

func main() {

	var eConfig executor.Config
	errorChan = make(chan error)
	log = logrus.New()
	configPath := prepare(&eConfig)

	execLib.Init(&eConfig, false)
	go execLib.Callbacker()
	action(configPath)

	defer execLib.Cleanup()

	wg.Wait()
	os.Exit(errorCode)

}

func cleanupIface(ifaceConfig common.UnitNetIfaceConfigUnified) (err error) {
	registerPath := fmt.Sprintf(`HKLM:System\ControlSet001\Services\Tcpip\Parameters\Interfaces\{%s}`, ifaceConfig.ConfigID)

	cmd := fmt.Sprintf(`Remove-ItemProperty -EA Ignore -Path '%s' -Name IPAddress,SubnetMask,DefaultGateway,DefaultGatewayMetric,DhcpIPAddress,DhcpSubnetMask,DhcpNameServer,DhcpDefaultGateway,DhcpSubnetMaskOpt`, registerPath)
	_, err = winps.ExecuteRaw(cmd)
	if err != nil {
		return
	}

	return nil
}

func toDHCP(ifaceConfig common.UnitNetIfaceConfigUnified) (err error) {

	var cmd string
	registerPath := fmt.Sprintf(`HKLM:System\ControlSet001\Services\Tcpip\Parameters\Interfaces\{%s}`, ifaceConfig.ConfigID)

	// cmd := fmt.Sprintf(`Remove-ItemProperty -EA Ignore -Path '%s' -Name IPAddress,SubnetMask,DefaultGateway,DefaultGatewayMetric`, registerPath)
	// _, err = winps.ExecuteRaw(cmd)
	// if err != nil {
	// 	return
	// }

	cmd = fmt.Sprintf(`Set-ItemProperty -Path '%s' -Type DWord -Name EnableDHCP -Value 1`, registerPath)
	_, err = winps.ExecuteRaw(cmd)
	if err != nil {
		return
	}

	cmd = fmt.Sprintf(`Set-ItemProperty -Path '%s' -Name NameServer ''`, registerPath)
	_, err = winps.ExecuteRaw(cmd)
	if err != nil {
		return
	}
	return nil
}

func toStatic(ifaceConfig common.UnitNetIfaceConfigUnified) (err error) {

	var cmd string
	registerPath := fmt.Sprintf(`HKLM:System\ControlSet001\Services\Tcpip\Parameters\Interfaces\{%s}`, ifaceConfig.ConfigID)

	// cmd := fmt.Sprintf(`Remove-ItemProperty -EA Ignore -Path '%s' -Name DhcpIPAddress,DhcpSubnetMask,DhcpNameServer,DhcpDefaultGateway,DhcpSubnetMaskOpt`, registerPath)
	// _, err = winps.ExecuteRaw(cmd)
	// if err != nil {
	// 	return
	// }

	cmd = fmt.Sprintf(`Set-ItemProperty -Path '%s' -Type DWord -Name EnableDHCP -Value 0`, registerPath)
	_, err = winps.ExecuteRaw(cmd)
	if err != nil {
		return
	}

	for _, cidrStr := range ifaceConfig.Addresses {

		var addr net.IP
		var cidr *net.IPNet

		addr, cidr, err = net.ParseCIDR(cidrStr)
		if err != nil {
			return
		}

		netmask := fmt.Sprintf("%d.%d.%d.%d", cidr.Mask[0], cidr.Mask[1], cidr.Mask[2], cidr.Mask[3])

		cmd = fmt.Sprintf(`Set-ItemProperty -Path '%s' -Type MultiString -Name IPAddress -Value "%s"`, registerPath, addr)
		_, err = winps.ExecuteRaw(cmd)
		if err != nil {
			return
		}

		cmd = fmt.Sprintf(`Set-ItemProperty -Path '%s' -Type MultiString -Name SubnetMask -Value "%s"`, registerPath, netmask)
		_, err = winps.ExecuteRaw(cmd)
		if err != nil {
			return
		}
	}

	cmd = fmt.Sprintf(`Set-ItemProperty -Path '%s'  -Type MultiString -Name DefaultGateway -Value "%s"`, registerPath, ifaceConfig.Gateway)
	_, err = winps.ExecuteRaw(cmd)
	if err != nil {
		return
	}
	cmd = fmt.Sprintf(`Set-ItemProperty -Path '%s'  -Type MultiString -Name DefaultGatewayMetric -Value 0`, registerPath)
	_, err = winps.ExecuteRaw(cmd)
	if err != nil {
		return
	}

	if len(ifaceConfig.DNS) > 0 {
		dnsString := strings.Join(ifaceConfig.DNS, ",")
		cmd = fmt.Sprintf(`Set-ItemProperty -Path '%s' -Name NameServer "%s"`, registerPath, dnsString)
		_, err = winps.ExecuteRaw(cmd)
		if err != nil {
			return
		}
	}

	return nil
}

func action(configPath string) {

	config, err := parseConfig(configPath)
	if err != nil {
		panic(err)
	}

	for _, iface := range config.Target.Network.Interfaces {
		ifaceConf := iface.Config.Unified
		cleanupIface(ifaceConf)
		if ifaceConf.DHCP {
			err = toDHCP(ifaceConf)
		} else {
			err = toStatic(ifaceConf)
		}
		if err != nil {
			panic(err)
		}
	}

}
