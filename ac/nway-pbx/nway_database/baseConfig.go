package nwaydb


import (
	_"github.com/nwaycn/Nway_ac2/ac/nway-pbx/config"
	"github.com/nwaycn/Nway_ac2/ac/nway-util/logger"
	_"github.com/nwaycn/Nway_ac2/ac/nway-pbx/nway_database/nway_connection"
	"github.com/nwaycn/Nway_ac2/ac/nway-pbx/nway_database/nway_db_base"
)

type BaseConfig struct{
	Config_name string
	Config_param string
}
type DbBaseConfig struct{
	dbbase nway_db_base.DbBase
	
}


func (d* DbBaseConfig)LoadConfig()(map[string]string , bool){
	logger.Info("load all config from base_config table")
	var configs map[string]string 
	var param,value string
	configs = make(map[string]string , 4)
	
	var strsql string
	strsql = "SELECT config_name, config_param FROM nway_callout_base_config"
	rows, bOp := d.dbbase.Query(strsql) //conn.Query(strsql)
	if bOp == false{
		logger.Error("query the data failed " )
		rows.Close()
		return configs,false
	}
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&param,&value)
		if err != nil {
			logger.Error("scan the data error,", err)
			
		}
		configs[param] = value
	}
	return configs, true
}