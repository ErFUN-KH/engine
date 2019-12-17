// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/misc/proto/misc.proto

package miscpb

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/gogo/protobuf/gogoproto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	elbix_dev_engine_pkg_grpcgw "elbix.dev/engine/pkg/grpcgw"
	gopkg_in_go_playground_validator_v9 "gopkg.in/go-playground/validator.v9"
	golang_org_x_net_context "golang.org/x/net/context"
	github_com_grpc_ecosystem_grpc_gateway_runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	elbix_dev_engine_pkg_assert "elbix.dev/engine/pkg/assert"
	google_golang_org_grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WrappedMiscSystemController interface {
	MiscSystemServer
	elbix_dev_engine_pkg_grpcgw.Controller
}

type wrappedMiscSystemServer struct {
	original MiscSystemServer
	v        *gopkg_in_go_playground_validator_v9.Validate
}

func (w *wrappedMiscSystemServer) Init(ctx golang_org_x_net_context.Context, conn *google_golang_org_grpc.ClientConn, mux *github_com_grpc_ecosystem_grpc_gateway_runtime.ServeMux) {
	cl := NewMiscSystemClient(conn)

	elbix_dev_engine_pkg_assert.Nil(RegisterMiscSystemHandlerClient(ctx, mux, cl))
}

func (w *wrappedMiscSystemServer) InitGRPC(ctx golang_org_x_net_context.Context, s *google_golang_org_grpc.Server) {
	RegisterMiscSystemServer(s, w)
}

func (w *wrappedMiscSystemServer) Version(ctx golang_org_x_net_context.Context, req *VersionRequest) (res *VersionResponse, err error) {
	ctx, err = elbix_dev_engine_pkg_grpcgw.ExecuteMiddleware(ctx, w.original)
	if err != nil {
		return nil, err
	}
	if err = w.v.StructCtx(ctx, req); err != nil {
		return nil, elbix_dev_engine_pkg_grpcgw.NewBadRequest(err, "validation failed")
	}

	res, err = w.original.Version(ctx, req)
	return
}

func (w *wrappedMiscSystemServer) Health(ctx golang_org_x_net_context.Context, req *HealthRequest) (res *HealthResponse, err error) {
	ctx, err = elbix_dev_engine_pkg_grpcgw.ExecuteMiddleware(ctx, w.original)
	if err != nil {
		return nil, err
	}
	if err = w.v.StructCtx(ctx, req); err != nil {
		return nil, elbix_dev_engine_pkg_grpcgw.NewBadRequest(err, "validation failed")
	}

	res, err = w.original.Health(ctx, req)
	return
}

func (w *wrappedMiscSystemServer) PublicKey(ctx golang_org_x_net_context.Context, req *PubKeyRequest) (res *PubKeyResponse, err error) {
	ctx, err = elbix_dev_engine_pkg_grpcgw.ExecuteMiddleware(ctx, w.original)
	if err != nil {
		return nil, err
	}
	if err = w.v.StructCtx(ctx, req); err != nil {
		return nil, elbix_dev_engine_pkg_grpcgw.NewBadRequest(err, "validation failed")
	}

	res, err = w.original.PublicKey(ctx, req)
	return
}

func NewWrappedMiscSystemServer(server MiscSystemServer) WrappedMiscSystemController {
	return &wrappedMiscSystemServer{
		original: server,
		v:        gopkg_in_go_playground_validator_v9.New(),
	}
}
func init() {
}
