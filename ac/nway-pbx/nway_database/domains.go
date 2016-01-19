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

package nwaydb


import (
	_"github.com/nwaycn/Nway_ac2/ac/nway-pbx/config"
	"github.com/nwaycn/Nway_ac2/ac/nway-util/logger"
	_"github.com/nwaycn/Nway_ac2/ac/nway-pbx/nway_database/nway_connection"
	"github.com/nwaycn/Nway_ac2/ac/nway-pbx/nway_database/nway_db_base"
	 
)
type Domain  struct{
  Domain_uuid int64
  Domain_name string
}

type DbDomain struct{
	dbbase nway_db_base.DbBase
}

func (d* DbDomain)GetDomainId(doamin_name string)(int64,bool){
	var strsql string
	strsql = "select domain_uuid from nway_callout_domains where domain_name='" + doamin_name + "'"
	rows, bOp := d.dbbase.Query(strsql) 
	
	if bOp == false{
		logger.Error("query the data failed " )
		rows.Close()
		return 0,false
	}
	var id int64 = 0
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&id)
		if err != nil {
			logger.Error("scan the data error,", err)
			return 0,false
		}
	}
	return id,true   
}