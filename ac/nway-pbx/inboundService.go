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
	"fmt"
	_ "fmt"
	_ "strconv"
	"strings"

	"github.com/nwaycn/Nway_ac2/ac/nway-pbx/config"
	"github.com/nwaycn/Nway_ac2/ac/nway-pbx/nway_database"
	. "github.com/nwaycn/Nway_ac2/ac/nway-util/esl/goesl"
	"github.com/nwaycn/Nway_ac2/ac/nway-util/logger"
)

var (
	bInbound  = true
	bAutoDail = true
)

type InboundService struct {
}

func (this *InboundService) Start() error {
	nwayconfig, result := pbxconfig.LoadConf()
	if result == false {
		logger.Error("load config file failed")
	}
	//c, err := eventsocket.Dial(nwayconfig.Fsconf.FsHost+":"+strconv.Itoa(nwayconfig.Fsconf.FsPort),
	//	nwayconfig.Fsconf.FsAuth)

	client, err := NewClient(nwayconfig.Fsconf.FsHost, uint(nwayconfig.Fsconf.FsPort), nwayconfig.Fsconf.FsAuth, nwayconfig.Fsconf.FsTimeout)

	if err != nil {
		logger.Error("Error while creating new client: %s", err)
		return err
	}

	Debug("Nway! New client: %q", client)

	// Apparently all is good... Let us now handle connection :)
	// We don't want this to be inside of new connection as who knows where it my lead us.
	// Remember that this is crutial part in handling incoming messages :)
	go client.Handle()

	for {
		msg, err := client.ReadMessage()

		if err != nil {

			// If it contains EOF, we really dont care...
			if !strings.Contains(err.Error(), "EOF") && err.Error() != "unexpected end of JSON input" {
				logger.Error("Error while reading Freeswitch message: %s", err)
			}
			//break
		}

		Debug("%s", msg)
	}
	return err
}

func (this *InboundService) AutoDialThread() {
	for bAutoDail == true {
		//循环获取任务
		//获取task
		var dbtask nwaydb.DbTask
		var tasks map[int64]nwaydb.Task
		var bSuccess bool
		tasks, bSuccess = dbtask.GetTasks()
		if bSuccess == true {
			//获取到了task
			//task := tasks[]
			for v, k := range tasks {
				fmt.Printf(v, k)
			}
		}
	}
}

func (this *InboundService) StartAutoDial() {
	//用作自动外呼部分的线程
	bAutoDail = true
	go this.AutoDialThread()

}

func (this *InboundService) StopAutoDial() {
	bAutoDail = false
}
