package main

import (
    "os"
    "log"
    "golang.org/x/net/websocket"
    "net/http"

    "words_craft/net"
    "words_craft/api"
    "words_craft/player"
)


func litsen(ws *websocket.Conn) {
    user := player.User{}
    user.Link = nil
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
        result, err = api.API_NUM_FUNC_MAP[api_num](&user, params)
        if err != nil {
            println(5, "result: ", api_num, " :", err)
            err = net.SendString(ws, api_num, err.Error())
            if err != nil {
                println(4, "e", err)
                continue
            }
        } else {
            println(7, "result", api_num, ":", result)
            if user.Link == nil {
                user.Link = ws
            }
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


