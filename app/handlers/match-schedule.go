package handlers

import (
	"context"
	"encoding/json"
	"match-schedule/app/services/core"
	"match-schedule/app/services/group"
	"match-schedule/proto"
)

type GrpcServer struct{}

func (g GrpcServer) GetSchedule(ctx context.Context, r *proto.MatchScheduleReq) (*proto.MatchScheduleResp, error) {
	var groups [][][]int32
	var err error
	if r.Amplitude == 0 {
		groups, err = group.GenGroups(r.PlayerNum, r.FieldNum, r.RoundNum, r.Mode)
	} else {
		groups, err = group.GenGroups(r.PlayerNum, r.FieldNum, r.RoundNum, r.Mode, core.WithAmplitude(r.Amplitude))
	}
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(groups)
	if err != nil {
		return nil, err
	}
	return &proto.MatchScheduleResp{
		Schedule: string(data),
	}, nil
}
