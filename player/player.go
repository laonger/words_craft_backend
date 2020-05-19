package player

import (
    "words_craft/model"
)


func Login(userName string, password string) (*User, error) {
    passRecord, ok := model.USER_PASS_MAP[userName]
    if !ok {
        return &User{}, ErrorNoUser{}
    }

    if (passRecord["pass"] != password) {
        return &User{}, ErrorWrongPass{}
    } else {
        userRecord := model.USER[passRecord["id"]]
        user := User{
            Id: userRecord.Id,
            UserName: userRecord.Id,
            NickName: userRecord.NickName,
            Level: userRecord.Level,
        }
        return &user, nil
    }
}

func Regist(userName string, password string, nickName string) (*User, error) {
    _, ok := model.USER_PASS_MAP[userName]
    if ok {
        return &User{}, ErrorUserExist{}
    } else {
        record := map[string]string {
            "pass": password,
            "id": userName,
        }
        model.USER_PASS_MAP[userName] = record
        userRecord := model.User {
            Id: userName,
            NickName: nickName,
            Level: 0,
        }
        model.USER[userName] = userRecord

        user := User{
            Id: userRecord.Id,
            UserName: userRecord.Id,
            NickName: userRecord.NickName,
            Level: userRecord.Level,
        }
        return &user, nil
    }
}

