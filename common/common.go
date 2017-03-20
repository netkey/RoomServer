package common


type ErrorCode uint16
const  (
	EC_NONE ErrorCode = iota
	EC_DB
	EC_PARAM
	EC_ServiceBusy

//user
	EC_PW_NOT_MATCH
	EC_USER_NOT_LOGIN
	EC_USER_GOLE_NOT_ENOUGH
	EC_WRONG_BETTING
	EC_WRONG_TICKET
	EC_GAME_GENERAL_ERROR
)




func (c ErrorCode) Code() (r uint16) {
	r = uint16(c)
	return
}


