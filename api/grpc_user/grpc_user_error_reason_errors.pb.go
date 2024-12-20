// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package grpc_user

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUserUnspecified(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_UNSPECIFIED.String() && e.Code == 500
}

func ErrorUserUnspecified(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_USER_UNSPECIFIED.String(), fmt.Sprintf(format, args...))
}

func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_NOT_FOUND.String() && e.Code == 500
}

func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsUserAlreadyExists(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_ALREADY_EXISTS.String() && e.Code == 500
}

func ErrorUserAlreadyExists(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_USER_ALREADY_EXISTS.String(), fmt.Sprintf(format, args...))
}

func IsInvalidArgument(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_INVALID_ARGUMENT.String() && e.Code == 500
}

func ErrorInvalidArgument(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_INVALID_ARGUMENT.String(), fmt.Sprintf(format, args...))
}
