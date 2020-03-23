package main

import (
    "os"
    "log"
    //"strings"
    "golang.org/x/net/websocket"
    "net/http"
    "encoding/json"

    api "word_scraft/api"
)


func litsen(ws *websocket.Conn) {
    println("listening...")
    var err error
    for {
        var request string
        var requestObj interface{}
        err = websocket.Message.Receive(ws, &request)
        if err != nil {
            log.Println(1, "e", err)
            ws.Close()
            break
        }
        if err := json.Unmarshal([]byte(request), &requestObj); err != nil {
            println(2, "e", err)
            continue
        }
        if result, err := api.API_NUM_FUNC_MAP[10001](&requestObj); err != nil {
            println(3, "e", err)
            continue
        } else {
            println(1, "result", result)
            responseJson, err := json.Marshal(result)
            if err != nil {
                println(4, "e", err)
                continue
            }
            if err = websocket.Message.Send(ws, string(responseJson)); err != nil {
                println(5, "e", err)
                continue
            }
            println(2, "responseJson", string(responseJson))
        }
    }
}

func main() {
    http.Handle("/", websocket.Handler(litsen))

    err := http.ListenAndServe(":9999", nil)
    if err != nil {
        println(err)
        os.Exit(1)
    }
}


