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
package nway_db_base

import (
	"errors"
	"strconv"
)

 
import(
	"database/sql"
	"github.com/nwaycn/Nway_ac2/ac/nway-pbx/nway_database/nway_connection"
	 
	"github.com/nwaycn/Nway_ac2/ac/nway-util/logger"
	_"database/sql/driver"
	
)

type DbBase struct{
	 
	ins dbBaser
}


func (d *DbBase)Query(sqlstring string)(*sql.Rows,bool){
	var dbinstance *nwayconnection.NwayConnect =nil
	dbinstance = nwayconnection.NewDb()
	NwayConn := dbinstance.GetConn()
	if NwayConn == nil{
		logger.Error("Get The Connect Failed")
		return nil,false
	}
	rows, err := NwayConn.Query(sqlstring)
	if err != nil {
		//log.Fatal("failed to scan", err)
		 
		logger.Error(err)
		return nil,false
	}
	return rows,true
	
}

func (d *DbBase)QureyPaging(sqlstring string , page, perpage int)(*sql.Rows,bool){
	sqlstring +=  " limit " + strconv.Itoa(perpage) + " offset " + strconv.Itoa((page-1)*perpage)
	var dbinstance *nwayconnection.NwayConnect =nil
	dbinstance = nwayconnection.NewDb()
	NwayConn := dbinstance.GetConn()
	if NwayConn == nil{
		logger.Error("Get The Connect Failed")
		return nil,false
	}
	rows, err := NwayConn.Query(sqlstring)
	if err != nil {
		//log.Fatal("failed to scan", err)
		 
		logger.Error(err)
		return nil,false
	}
	return rows,true
}

func (d *DbBase)Exec(sqlstring string)(bool,error){
	var dbinstance *nwayconnection.NwayConnect =nil
	dbinstance = nwayconnection.NewDb()
	NwayConn := dbinstance.GetConn()
	if NwayConn == nil{
		logger.Error("Get The Connect Failed")
		err := errors.New("connect failed")
		return false,err
	}
	_, err := NwayConn.Exec(sqlstring)
	if err != nil {
		//log.Fatal("failed to scan", err)
		 
		logger.Error(err)
		return false,err
	}
	return true,nil
}