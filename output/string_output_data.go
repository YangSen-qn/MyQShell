package output

type StringOutputData struct {
	info string
}

func NewStringOutputData(info string) *StringOutputData {
	return &StringOutputData{info: info}
}

func (p *StringOutputData) String() string {
	return p.info
}

func OutputDebugString(output IOutput, message string) {
	OutputDebug(output, NewStringOutputData(message))
}
