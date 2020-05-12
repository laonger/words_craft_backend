package game

import (
    "time"
    "golang.org/x/net/websocket"
)

type Word struct {
    Options []string `json: "chinese"`
    Question string `json: "word"`
    Answer int `json: "answer"`
}

type questionStruct struct {
    question string
    options []string
    aAnswer int
    aTime time.Duration
    bAnswer int
    bTime time.Duration
    answer int
}

type RoomStruct struct {
    RoomId string
    Level string
    UserA string    // userId
    UserB string    // userId
    UserALink *websocket.Conn
    UserBLink *websocket.Conn
    Questions []questionStruct
    StartTime time.Time
    QuestionNum int
    QuestionAmount int
}
