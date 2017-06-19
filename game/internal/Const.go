package internal



const(
	_None_type = int32(0)
	MJPAI_ZFB
	MJPAI_FENG
	MJPAI_WAN
	MJPAI_TIAO
	MJPAI_BING
	MJPAI_HUA
)


type MJPai struct {
	M_Type 		int32
	M_Value 	int32
}

const(
	TotalMarjong = 136
	TotalType = 34
	MarjonInHand = 13
)

const(
	HuType_PingHu 	= 1
	HuType_ZiMo 	= 2

)