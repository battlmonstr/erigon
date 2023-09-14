package sentry

import (
	"context"
	grpcSentry "github.com/ledgerwatch/erigon-lib/gointerfaces/sentry"
	"github.com/ledgerwatch/erigon-lib/gointerfaces/types"
)

type InboundMessageResult struct {
	Message *grpcSentry.InboundMessage
	Error   error
}

type PeerEventResult struct {
	Event *grpcSentry.PeerEvent
	Error error
}

type Service interface {
	SetStatus(ctx context.Context, in *grpcSentry.StatusData) (*grpcSentry.SetStatusReply, error)
	PenalizePeer(ctx context.Context, in *grpcSentry.PenalizePeerRequest) error
	PeerMinBlock(ctx context.Context, in *grpcSentry.PeerMinBlockRequest) error
	HandShake(ctx context.Context) (*grpcSentry.HandShakeReply, error)
	SendMessageByMinBlock(ctx context.Context, in *grpcSentry.SendMessageByMinBlockRequest) (*grpcSentry.SentPeers, error)
	SendMessageById(ctx context.Context, in *grpcSentry.SendMessageByIdRequest) (*grpcSentry.SentPeers, error)
	SendMessageToRandomPeers(ctx context.Context, in *grpcSentry.SendMessageToRandomPeersRequest) (*grpcSentry.SentPeers, error)
	SendMessageToAll(ctx context.Context, in *grpcSentry.OutboundMessageData) (*grpcSentry.SentPeers, error)
	Messages(ctx context.Context, in *grpcSentry.MessagesRequest) (<-chan InboundMessageResult, error)
	Peers(ctx context.Context) (*grpcSentry.PeersReply, error)
	PeerCount(ctx context.Context, in *grpcSentry.PeerCountRequest) (*grpcSentry.PeerCountReply, error)
	PeerById(ctx context.Context, in *grpcSentry.PeerByIdRequest) (*grpcSentry.PeerByIdReply, error)
	PeerEvents(ctx context.Context, in *grpcSentry.PeerEventsRequest) (<-chan PeerEventResult, error)
	AddPeer(ctx context.Context, in *grpcSentry.AddPeerRequest) (*grpcSentry.AddPeerReply, error)
	NodeInfo(ctx context.Context) (*types.NodeInfoReply, error)
}
