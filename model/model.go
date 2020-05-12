package model

import (
    "time"
    "context"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func test() error {
   ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
   client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
   if err != nil {
   }
   println(client);
   return nil
}


var USER_PASS_MAP = map[string]map[string]string {
    "aaa": { "pass": "111", "id": "aabba"},
}

type User struct {
    Id string
    Name string
    Level int
}
var USER = map[string]User {
    "aabba": User{
        Id: "aabba",
        Name: "哟",
        Level: 1,
    },
}

type Word struct {
    En string
    Cn string
    Series string
}

var WORDS = map[string]Word {
    "apple": Word {
        "apple",
        "苹果",
        "primary",
    },
    "banana": Word{
        "banana",
        "香蕉",
        "primary",
    },
    "plum": Word{
        "plum",
        "李子",
        "primary",
    },
    "pear": Word{
        "pear",
        "梨",
        "primary",
    },
}
