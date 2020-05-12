package api


import (
    "golang.org/x/net/websocket"
)

var API_NUM_FUNC_MAP = map[int] func(map[string]interface{})(interface{}, error) {
    10001: apiTest,
    10002: login,
    10003: regist,
}

//9xxxx
var API_NUM_FUNC_WITH_WS_MAP =  map[int] func(*websocket.Conn, map[string]interface{})(interface{}, error) {
    90001: enterGame,
}
