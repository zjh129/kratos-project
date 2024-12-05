// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package http_app

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsGreeterUnspecified(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GREETER_UNSPECIFIED.String() && e.Code == 500
}

func ErrorGreeterUnspecified(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_GREETER_UNSPECIFIED.String(), fmt.Sprintf(format, args...))
}

func IsUserLoginFailed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_LOGIN_FAILED.String() && e.Code == 401
}

func ErrorUserLoginFailed(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_USER_LOGIN_FAILED.String(), fmt.Sprintf(format, args...))
}

func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_NOT_FOUND.String() && e.Code == 404
}

func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}