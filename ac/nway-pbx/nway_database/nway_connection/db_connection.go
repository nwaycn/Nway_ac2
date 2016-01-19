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
package nwayconnection

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nwaycn/Nway_ac2/ac/nway-pbx/config"
	"github.com/nwaycn/Nway_ac2/ac/nway-util/logger"
	
	"sync"
	
)

type NwayConnect struct {
	NwayDb *sql.DB
}

var _init_db sync.Once
var _instance_ *NwayConnect

func NewDb() *NwayConnect {
	_init_db.Do(func() {
		_instance_ = new(NwayConnect)

	})
	return _instance_
}

func (this *NwayConnect) Init() bool {
	//	result := true
	var err error
	var conf *pbxconfig.NwayConfig
	var isload bool
	conf, isload = pbxconfig.LoadConf()
	if isload == true {

		dbstring := conf.Dbconf.ConnString

		this.NwayDb, err = sql.Open("postgres", dbstring)
		this.NwayDb.SetMaxIdleConns(4)
		this.NwayDb.SetMaxOpenConns(10)
		if err != nil {
			//log.Fatalf()
			logger.Error("Got error when connect database, the error is '%v'", err)
			fmt.Println(dbstring)
			return false
		}

		return true
	} else {
		return false
	}
}

func (this *NwayConnect) GetConn() *sql.DB {
	//保证所有的连接都是有效的
	strsql := "select now();"
	rows, err := this.NwayDb.Query(strsql)
	if err != nil {
		logger.Error("failed to get a connect '%v'", err)

		return nil
	}
	rows.Close()
	return this.NwayDb
}
