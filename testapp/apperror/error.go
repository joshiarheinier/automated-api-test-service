package apperror

import "strconv"

const (
	ErrFailedToDecodeConfigurationFile = "Failed to decode configuration file: %v\n"
)

type V1Error struct {
	PublicMessage	string	`json:"public_message"`
	DebugMessage	string	`json:"debug_message"`
	HttpCode		int		`json:"http_code"`
}

var errorData  = make(map[string]*V1Error)

func (e *V1Error) Error() string {
	errorStr := "public_message:" + e.PublicMessage + ";" +
		"" + e.DebugMessage + ";" +
		"" + strconv.Itoa(e.HttpCode)
	return errorStr
}

func NewV1Error(code string, args ...interface{}) *V1Error {
	return errorData[code]
}