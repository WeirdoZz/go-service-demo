package server

import (
	"context"
	"encoding/json"
	"tag-service/pkg/bapi"
	"tag-service/pkg/errcode"
	pb "tag-service/proto"
)

type TagService struct{}

// NewTagService 新建一个TagService
func NewTagService() *TagService {
	return &TagService{}
}

func (t *TagService) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewAPI("http:localhost:9999")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	}

	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, errcode.TogRPCError(errcode.Fail)
	}

	return &tagList, nil
}
