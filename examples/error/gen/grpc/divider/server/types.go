// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// divider gRPC server types
//
// Command:
// $ goa gen goa.design/goa/examples/error/design -o
// $(GOPATH)/src/goa.design/goa/examples/error

package server

import (
	dividersvc "goa.design/goa/examples/error/gen/divider"
	dividerpb "goa.design/goa/examples/error/gen/grpc/divider"
)

// NewIntOperands builds the payload of the "integer_divide" endpoint of the
// "divider" service from the gRPC request type.
func NewIntOperands(p *dividerpb.IntOperands) *dividersvc.IntOperands {
	v := &dividersvc.IntOperands{
		A: int(p.A),
		B: int(p.B),
	}
	return v
}

// NewIntegerDivideResponse builds the gRPC response type from the result of
// the "integer_divide" endpoint of the "divider" service.
func NewIntegerDivideResponse(res int) *dividerpb.IntegerDivideResponse {
	v := &dividerpb.IntegerDivideResponse{
		Field: int32(res),
	}
	return v
}

// NewFloatOperands builds the payload of the "divide" endpoint of the
// "divider" service from the gRPC request type.
func NewFloatOperands(p *dividerpb.FloatOperands) *dividersvc.FloatOperands {
	v := &dividersvc.FloatOperands{
		A: p.A,
		B: p.B,
	}
	return v
}

// NewDivideResponse builds the gRPC response type from the result of the
// "divide" endpoint of the "divider" service.
func NewDivideResponse(res float64) *dividerpb.DivideResponse {
	v := &dividerpb.DivideResponse{
		Field: res,
	}
	return v
}