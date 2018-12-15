// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/misc/proto/misc.proto

package miscpb

import github_com_fzerorubigd_balloon_pkg_grpcgw "github.com/fzerorubigd/balloon/pkg/grpcgw"
import gopkg_in_go_playground_validator_v9 "gopkg.in/go-playground/validator.v9"
import golang_org_x_net_context "golang.org/x/net/context"
import github_com_fullstorydev_grpchan_inprocgrpc "github.com/fullstorydev/grpchan/inprocgrpc"
import github_com_grpc_ecosystem_grpc_gateway_runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
import github_com_fzerorubigd_balloon_pkg_assert "github.com/fzerorubigd/balloon/pkg/assert"
import github_com_fzerorubigd_balloon_pkg_log "github.com/fzerorubigd/balloon/pkg/log"
import github_com_pkg_errors "github.com/pkg/errors"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WrappedMiscSystemController interface {
	MiscSystemServer
	github_com_fzerorubigd_balloon_pkg_grpcgw.Controller
}

type wrappedMiscSystemServer struct {
	original MiscSystemServer
	v        *gopkg_in_go_playground_validator_v9.Validate
}

func (w *wrappedMiscSystemServer) Init(ctx golang_org_x_net_context.Context, ch *github_com_fullstorydev_grpchan_inprocgrpc.Channel, mux *github_com_grpc_ecosystem_grpc_gateway_runtime.ServeMux) {
	RegisterHandlerMiscSystem(ch, w)
	cl := NewMiscSystemChannelClient(ch)

	github_com_fzerorubigd_balloon_pkg_assert.Nil(RegisterMiscSystemHandlerClient(ctx, mux, cl))
}

func (w *wrappedMiscSystemServer) Version(ctx golang_org_x_net_context.Context, req *VersionRequest) (res *VersionResponse, err error) {
	github_com_fzerorubigd_balloon_pkg_log.Info("MiscSystem.Version request")
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		github_com_fzerorubigd_balloon_pkg_log.Error("Recovering from panic", github_com_fzerorubigd_balloon_pkg_log.Any("panic", e))
		res, err = nil, github_com_pkg_errors.New("internal server error")
	}()
	ctx, err = github_com_fzerorubigd_balloon_pkg_grpcgw.ExecuteMiddleware(ctx, w.original)
	if err != nil {
		return nil, err
	}
	if err = w.v.Struct(req); err != nil {
		return nil, err
	}

	res, err = w.original.Version(ctx, req)
	return
}

func NewWrappedMiscSystemServer(server MiscSystemServer) WrappedMiscSystemController {
	return &wrappedMiscSystemServer{
		original: server,
		v:        gopkg_in_go_playground_validator_v9.New(),
	}
}

/*
map[string]string{}
*/
