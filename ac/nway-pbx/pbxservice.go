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

package pbx

import (
	//"fmt"
	"github.com/nwaycn/Nway_ac2/ac/nway-util/logger"
	"runtime"
)

func StartInboundService() bool {
	var inbounds InboundService 
	err := inbounds.Start()
	if err != nil{
		return false
	}
	return true
}

func StartOutboundService() bool {
	var outbounds OutboundService
	err := outbounds.Start()
	if err != nil {
		return false
	}
	return true
}

func StartService() bool {
	runtime.GOMAXPROCS(runtime.NumCPU())
	logger.SetConsole(true)
	logger.SetRollingDaily(".", "nwaypbx.log")
	logger.SetLevel(logger.INFO)
	if StartOutboundService() {
		logger.Debug("Outbound service has started")
	} else {
		logger.Debug("Outbound service start failed")
		return false
	}

	if StartInboundService() {
		logger.Debug("Inbound service has started")
	} else {
		logger.Debug("Inbound service start failed")
		return false
	}
	return true
}
