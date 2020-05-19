package player

import (
    "golang.org/x/net/websocket"
)

type User struct {
    Id string
    NickName string
    UserName string
    RoomId string
    Link *websocket.Conn
    Level int

}


