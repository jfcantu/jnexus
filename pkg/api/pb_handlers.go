package api

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jfcantu/jnexus/pb"
	"github.com/sirupsen/logrus"
)

// UpdateLinkStatus handles a link status update from the client
func (s *Server) UpdateLinkStatus(ctx context.Context, m *pb.LinkStatus) (*empty.Empty, error) {
	logrus.Debug(m)
	// Handle netsplit notifications slightly differently
	if m.Status == pb.LinkState_INACTIVE {
		linkData, err := s.DBClient.GetLinkStatus(m.Server1, m.Server2)
		if err != nil {
			return nil, err
		}

		// If this is a link that's not in the route map - delete it rather than retain it
		if linkData.Type == "SECONDARY" {
			if err := s.DBClient.DeleteLink(m.Server1, m.Server2); err != nil {
				return nil, err
			}
			return &empty.Empty{}, nil
		}

		// Otherwise, mark it as INACTIVE
		s.DBClient.UpdateLink(m.Server1, m.Server2, "INACTIVE")
		return &empty.Empty{}, nil
	}

	if err := s.DBClient.UpdateLink(m.Server1, m.Server2, m.Status.String()); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
