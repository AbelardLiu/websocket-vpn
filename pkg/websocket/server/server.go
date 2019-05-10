package server

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"time"
	"lll.github.com/websocket-vpn/pkg/websocket/common"
)

func svrConnHandler(conn *websocket.Conn) {
	request := make([]byte, 128);
	defer conn.Close();
	for {
		readLen, err := conn.Read(request)
		if common.CheckErr(err, "Read") {
			break;
		}

		//socket被关闭了
		if readLen == 0 {
			fmt.Println("Client connection close!");
			break;
		} else {
			//输出接收到的信息
			fmt.Println(string(request[:readLen]))

			time.Sleep(time.Second);
			//发送
			conn.Write([]byte("World !"));
		}

		request = make([]byte, 128);
	}
}

func Run() error {
	http.Handle("/echo", websocket.Handler(svrConnHandler));
	err := http.ListenAndServe(":6666", nil);
	common.CheckErr(err, "ListenAndServe");
	fmt.Println("Func finish.");

	return nil
}
