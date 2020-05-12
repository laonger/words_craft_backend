package game

import (
    "fmt"
    "time"
    "errors"
    "golang.org/x/net/websocket"
    //"words_craft/model"
)


// userId: roomId
var PlayerRoomMap = map[string]string {
}

// 正在进行游戏的房间
var Rooms = map[string]*RoomStruct {
}
var roomsLock = false

// 等代配对的房间
var tempRooms = []*RoomStruct {}
var tempRoomsWriteLock = false


func makeRoomId(userId string) string{
    return userId
}

func newQuestions(level string) ([]questionStruct, error){
    return []questionStruct{
            questionStruct {
                question: "apple",
                options: []string{ "苹果", "香蕉", "李子", "梨", "没有正确答案" },
                aAnswer: -1,
                aTime: 0,
                bAnswer: -1,
                bTime: 0,
                answer: 0,
            },
            questionStruct {
                question: "banana",
                options: []string{ "苹果", "香蕉", "李子", "梨", "没有正确答案" },
                aAnswer: -1,
                aTime: 0,
                bAnswer: -1,
                bTime: 0,
                answer: 1,
            },
            questionStruct {
                question: "pear",
                options: []string{ "苹果", "香蕉", "李子", "梨", "没有正确答案" },
                aAnswer: -1,
                aTime: 0,
                bAnswer: -1,
                bTime: 0,
                answer: 2,
            },
    }, nil
}

func NextQuestion(room *RoomStruct) (Word, error) {

    if (room.QuestionNum == len(room.Questions)-1) {
        return Word{}, errors.New("超出上限")
    }

    room.QuestionNum = room.QuestionNum+1
    question := &room.Questions[room.QuestionNum]

    word := Word{
        Options: question.options,
        Question: question.question,
        Answer: question.answer,
    }
    return word, nil
}



// 如果加入的是userA则返回1，如果是userB则返回2
func JoinGame(userId string, userLink *websocket.Conn) (aORb int, room *RoomStruct, err error) {
    tempRoomsWriteLock = true
    if (len(tempRooms) > 0){
        room = tempRooms[0]
        tempRooms = tempRooms[1:]
        if (len(tempRooms) == 0){
            tempRooms = []*RoomStruct{}
        }
        room.UserB = userId
        room.UserBLink = userLink
        questions, _ := newQuestions(room.Level)
        room.Questions = questions
        room.StartTime = time.Now()

        roomsLock = true
        Rooms[room.RoomId] = room
        roomsLock = false

        fmt.Println(room)
        tempRoomsWriteLock = false
        return 2, room, nil
    } else {
        room = &RoomStruct{
            RoomId: makeRoomId(userId),
            Level: "",
            UserA: userId,
            UserB: "",
            UserALink: userLink,
            UserBLink: new(websocket.Conn),
            Questions: []questionStruct{},
            StartTime: time.Now(),
            QuestionNum: -1,
        }
        tempRooms = append(tempRooms, room)
        tempRoomsWriteLock = false
        return 1, room, nil
    }
}

