// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// divider gRPC client types
//
// Command:
// $ goa gen goa.design/goa/examples/error/design -o
// $(GOPATH)/src/goa.design/goa/examples/error

package client

import (
	dividersvc "goa.design/goa/examples/error/gen/divider"
	"goa.design/goa/examples/error/gen/grpc/divider/pb"
)

// NewIntegerDivideRequest builds the gRPC request type from the payload of the
// "integer_divide" endpoint of the "divider" service.
func NewIntegerDivideRequest(payload *dividersvc.IntOperands) *pb.IntegerDivideRequest {
	message := &pb.IntegerDivideRequest{}
	aptr := int32(payload.A)
	message.A = aptr
	bptr := int32(payload.B)
	message.B = bptr
	return message
}

// NewIntegerDivideResponse builds the result type of the "integer_divide"
// endpoint of the "divider" service from the gRPC response type.
func NewIntegerDivideResponse(message *pb.IntegerDivideResponse) int {
	result := int(message.Field)
	return result
}

// NewDivideRequest builds the gRPC request type from the payload of the
// "divide" endpoint of the "divider" service.
func NewDivideRequest(payload *dividersvc.FloatOperands) *pb.DivideRequest {
	message := &pb.DivideRequest{
		A: payload.A,
		B: payload.B,
	}
	return message
}

// NewDivideResponse builds the result type of the "divide" endpoint of the
// "divider" service from the gRPC response type.
func NewDivideResponse(message *pb.DivideResponse) float64 {
	result := message.Field
	return result
}