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


type Cdr struct{
  Id int64
  Domain_id int64
  Accountcode string
  Xml_cdr string
  Caller_id_name string
  Caller_id_number string
  Destination_number string
  Star_epoch int32
  Start_stamp string
  A_answer_stamp string
  A_answer_epoch int32
  A_end_epoch int32
  A_end_stamp string
  Duration int32
  Dduration int32
  Billsec int32
  Recording_file string
  Conference_name string
  Conference_id int64
  Digites_dialed string
  Hangup_cause string
  Hangup_cause_id int64
  Waitsec int32
  Call_gateway_id int64
  B_answer_stamp string
  B_answer_epoch int32
  B_end_stamp string
  B_end_epoch int32
  Hangup_direction int32  
  A_leg_called bool   
  B_leg_called bool
  Called_number string
  Auto_callout bool
  Task_id int64         
  Input_key string
  Domain_name string
}

type DbCdr struct{
	dbbase nway_db_base.DbBase
}

func (d *DbCdr)InsertCdr(c* Cdr)  (int64, bool){
	
	var sqlstr string
	sqlstr = "SELECT insertnewcdr('" + c.Domain_name + "',"
    sqlstr += "'" + c.Caller_id_name + "',"
	sqlstr += "'" + c.Caller_id_number + "',"
	sqlstr += "'" + c.Called_number + "',"
	if c.Auto_callout {
		sqlstr += "True,"
	}else{
		sqlstr += "False,"
	}
	if c.Task_id > -1{
		 
		sqlstr += strconv.FormatInt(c.Task_id,10) + ")"
	}
	
	rows, bOp := d.dbbase.Query(sqlstr)    //conn.Query(sqlstr)
	if bOp == false{
		logger.Error("get the query connection failed! ")
		rows.Close()
		return 0,false
	}
	var mynum int64 = 0
	defer rows.Close()
	for rows.Next(){
		err := rows.Scan(&mynum)
		if err != nil{
			logger.Error("Insert a new cdr failed! ",err)
			return 0,false
		} 
		
	}
	return mynum,true
	
}
func (d* DbCdr)ALegAnwser(cdrid int64) bool{
	
	var sqlstr string
	sqlstr = "update nway_callout_cdr set a_answer_stamp=current_timestamp, a_leg_called=True where id="+ strconv.FormatInt(cdrid,10)
	bOp, err := d.dbbase.Exec(sqlstr)
	 
	if (err != nil && bOp==false) {
		logger.Error("update Aleg answer failed!", err)
		return false
	}
	return true
}

func (d* DbCdr)BlegAnswer(cdrid int64) bool {
	var sqlstr string
	sqlstr = "update nway_callout_cdr set b_answer_stamp=current_timestamp, b_leg_called=True where id="+ strconv.FormatInt(cdrid,10)
	bOp, err := d.dbbase.Exec(sqlstr)
	 
	if (err != nil && bOp==false) {
		logger.Error("update Aleg answer failed!", err)
		return false
	}
	return true
}

func (d* DbCdr)AlegHangup(cdrid int64) bool{
	var sqlstr string
	sqlstr = "update nway_callout_cdr set a_end_stamp=current_timestamp  where id="+ strconv.FormatInt(cdrid,10)
	bOp, err := d.dbbase.Exec(sqlstr)
	 
	if (err != nil && bOp==false) {
		logger.Error("update Aleg answer failed!", err)
		return false
	}
	return true
}

func (d* DbCdr)BlegHangup(cdrid int64) bool{
	var sqlstr string
	sqlstr = "update nway_callout_cdr set b_end_stamp=current_timestamp  where id="+ strconv.FormatInt(cdrid,10)
	bOp, err := d.dbbase.Exec(sqlstr)
	 
	if (err != nil && bOp==false) {
		logger.Error("update Aleg answer failed!", err)
		return false
	}
	return true
}