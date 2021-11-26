// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: oidb0x6d6.proto

package oidb

type DeleteFileReqBody struct {
	GroupCode      int64  `protobuf:"varint,1,opt"`
	AppId          int32  `protobuf:"varint,2,opt"`
	BusId          int32  `protobuf:"varint,3,opt"`
	ParentFolderId string `protobuf:"bytes,4,opt"`
	FileId         string `protobuf:"bytes,5,opt"`
}

func (x *DeleteFileReqBody) GetGroupCode() int64 {
	if x != nil {
		return x.GroupCode
	}
	return 0
}

func (x *DeleteFileReqBody) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *DeleteFileReqBody) GetBusId() int32 {
	if x != nil {
		return x.BusId
	}
	return 0
}

func (x *DeleteFileReqBody) GetParentFolderId() string {
	if x != nil {
		return x.ParentFolderId
	}
	return ""
}

func (x *DeleteFileReqBody) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

type DeleteFileRspBody struct {
	RetCode       int32  `protobuf:"varint,1,opt"`
	RetMsg        string `protobuf:"bytes,2,opt"`
	ClientWording string `protobuf:"bytes,3,opt"`
}

func (x *DeleteFileRspBody) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *DeleteFileRspBody) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

func (x *DeleteFileRspBody) GetClientWording() string {
	if x != nil {
		return x.ClientWording
	}
	return ""
}

type DownloadFileReqBody struct {
	GroupCode        int64  `protobuf:"varint,1,opt"`
	AppId            int32  `protobuf:"varint,2,opt"`
	BusId            int32  `protobuf:"varint,3,opt"`
	FileId           string `protobuf:"bytes,4,opt"`
	BoolThumbnailReq bool   `protobuf:"varint,5,opt"`
	UrlType          int32  `protobuf:"varint,6,opt"`
	BoolPreviewReq   bool   `protobuf:"varint,7,opt"`
}

func (x *DownloadFileReqBody) GetGroupCode() int64 {
	if x != nil {
		return x.GroupCode
	}
	return 0
}

func (x *DownloadFileReqBody) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *DownloadFileReqBody) GetBusId() int32 {
	if x != nil {
		return x.BusId
	}
	return 0
}

func (x *DownloadFileReqBody) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *DownloadFileReqBody) GetBoolThumbnailReq() bool {
	if x != nil {
		return x.BoolThumbnailReq
	}
	return false
}

func (x *DownloadFileReqBody) GetUrlType() int32 {
	if x != nil {
		return x.UrlType
	}
	return 0
}

func (x *DownloadFileReqBody) GetBoolPreviewReq() bool {
	if x != nil {
		return x.BoolPreviewReq
	}
	return false
}

type DownloadFileRspBody struct {
	RetCode       int32  `protobuf:"varint,1,opt"`
	RetMsg        string `protobuf:"bytes,2,opt"`
	ClientWording string `protobuf:"bytes,3,opt"`
	DownloadIp    string `protobuf:"bytes,4,opt"`
	DownloadDns   []byte `protobuf:"bytes,5,opt"`
	DownloadUrl   []byte `protobuf:"bytes,6,opt"`
	Sha           []byte `protobuf:"bytes,7,opt"`
	Sha3          []byte `protobuf:"bytes,8,opt"`
	Md5           []byte `protobuf:"bytes,9,opt"`
	CookieVal     []byte `protobuf:"bytes,10,opt"`
	SaveFileName  string `protobuf:"bytes,11,opt"`
	PreviewPort   int32  `protobuf:"varint,12,opt"`
}

func (x *DownloadFileRspBody) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *DownloadFileRspBody) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

func (x *DownloadFileRspBody) GetClientWording() string {
	if x != nil {
		return x.ClientWording
	}
	return ""
}

func (x *DownloadFileRspBody) GetDownloadIp() string {
	if x != nil {
		return x.DownloadIp
	}
	return ""
}

func (x *DownloadFileRspBody) GetDownloadDns() []byte {
	if x != nil {
		return x.DownloadDns
	}
	return nil
}

func (x *DownloadFileRspBody) GetDownloadUrl() []byte {
	if x != nil {
		return x.DownloadUrl
	}
	return nil
}

func (x *DownloadFileRspBody) GetSha() []byte {
	if x != nil {
		return x.Sha
	}
	return nil
}

func (x *DownloadFileRspBody) GetSha3() []byte {
	if x != nil {
		return x.Sha3
	}
	return nil
}

func (x *DownloadFileRspBody) GetMd5() []byte {
	if x != nil {
		return x.Md5
	}
	return nil
}

func (x *DownloadFileRspBody) GetCookieVal() []byte {
	if x != nil {
		return x.CookieVal
	}
	return nil
}

func (x *DownloadFileRspBody) GetSaveFileName() string {
	if x != nil {
		return x.SaveFileName
	}
	return ""
}

func (x *DownloadFileRspBody) GetPreviewPort() int32 {
	if x != nil {
		return x.PreviewPort
	}
	return 0
}

type MoveFileReqBody struct {
	GroupCode      int64  `protobuf:"varint,1,opt"`
	AppId          int32  `protobuf:"varint,2,opt"`
	BusId          int32  `protobuf:"varint,3,opt"`
	FileId         string `protobuf:"bytes,4,opt"`
	ParentFolderId string `protobuf:"bytes,5,opt"`
	DestFolderId   string `protobuf:"bytes,6,opt"`
}

func (x *MoveFileReqBody) GetGroupCode() int64 {
	if x != nil {
		return x.GroupCode
	}
	return 0
}

func (x *MoveFileReqBody) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *MoveFileReqBody) GetBusId() int32 {
	if x != nil {
		return x.BusId
	}
	return 0
}

func (x *MoveFileReqBody) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *MoveFileReqBody) GetParentFolderId() string {
	if x != nil {
		return x.ParentFolderId
	}
	return ""
}

func (x *MoveFileReqBody) GetDestFolderId() string {
	if x != nil {
		return x.DestFolderId
	}
	return ""
}

type MoveFileRspBody struct {
	RetCode        int32  `protobuf:"varint,1,opt"`
	RetMsg         string `protobuf:"bytes,2,opt"`
	ClientWording  string `protobuf:"bytes,3,opt"`
	ParentFolderId string `protobuf:"bytes,4,opt"`
}

func (x *MoveFileRspBody) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *MoveFileRspBody) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

func (x *MoveFileRspBody) GetClientWording() string {
	if x != nil {
		return x.ClientWording
	}
	return ""
}

func (x *MoveFileRspBody) GetParentFolderId() string {
	if x != nil {
		return x.ParentFolderId
	}
	return ""
}

type RenameFileReqBody struct {
	GroupCode      int64  `protobuf:"varint,1,opt"`
	AppId          int32  `protobuf:"varint,2,opt"`
	BusId          int32  `protobuf:"varint,3,opt"`
	FileId         string `protobuf:"bytes,4,opt"`
	ParentFolderId string `protobuf:"bytes,5,opt"`
	NewFileName    string `protobuf:"bytes,6,opt"`
}

func (x *RenameFileReqBody) GetGroupCode() int64 {
	if x != nil {
		return x.GroupCode
	}
	return 0
}

func (x *RenameFileReqBody) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *RenameFileReqBody) GetBusId() int32 {
	if x != nil {
		return x.BusId
	}
	return 0
}

func (x *RenameFileReqBody) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *RenameFileReqBody) GetParentFolderId() string {
	if x != nil {
		return x.ParentFolderId
	}
	return ""
}

func (x *RenameFileReqBody) GetNewFileName() string {
	if x != nil {
		return x.NewFileName
	}
	return ""
}

type RenameFileRspBody struct {
	RetCode       int32  `protobuf:"varint,1,opt"`
	RetMsg        string `protobuf:"bytes,2,opt"`
	ClientWording string `protobuf:"bytes,3,opt"`
}

func (x *RenameFileRspBody) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *RenameFileRspBody) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

func (x *RenameFileRspBody) GetClientWording() string {
	if x != nil {
		return x.ClientWording
	}
	return ""
}

type D6D6ReqBody struct {
	UploadFileReq   *UploadFileReqBody   `protobuf:"bytes,1,opt"`
	ResendFileReq   *ResendReqBody       `protobuf:"bytes,2,opt"`
	DownloadFileReq *DownloadFileReqBody `protobuf:"bytes,3,opt"`
	DeleteFileReq   *DeleteFileReqBody   `protobuf:"bytes,4,opt"`
	RenameFileReq   *RenameFileReqBody   `protobuf:"bytes,5,opt"`
	MoveFileReq     *MoveFileReqBody     `protobuf:"bytes,6,opt"`
}

func (x *D6D6ReqBody) GetUploadFileReq() *UploadFileReqBody {
	if x != nil {
		return x.UploadFileReq
	}
	return nil
}

func (x *D6D6ReqBody) GetResendFileReq() *ResendReqBody {
	if x != nil {
		return x.ResendFileReq
	}
	return nil
}

func (x *D6D6ReqBody) GetDownloadFileReq() *DownloadFileReqBody {
	if x != nil {
		return x.DownloadFileReq
	}
	return nil
}

func (x *D6D6ReqBody) GetDeleteFileReq() *DeleteFileReqBody {
	if x != nil {
		return x.DeleteFileReq
	}
	return nil
}

func (x *D6D6ReqBody) GetRenameFileReq() *RenameFileReqBody {
	if x != nil {
		return x.RenameFileReq
	}
	return nil
}

func (x *D6D6ReqBody) GetMoveFileReq() *MoveFileReqBody {
	if x != nil {
		return x.MoveFileReq
	}
	return nil
}

type ResendReqBody struct {
	GroupCode int64  `protobuf:"varint,1,opt"`
	AppId     int32  `protobuf:"varint,2,opt"`
	BusId     int32  `protobuf:"varint,3,opt"`
	FileId    string `protobuf:"bytes,4,opt"`
	Sha       []byte `protobuf:"bytes,5,opt"`
}

func (x *ResendReqBody) GetGroupCode() int64 {
	if x != nil {
		return x.GroupCode
	}
	return 0
}

func (x *ResendReqBody) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *ResendReqBody) GetBusId() int32 {
	if x != nil {
		return x.BusId
	}
	return 0
}

func (x *ResendReqBody) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *ResendReqBody) GetSha() []byte {
	if x != nil {
		return x.Sha
	}
	return nil
}

type ResendRspBody struct {
	RetCode       int32  `protobuf:"varint,1,opt"`
	RetMsg        string `protobuf:"bytes,2,opt"`
	ClientWording string `protobuf:"bytes,3,opt"`
	UploadIp      string `protobuf:"bytes,4,opt"`
	FileKey       []byte `protobuf:"bytes,5,opt"`
	CheckKey      []byte `protobuf:"bytes,6,opt"`
}

func (x *ResendRspBody) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *ResendRspBody) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

func (x *ResendRspBody) GetClientWording() string {
	if x != nil {
		return x.ClientWording
	}
	return ""
}

func (x *ResendRspBody) GetUploadIp() string {
	if x != nil {
		return x.UploadIp
	}
	return ""
}

func (x *ResendRspBody) GetFileKey() []byte {
	if x != nil {
		return x.FileKey
	}
	return nil
}

func (x *ResendRspBody) GetCheckKey() []byte {
	if x != nil {
		return x.CheckKey
	}
	return nil
}

type D6D6RspBody struct {
	UploadFileRsp   *UploadFileRspBody   `protobuf:"bytes,1,opt"`
	ResendFileRsp   *ResendRspBody       `protobuf:"bytes,2,opt"`
	DownloadFileRsp *DownloadFileRspBody `protobuf:"bytes,3,opt"`
	DeleteFileRsp   *DeleteFileRspBody   `protobuf:"bytes,4,opt"`
	RenameFileRsp   *RenameFileRspBody   `protobuf:"bytes,5,opt"`
	MoveFileRsp     *MoveFileRspBody     `protobuf:"bytes,6,opt"`
}

func (x *D6D6RspBody) GetUploadFileRsp() *UploadFileRspBody {
	if x != nil {
		return x.UploadFileRsp
	}
	return nil
}

func (x *D6D6RspBody) GetResendFileRsp() *ResendRspBody {
	if x != nil {
		return x.ResendFileRsp
	}
	return nil
}

func (x *D6D6RspBody) GetDownloadFileRsp() *DownloadFileRspBody {
	if x != nil {
		return x.DownloadFileRsp
	}
	return nil
}

func (x *D6D6RspBody) GetDeleteFileRsp() *DeleteFileRspBody {
	if x != nil {
		return x.DeleteFileRsp
	}
	return nil
}

func (x *D6D6RspBody) GetRenameFileRsp() *RenameFileRspBody {
	if x != nil {
		return x.RenameFileRsp
	}
	return nil
}

func (x *D6D6RspBody) GetMoveFileRsp() *MoveFileRspBody {
	if x != nil {
		return x.MoveFileRsp
	}
	return nil
}

type UploadFileReqBody struct {
	GroupCode          int64  `protobuf:"varint,1,opt"`
	AppId              int32  `protobuf:"varint,2,opt"`
	BusId              int32  `protobuf:"varint,3,opt"`
	Entrance           int32  `protobuf:"varint,4,opt"`
	ParentFolderId     string `protobuf:"bytes,5,opt"`
	FileName           string `protobuf:"bytes,6,opt"`
	LocalPath          string `protobuf:"bytes,7,opt"`
	Int64FileSize      int64  `protobuf:"varint,8,opt"`
	Sha                []byte `protobuf:"bytes,9,opt"`
	Sha3               []byte `protobuf:"bytes,10,opt"`
	Md5                []byte `protobuf:"bytes,11,opt"`
	SupportMultiUpload bool   `protobuf:"varint,15,opt"`
}

func (x *UploadFileReqBody) GetGroupCode() int64 {
	if x != nil {
		return x.GroupCode
	}
	return 0
}

func (x *UploadFileReqBody) GetAppId() int32 {
	if x != nil {
		return x.AppId
	}
	return 0
}

func (x *UploadFileReqBody) GetBusId() int32 {
	if x != nil {
		return x.BusId
	}
	return 0
}

func (x *UploadFileReqBody) GetEntrance() int32 {
	if x != nil {
		return x.Entrance
	}
	return 0
}

func (x *UploadFileReqBody) GetParentFolderId() string {
	if x != nil {
		return x.ParentFolderId
	}
	return ""
}

func (x *UploadFileReqBody) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadFileReqBody) GetLocalPath() string {
	if x != nil {
		return x.LocalPath
	}
	return ""
}

func (x *UploadFileReqBody) GetInt64FileSize() int64 {
	if x != nil {
		return x.Int64FileSize
	}
	return 0
}

func (x *UploadFileReqBody) GetSha() []byte {
	if x != nil {
		return x.Sha
	}
	return nil
}

func (x *UploadFileReqBody) GetSha3() []byte {
	if x != nil {
		return x.Sha3
	}
	return nil
}

func (x *UploadFileReqBody) GetMd5() []byte {
	if x != nil {
		return x.Md5
	}
	return nil
}

func (x *UploadFileReqBody) GetSupportMultiUpload() bool {
	if x != nil {
		return x.SupportMultiUpload
	}
	return false
}

type UploadFileRspBody struct {
	RetCode       int32    `protobuf:"varint,1,opt"`
	RetMsg        string   `protobuf:"bytes,2,opt"`
	ClientWording string   `protobuf:"bytes,3,opt"`
	UploadIp      string   `protobuf:"bytes,4,opt"`
	ServerDns     string   `protobuf:"bytes,5,opt"`
	BusId         int32    `protobuf:"varint,6,opt"`
	FileId        string   `protobuf:"bytes,7,opt"`
	FileKey       []byte   `protobuf:"bytes,8,opt"`
	CheckKey      []byte   `protobuf:"bytes,9,opt"`
	BoolFileExist bool     `protobuf:"varint,10,opt"`
	UploadIpLanV4 []string `protobuf:"bytes,12,rep"`
	UploadIpLanV6 []string `protobuf:"bytes,13,rep"`
	UploadPort    int32    `protobuf:"varint,14,opt"`
}

func (x *UploadFileRspBody) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *UploadFileRspBody) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

func (x *UploadFileRspBody) GetClientWording() string {
	if x != nil {
		return x.ClientWording
	}
	return ""
}

func (x *UploadFileRspBody) GetUploadIp() string {
	if x != nil {
		return x.UploadIp
	}
	return ""
}

func (x *UploadFileRspBody) GetServerDns() string {
	if x != nil {
		return x.ServerDns
	}
	return ""
}

func (x *UploadFileRspBody) GetBusId() int32 {
	if x != nil {
		return x.BusId
	}
	return 0
}

func (x *UploadFileRspBody) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *UploadFileRspBody) GetFileKey() []byte {
	if x != nil {
		return x.FileKey
	}
	return nil
}

func (x *UploadFileRspBody) GetCheckKey() []byte {
	if x != nil {
		return x.CheckKey
	}
	return nil
}

func (x *UploadFileRspBody) GetBoolFileExist() bool {
	if x != nil {
		return x.BoolFileExist
	}
	return false
}

func (x *UploadFileRspBody) GetUploadIpLanV4() []string {
	if x != nil {
		return x.UploadIpLanV4
	}
	return nil
}

func (x *UploadFileRspBody) GetUploadIpLanV6() []string {
	if x != nil {
		return x.UploadIpLanV6
	}
	return nil
}

func (x *UploadFileRspBody) GetUploadPort() int32 {
	if x != nil {
		return x.UploadPort
	}
	return 0
}
