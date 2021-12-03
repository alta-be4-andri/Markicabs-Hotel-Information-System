package response

import (
	"net/http"
)

// function response false param
func FalseParamResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "False Param",
	}
	return result
}

// function response bad request
func BadRequestResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Bad Request",
	}
	return result
}

// function response access forbidden
func AccessForbiddenResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Access Forbidden",
	}
	return result
}

// function response success dengan paramater
func SuccessResponseData(data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
		"data":    data,
	}
	return result
}

// function response success tanpa parameter
func SuccessResponseNonData() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Successful Operation",
	}
	return result
}

// function response login failure
func LoginFailedResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Email or Password Invalid",
	}
	return result
}

// function response login success
func LoginSuccessResponse(data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Login Success",
		"data":    data,
	}
	return result
}

// function response success dengan paramater
func ReservationSuccessResponse(data interface{}) map[string]interface{} {
	result := map[string]interface{}{
		"code":           http.StatusOK,
		"message":        "Your Resevation Success",
		"reservation_id": data,
	}
	return result
}

// function response check room availability failed
func CheckFailedResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Room not available",
	}
	return result
}

// function response check room availability success
func CheckSuccessResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Room available",
	}
	return result
}

func PasswordCannotLess5() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "password cannot less than 5 character",
	}
	return result
}

func NameCannotEmpty() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "username cannot be empty",
	}
	return result
}

func EmailCannotEmpty() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "email cannot be empty",
	}
	return result
}

func IsExist() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": "Email or Phone Number is Exist",
	}
	return result
}

func FormatEmailInvalid() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Email must contain email format",
	}
	return result
}

// function response homestay not found
func HomestayNotFoundResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Homestay not found",
	}
	return result
}

// function response room not found
func RoomNotFoundResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Room not found",
	}
	return result
}

// function response for date invalid, if check out date earlier than check in
func DateInvalidResponse() map[string]interface{} {
	result := map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": "Check In and Check Out Date Invalid",
	}
	return result
}
