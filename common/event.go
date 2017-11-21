package common

const (
	AUTH_LOGIN    = "AUTH_LOGIN"    // map[string]string{username, password}
	AUTH_SIGNUP   = "AUTH_SIGNUP"   // map[string]string{username, password}
	AUTH_BIND     = "AUTH_BIND"     // token string
	AUTH_ERROR    = "AUTH_ERROR"    // string
	AUTH_SUCCESS  = "AUTH_SUCCESS"  // string
	TOKEN_RECEIPT = "TOKEN_RECEIPT" // string

	GAME_RECEIPT      = "GAME_RECEIPT"      // Game{}
	GAME_RECEIPT_LIST = "GAME_RECEIPT_LIST" // []Game{}
	GAME_CREATE       = "GAME_CREATE"       // string
	GAME_SELECT       = "GAME_SELECT"       // int
	GAME_ERROR        = "GAME_ERROR"        // string

	MAP_SEARCH = "MAP_SEARCH" // -
)
