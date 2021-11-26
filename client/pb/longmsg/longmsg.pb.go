// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: longmsg.proto

package longmsg

type LongMsgDeleteReq struct {
	MsgResid []byte `protobuf:"bytes,1,opt"`
	MsgType  int32  `protobuf:"varint,2,opt"`
}

func (x *LongMsgDeleteReq) GetMsgResid() []byte {
	if x != nil {
		return x.MsgResid
	}
	return nil
}

func (x *LongMsgDeleteReq) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

type LongMsgDeleteRsp struct {
	Result   int32  `protobuf:"varint,1,opt"`
	MsgResid []byte `protobuf:"bytes,2,opt"`
}

func (x *LongMsgDeleteRsp) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *LongMsgDeleteRsp) GetMsgResid() []byte {
	if x != nil {
		return x.MsgResid
	}
	return nil
}

type LongMsgDownReq struct {
	SrcUin    int32  `protobuf:"varint,1,opt"`
	MsgResid  []byte `protobuf:"bytes,2,opt"`
	MsgType   int32  `protobuf:"varint,3,opt"`
	NeedCache int32  `protobuf:"varint,4,opt"`
}

func (x *LongMsgDownReq) GetSrcUin() int32 {
	if x != nil {
		return x.SrcUin
	}
	return 0
}

func (x *LongMsgDownReq) GetMsgResid() []byte {
	if x != nil {
		return x.MsgResid
	}
	return nil
}

func (x *LongMsgDownReq) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *LongMsgDownReq) GetNeedCache() int32 {
	if x != nil {
		return x.NeedCache
	}
	return 0
}

type LongMsgDownRsp struct {
	Result     int32  `protobuf:"varint,1,opt"`
	MsgResid   []byte `protobuf:"bytes,2,opt"`
	MsgContent []byte `protobuf:"bytes,3,opt"`
}

func (x *LongMsgDownRsp) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *LongMsgDownRsp) GetMsgResid() []byte {
	if x != nil {
		return x.MsgResid
	}
	return nil
}

func (x *LongMsgDownRsp) GetMsgContent() []byte {
	if x != nil {
		return x.MsgContent
	}
	return nil
}

type LongMsgUpReq struct {
	MsgType    int32  `protobuf:"varint,1,opt"`
	DstUin     int64  `protobuf:"varint,2,opt"`
	MsgId      int32  `protobuf:"varint,3,opt"`
	MsgContent []byte `protobuf:"bytes,4,opt"`
	StoreType  int32  `protobuf:"varint,5,opt"`
	MsgUkey    []byte `protobuf:"bytes,6,opt"`
	NeedCache  int32  `protobuf:"varint,7,opt"`
}

func (x *LongMsgUpReq) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *LongMsgUpReq) GetDstUin() int64 {
	if x != nil {
		return x.DstUin
	}
	return 0
}

func (x *LongMsgUpReq) GetMsgId() int32 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *LongMsgUpReq) GetMsgContent() []byte {
	if x != nil {
		return x.MsgContent
	}
	return nil
}

func (x *LongMsgUpReq) GetStoreType() int32 {
	if x != nil {
		return x.StoreType
	}
	return 0
}

func (x *LongMsgUpReq) GetMsgUkey() []byte {
	if x != nil {
		return x.MsgUkey
	}
	return nil
}

func (x *LongMsgUpReq) GetNeedCache() int32 {
	if x != nil {
		return x.NeedCache
	}
	return 0
}

type LongMsgUpRsp struct {
	Result   int32  `protobuf:"varint,1,opt"`
	MsgId    int32  `protobuf:"varint,2,opt"`
	MsgResid []byte `protobuf:"bytes,3,opt"`
}

func (x *LongMsgUpRsp) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *LongMsgUpRsp) GetMsgId() int32 {
	if x != nil {
		return x.MsgId
	}
	return 0
}

func (x *LongMsgUpRsp) GetMsgResid() []byte {
	if x != nil {
		return x.MsgResid
	}
	return nil
}

type LongReqBody struct {
	Subcmd       int32               `protobuf:"varint,1,opt"`
	TermType     int32               `protobuf:"varint,2,opt"`
	PlatformType int32               `protobuf:"varint,3,opt"`
	MsgUpReq     []*LongMsgUpReq     `protobuf:"bytes,4,rep"`
	MsgDownReq   []*LongMsgDownReq   `protobuf:"bytes,5,rep"`
	MsgDelReq    []*LongMsgDeleteReq `protobuf:"bytes,6,rep"`
	AgentType    int32               `protobuf:"varint,10,opt"`
}

func (x *LongReqBody) GetSubcmd() int32 {
	if x != nil {
		return x.Subcmd
	}
	return 0
}

func (x *LongReqBody) GetTermType() int32 {
	if x != nil {
		return x.TermType
	}
	return 0
}

func (x *LongReqBody) GetPlatformType() int32 {
	if x != nil {
		return x.PlatformType
	}
	return 0
}

func (x *LongReqBody) GetMsgUpReq() []*LongMsgUpReq {
	if x != nil {
		return x.MsgUpReq
	}
	return nil
}

func (x *LongReqBody) GetMsgDownReq() []*LongMsgDownReq {
	if x != nil {
		return x.MsgDownReq
	}
	return nil
}

func (x *LongReqBody) GetMsgDelReq() []*LongMsgDeleteReq {
	if x != nil {
		return x.MsgDelReq
	}
	return nil
}

func (x *LongReqBody) GetAgentType() int32 {
	if x != nil {
		return x.AgentType
	}
	return 0
}

type LongRspBody struct {
	Subcmd     int32               `protobuf:"varint,1,opt"`
	MsgUpRsp   []*LongMsgUpRsp     `protobuf:"bytes,2,rep"`
	MsgDownRsp []*LongMsgDownRsp   `protobuf:"bytes,3,rep"`
	MsgDelRsp  []*LongMsgDeleteRsp `protobuf:"bytes,4,rep"`
}

func (x *LongRspBody) GetSubcmd() int32 {
	if x != nil {
		return x.Subcmd
	}
	return 0
}

func (x *LongRspBody) GetMsgUpRsp() []*LongMsgUpRsp {
	if x != nil {
		return x.MsgUpRsp
	}
	return nil
}

func (x *LongRspBody) GetMsgDownRsp() []*LongMsgDownRsp {
	if x != nil {
		return x.MsgDownRsp
	}
	return nil
}

func (x *LongRspBody) GetMsgDelRsp() []*LongMsgDeleteRsp {
	if x != nil {
		return x.MsgDelRsp
	}
	return nil
}
