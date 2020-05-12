package api

import (
    "fmt"
    //"time"
    "golang.org/x/net/websocket"
    model "words_craft/model"
    game "words_craft/game"
    net "words_craft/net"
)

func apiTest(params map[string]interface{}) (interface{}, error) {
    fmt.Println("api", params)
    return params, nil
}

func regist(params map[string]interface{}) (interface{}, error) {
    name := params["name"].(string)
    pass := params["pass"].(string)
    if (name == "" || pass=="") {
        return nil, API_Error{10003001}
    }
    _, ok := model.USER_PASS_MAP[name]
    if ok {
        return nil, API_Error{10003001}
    } else {
        record := map[string]string {
            "pass": pass,
            "id": name,
        }
        model.USER_PASS_MAP[name] = record
        player := model.User {
            Id: name,
            Name: name,
            Level: 0,
        }
        model.USER[name] = player
        return player, nil
    }
    return nil, nil
}

// params = {"name": string, "pass": string}
func login(params map[string]interface{}) (interface{}, error) {
    name := params["name"].(string)
    pass := params["pass"].(string)

    pass_record := model.USER_PASS_MAP[name]

    if (name == "" || pass=="" || pass_record["pass"] != pass) {
        return nil, API_Error{10002001}
    } else {
        player_record := model.USER[pass_record["id"]]
        player := API_Player{
            Id: player_record.Id,
            Name: player_record.Name,
            Level: player_record.Level,
        }
        return player, nil
    }
}

func enterGame(ws *websocket.Conn, params map[string]interface{}) (interface{}, error){
    /*
    word_1 := 
    word_2 := WordStruct{
        [ ]string{ "学校", "大学", "日记", "汽车", "没有正确答案" },
        "school",
        0,
    }
    */
    var gameData API_Game
    if aORb, room, err := game.JoinGame("aaaa", ws); aORb == 1 {
        if err != nil {
        }
        gameData = API_Game{
            RoomNum: room.RoomId,
            NeedWait: true,
        }
    } else {
        word, _ := game.NextQuestion(room)
        otherGameData := API_Game{
            Other: API_Player{ room.UserA, "我", 1},
            Word: word,
            Num: room.QuestionNum +1,
            RoomNum: room.RoomId,
            QuestionAmount: len(room.Questions),
            Timestamp: room.StartTime.Unix(),
            NeedWait: false,
        }
        net.Send(room.UserALink, 90001, otherGameData)

        gameData = API_Game{
            Other: API_Player{ room.UserA, "他", 1},
            Word: word,
            Num: room.QuestionNum +1,
            RoomNum: room.RoomId,
            QuestionAmount: len(room.Questions),
            Timestamp: room.StartTime.Unix(),
            NeedWait: false,
        }

    }
    return gameData, nil
}
