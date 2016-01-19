/*************************************************************************
based for NwayPBX-go
Copyright (C) 2015-, Li hao <lihao@nway.com.cn>
License： GPL
author: Li hao
email: lihao@nway.com.cn
The Initial Developer of the Original Code is
Li hao<lihao@nway.com.cn>
Portions created by the Initial Developer are Copyright (C)
the Initial Developer. All Rights Reserved.
Contributor(s):
**************************************************************************/
package pbxconfig

//单例模式，用于加快访问速度
import (
	//"go_fs/fs_control/conf"
	"github.com/nwaycn/Nway_ac2/ac/nway-util/logger"
	"sync"

	//"fmt"
)

type DbConf struct {
	ConnString string
}

type FsConf struct {
	FsHost    string
	FsPort    int
	FsAuth    string
	FsTimeout int
}

type FsOutboundConf struct {
	ObPort string
	ObPPort string //the port of public outbound connection
}

type NwayConfig struct {
	Dbconf         DbConf
	Fsconf         FsConf
	Fsoutboundconf FsOutboundConf
}

var _init_conf sync.Once
var _instanceconf_ *NwayConfig
var isLoad bool = false

func LoadConf() (*NwayConfig, bool) {
	_init_conf.Do(func() {
		_instanceconf_ = new(NwayConfig)
		isLoad = _instanceconf_.Load()

	})
	return _instanceconf_, isLoad
}

func (this *NwayConfig) Load() bool {
	conf, result := SetConfig("conf.ini")
	if result == false {
		logger.Debug("load the config file failed")
		return false
	}

	this.Dbconf.ConnString, result = conf.GetValue("database", "connstring")
	if result == false {
		logger.Error("load database connect string failed")
		return false
	}
	this.Fsconf.FsHost, result = conf.GetValue("freeswitch", "host")
	if result == false {
		logger.Error("load freeswitch host string failed")
		return false
	}
	this.Fsconf.FsAuth, result = conf.GetValue("freeswitch", "auth")
	if result == false {
		logger.Error("load freeswitch auth string failed")
		return false
	}
	this.Fsconf.FsPort = conf.GetInt("freeswitch", "port", 8021)

	this.Fsconf.FsTimeout = conf.GetInt("freeswitch", "timeout", 500)

	this.Fsoutboundconf.ObPort, result = conf.GetValue("freeswitch", "outboundport")
	this.Fsoutboundconf.ObPPort, result = conf.GetValue("freeswitch", "pub-outboundport")

	return true

}
