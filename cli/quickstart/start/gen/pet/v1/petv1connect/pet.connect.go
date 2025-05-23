// Copyright 2020-2025 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: pet/v1/pet.proto

package petv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/bufbuild/buf-examples/gen/pet/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// PetStoreName is the fully-qualified name of the PetStore service.
	PetStoreName = "pet.v1.PetStore"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PetStoreGetPetProcedure is the fully-qualified name of the PetStore's GetPet RPC.
	PetStoreGetPetProcedure = "/pet.v1.PetStore/GetPet"
	// PetStorePutPetProcedure is the fully-qualified name of the PetStore's PutPet RPC.
	PetStorePutPetProcedure = "/pet.v1.PetStore/PutPet"
	// PetStoreDeletePetProcedure is the fully-qualified name of the PetStore's DeletePet RPC.
	PetStoreDeletePetProcedure = "/pet.v1.PetStore/DeletePet"
)

// PetStoreClient is a client for the pet.v1.PetStore service.
type PetStoreClient interface {
	GetPet(context.Context, *connect.Request[v1.GetPetRequest]) (*connect.Response[v1.GetPetResponse], error)
	PutPet(context.Context, *connect.Request[v1.PutPetRequest]) (*connect.Response[v1.PutPetResponse], error)
	DeletePet(context.Context, *connect.Request[v1.DeletePetRequest]) (*connect.Response[v1.DeletePetResponse], error)
}

// NewPetStoreClient constructs a client for the pet.v1.PetStore service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPetStoreClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) PetStoreClient {
	baseURL = strings.TrimRight(baseURL, "/")
	petStoreMethods := v1.File_pet_v1_pet_proto.Services().ByName("PetStore").Methods()
	return &petStoreClient{
		getPet: connect.NewClient[v1.GetPetRequest, v1.GetPetResponse](
			httpClient,
			baseURL+PetStoreGetPetProcedure,
			connect.WithSchema(petStoreMethods.ByName("GetPet")),
			connect.WithClientOptions(opts...),
		),
		putPet: connect.NewClient[v1.PutPetRequest, v1.PutPetResponse](
			httpClient,
			baseURL+PetStorePutPetProcedure,
			connect.WithSchema(petStoreMethods.ByName("PutPet")),
			connect.WithClientOptions(opts...),
		),
		deletePet: connect.NewClient[v1.DeletePetRequest, v1.DeletePetResponse](
			httpClient,
			baseURL+PetStoreDeletePetProcedure,
			connect.WithSchema(petStoreMethods.ByName("DeletePet")),
			connect.WithClientOptions(opts...),
		),
	}
}

// petStoreClient implements PetStoreClient.
type petStoreClient struct {
	getPet    *connect.Client[v1.GetPetRequest, v1.GetPetResponse]
	putPet    *connect.Client[v1.PutPetRequest, v1.PutPetResponse]
	deletePet *connect.Client[v1.DeletePetRequest, v1.DeletePetResponse]
}

// GetPet calls pet.v1.PetStore.GetPet.
func (c *petStoreClient) GetPet(ctx context.Context, req *connect.Request[v1.GetPetRequest]) (*connect.Response[v1.GetPetResponse], error) {
	return c.getPet.CallUnary(ctx, req)
}

// PutPet calls pet.v1.PetStore.PutPet.
func (c *petStoreClient) PutPet(ctx context.Context, req *connect.Request[v1.PutPetRequest]) (*connect.Response[v1.PutPetResponse], error) {
	return c.putPet.CallUnary(ctx, req)
}

// DeletePet calls pet.v1.PetStore.DeletePet.
func (c *petStoreClient) DeletePet(ctx context.Context, req *connect.Request[v1.DeletePetRequest]) (*connect.Response[v1.DeletePetResponse], error) {
	return c.deletePet.CallUnary(ctx, req)
}

// PetStoreHandler is an implementation of the pet.v1.PetStore service.
type PetStoreHandler interface {
	GetPet(context.Context, *connect.Request[v1.GetPetRequest]) (*connect.Response[v1.GetPetResponse], error)
	PutPet(context.Context, *connect.Request[v1.PutPetRequest]) (*connect.Response[v1.PutPetResponse], error)
	DeletePet(context.Context, *connect.Request[v1.DeletePetRequest]) (*connect.Response[v1.DeletePetResponse], error)
}

// NewPetStoreHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPetStoreHandler(svc PetStoreHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	petStoreMethods := v1.File_pet_v1_pet_proto.Services().ByName("PetStore").Methods()
	petStoreGetPetHandler := connect.NewUnaryHandler(
		PetStoreGetPetProcedure,
		svc.GetPet,
		connect.WithSchema(petStoreMethods.ByName("GetPet")),
		connect.WithHandlerOptions(opts...),
	)
	petStorePutPetHandler := connect.NewUnaryHandler(
		PetStorePutPetProcedure,
		svc.PutPet,
		connect.WithSchema(petStoreMethods.ByName("PutPet")),
		connect.WithHandlerOptions(opts...),
	)
	petStoreDeletePetHandler := connect.NewUnaryHandler(
		PetStoreDeletePetProcedure,
		svc.DeletePet,
		connect.WithSchema(petStoreMethods.ByName("DeletePet")),
		connect.WithHandlerOptions(opts...),
	)
	return "/pet.v1.PetStore/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PetStoreGetPetProcedure:
			petStoreGetPetHandler.ServeHTTP(w, r)
		case PetStorePutPetProcedure:
			petStorePutPetHandler.ServeHTTP(w, r)
		case PetStoreDeletePetProcedure:
			petStoreDeletePetHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPetStoreHandler returns CodeUnimplemented from all methods.
type UnimplementedPetStoreHandler struct{}

func (UnimplementedPetStoreHandler) GetPet(context.Context, *connect.Request[v1.GetPetRequest]) (*connect.Response[v1.GetPetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pet.v1.PetStore.GetPet is not implemented"))
}

func (UnimplementedPetStoreHandler) PutPet(context.Context, *connect.Request[v1.PutPetRequest]) (*connect.Response[v1.PutPetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pet.v1.PetStore.PutPet is not implemented"))
}

func (UnimplementedPetStoreHandler) DeletePet(context.Context, *connect.Request[v1.DeletePetRequest]) (*connect.Response[v1.DeletePetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("pet.v1.PetStore.DeletePet is not implemented"))
}
