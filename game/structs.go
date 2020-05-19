package game

import (
    "time"
    "words_craft/player"
)

var TIME_COUNTDOWN = 10

type Word struct {
    Options []string `json: "chinese"`
    Question string `json: "word"`
    Answer int `json: "answer"`
}

type questionStruct struct {
    question string
    options []string
    aAnswer int
    bAnswer int
    aTime time.Duration
    bTime time.Duration
    aUseTime int
    bUseTime int
    aScore int  // 此题获得分数
    bScore int
    answer int
}

type RoomStruct struct {
    RoomId string
    Level string
    UserA *player.User    // userId
    UserB *player.User    // userId
    Questions []questionStruct
    StartTime time.Time
    QuestionNum int
    QuestionAmount int
}
