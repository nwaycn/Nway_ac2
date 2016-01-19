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

package main

import (
	_ "github.com/nwaycn/Nway_ac2/routers"
	"github.com/astaxie/beego"
	"github.com/nwaycn/Nway_ac2/ac/nway-pbx/nway_database/nway_connection"
	_"github.com/nwaycn/Nway_ac2/ac/nway-pbx/nway_database/nway_db_base"
	"github.com/nwaycn/Nway_ac2/ac/nway-util/logger"
)

func main() {
	var dbinstance *nwayconnection.NwayConnect =nil
	dbinstance = nwayconnection.NewDb()
	if dbinstance == nil{
		logger.Error("Database create instance failed")
		return
	}
	bInit := dbinstance.Init()
	if bInit == false{
		logger.Error("Database create connection failed")
		return 
	}
	 
	
	beego.Run()
}

