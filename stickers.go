package telegram

// Sticker struct
type Sticker struct {
	FileID       string        `json:"file_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	IsAnimated   bool          `json:"is_animated"`
	Thumb        *PhotoSize    `json:"thumb,omitempty"`
	Emoji        *string       `json:"emoji,omitempty"`
	SetName      *string       `json:"set_name,omitempty"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	FileSize     *int          `json:"file_size,omitempty"`
}

// StickerSet struct
type StickerSet struct {
	Name          string    `json:"name"`
	Title         string    `json:"title"`
	IsAnimated    bool      `json:"is_animated"`
	ContainsMasks bool      `json:"contains_masks"`
	Stickers      []Sticker `json:"stickers"`
}

// MaskPosition struct
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float32 `json:"x_shift"`
	YShift float32 `json:"y_shift"`
	Scale  float32 `json:"scale"`
}

// SendStickerResponse interface
type SendStickerResponse interface {
	Response
	GetMessage() *Message
}

type sendStickerResponse struct {
	response
	Result *Message `json:"result,omitempty"`
}

func (r *sendStickerResponse) GetMessage() *Message {
	return r.Result
}

func (b *bot) SendSticker(options ...Option) (SendStickerResponse, error) {
	var res sendStickerResponse
	err := doRequest(b.token, "sendSticker", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GetStickerSetResponse interface
type GetStickerSetResponse interface {
	Response
	GetStickerSet() *StickerSet
}

type getStickerSetResponse struct {
	response
	Result *StickerSet `json:"result,omitempty"`
}

func (r *getStickerSetResponse) GetStickerSet() *StickerSet {
	return r.Result
}

func (b *bot) GetStickerSet(options ...Option) (GetStickerSetResponse, error) {
	var res getStickerSetResponse
	err := doRequest(b.token, "getStickerSet", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// UploadStickerFileResponse interface
type UploadStickerFileResponse interface {
	Response
	GetUploadedFile() *File
}

type uploadStickerFileResponse struct {
	response
	Result *File `json:"result,omitempty"`
}

func (r *uploadStickerFileResponse) GetUploadedFile() *File {
	return r.Result
}

func (b *bot) UploadStickerFile(options ...Option) (UploadStickerFileResponse, error) {
	var res uploadStickerFileResponse
	err := doRequest(b.token, "uploadStickerFile", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// CreateNewStickerSetResponse interface
type CreateNewStickerSetResponse interface {
	Response
}

type createNewStickerSetResponse struct {
	response
}

func (b *bot) CreateNewStickerSet(options ...Option) (CreateNewStickerSetResponse, error) {
	var res createNewStickerSetResponse
	err := doRequest(b.token, "createNewStickerSet", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// AddStickerToSetResponse interface
type AddStickerToSetResponse interface {
	Response
}

type addStickerToSetResponse struct {
	response
}

func (b *bot) AddStickerToSet(options ...Option) (AddStickerToSetResponse, error) {
	var res addStickerToSetResponse
	err := doRequest(b.token, "addStickerToSet", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// SetStickerPositionInSetResponse interface
type SetStickerPositionInSetResponse interface {
	Response
}

type setStickerPositionInSetResponse struct {
	response
}

func (b *bot) SetStickerPositionInSet(options ...Option) (SetStickerPositionInSetResponse, error) {
	var res setStickerPositionInSetResponse
	err := doRequest(b.token, "setStickerPositionInSet", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// DeleteStickerFromSetResponse interface
type DeleteStickerFromSetResponse interface {
	Response
}

type deleteStickerFromSetResponse struct {
	response
}

func (b *bot) DeleteStickerFromSet(options ...Option) (DeleteStickerFromSetResponse, error) {
	var res deleteStickerFromSetResponse
	err := doRequest(b.token, "deleteStickerFromSet", &res, options...)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
