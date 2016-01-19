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

type Extension struct {
	Id                         int64
	Domain_id                  int64
	Extension_name             string //(50) NOT NULL, -- 分机名称
	Extension_number           string //(50) NOT NULL, -- 分机号码
	Callout_number             string //(50), -- 外呼时的号码
	Extension_type             int64  // -- 分机类型
	Group_id                   int64  // -- 和extension_groups中的id对应
	Extension_pswd             string //(130),
	Extension_login_state      string //(50) DEFAULT 'login'::string  //,
	Extension_reg_state        string //(50), -- 注册状态
	Callout_gateway            int64  //
	Is_allow_callout           bool
	Limit_max                  int64
	Limit_destination          string //
	Mailbox                    string //
	Vm_password                string //
	Accountcode                string //
	Effective_caller_id_name   string //
	Effective_caller_id_number string //
	Outbound_caller_id_name    string //
	Outbound_caller_id_number  string //
	Call_group                 string //
	Hold_music                 string //
	Call_state                 int32  // -- 该分机的通话状态，空闲为0，正在通话中1
	Is_record                  bool   // -- 是否录音
	Curr_talking               int32  // -- 当前并发数，包括处理呼出或振铃中等
	Domain_name                string //(50),
}

type DbExtension struct {
	dbbase nway_db_base.DbBase
}

func (d *DbExtension) GetExtensionByDomain(domain_name string) (map[int64]Extension, bool) {
	var exts map[int64]Extension
	exts = make(map[int64]Extension, 0)

	sqlstr := "SELECT id, domain_id, extension_name, extension_number, callout_number, " +
		"extension_type, group_id, extension_pswd, extension_login_state, " +
		"extension_reg_state, callout_gateway, is_allow_callout, limit_max, " +
		"limit_destination, mailbox, vm_password, accountcode, effective_caller_id_name, " +
		"effective_caller_id_number, outbound_caller_id_name, outbound_caller_id_number," +
		"call_group, hold_music, call_state, is_record, curr_talking, " +
		"domain_name" +
		"FROM nway_callout_extensions a,nway_callout_domains b " +
		" where a.domain_id=b.domain_uuid and b.domain_name='" +
		domain_name + "'"

	rows, bOp := d.dbbase.Query(sqlstr)

	if bOp == false {
		logger.Error("query the data failed ")
		rows.Close()
		return exts, false
	}

	var ext Extension

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ext.Id,
			&ext.Domain_id,
			&ext.Extension_name,
			&ext.Extension_number,
			&ext.Callout_number,
			&ext.Extension_type,
			&ext.Group_id,
			&ext.Extension_pswd,
			&ext.Extension_login_state,
			&ext.Extension_reg_state,
			&ext.Callout_gateway,
			&ext.Is_allow_callout,
			&ext.Limit_max,
			&ext.Limit_destination,
			&ext.Mailbox,
			&ext.Vm_password,
			&ext.Accountcode,
			&ext.Effective_caller_id_name,
			&ext.Effective_caller_id_number,
			&ext.Outbound_caller_id_name,
			&ext.Outbound_caller_id_number,
			&ext.Call_group,
			&ext.Hold_music,
			&ext.Call_state,
			&ext.Is_record,
			&ext.Curr_talking,
			&ext.Domain_name)
		if err != nil {
			logger.Error("scan the data error,", err)
			return exts, false
		}
		exts[ext.Id] = ext
	}
	return exts, true
}

func (d *DbExtension) GetExtensionByGroup(group_id int64) (map[int64]Extension, bool) {
	var exts map[int64]Extension
	exts = make(map[int64]Extension, 0)

	sqlstr := "SELECT id, domain_id, extension_name, extension_number, callout_number, " +
		"extension_type, group_id, extension_pswd, extension_login_state, " +
		"extension_reg_state, callout_gateway, is_allow_callout, limit_max, " +
		"limit_destination, mailbox, vm_password, accountcode, effective_caller_id_name, " +
		"effective_caller_id_number, outbound_caller_id_name, outbound_caller_id_number," +
		"call_group, hold_music, call_state, is_record, curr_talking, " +
		"domain_name" +
		"FROM nway_callout_extensions   " +
		" where  group_id=" + strconv.FormatInt(group_id, 10)

	rows, bOp := d.dbbase.Query(sqlstr)

	if bOp == false {
		logger.Error("query the data failed ")
		rows.Close()
		return exts, false
	}

	var ext Extension

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ext.Id,
			&ext.Domain_id,
			&ext.Extension_name,
			&ext.Extension_number,
			&ext.Callout_number,
			&ext.Extension_type,
			&ext.Group_id,
			&ext.Extension_pswd,
			&ext.Extension_login_state,
			&ext.Extension_reg_state,
			&ext.Callout_gateway,
			&ext.Is_allow_callout,
			&ext.Limit_max,
			&ext.Limit_destination,
			&ext.Mailbox,
			&ext.Vm_password,
			&ext.Accountcode,
			&ext.Effective_caller_id_name,
			&ext.Effective_caller_id_number,
			&ext.Outbound_caller_id_name,
			&ext.Outbound_caller_id_number,
			&ext.Call_group,
			&ext.Hold_music,
			&ext.Call_state,
			&ext.Is_record,
			&ext.Curr_talking,
			&ext.Domain_name)
		if err != nil {
			logger.Error("scan the data error,", err)
			return exts, false
		}
		exts[ext.Id] = ext
	}
	return exts, true
}

func (d *DbExtension) GetExtById(id int64) (Extension, bool) {
	var ext Extension
	sqlstr := "SELECT id, domain_id, extension_name, extension_number, callout_number, " +
		"extension_type, group_id, extension_pswd, extension_login_state, " +
		"extension_reg_state, callout_gateway, is_allow_callout, limit_max, " +
		"limit_destination, mailbox, vm_password, accountcode, effective_caller_id_name, " +
		"effective_caller_id_number, outbound_caller_id_name, outbound_caller_id_number," +
		"call_group, hold_music, call_state, is_record, curr_talking, " +
		"domain_name" +
		"FROM nway_callout_extensions   " +
		" where  id=" + strconv.FormatInt(id, 10)

	rows, bOp := d.dbbase.Query(sqlstr)

	if bOp == false {
		logger.Error("query the data failed ")
		rows.Close()
		return ext, false
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ext.Id,
			&ext.Domain_id,
			&ext.Extension_name,
			&ext.Extension_number,
			&ext.Callout_number,
			&ext.Extension_type,
			&ext.Group_id,
			&ext.Extension_pswd,
			&ext.Extension_login_state,
			&ext.Extension_reg_state,
			&ext.Callout_gateway,
			&ext.Is_allow_callout,
			&ext.Limit_max,
			&ext.Limit_destination,
			&ext.Mailbox,
			&ext.Vm_password,
			&ext.Accountcode,
			&ext.Effective_caller_id_name,
			&ext.Effective_caller_id_number,
			&ext.Outbound_caller_id_name,
			&ext.Outbound_caller_id_number,
			&ext.Call_group,
			&ext.Hold_music,
			&ext.Call_state,
			&ext.Is_record,
			&ext.Curr_talking,
			&ext.Domain_name)
		if err != nil {
			logger.Error("scan the data error,", err)
			return ext, false
		}

	}
	return ext, true
}

func (d *DbExtension) GetExtByExtNumber(ext_number string) (Extension, bool) {
	var ext Extension
	sqlstr := "SELECT id, domain_id, extension_name, extension_number, callout_number, " +
		"extension_type, group_id, extension_pswd, extension_login_state, " +
		"extension_reg_state, callout_gateway, is_allow_callout, limit_max, " +
		"limit_destination, mailbox, vm_password, accountcode, effective_caller_id_name, " +
		"effective_caller_id_number, outbound_caller_id_name, outbound_caller_id_number," +
		"call_group, hold_music, call_state, is_record, curr_talking, " +
		"domain_name" +
		"FROM nway_callout_extensions   " +
		" where  extension_number='" + ext_number + "'"

	rows, bOp := d.dbbase.Query(sqlstr)

	if bOp == false {
		logger.Error("query the data failed ")
		rows.Close()
		return ext, false
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ext.Id,
			&ext.Domain_id,
			&ext.Extension_name,
			&ext.Extension_number,
			&ext.Callout_number,
			&ext.Extension_type,
			&ext.Group_id,
			&ext.Extension_pswd,
			&ext.Extension_login_state,
			&ext.Extension_reg_state,
			&ext.Callout_gateway,
			&ext.Is_allow_callout,
			&ext.Limit_max,
			&ext.Limit_destination,
			&ext.Mailbox,
			&ext.Vm_password,
			&ext.Accountcode,
			&ext.Effective_caller_id_name,
			&ext.Effective_caller_id_number,
			&ext.Outbound_caller_id_name,
			&ext.Outbound_caller_id_number,
			&ext.Call_group,
			&ext.Hold_music,
			&ext.Call_state,
			&ext.Is_record,
			&ext.Curr_talking,
			&ext.Domain_name)
		if err != nil {
			logger.Error("scan the data error,", err)
			return ext, false
		}

	}
	return ext, true
}

func (d *DbExtension) GetExtId(ext string, call_numberid int64) (int64, bool) {

	var extid int64

	sqlstr := "SELECT a.id FROM nway_callout_extensions a,nway_callout_numbers b  " +
		" where  a.extension_number='" + ext +
		"' and a.domain_id=b.domain_id and b.id=" +
		strconv.FormatInt(call_numberid, 10)

	rows, bOp := d.dbbase.Query(sqlstr)

	if bOp == false {
		logger.Error("query the data failed ")
		rows.Close()
		return 0, false
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&extid)
		if err != nil {
			logger.Error("scan the data error,", err)
			return 0, false
		}

	}
	return extid, true
}
