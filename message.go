package tomqclient

const (
	MESSAGE_WAITING_NACK = 1
)

type message struct {
	id    string
	queue string
	body  []byte
}

func newMessage(id string, queueName string, body []byte) *message {
	m := message{
		id:    id,
		queue: queueName,
		body:  body,
	}
	return &m
}
