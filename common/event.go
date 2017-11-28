package common

const (
	AUTH_LOGIN    = "AUTH_LOGIN"    // map[string]string{username, password}
	AUTH_SIGNUP   = "AUTH_SIGNUP"   // map[string]string{username, password}
	AUTH_BIND     = "AUTH_BIND"     // token string
	AUTH_ERROR    = "AUTH_ERROR"    // string
	AUTH_SUCCESS  = "AUTH_SUCCESS"  // string
	TOKEN_RECEIPT = "TOKEN_RECEIPT" // string

	GAME_CREATE = "GAME_CREATE" // string
	GAME_SELECT = "GAME_SELECT" // int
	GAME_ERROR  = "GAME_ERROR"  // string

	GAME_RECEIPT      = "GAME_RECEIPT"      // Game{}
	GAME_RECEIPT_LIST = "GAME_RECEIPT_LIST" // []Game{}
	// GAME_RECEIPT_MESSAGE = "GAME_RECEIPT_MESSAGE" // {type, string}

	// Map
	MAP_SEARCH  = "MAP_SEARCH"  // -
	MAP_REQUEST = "MAP_REQUEST" // -
	MAP_RECEIPT = "MAP_RECEIPT" // (Map{})

	// Quest
	QUESTS_RECEIPT = "QUESTS_RECEIPT"
	QUEST_HANDLE   = "QUEST_HANDLE"
	QUEST_MARK     = "QUEST_SAVE"

	// Bag
	BAG_RECEIPT = "BAG_RECEIPT" // (Bag{})
	BAG_REQUEST = "BAG_REQUEST" // (ID int)

	// Item
	ITEM_USE = "ITEM_USE" // (ID int)
)
