package frontend

import (
	"context"

	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type EngineBackend interface {
	GetPayloadV1(ctx context.Context, payloadID eth.PayloadID) (*eth.ExecutionPayload, error)
	GetPayloadV2(ctx context.Context, payloadID eth.PayloadID) (*eth.ExecutionPayloadEnvelope, error)
	GetPayloadV3(ctx context.Context, payloadID eth.PayloadID) (*eth.ExecutionPayloadEnvelope, error)
	GetPayloadV4(ctx context.Context, payloadID eth.PayloadID) (*eth.ExecutionPayloadEnvelope, error)

	ForkchoiceUpdatedV1(ctx context.Context, state *eth.ForkchoiceState, attr *eth.PayloadAttributes) (*eth.ForkchoiceUpdatedResult, error)
	ForkchoiceUpdatedV2(ctx context.Context, state *eth.ForkchoiceState, attr *eth.PayloadAttributes) (*eth.ForkchoiceUpdatedResult, error)
	ForkchoiceUpdatedV3(ctx context.Context, state *eth.ForkchoiceState, attr *eth.PayloadAttributes) (*eth.ForkchoiceUpdatedResult, error)

	NewPayloadV1(ctx context.Context, payload *eth.ExecutionPayload) (*eth.PayloadStatusV1, error)
	NewPayloadV2(ctx context.Context, payload *eth.ExecutionPayload) (*eth.PayloadStatusV1, error)
	NewPayloadV3(ctx context.Context, payload *eth.ExecutionPayload, versionedHashes []common.Hash, beaconRoot *common.Hash) (*eth.PayloadStatusV1, error)
	NewPayloadV4(ctx context.Context, payload *eth.ExecutionPayload, versionedHashes []common.Hash, beaconRoot *common.Hash, executionRequests []hexutil.Bytes) (*eth.PayloadStatusV1, error)
}

type EngineFrontend struct {
	EngineBackend
}

func NewEngineFrontend(b EngineBackend) *EngineFrontend {
	return &EngineFrontend{EngineBackend: b}
}
