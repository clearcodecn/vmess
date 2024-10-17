package vmess

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

const testJson = `{
"v": "2",                                       
"ps": "aliasName",                              
"add": "111.111.111.111",                       
"port": "32000",                                
"id": "1386f85e-657b-4d6e-9d56-78badb75e1fd",   
"aid": "100",                                   
"net": "tcp",                                   
"type": "none",                                 
"host": "www.bbb.com",                          
"path": "/",                                    
"tls": "tls"                                    
}`

func TestVmess(t *testing.T) {
	var link Link
	json.Unmarshal([]byte(testJson), &link)

	vmess := link.ToVmessLink()
	fmt.Println(vmess)

	link2, err := FromVmessLink(vmess)
	if err != nil {
		t.Fatal(err)
		return
	}

	if !reflect.DeepEqual(link2, &link) {
		fmt.Println("test not pass")
	}

}
