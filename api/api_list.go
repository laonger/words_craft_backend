package api

var API_NUM_FUNC_MAP = map[int] func(interface{})(interface{}, error) {
    10001: enterGame,
}
