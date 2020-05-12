package main

import (
    "os"
    "log"
    "golang.org/x/net/websocket"
    "net/http"

    net "words_craft/net"
    api "words_craft/api"
)


func litsen(ws *websocket.Conn) {
    user = nil
    for {
        var err error
        var request string
        var result interface{}
        err = websocket.Message.Receive(ws, &request)
        if err != nil {
            log.Println(1, "e", err)
            ws.Close()
            break
        }
        api_num, params, err := net.DecodeRequest(request)
        if (api_num == 6 || api_num == 2){
            continue
        }
        // 运行处理函数
        if (api_num > 90000){
            result, err = api.API_NUM_FUNC_WITH_WS_MAP[api_num](ws, params)
        } else {
            result, err = api.API_NUM_FUNC_MAP[api_num](params)
        }
        if err != nil {
            println(5, "result: ", api_num, " :", err)
            err = net.SendString(ws, api_num, err.Error())
            if err != nil {
                println(4, "e", err)
                continue
            }
        } else {
            println(7, "result", api_num, ":", result)
            err = net.Send(ws, api_num, result)
            if err != nil {
                println(4, "e", err)
                continue
            }
        }
    }
}

func main() {
    http.Handle("/", websocket.Handler(litsen))

    err := http.ListenAndServe(":9999", nil)
    if err != nil {
        println(err)
        os.Exit(1)
    } else {
        println("listening...")
    }
}


