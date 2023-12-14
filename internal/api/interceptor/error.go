package interceptor

import (
	"context"
	"errors"
	"fmt"
	"users/internal/common/sys"
	appCode "users/internal/common/sys/code"
	"users/internal/common/sys/validate"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCStatusInterface interface {
	GRPCStatus() *status.Status
}

func ErrorCodesInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res interface{}, err error) {
	res, err = handler(ctx, req)
	if nil == err {
		return res, nil
	}

	fmt.Printf(color.RedString("error: %s\n", err.Error()))

	switch {
	case sys.IsCommonError(err):
		commEr := sys.GetCommonError(err)
		code := toGRPCCode(commEr.Code())

		err = status.Error(code, commEr.Error())

	case validate.IsValidationError(err):
		err = status.Error(codes.InvalidArgument, err.Error())

	default:
		var se GRPCStatusInterface
		if errors.As(err, &se) {
			return nil, se.GRPCStatus().Err()
		} else {
			if errors.Is(err, context.DeadlineExceeded) {
				err = status.Error(codes.DeadlineExceeded, err.Error())
			} else if errors.Is(err, context.Canceled) {
				err = status.Error(codes.Canceled, err.Error())
			} else {
				err = status.Error(codes.Internal, "internal error")
			}
		}
	}

	return res, err
}

func toGRPCCode(code appCode.ErrorCode) codes.Code {
	var res codes.Code

	switch code {
	case code.OK:
		res = codes.OK
	case code.Canceled:
		res = codes.Canceled
	case code.InvalidArgument:
		res = codes.InvalidArgument
	case code.DeadlineExceeded:
		res = codes.DeadlineExceeded
	case code.NotFound:
		res = codes.NotFound
	case code.AlreadyExists:
		res = codes.AlreadyExists
	case code.PermissionDenied:
		res = codes.PermissionDenied
	case code.ResourceExhausted:
		res = codes.ResourceExhausted
	case code.FailedPrecondition:
		res = codes.FailedPrecondition
	case code.Aborted:
		res = codes.Aborted
	case code.OutOfRange:
		res = codes.OutOfRange
	case code.Unimplemented:
		res = codes.Unimplemented
	case code.Internal:
		res = codes.Internal
	case code.Unavailable:
		res = codes.Unavailable
	case code.DataLoss:
		res = codes.DataLoss
	case code.Unauthenticated:
		res = codes.Unauthenticated
	default:
		res = codes.Unknown
	}

	return res
}
