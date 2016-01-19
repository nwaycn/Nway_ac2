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
package pbxmodule


import (
	_"github.com/nwaycn/Nway_ac2/ac/nway-util/logger"
)

type ModuleBase interface{
	ModuleName() string     //module name
	Path() string
}

