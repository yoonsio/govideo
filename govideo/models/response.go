package models

import "sync"

// ErrResponse is json response message for all server errors
type ErrResponse struct {
	Msg  string
	Code int
}

var errResponsePool = sync.Pool{
	New: func() interface{} {
		return &ErrResponse{}
	},
}

// GetErrResponse gets ErrResponse struct from sync pool
func GetErrResponse() *ErrResponse {
	return errResponsePool.Get().(*ErrResponse)
}

// RecycleErrResponse puts back ErrResponse struct back into sync pool
func RecycleErrResponse(errResponse *ErrResponse) {
	errResponsePool.Put(errResponse)
}

// SuccessResponse is json response message for successful execution
type SuccessResponse struct {
	Msg string
}

var successResponsePool = sync.Pool{
	New: func() interface{} {
		return &SuccessResponse{}
	},
}

// GetSuccessResponse gets SuccessResponse struct from sync pool
func GetSuccessResponse() *SuccessResponse {
	return successResponsePool.Get().(*SuccessResponse)
}

// RecycleSuccessResponse puts back SuccessResponse struct back into sync pool
func RecycleSuccessResponse(successResponse *SuccessResponse) {
	successResponsePool.Put(successResponse)
}
