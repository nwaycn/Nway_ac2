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
	_ "strconv"

	_ "github.com/nwaycn/Nway_ac2/ac/nway-pbx/config"
	_ "github.com/nwaycn/Nway_ac2/ac/nway-pbx/nway_database/nway_connection"
	"github.com/nwaycn/Nway_ac2/ac/nway-pbx/nway_database/nway_db_base"
	"github.com/nwaycn/Nway_ac2/ac/nway-util/logger"
)

type Task struct {
	Id                      int64
	Domain_id               int64
	Callout_name            string
	Number_group_id         int64
	Number_group_uploadfile string
	Run_position            int64
	Time_rule_id            int64
	Start_time              string
	Stop_time               string
	Ring_id                 int64
	After_ring_play         int32
	Ring_timeout            int32
	Group_id                int64
	Call_project_id         int64  //
	Concurr_type_id         int64  // -- 并发类型，0为按在线坐席数量的比例，1为指定值
	Concurr_number          int32  //-- 并发倍数，按并发类型处理并发数
	Callout_state_id        int64  //
	Total_number            int32  //
	Wait_number             int32  // -- 等待数量
	Success_number          int32  // -- 接通数量
	Failed_number           int32  // -- 接通失败数量
	Cancel_number           int32  // -- 取消的数量
	Has_parse_from_file     bool   // -- 当上传了文件后，是否从文件中解析了内容插到数据表中，解析结束后置为true
	Callout_gateway_id      int64  //
	Max_concurr_number      int32  // -- 最大并发数，前一个concurr_number为并发倍数
	Second_ring_id          int64  // -- 由after_ring_play定为播放彩铃生效
	Second_after_ring_opt   int32  // -- 第二次再播放后的操作，和call_after_opt对应
	After_ring_key          string // -- 播放语音时按键中止播放
	After_key_opt_id        int32  // -- 按键后的操作，和call_after_opt对应
	Outside_line_number     string // -- 外呼时，如手机上显示的来电号码，需运营商许可通过
}
type DbTask struct {
	dbbase nway_db_base.DbBase
}

func (d *DbTask) GetTasks() (map[int64]Task, bool) {
	var tasks map[int64]Task
	tasks = make(map[int64]Task, 0)

	//
	sqlstr := "SELECT id, domain_id, callout_name, number_group_id, number_group_uploadfile, " +
		"run_position, time_rule_id, start_time, stop_time, ring_id, after_ring_play, " +
		"ring_timeout, group_id, call_project_id, concurr_type_id, concurr_number, " +
		"callout_state_id, total_number, wait_number, success_number, " +
		"failed_number, cancel_number, has_parse_from_file, callout_gateway_id, " +
		"max_concurr_number, second_ring_id, second_after_ring_opt, after_ring_key, " +
		"after_key_opt_id, outside_line_number " +
		"FROM nway_callout_tasks" +
		" where  where t.start_time < now() and t.stop_time >now() " +
		" and (callout_state_id=1 or callout_state_id=2 or callout_state_id=4) and has_parse_from_file=True "

	rows, bOp := d.dbbase.Query(sqlstr)

	if bOp == false {
		logger.Error("query the data failed ")
		rows.Close()
		return tasks, false
	}

	var task Task

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&task.Id,
			&task.Domain_id,
			&task.Callout_name,
			&task.Number_group_id,
			&task.Number_group_uploadfile,
			&task.Run_position,
			&task.Time_rule_id,
			&task.Start_time,
			&task.Stop_time,
			&task.Ring_id,
			&task.After_ring_play,
			&task.Ring_timeout,
			&task.Group_id,
			&task.Call_project_id,
			&task.Concurr_type_id,
			&task.Concurr_number,
			&task.Callout_state_id,
			&task.Total_number,
			&task.Wait_number,
			&task.Success_number,
			&task.Failed_number,
			&task.Cancel_number,
			&task.Has_parse_from_file,
			&task.Callout_gateway_id,
			&task.Max_concurr_number,
			&task.Second_ring_id,
			&task.Second_after_ring_opt,
			&task.After_ring_key,
			&task.After_key_opt_id,
			&task.Outside_line_number)
		if err != nil {
			logger.Error("scan the data error,", err)
			return tasks, false
		}
		tasks[task.Id] = task
	}
	return tasks, true
}
