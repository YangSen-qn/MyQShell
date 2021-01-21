package message

type StringOutputMessage struct {
	info string
}

func NewStringOutputData(info string) *StringOutputMessage {
	return &StringOutputMessage{info: info}
}

func (message *StringOutputMessage) String() string {
	return message.info
}

