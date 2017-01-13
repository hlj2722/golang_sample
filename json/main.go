package main

import(
	"encoding/json"
	"fmt"
)

var str = []byte(`
    {
  "msg-type": "response",
  "payload": {
    "members": {
      "8a0ec22a": {
        "ip": "192.168.100.227",
        "role": "master",
        "state": "online"
      },
      "aaaaaaa": {
        "ip": "192.168.100.228",
        "role": "master",
        "state": "online"
      }
    },
    "result": "OK"
  },
  "src_ip": "localhost",
  "target": "client",
  "version": "1.0"
}`)

func main() {
	resp := map[string]interface{}{}
	json.Unmarshal(str, &resp)
	fmt.Printf("src_ip is %v or %s\n", resp["src_ip"], resp["src_ip"].(string))
	payload := resp["payload"].(map[string]interface{})
	fmt.Printf("payload.members is %v \n", payload["members"])
	fmt.Printf("payload.result is %v or %s\n", payload["result"], payload["result"])

	members := payload["members"].(map[string]interface{})
	fmt.Printf("members.8a0ec22a is %s\n", members["8a0ec22a"])

	for k, v := range payload["members"].(map[string]interface{}) {
		fmt.Printf("k=%v, v=%v\n", k, v)
		mem := members[k].(map[string]interface{})
		fmt.Println("mem.ip:", mem["ip"])
	}

}
