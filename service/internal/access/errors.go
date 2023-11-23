package access

import "errors"

var ErrInvalidOtp = errors.New("invalid otp")
var ErrTokenGenerationFailed = errors.New("token generation failed")
