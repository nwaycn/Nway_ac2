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

/*
package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("^86([1][358][0-9]{9})$")
	fmt.Println(re.FindAllStringSubmatch("8618621575908", -1))
	
	fmt.Println(re.SubexpNames()[1])
}
*/

package nwayregex

import (
	"regexp"
	//"regexp/syntax"
)

type nwayRegex struct{
	M map[int64]regexp.Regexp
}

func NewRegex() *nwayRegex{
	var nr nwayRegex;
	nr.M =  make(map[int64]regexp.Regexp)
	return &nr
}
func CleanRegex(n *nwayRegex){
	for k,_ := range n.M{
		delete(n.M ,k)
	}
}

func (n *nwayRegex)AddMatchString(id int64, matchstring string){
	n.M[id] = *regexp.MustCompile(matchstring)
}

func (n *nwayRegex)Match(inputstring string)(bool, int64, *regexp.Regexp){
	for k,v := range n.M{
		if len(v.FindAllStringSubmatch(inputstring, -1))>0{
			//表示找到了这个匹配的
			return true,k,&v
		}
	}
	return false,0,nil
}