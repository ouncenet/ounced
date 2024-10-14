package server

import (
	"context"
	"github.com/ouncenet/ounced/cmd/ouncewallet/daemon/pb"
	"github.com/ouncenet/ounced/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
