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

type PlayerStruct struct {
    Id string `json: "id"`
    Name string `json: "name"`
}

type WordStruct struct {
    Chinese []string `json: "chinese"`
    Word string `json: "word"`
    Answer int `json: "answer"`
    Num int `json: "num"`
}

type GameData struct {
    Own PlayerStruct `json: "own"`
    Other PlayerStruct `json: "other"`
    Word WordStruct `json: "word"`
    RoomNum int `json: "room_num"`
    QuestionAmount int `json: "question_amount"`
    Timestamp int `json: "timestamp"`
}
