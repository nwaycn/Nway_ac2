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

package pbx

import (
	"fmt"
	_ "fmt"
	_ "strconv"
	"strings"

	"github.com/nwaycn/Nway_ac2/ac/nway-pbx/config"
	"github.com/nwaycn/Nway_ac2/ac/nway-pbx/nway_database"
	. "github.com/nwaycn/Nway_ac2/ac/nway-util/esl/goesl"
	"github.com/nwaycn/Nway_ac2/ac/nway-util/logger"
)

/*
说明：
用于处理task这部分，不足则补齐，暂停也需要把当前的任务呼叫清空
以及处理号码相关
*/
type CallNumberManager struct {
	Task_id     int64
	Callnumbers map[int64]nwaydb.CallNumber
	Call_task   nwaydb.Task
}

type TaskManager struct {
	Tasks map[int64]CallNumberManager
}
