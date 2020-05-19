package api


import (
    "golang.org/x/net/websocket"

    "words_craft/player"
)

var API_NUM_FUNC_MAP = map[int] func(*player.User, map[string]interface{})(interface{}, error) {
    10001: apiTest,
    10002: login,
    10003: regist,

    10101: enterGame,
    10102: commitAnswer,
}

//9xxxx
var API_NUM_FUNC_WITH_WS_MAP =  map[int] func(*player.User, *websocket.Conn, map[string]interface{})(interface{}, error) {
}
