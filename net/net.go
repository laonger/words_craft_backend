package net

import (
    "strings"
    "strconv"
    "encoding/json"
    "golang.org/x/net/websocket"
)

func DecodeRequest(requestString string) (api_num int, requestData map[string]interface{}, err error) {
    var requestObj interface{}
    l := strings.Split(requestString, "_%%_")
    api_num_string := l[0]
    api_num, err = strconv.Atoi(api_num_string)
    if err != nil {
        println(6, "e", err)
        return 6, nil, nil
    }
    println("byte", []byte(l[1]))
    if err := json.Unmarshal([]byte(l[1]), &requestObj); err != nil {
        println(2, "e", err)
        return 2, nil, nil
    }
    // 运行处理函数
    requestData = requestObj.(map[string]interface{})
    return api_num, requestData, nil
}

func encodeResponse(api_num int, requestData interface{}) (string, error) {
    responseJson, err := json.Marshal(requestData)
    if err != nil {
        println(4, "e", err)
        return "4", nil
    }
    responseString := strconv.FormatInt(int64(api_num), 10)+"_%%_"+string(responseJson)
    return responseString, nil
}

func Send(ws *websocket.Conn, api_num int, data interface{}) error {
    responseString, err := encodeResponse(api_num, data)
    if err = websocket.Message.Send(ws, responseString); err != nil {
        println(5, "e", err)
        return nil
    }
    println(2, "responseJson", string(responseString))
    return nil
}

func SendString(ws *websocket.Conn, api_num int, s string) error {
    responseString := strconv.FormatInt(int64(api_num), 10)+"_%%_"+s
    if err := websocket.Message.Send(ws, responseString); err != nil {
        println(5, "e", err)
        return nil
    }
    println(2, "responseJson", string(responseString))
    return nil
}

