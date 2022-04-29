/**
  @author: ZYL
  @date:
  @note
*/
package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)

var node *sf.Node

// startTime 起始时间戳 "2022-04-02"
// machineID 主机id 1
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}
func GenID() int64 {
	return node.Generate().Int64()
}

//func main() {
//	if err := Init("2022-04-01", 1); err != nil {
//		fmt.Printf("init failed, err:%v\n", err)
//		return
//	}
//	id := GenID()
//	fmt.Println(id)
//}
