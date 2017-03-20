package gate

type Agent interface {
	WriteMsg(msg interface{},code uint16)
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
}
