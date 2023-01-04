package info

import (
	"context"

	info_pb "github.com/manhrev/IOT/server/pkg/api"
	"github.com/manhrev/IOT/server/pkg/ent"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *infoServer) ListGroup(ctx context.Context, request *info_pb.ListGroupRequest) (*info_pb.ListGroupReply, error) {
	groupList, err := s.repository.Group.List(ctx)
	if err != nil {
		return nil, err
	}
	return &info_pb.ListGroupReply{
		Info: convertEntGroupToGroup(groupList),
	}, nil
}

func convertEntGroupToGroup(entGroupList []*ent.Group) []*info_pb.GroupInfo {
	groupListPb := []*info_pb.GroupInfo{}
	for _, entGroup := range entGroupList {
		feedNameList := []string{}
		for _, feedname := range entGroup.Edges.Feeds {
			feedNameList = append(feedNameList, feedname.FeedName)
		}
		groupPb := &info_pb.GroupInfo{
			GroupId:   int64(entGroup.ID),
			GroupName: (entGroup.GroupName),
			FeedList:  feedNameList,
			CreatedAt: timestamppb.New(entGroup.CreatedAt),
		}
		groupListPb = append(groupListPb, groupPb)
	}
	return groupListPb
}
