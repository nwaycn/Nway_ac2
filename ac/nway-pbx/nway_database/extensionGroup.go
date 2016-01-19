
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

type ExtensionGroup struct{
	Id int64
  Domain_id int64
  Group_name string
  Group_description string
  Call_order_id int64
  Call_order_crycle int32
}

type DbExtensionGroup struct{
	dbbase nway_db_base.DbBase
}

func (d* DbExtensionGroup)GetGroupByDomain(domain_name string)(map[int64]ExtensionGroup,bool){
	var groups map[int64]ExtensionGroup
	groups = make(map[int64]ExtensionGroup,0)
	
	sqlstr := "SELECT a.id, a.domain_id, a.group_name, a.call_order_id," +
       " a.call_order_crycle " +
  		" FROM nway_callout_extension_groups a,nway_callout_domains b " +
		" where a.domain_id=b.domain_uuid and b.domain_name='" +
		domain_name + "'"
		
	rows, bOp := d.dbbase.Query(sqlstr) 
	
	if bOp == false{
		logger.Error("query the data failed " )
		rows.Close()
		return groups,false
	}
	 
	var group ExtensionGroup
	
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&group.Id, &group.Domain_id, &group.Group_name,
			&group.Call_order_id,&group.Call_order_crycle)
		if err != nil {
			logger.Error("scan the data error,", err)
			return groups,false
		}
		groups[group.Id] = group
	}
	return groups,true
	
}