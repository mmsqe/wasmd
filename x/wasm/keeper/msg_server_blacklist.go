package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"github.com/CosmWasm/wasmd/x/wasm/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) AddBlacklistMsgs(ctx context.Context, msg *types.MsgAddBlacklistMsgs) (*types.MsgAddBlacklistMsgsResponse, error) {
	if err := msg.Validate(); err != nil {
		return nil, err
	}

	authority := k.keeper.GetAuthority()
	if authority != msg.Authority {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid authority; expected %s, got %s", authority, msg.Authority)
	}

	for _, msg := range msg.BlacklistMsgs {
		hasMsg, err := k.keeper.blacklistMsgs.Has(ctx, msg)
		if err != nil {
			return nil, err
		}
		if hasMsg {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "msg %s has already been blacklisted", msg)
		}

		if err := k.keeper.blacklistMsgs.Set(ctx, msg); err != nil {
			return nil, err
		}
	}

	return &types.MsgAddBlacklistMsgsResponse{}, nil
}

func (k msgServer) RemoveBlacklistMsgs(ctx context.Context, msg *types.MsgRemoveBlacklistMsgs) (*types.MsgRemoveBlacklistMsgsResponse, error) {
	if err := msg.Validate(); err != nil {
		return nil, err
	}

	authority := k.keeper.GetAuthority()
	if authority != msg.Authority {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid authority; expected %s, got %s", authority, msg.Authority)
	}

	for _, msg := range msg.BlacklistMsgs {
		hasMsg, err := k.keeper.blacklistMsgs.Has(ctx, msg)
		if err != nil {
			return nil, err
		}
		if !hasMsg {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "blacklist msg %s is not blacklisted", msg)
		}

		if err := k.keeper.blacklistMsgs.Remove(ctx, msg); err != nil {
			return nil, err
		}
	}

	return &types.MsgRemoveBlacklistMsgsResponse{}, nil
}
