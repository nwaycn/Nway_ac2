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
	"strconv"
	_"github.com/nwaycn/Nway_ac2/ac/nway-pbx/config"
	"github.com/nwaycn/Nway_ac2/ac/nway-util/logger"
	_"github.com/nwaycn/Nway_ac2/ac/nway-pbx/nway_database/nway_connection"
	"github.com/nwaycn/Nway_ac2/ac/nway-pbx/nway_database/nway_db_base"
	 
)

type Gateway struct{
  Id  int64
  Gateway_name  string
  Gateway_url string
  Call_prefix string
  Max_call int32
  Added_autodial bool
  Curr_talking int32
  Domain_id int64
}

type DbGateway struct{
	dbbase nway_db_base.DbBase
}


func (d* DbGateway)GetGateways(domain_id int64)( map[int64]Gateway, bool){
	var gateways map[int64]Gateway
	gateways = make(map[int64]Gateway,0)
	
	//
	sqlstr := "SELECT id, gateway_name, gateway_url, call_prefix, max_call, added_autodial," + 
          "curr_talking, domain_id " +
  		" FROM nway_callout_gateways " +
		" where  domain_id=" + strconv.FormatInt(domain_id,10)
		 
		
	rows, bOp := d.dbbase.Query(sqlstr) 
	
	if bOp == false{
		logger.Error("query the data failed " )
		rows.Close()
		return gateways,false
	}
	 
	var gwt Gateway
	
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&gwt.Id,&gwt.Gateway_name, &gwt.Gateway_url, &gwt.Call_prefix, &gwt.Max_call,
		    &gwt.Added_autodial, &gwt.Curr_talking, &gwt.Domain_id)
		if err != nil {
			logger.Error("scan the data error,", err)
			return gateways,false
		}
		gateways[gwt.Id] = gwt
	}
	return gateways,true

}
func (d* DbGateway)GetGatewayById(gw_id int64) (Gateway,bool){
	 var gwt Gateway
	
	//
	sqlstr := "SELECT id, gateway_name, gateway_url, call_prefix, max_call, added_autodial," + 
          "curr_talking, domain_id " +
  		" FROM nway_callout_gateways " +
		" where  id=" + strconv.FormatInt(gw_id,10)
		 
		
	rows, bOp := d.dbbase.Query(sqlstr) 
	
	if bOp == false{
		logger.Error("query the data failed " )
		rows.Close()
		return gwt,false
	}
	 
	
	
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&gwt.Id,&gwt.Gateway_name, &gwt.Gateway_url, &gwt.Call_prefix, &gwt.Max_call,
		    &gwt.Added_autodial, &gwt.Curr_talking, &gwt.Domain_id)
		if err != nil {
			logger.Error("scan the data error,", err)
			return gwt,false
		}
		 
	}
	return gwt,true
}
