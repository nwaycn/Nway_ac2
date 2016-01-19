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
	"strconv"
)
type DomainSetting struct{
  Domain_uuid int64
  Domain_setting_uuid int64
  Domain_setting_category string
  Domain_setting_subcategory string
  Domain_setting_name string
  Domain_setting_value string
  Domain_setting_enabled string
  Domain_setting_description string
}

type DbDomainSetting struct{
	dbbase nway_db_base.DbBase
}

func (d* DbDomainSetting)GetDomainSetting(domain_id int64) (map[int64]DomainSetting,bool){
	var dss map[int64]DomainSetting
	dss = make(map[int64]DomainSetting, 4)
	var sqlstr string
	sqlstr = "SELECT  domain_setting_uuid, domain_setting_category, domain_setting_subcategory," +
       "domain_setting_name, domain_setting_value, domain_setting_enabled," +
       "FROM nway_callout_domain_settings where domain_uuid=" + strconv.FormatInt(domain_id,10)
    rows, bOp := d.dbbase.Query(sqlstr) //conn.Query(strsql)
	if bOp == false{
		logger.Error("query the data failed " )
		rows.Close()
		return dss,false
	}
	   
	var(
		Domain_setting_uuid int64
	  Domain_setting_category string
	  Domain_setting_subcategory string
	  Domain_setting_name string
	  Domain_setting_value string
	  Domain_setting_enabled string
	)
	
	  
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&Domain_setting_uuid,&Domain_setting_category,
		        &Domain_setting_subcategory,&Domain_setting_name,&Domain_setting_value,
				&Domain_setting_enabled)
		if err != nil {
			logger.Error("scan the data error,", err)
			
		}
		var ds DomainSetting 
		ds.Domain_setting_category = Domain_setting_category
		ds.Domain_setting_enabled = Domain_setting_enabled
		ds.Domain_setting_name = Domain_setting_name
		ds.Domain_setting_subcategory = Domain_setting_subcategory
		ds.Domain_setting_uuid = Domain_setting_uuid
		ds.Domain_uuid = domain_id
		dss[Domain_setting_uuid] = ds
	}
    
	return dss,true
}

