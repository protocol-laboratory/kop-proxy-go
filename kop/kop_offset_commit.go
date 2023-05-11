package kop

import (
	"fmt"
	"github.com/protocol-laboratory/kafka-codec-go/codec"
	"github.com/protocol-laboratory/kop-proxy-go/constant"
	"github.com/sirupsen/logrus"
)

func (b *Broker) OffsetCommitVersion(ctx *NetworkContext, req *codec.OffsetCommitReq) (*codec.OffsetCommitResp, error) {
	if !b.checkSasl(ctx) {
		return nil, fmt.Errorf("connection is not authed")
	}
	logrus.Debugf("offset commit req: %+v", req)
	resp := make([]*codec.OffsetCommitTopicResp, len(req.TopicReqList))
	for i, topicReq := range req.TopicReqList {
		if !b.checkSaslTopic(ctx, topicReq.Topic, constant.ConsumerPermissionType) {
			return nil, fmt.Errorf("check topic read permission failed:  %s", topicReq.Topic)
		}
		f := &codec.OffsetCommitTopicResp{
			Topic:             topicReq.Topic,
			PartitionRespList: make([]*codec.OffsetCommitPartitionResp, len(topicReq.PartitionReqList)),
		}
		for j, partitionReq := range topicReq.PartitionReqList {
			var err error
			f.PartitionRespList[j], err = b.OffsetCommitPartitionAction(ctx.Addr, topicReq.Topic, req.ClientId, partitionReq)
			if err != nil {
				return nil, fmt.Errorf("get offset commit partition failed: %v", err)
			}
		}
		resp[i] = f
	}
	return &codec.OffsetCommitResp{
		BaseResp: codec.BaseResp{
			CorrelationId: req.CorrelationId,
		},
		TopicRespList: resp,
	}, nil
}
