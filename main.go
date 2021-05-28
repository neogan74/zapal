package main

import (
	"fmt"
	"time"

	"github.com/neogan74/go-zabbix"
)

func main() {
	fmt.Println("Hello")

	//	session, err := zabbix.NewSession("http://zabbix/api_jsonrpc.php", "Admin", "zabbix")

	cache := zabbix.NewSessionFileCache().SetFilePath("./zabbix_session")
	session, err := zabbix.CreateClient("http://192.168.31.147/zabbix/api_jsonrpc.php").
		WithCache(cache).
		WithCredentials("Admin", "zabbix").
		WithTimeout(5 * time.Second).
		Connect()

	if err != nil {
		panic(err)
	}
	params := zabbix.HostgroupGetParams{}
	host_groups, err := session.GetHostgroups(params)
	if err != nil {
		panic(err)
	}

	for i, hg := range host_groups {
		fmt.Println(i, hg)
	}

	fmt.Println("Connected to Zabbix API v%s", session.Version())
}
