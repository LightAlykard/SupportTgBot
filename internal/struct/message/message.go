package messege

type BotMessage struct {
	Message struct {
		Message_id int
		From       struct {
			Username string
			Id       int
		}
		Chat struct {
			Id int
		}
		Text string
	}
}

type BotSendMessageID struct {
	Result struct {
		Message_id int
	}
}
