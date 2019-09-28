package fc

// OSSEvent represents the oss event type in oss trigger
type OSSEvent string

const (
	OSSEventObjectCreatedAll                     OSSEvent = "oss:ObjectCreated:*"
	OSSEventObjectCreatedPutObject               OSSEvent = "oss:ObjectCreated:PutObject"
	OSSEventObjectCreatedPutSymlink              OSSEvent = "oss:ObjectCreated:PutSymlink"
	OSSEventObjectCreatedPostObject              OSSEvent = "oss:ObjectCreated:PostObject"
	OSSEventObjectCreatedCopyObject              OSSEvent = "oss:ObjectCreated:CopyObject"
	OSSEventObjectCreatedInitiateMultipartUpload OSSEvent = "oss:ObjectCreated:InitiateMultipartUpload"
	OSSEventObjectCreatedUploadPart              OSSEvent = "oss:ObjectCreated:UploadPart"
	OSSEventObjectCreatedUploadPartCopy          OSSEvent = "oss:ObjectCreated:UploadPartCopy"
	OSSEventObjectCreatedCompleteMultipartUpload OSSEvent = "oss:ObjectCreated:CompleteMultipartUpload"
	OSSEventObjectCreatedAppendObject            OSSEvent = "oss:ObjectCreated:AppendObject"
	OSSEventObjectRemovedDeleteObject            OSSEvent = "oss:ObjectRemoved:DeleteObject"
	OSSEventObjectRemovedDeleteObjects           OSSEvent = "oss:ObjectRemoved:DeleteObjects"
	OSSEventObjectRemovedAbortMultipartUpload    OSSEvent = "oss:ObjectRemoved:AbortMultipartUpload"
	OSSEventObjectReplicationObjectCreated       OSSEvent = "oss:ObjectReplication:ObjectCreated"
	OSSEventObjectReplicationObjectRemoved       OSSEvent = "oss:ObjectReplication:ObjectRemoved"
	OSSEventObjectReplicationObjectModified      OSSEvent = "oss:ObjectReplication:ObjectModified"
)

const (
	// HTTPHeaderRequestID get request ID
	HTTPHeaderRequestID = "X-Fc-Request-Id"

	// HTTPHeaderInvocationType stores the invocation type.
	HTTPHeaderInvocationType = "X-Fc-Invocation-Type"

	// HTTPHeaderAccountID stores the account ID
	HTTPHeaderAccountID = "X-Fc-Account-Id"

	// HTTPHeaderFCErrorType get the error type when invoke function
	HTTPHeaderFCErrorType = "X-Fc-Error-Type"

	// HTTPHeaderSecurityToken is the header key for STS security token
	HTTPHeaderSecurityToken = "X-Fc-Security-Token"

	// HTTPHeaderInvocationLogType is the header key for invoke function log type
	HTTPHeaderInvocationLogType = "X-Fc-Log-Type"

	// HTTPHeaderInvocationLogResult is the header key for invoke function log result
	HTTPHeaderInvocationLogResult = "X-Fc-Log-Result"

	// HTTPHeaderEtag get the etag of the resource
	HTTPHeaderEtag = "Etag"

	//HTTPHeaderPrefix :Prefix string in headers
	HTTPHeaderPrefix = "x-fc-"

	//HTTPHeaderContentMD5 :Key in request headers
	HTTPHeaderContentMD5 = "Content-MD5"

	//HTTPHeaderContentType :Key in request headers
	HTTPHeaderContentType = "Content-Type"

	// HTTPHeaderUserAgent : Key in request headers
	HTTPHeaderUserAgent = "User-Agent"

	//HTTPHeaderDate :Key in request headers
	HTTPHeaderDate = "Date"
)

// Supported api versions
const (
	APIVersionV1 = "2016-08-15"
)
