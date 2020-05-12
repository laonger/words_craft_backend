package api

import (
    "encoding/json"
)

type API_Error struct {
    errorNum int64
}

func (ae API_Error) Error() string {
    result := API_ErrorMessage{
        IsError: true,
        Message: ERROR_MESSAGES[ae.errorNum],
        ErrorNum: ae.errorNum,
    }
    responseJson, err := json.Marshal(result)
    if err != nil {
        println(4, "e", err)
        return "4"
    }

    return string(responseJson)
}




var ERROR_MESSAGES = map[int64]string {
    10003001: "用户名密码错误",
    10003002: "用户已存在",
    10002001: "用户名密码错误",
}

