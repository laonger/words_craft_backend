package game

import (
    "fmt"
    "time"
    "errors"
    //"words_craft/model"
    "words_craft/player"
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
                answer: 0,
                aAnswer: -2,
                aUseTime : -1,
                bAnswer: -2,
                bTime: 0,
                bUseTime: -1,
                aScore: -1,
                bScore: -1,
            },
            questionStruct {
                question: "banana",
                options: []string{ "苹果", "香蕉", "李子", "梨", "没有正确答案" },
                answer: 1,
                aAnswer: -2,
                aTime: 0,
                aUseTime : -1,
                bAnswer: -2,
                bTime: 0,
                bUseTime: -1,
                aScore: -1,
                bScore: -1,
            },
            questionStruct {
                question: "pear",
                options: []string{ "苹果", "香蕉", "李子", "梨", "没有正确答案" },
                answer: 3,
                aAnswer: -2,
                aTime: 0,
                aUseTime : -1,
                bAnswer: -2,
                bTime: 0,
                bUseTime: -1,
                aScore: -1,
                bScore: -1,
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
func JoinGame(user *player.User) (aORb int, room *RoomStruct, err error) {
    tempRoomsWriteLock = true
    if (len(tempRooms) > 0){
        room = tempRooms[0]
        tempRooms = tempRooms[1:]
        if (len(tempRooms) == 0){
            tempRooms = []*RoomStruct{}
        }
        room.UserB = user
        user.RoomId = room.RoomId
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
            RoomId: makeRoomId(user.Id),
            Level: "",
            UserA: user,
            UserB: nil,
            Questions: []questionStruct{},
            StartTime: time.Now(),
            QuestionNum: -1,
        }
        user.RoomId = room.RoomId
        tempRooms = append(tempRooms, room)
        tempRoomsWriteLock = false
        return 1, room, nil
    }
}

func CommitAnswer(user *player.User, answer int, useTime int) (needWait bool, getScore int, otherScore int, realAnswer string, err error){
    needWait = true
    room := Rooms[user.RoomId]
    getScore = 0
    println("gamegamegamemgae: ", answer, useTime, room.QuestionNum, len(room.Questions))
    question := &room.Questions[room.QuestionNum]
    if answer == question.answer {
        getScore = (TIME_COUNTDOWN - useTime)*10
    }
    if user.Id == room.UserA.Id {
        question.aAnswer = answer
        question.aUseTime = useTime
        question.aScore = getScore
        otherScore = question.bScore
    } else {
        question.bAnswer = answer
        question.bUseTime = useTime
        question.bScore = getScore
        otherScore = question.aScore
    }
    if (question.aAnswer != -2 && question.bAnswer != -2){
        needWait = false
    }
    return needWait, getScore, otherScore, question.options[question.answer], nil

}


