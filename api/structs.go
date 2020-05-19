package api
/*
    'players': {
        'own': {
            'id': 'nnnn',
            'name': '我'
        },
        'other': {
            'id': 'aaaa',
            'name': '你'
        }
    },
    'words': [
        {
            'word': 'apple',
            'chinese': [
                '苹果',
                '香蕉',
                '李子',
                '梨',
                '没有正确答案'
            ],
            'answer': 0
        },
        {
            'word': 'school',
            'chinese': [
                '学校',
                '大学',
                '日记',
                '汽车',
                '没有正确答案'
            ],
            'answer': 0
        }
    ]
*/

import "words_craft/game"

type API_ErrorMessage struct {
    Message string  `json: "message"`
    ErrorNum int64    `json: "error_num"`
    IsError bool     `json: "is_error"`
}

type API_Player struct {
    Id string `json: "id"`
    Name string `json: "name"`
    Level int `json: "level"`
}


type API_Game struct {
    Other API_Player `json: "other"`
    Word game.Word `json: "word"`
    RoomId string `json: "room_num"`
    QuestionAmount int `json: "question_amount"`
    Num int `json: "num"`
    Timestamp int64 `json: "timestamp"`
    NeedWait bool `json: "need_wait"`
}

type API_UserNamePass struct {
    Name string `json: "name"`
    Pass string `json: "pass"`
}

type API_ResponseCommitAnswer struct {
    Word game.Word `json: "word"`
    Answer string `json: "answer"`
    Num int `json: "num"`
    NeedWait bool `json: "need_wait"`
    GetScore int `json: "getScore"`
    OtherScore int `json: "otherScore"`
}
