package model

import (
	"client/common/message"
	"net"
)

//因为在客户端很难躲地方会使用CurUser，我们将其作为一个全局的变量
type CurUser struct {
	Conn net.Conn
	message.User
}
