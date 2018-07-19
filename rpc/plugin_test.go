package rpc

// import (
// 	"fmt"
// 	"net"
// 	"net/rpc"
// 	"net/rpc/jsonrpc"
// 	"os"
// 	"testing"
// 	"time"
// )

// const hbsServer = "192.168.20.163:6030"

// func TestConcurency(t *testing.T) {
// 	go func() {
// 		fmt.Println("concurency excecution!")
// 	}()
// 	fmt.Println("heheh")
// }

// func TestPluginExecuteStatusReport(t *testing.T) {

// 	timeout := time.Duration(1) * time.Second
// 	client, err := JsonRpcClient("tcp", hbsServer, timeout)
// 	if err != nil {
// 		fmt.Println("dial failed:", err)
// 		os.Exit(1)
// 	}
// 	req := PluginStatusRequest{
// 		IP:           "192.168.20.163",
// 		AgentVersion: "1.1",
// 		Plugin: PluginExecInfo{
// 			Name:    "test2",
// 			Version: "1.2.1",
// 			Status:  1,
// 		},
// 		Timestamp: time.Now().Unix(),
// 	}

// 	var res SimpleResponse
// 	count := 1
// 	for {
// 		if count == 10000 {
// 			break
// 		}
// 		go func(client *rpc.Client) {
// 			client.Call("Plugin.ReportStatus", req, &res)
// 		}(client)
// 		count++
// 	}
// 	time.Sleep(time.Millisecond * 10000)
// 	defer client.Close()
// }

// func TestPlugin(t *testing.T) {
// 	timeout := time.Duration(1) * time.Second
// 	client, err := JsonRpcClient("tcp", hbsServer, timeout)
// 	if err != nil {
// 		fmt.Println("dial failed:", err)
// 		os.Exit(1)
// 	}
// 	req := SimplePluginRequest{
// 		IP:        "127.0.0.1",
// 		Timestamp: time.Now().Unix(),
// 	}
// 	var res PluginListResponse
// 	err = client.Call("Plugin.GetExecPluginList", req, &res)

// 	// err := c.Call("Plugin.GetExecPluginList", req, &resp)
// 	if err != nil {
// 		fmt.Println("Send to \"hbs\" failed:", err, "resp:", res)
// 		return
// 	}
// 	// fmt.Println(res.Plugins)
// 	for _, plugin := range res.Plugins {
// 		fmt.Println(plugin)
// 	}
// 	fmt.Println(res.Plugins)
// 	defer client.Close()
// }

// func TestAgentReportStatus(t *testing.T) {
// 	// var response PluginListResponse
// 	timeout := time.Duration(1) * time.Second
// 	client, err := JsonRpcClient("tcp", hbsServer, timeout)
// 	if err != nil {
// 		fmt.Println("dial failed:", err)
// 		os.Exit(1)
// 	}
// 	defer client.Close()

// 	var response SimpleResponse

// 	plugin := PluginExecInfo{
// 		Name:    "mac",
// 		Version: "2.0",
// 		Status:  StatusOK,
// 	}
// 	request := PluginStatusRequest{
// 		IP:           "127.0.0.1",
// 		AgentVersion: "2.0",
// 		Plugin:       plugin,
// 		Timestamp:    time.Now().Unix(),
// 	}

// 	err = client.Call("Plugin.ReportStatus", request, &response)
// 	if err != nil {
// 		fmt.Println("rpc call failed", err)
// 	}
// 	fmt.Println(response)
// }

// func JsonRpcClient(network, address string, timeout time.Duration) (*rpc.Client, error) {
// 	conn, err := net.DialTimeout(network, address, timeout)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return jsonrpc.NewClient(conn), err
// }
