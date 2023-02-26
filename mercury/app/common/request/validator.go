package request

type Validator interface {
	GetMessages() ValidatorMessages
}

type ValidatorMessages map[string]string
