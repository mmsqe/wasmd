package types

import (
	"errors"
)

func NewMsgAddBlacklistMsgs(creator string, blacklistMsgs []string) *MsgAddBlacklistMsgs {
	return &MsgAddBlacklistMsgs{
		Authority:     creator,
		BlacklistMsgs: blacklistMsgs,
	}
}

func (m MsgAddBlacklistMsgs) Validate() error {
	if m.Authority == "" {
		return errors.New("authority cannot be empty")
	}
	if len(m.BlacklistMsgs) == 0 || m.BlacklistMsgs == nil {
		return errors.New("blacklistMsgs cannot be empty")
	}
	return nil
}

func NewMsgRemoveBlacklistMsgs(creator string, blacklistMsgs []string) *MsgRemoveBlacklistMsgs {
	return &MsgRemoveBlacklistMsgs{
		Authority:     creator,
		BlacklistMsgs: blacklistMsgs,
	}
}

func (m MsgRemoveBlacklistMsgs) Validate() error {
	if m.Authority == "" {
		return errors.New("authority cannot be empty")
	}
	if len(m.BlacklistMsgs) == 0 || m.BlacklistMsgs == nil {
		return errors.New("blacklistMsgs cannot be empty")
	}
	return nil
}
