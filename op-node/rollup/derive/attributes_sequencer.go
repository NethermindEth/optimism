package derive

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type L1OriginSelectorIface interface {
	FindL1Origin(ctx context.Context, l2Head eth.L2BlockRef) (eth.L1BlockRef, error)
}

type AttributesSequencer struct {
	log     log.Logger
	metrics Metrics

	l1OriginSelector      L1OriginSelectorIface
	attrBuilder           AttributesBuilder
	broadcastPayloadAttrs func(id string, data []byte)
}

func NewAttributesSequencer(log log.Logger, l1OriginSelector L1OriginSelectorIface, attrBuilder AttributesBuilder, broadcastPayloadAttrs func(id string, data []byte), metrics Metrics) *AttributesSequencer {
	return &AttributesSequencer{
		log:     log,
		metrics: metrics,

		l1OriginSelector: l1OriginSelector,
		attrBuilder:      attrBuilder,
	}
}

func (as *AttributesSequencer) PreparePayloadAttributes(ctx context.Context, l2Head eth.L2BlockRef) (*eth.PayloadAttributes, error) {
	// Figure out which L1 origin block we're going to be building on top of.
	l1Origin, err := as.l1OriginSelector.FindL1Origin(ctx, l2Head)
	if err != nil {
		as.log.Error("Error finding next L1 Origin", "err", err)
		return nil, err
	}

	if !(l2Head.L1Origin.Hash == l1Origin.ParentHash || l2Head.L1Origin.Hash == l1Origin.Hash) {
		// TODO: use different metrics for this
		// as.metrics.RecordSequencerInconsistentL1Origin(l2Head.L1Origin, l1Origin.ID())
		return nil, NewResetError(fmt.Errorf("cannot build new L2 block with L1 origin %s (parent L1 %s) on current L2 head %s with L1 origin %s", l1Origin, l1Origin.ParentHash, l2Head, l2Head.L1Origin))
	}

	as.log.Info("creating new block", "parent", l2Head, "l1Origin", l1Origin)

	fetchCtx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()

	attrs, err := as.attrBuilder.PreparePayloadAttributes(fetchCtx, l2Head, l1Origin.ID())
	if err != nil {
		return nil, err
	}

	txs := make(types.Transactions, 0, len(attrs.Transactions))
	for i, tx := range attrs.Transactions {
		txs[i].UnmarshalBinary(tx)
	}

	attrsEvent := &eth.BuilderPayloadAttributes{
		Timestamp:             attrs.Timestamp,
		Random:                common.Hash(attrs.PrevRandao),
		SuggestedFeeRecipient: attrs.SuggestedFeeRecipient,
		Slot:                  l2Head.Number,
		HeadHash:              l2Head.Hash,
		Transactions:          txs,
		GasLimit:              uint64(*attrs.GasLimit),
	}

	attrsData, err := json.Marshal(attrsEvent)
	if err != nil {
		return nil, err
	}

	as.log.Info("broadcasting new payload attributes", "json", attrsData)
	as.broadcastPayloadAttrs("payload_attributes", attrsData)
	return attrs, nil
}
