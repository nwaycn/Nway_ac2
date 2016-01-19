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
	"errors"

	"github.com/nwaycn/Nway_ac2/ac/nway-pbx/config"
	. "github.com/nwaycn/Nway_ac2/ac/nway-util/esl/goesl"
	"github.com/nwaycn/Nway_ac2/ac/nway-util/logger"
	"github.com/nwaycn/Nway_ac2/ac/nway-util/regex"
)

type OutboundService struct {
}

//一般outbound用于走呼入流程会更好
func (this *OutboundService) Start() error {
	nwayconfig, result := pbxconfig.LoadConf()
	if result == false {
		logger.Error("load config file failed")
	}
	if len(nwayconfig.Fsoutboundconf.ObPort) > 0 && len(nwayconfig.Fsoutboundconf.ObPPort) > 0 {
		s, err := NewOutboundServer(":" + nwayconfig.Fsoutboundconf.ObPort)
		if err != nil {
			logger.Error("Start OutboundService Failed ", err)
		} else {
			go handle(s)
			s.Start()
		}

	} else {
		return errors.New("Not Found the outboud port or public outbound port")
	}

	return nil
}
func handle(s *OutboundServer) {
	for {

		select {

		case conn := <-s.Conns:
			Notice("New incomming connection: %v", conn)

			if err := conn.Connect(); err != nil {
				logger.Error("Got error while accepting connection: %s", err)
				//break
			}

		default:
			// YabbaDabbaDooooo!
			//Flintstones. Meet the Flintstones. They're the modern stone age family. From the town of Bedrock,
			// They're a page right out of history. La la,lalalalala la :D
		}
	}
}

//处理outound的函数
func process_outbound(c *SocketConnection) error {

	msg, err := c.ReadMessage()
	if err != nil {
		logger.Error("Read esl Message error ", err)
	}

	err = c.Send("linger")
	if err != nil {
		logger.Error(err)
	}
	caller_uuid := msg.GetHeader("Caller-Unique-ID")
	content_type := msg.GetHeader("Content-Type")
	event_name := msg.GetHeader("Event-Name")
	call_direction := msg.GetHeader("Call-Direction")
	b_uuid := msg.GetHeader("Unique-ID")
	channel_state := msg.GetHeader("Channel-State")
	caller_destination_number := msg.GetHeader("Caller-Destination-Number")
	caller_username := msg.GetHeader("Caller-Username")
	caller_number := msg.GetHeader("Caller-Caller-ID-Number")
	channel_call_state := msg.GetHeader("Channel-Call-State")

	gateway_name := msg.GetHeader("Gateway-Name")
	if (len(gateway_name) > 0) && (call_direction == "inbound") {
		//由外线呼入的
		Debug("caller id ", caller_uuid)
		Debug("content_type ", content_type)
		Debug("event_name ", event_name)
		Debug("b_uuid ", b_uuid)
		Debug("channel_state ", channel_state)
		Debug("caller_destination_number ", caller_destination_number)
		Debug("caller_username ", caller_username)
		Debug("caller_number ", caller_number)
		Debug("channel_call_state ", channel_call_state)
		//通过呼入号码查找外线呼入的号码或ivr等
		outbRegex := nwayregex.NewRegex()
		//Debug(outbRegex)
		_, _, _ = outbRegex.Match("18621575908")

	} else {
		//由内呼外的

	}

	return err
}

/*

func main() {
	eventsocket.ListenAndServe(":9090", handler)
}

func handler(c *eventsocket.Connection) {
	fmt.Println("new client:", c.RemoteAddr())
	c.Send("connect")
	c.Send("myevents")
	c.Execute("answer", "", false)
	ev, err := c.Execute("playback", audioFile, true)
	if err != nil {
		log.Fatal(err)
	}
	ev.PrettyPrint()
	for {
		ev, err = c.ReadEvent()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("\nNew event")
		ev.PrettyPrint()
		if ev.Get("Application") == "playback" {
			if ev.Get("Application-Response") == "FILE PLAYED" {
				c.Send("exit")
			}
		}
	}
}
*/
