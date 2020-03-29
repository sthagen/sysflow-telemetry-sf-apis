package sfgo

type SFObjectType = UnionSFHeaderContainerProcessFileProcessEventNetworkFlowFileFlowFileEventNetworkEventTypeEnum
type SFObject = UnionSFHeaderContainerProcessFileProcessEventNetworkFlowFileFlowFileEventNetworkEvent

const (
	SF_HEADER    SFObjectType = UnionSFHeaderContainerProcessFileProcessEventNetworkFlowFileFlowFileEventNetworkEventTypeEnumSFHeader
	SF_CONT      SFObjectType = UnionSFHeaderContainerProcessFileProcessEventNetworkFlowFileFlowFileEventNetworkEventTypeEnumContainer
	SF_PROCESS   SFObjectType = UnionSFHeaderContainerProcessFileProcessEventNetworkFlowFileFlowFileEventNetworkEventTypeEnumProcess
	SF_FILE      SFObjectType = UnionSFHeaderContainerProcessFileProcessEventNetworkFlowFileFlowFileEventNetworkEventTypeEnumFile
	SF_PROC_EVT  SFObjectType = UnionSFHeaderContainerProcessFileProcessEventNetworkFlowFileFlowFileEventNetworkEventTypeEnumProcessEvent
	SF_NET_FLOW  SFObjectType = UnionSFHeaderContainerProcessFileProcessEventNetworkFlowFileFlowFileEventNetworkEventTypeEnumNetworkFlow
	SF_FILE_FLOW SFObjectType = UnionSFHeaderContainerProcessFileProcessEventNetworkFlowFileFlowFileEventNetworkEventTypeEnumFileFlow
	SF_FILE_EVT  SFObjectType = UnionSFHeaderContainerProcessFileProcessEventNetworkFlowFileFlowFileEventNetworkEventTypeEnumFileEvent
	SF_NET_EVT   SFObjectType = UnionSFHeaderContainerProcessFileProcessEventNetworkFlowFileFlowFileEventNetworkEventTypeEnumNetworkEvent
)
