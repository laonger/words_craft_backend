package api

func enterGame(requestObj interface{}) (interface{}, error){
    /*
    word_1 := 
    word_2 := WordStruct{
        []string{ "学校", "大学", "日记", "汽车", "没有正确答案" },
        "school",
        0,
    }
    */
    gameData := GameData{
        PlayerStruct{ "aaaa", "你" },
        PlayerStruct{ "bbbb", "他" },
        WordStruct{
            Chinese: []string{ "苹果", "香蕉", "李子", "梨", "没有正确答案" },
            Word: "apple",
            Answer: 0,
            Num: 1,
        },
        1,
        2,
        0,
    }
    return &gameData, nil
}
