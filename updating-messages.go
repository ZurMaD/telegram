package telegram

// EditMessageTextResponse interface
type EditMessageTextResponse interface {
	Response
	GetEditedMessage() *Message
}

type editMessageTextResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *editMessageTextResponse) GetEditedMessage() *Message {
	return r.Result
}

func (b *bot) EditMessageText(options ...Option) (EditMessageTextResponse, error) {
	var res editMessageTextResponse
	err := doRequest(b.token, "editMessageText", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// EditMessageCaptionResponse interface
type EditMessageCaptionResponse interface {
	Response
	GetEditedMessage() *Message
}

type editMessageCaptionResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *editMessageCaptionResponse) GetEditedMessage() *Message {
	return r.Result
}

func (b *bot) EditMessageCaption(options ...Option) (EditMessageCaptionResponse, error) {
	var res editMessageCaptionResponse
	err := doRequest(b.token, "editMessageCaption", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// EditMessageMediaResponse interface
type EditMessageMediaResponse interface {
	Response
	GetEditedMessage() *Message
}

type editMessageMediaResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *editMessageMediaResponse) GetEditedMessage() *Message {
	return r.Result
}

func (b *bot) EditMessageMedia(options ...Option) (EditMessageMediaResponse, error) {
	var res editMessageMediaResponse
	err := doRequest(b.token, "editMessageMedia", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// EditMessageReplyMarkupResponse interface
type EditMessageReplyMarkupResponse interface {
	Response
	GetEditedMessage() *Message
}

type editMessageReplyMarkupResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *editMessageReplyMarkupResponse) GetEditedMessage() *Message {
	return r.Result
}

func (b *bot) EditMessageReplyMarkup(options ...Option) (EditMessageReplyMarkupResponse, error) {
	var res editMessageReplyMarkupResponse
	err := doRequest(b.token, "editMessageReplyMarkup", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// StopPollResponse interface
type StopPollResponse interface {
	Response
	GetStoppedPoll() *Poll
}

type stopPollResponse struct {
	response
	Result *Poll `json:"result,omitempty"`
}

func (r *stopPollResponse) GetStoppedPoll() *Poll {
	return r.Result
}

func (b *bot) StopPoll(options ...Option) (StopPollResponse, error) {
	var res stopPollResponse
	err := doRequest(b.token, "stopPoll", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// DeleteMessageResponse interface
type DeleteMessageResponse interface {
	Response
}

type deleteMessageResponse struct {
	response
}

func (b *bot) DeleteMessage(options ...Option) (DeleteMessageResponse, error) {
	var res deleteMessageResponse
	err := doRequest(b.token, "deleteMessage", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
