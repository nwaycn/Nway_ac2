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

type Ring struct {
	Id               int64
	Ring_name        string
	Ring_path        string
	Ring_description string
	Ring_category    int64
	Domain_id        int64
}

type DbRing struct {
	dbbase nway_db_base.DbBase
}

func (d *DbRing) GetPathById(ring_id int64) (string, bool) {
	sqlstr := "SELECT  ring_name " +
		" FROM nway_callout_rings " +
		" where  id=" + strconv.FormatInt(ring_id, 10)

	rows, bOp := d.dbbase.Query(sqlstr)

	if bOp == false {
		logger.Error("query the data failed ")
		rows.Close()
		return "", false
	}
	var ringpath string
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ringpath)
		if err != nil {
			logger.Error("scan the data error,", err)
			return "", false
		}
	}
	return ringpath, true
}
