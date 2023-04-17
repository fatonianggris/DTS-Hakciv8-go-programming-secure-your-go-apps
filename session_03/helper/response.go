package helper

import (
	"net/http"
)

func ResponseSuccessNoData(code int) map[string]interface{} {
	status := map[string]interface{}{
		"code":   code,
		"status": http.StatusText(code),
	}

	return status
}

func ResponseSuccessWithData(code int, data interface{}) map[string]interface{} {
	status := map[string]interface{}{
		"code":   code,
		"status": http.StatusText(code),
		"data":   data,
	}

	return status
}

func ResponseError(code int) map[string]interface{} {
	status := map[string]interface{}{
		"code":  code,
		"error": http.StatusText(code),
	}

	return status
}
