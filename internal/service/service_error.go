package service

type ServiceErrorType int

const (
	ValidationError ServiceErrorType = iota
	OtherError
)

type ServiceError struct {
	Type    ServiceErrorType
	Message string
}

func (e *ServiceError) Error() string {
	return e.Message
}

var (
	ErrDeviceNotFound = &ServiceError{Type: OtherError, Message: "Device not found"}

	ErrDeviceIDEmpty = &ServiceError{Type: ValidationError, Message: "Device ID cannot be empty"}
	ErrToEmpty       = &ServiceError{Type: ValidationError, Message: "To cannot be empty"}
	ErrTypeEmpty     = &ServiceError{Type: ValidationError, Message: "Type cannot be empty"}
	// Image
	ErrImageUrlEmpty      = &ServiceError{Type: ValidationError, Message: "Image url cannot be empty"}
	ErrImageCaptionEmpty  = &ServiceError{Type: ValidationError, Message: "Image caption cannot be empty"}
	ErrImageMimetypeEmpty = &ServiceError{Type: ValidationError, Message: "Image mimetype cannot be empty"}
	// Video
	ErrVideoUrlEmpty      = &ServiceError{Type: ValidationError, Message: "Video url cannot be empty"}
	ErrVideoCaptionEmpty  = &ServiceError{Type: ValidationError, Message: "Video caption cannot be empty"}
	ErrVideoMimetypeEmpty = &ServiceError{Type: ValidationError, Message: "Video mimetype cannot be empty"}
	// Document
	ErrDocumentUrlEmpty      = &ServiceError{Type: ValidationError, Message: "Document url cannot be empty"}
	ErrDocumentFilenameEmpty = &ServiceError{Type: ValidationError, Message: "Document filename cannot be empty"}
	ErrDocumentMimetypeEmpty = &ServiceError{Type: ValidationError, Message: "Document mimetype cannot be empty"}
	ErrDocumentTitleEmpty    = &ServiceError{Type: ValidationError, Message: "Document title cannot be empty"}
	// Audio
	ErrAudioUrlEmpty      = &ServiceError{Type: ValidationError, Message: "Audio url cannot be empty"}
	ErrAudioMimetypeEmpty = &ServiceError{Type: ValidationError, Message: "Audio mimetype cannot be empty"}
	// Location
	ErrLocationLatitudeEmpty  = &ServiceError{Type: ValidationError, Message: "Location latitude cannot be empty"}
	ErrLocationLongitudeEmpty = &ServiceError{Type: ValidationError, Message: "Location longitude cannot be empty"}
	ErrLocationNameEmpty      = &ServiceError{Type: ValidationError, Message: "Location name cannot be empty"}
	ErrLocationAddressEmpty   = &ServiceError{Type: ValidationError, Message: "Location address cannot be empty"}
)
