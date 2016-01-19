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

	_ "github.com/nwaycn/Nway_ac2/ac/nway-pbx/config"
	_ "github.com/nwaycn/Nway_ac2/ac/nway-pbx/nway_database/nway_connection"
	"github.com/nwaycn/Nway_ac2/ac/nway-pbx/nway_database/nway_db_base"
	"github.com/nwaycn/Nway_ac2/ac/nway-util/logger"
)

type CallNumber struct {
	Id                  int64
	Domain_id           int64
	Group_id            int64
	Number              string
	Is_called           int32
	Call_state          int32
	Start_time          string
	Answer_time         string
	Hangup_time         string
	Hangup_reason_id    int64
	Answer_extension_id int64
	Record_file         string
	Wait_sec            int32
	Cdr_id              int64
	Call_out_task_id    int64
	Create_time         string
}

type DbCallNumber struct {
	dbbase nway_db_base.DbBase
}

func (d *DbCallNumber) GetNumbers(taskid int64, runpostion int64) (map[int64]CallNumber, bool) {
	var numbers map[int64]CallNumber
	numbers = make(map[int64]CallNumber, 0)

	sqlstr := "SELECT id, domain_id, group_id, 'number', is_called, call_state, start_time," +
		"answer_time, hangup_time, hangup_reason_id, answer_extension_id, " +
		"record_file, wait_sec, cdr_id, call_out_task_id, create_time" +
		"FROM nway_callout_numbers" +
		" where  call_out_task_id=" + strconv.FormatInt(taskid, 10) +
		" and id > " + strconv.FormatInt(runpostion, 10)

	rows, bOp := d.dbbase.Query(sqlstr)

	if bOp == false {
		logger.Error("query the data failed ")
		rows.Close()
		return numbers, false
	}
	var nmb CallNumber
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&nmb.Id,
			&nmb.Domain_id,
			&nmb.Group_id,
			&nmb.Number,
			&nmb.Is_called,
			&nmb.Call_state,
			&nmb.Start_time,
			&nmb.Answer_time,
			&nmb.Hangup_time,
			&nmb.Hangup_reason_id,
			&nmb.Answer_extension_id,
			&nmb.Record_file,
			&nmb.Wait_sec,
			&nmb.Cdr_id,
			&nmb.Call_out_task_id,
			&nmb.Create_time)
		if err != nil {
			logger.Error("scan the data error,", err)
			return numbers, false
		}
		numbers[nmb.Id] = nmb
	}
	return numbers, true

}

func (d *DbCallNumber) SetStartCall(id int64) bool {

	sqlstr := "update nway_callout_numbers set call_state=1," +
		"Start_time=current_timestamp,Is_called=True " +
		" where id=" + strconv.FormatInt(id, 10)
	bSuccess, err := d.dbbase.Exec(sqlstr)
	if err != nil {
		logger.Error("update call error,", err)
		return false
	}
	return bSuccess

}

func (d *DbCallNumber) SetAnswerCall(id int64, ext string) bool {
	var dbext DbExtension
	ext_id, mybool := dbext.GetExtId(ext, id)
	if mybool == false {
		logger.Error("the ext id not found")
		return false
	}
	sqlstr := "update nway_callout_numbers set call_state=2,Answer_time=current_timestamp, " +
		"Answer_extension_id=" + strconv.FormatInt(ext_id, 10) +
		" where id=" + strconv.FormatInt(id, 10)
	bSuccess, err := d.dbbase.Exec(sqlstr)
	if err != nil {
		logger.Error("update call error,", err)
		return false
	}
	return bSuccess
}

func (d *DbCallNumber) SetHangupCall(id int64, reson_id int) bool {

	sqlstr := "update nway_callout_numbers set call_state=1," +
		"Hangup_time=current_timestamp,Hangup_reason_id=" + strconv.Itoa(reson_id) +
		" where id=" + strconv.FormatInt(id, 10)
	bSuccess, err := d.dbbase.Exec(sqlstr)
	if err != nil {
		logger.Error("update call error,", err)
		return false
	}
	return bSuccess

}
