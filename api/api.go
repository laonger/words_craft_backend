package api

import (
    "fmt"
    "strconv"
    //"time"
    "errors"
    //"words_craft/model"
    "words_craft/net"
    "words_craft/game"
    "words_craft/player"
)

func apiTest(user *player.User, params map[string]interface{}) (interface{}, error) {
    fmt.Println("api", params)
    return params, nil
}

func regist(user *player.User, params map[string]interface{}) (interface{}, error) {
    name := params["name"].(string)
    pass := params["pass"].(string)
    nick := params["nick_name"].(string)

    if (name == "" || pass=="" || nick=="") {
        return nil, API_Error{10003001}
    }
    userRecord, err := player.Regist(name, pass, nick)
    if (errors.Is(err, player.ErrorUserExist{})) {
        return nil, API_Error{10003001}
    } else {
        user.Id = userRecord.Id
        user.UserName = userRecord.UserName
        user.NickName = userRecord.NickName
        user.Level = userRecord.Level
        user.Link = nil
        return API_Player {
            Id: user.Id,
            Name: user.NickName,
            Level: user.Level,
        }, nil
    }
    return nil, nil
}

// params = {"name": string, "pass": string}
func login(user *player.User, params map[string]interface{}) (interface{}, error) {
    name := params["name"].(string)
    pass := params["pass"].(string)
    if (name == "" || pass=="") {
        return nil, API_Error{10002001}
    }

    userRecord, err := player.Login(name, pass)

    if (errors.Is(err, player.ErrorWrongPass{}) || errors.Is(err, player.ErrorNoUser{})){
        return nil, API_Error{10002001}
    } else {
        user.Id = userRecord.Id
        user.UserName = userRecord.UserName
        user.NickName = userRecord.NickName
        user.Level = userRecord.Level
        user.Link = nil
        return API_Player {
            Id: user.Id,
            Name: user.NickName,
            Level: user.Level,
        }, nil
    }
}

func enterGame(user *player.User, params map[string]interface{}) (interface{}, error){
    /*
    word_1 := 
    word_2 := WordStruct{
        [ ]string{ "学校", "大学", "日记", "汽车", "没有正确答案" },
        "school",
        0,
    }
    */
    var gameData API_Game
    if aORb, room, err := game.JoinGame(user); aORb == 1 {
        if err != nil {
        }
        gameData = API_Game{
            RoomId: room.RoomId,
            NeedWait: true,
        }
    } else {
        word, _ := game.NextQuestion(room)
        otherGameData := API_Game{
            Other: API_Player{ user.Id, user.NickName, user.Level},
            Word: word,
            Num: room.QuestionNum,
            RoomId: room.RoomId,
            QuestionAmount: len(room.Questions),
            Timestamp: room.StartTime.Unix(),
            NeedWait: false,
        }
        net.Send(room.UserA.Link, 10101, otherGameData)

        gameData = API_Game{
            Other: API_Player{ room.UserA.Id, room.UserA.NickName, room.UserA.Level},
            Word: word,
            Num: room.QuestionNum,
            RoomId: room.RoomId,
            QuestionAmount: len(room.Questions),
            Timestamp: room.StartTime.Unix(),
            NeedWait: false,
        }

    }
    return gameData, nil
}

func commitAnswer(user *player.User, params map[string]interface{}) (interface{}, error) {
    answer, _ := strconv.Atoi(params["answer"].(string))
    useTime, _ := strconv.Atoi(params["useTime"].(string))
    needWait, getScore, otherScore, realAnswer, err := game.CommitAnswer(user, answer, useTime)
    if err != nil {
        return nil, nil
    } else {
        if needWait {
            return API_ResponseCommitAnswer{
                NeedWait: needWait,
                GetScore: getScore,
            }, nil
        } else {
            room := game.Rooms[user.RoomId]
            word, _ := game.NextQuestion(room)

            //// 给另对方发送新的问题
            var other *player.User
            if room.UserA.Id == user.Id {
                other = room.UserB
            } else {
                other = room.UserA
            }
            net.Send(other.Link, 10102, API_ResponseCommitAnswer{
                Answer: realAnswer,
                GetScore: otherScore,
                NeedWait: needWait,
                Num: room.QuestionNum,
                OtherScore: getScore,
                Word: word,
            })
            ///////

            return API_ResponseCommitAnswer{
                Answer: realAnswer,
                NeedWait: needWait,
                GetScore: getScore,
                OtherScore: otherScore,
                Num: room.QuestionNum,
                Word: word,
            }, nil
        }
    }
}

