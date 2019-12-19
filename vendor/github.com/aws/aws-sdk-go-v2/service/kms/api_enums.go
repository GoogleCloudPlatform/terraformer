// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kms

type AlgorithmSpec string

// Enum values for AlgorithmSpec
const (
	AlgorithmSpecRsaesPkcs1V15   AlgorithmSpec = "RSAES_PKCS1_V1_5"
	AlgorithmSpecRsaesOaepSha1   AlgorithmSpec = "RSAES_OAEP_SHA_1"
	AlgorithmSpecRsaesOaepSha256 AlgorithmSpec = "RSAES_OAEP_SHA_256"
)

func (enum AlgorithmSpec) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AlgorithmSpec) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ConnectionErrorCodeType string

// Enum values for ConnectionErrorCodeType
const (
	ConnectionErrorCodeTypeInvalidCredentials       ConnectionErrorCodeType = "INVALID_CREDENTIALS"
	ConnectionErrorCodeTypeClusterNotFound          ConnectionErrorCodeType = "CLUSTER_NOT_FOUND"
	ConnectionErrorCodeTypeNetworkErrors            ConnectionErrorCodeType = "NETWORK_ERRORS"
	ConnectionErrorCodeTypeInternalError            ConnectionErrorCodeType = "INTERNAL_ERROR"
	ConnectionErrorCodeTypeInsufficientCloudhsmHsms ConnectionErrorCodeType = "INSUFFICIENT_CLOUDHSM_HSMS"
	ConnectionErrorCodeTypeUserLockedOut            ConnectionErrorCodeType = "USER_LOCKED_OUT"
)

func (enum ConnectionErrorCodeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ConnectionErrorCodeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ConnectionStateType string

// Enum values for ConnectionStateType
const (
	ConnectionStateTypeConnected     ConnectionStateType = "CONNECTED"
	ConnectionStateTypeConnecting    ConnectionStateType = "CONNECTING"
	ConnectionStateTypeFailed        ConnectionStateType = "FAILED"
	ConnectionStateTypeDisconnected  ConnectionStateType = "DISCONNECTED"
	ConnectionStateTypeDisconnecting ConnectionStateType = "DISCONNECTING"
)

func (enum ConnectionStateType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ConnectionStateType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DataKeySpec string

// Enum values for DataKeySpec
const (
	DataKeySpecAes256 DataKeySpec = "AES_256"
	DataKeySpecAes128 DataKeySpec = "AES_128"
)

func (enum DataKeySpec) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DataKeySpec) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ExpirationModelType string

// Enum values for ExpirationModelType
const (
	ExpirationModelTypeKeyMaterialExpires       ExpirationModelType = "KEY_MATERIAL_EXPIRES"
	ExpirationModelTypeKeyMaterialDoesNotExpire ExpirationModelType = "KEY_MATERIAL_DOES_NOT_EXPIRE"
)

func (enum ExpirationModelType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ExpirationModelType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type GrantOperation string

// Enum values for GrantOperation
const (
	GrantOperationDecrypt                         GrantOperation = "Decrypt"
	GrantOperationEncrypt                         GrantOperation = "Encrypt"
	GrantOperationGenerateDataKey                 GrantOperation = "GenerateDataKey"
	GrantOperationGenerateDataKeyWithoutPlaintext GrantOperation = "GenerateDataKeyWithoutPlaintext"
	GrantOperationReEncryptFrom                   GrantOperation = "ReEncryptFrom"
	GrantOperationReEncryptTo                     GrantOperation = "ReEncryptTo"
	GrantOperationCreateGrant                     GrantOperation = "CreateGrant"
	GrantOperationRetireGrant                     GrantOperation = "RetireGrant"
	GrantOperationDescribeKey                     GrantOperation = "DescribeKey"
)

func (enum GrantOperation) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum GrantOperation) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type KeyManagerType string

// Enum values for KeyManagerType
const (
	KeyManagerTypeAws      KeyManagerType = "AWS"
	KeyManagerTypeCustomer KeyManagerType = "CUSTOMER"
)

func (enum KeyManagerType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum KeyManagerType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type KeyState string

// Enum values for KeyState
const (
	KeyStateEnabled         KeyState = "Enabled"
	KeyStateDisabled        KeyState = "Disabled"
	KeyStatePendingDeletion KeyState = "PendingDeletion"
	KeyStatePendingImport   KeyState = "PendingImport"
	KeyStateUnavailable     KeyState = "Unavailable"
)

func (enum KeyState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum KeyState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type KeyUsageType string

// Enum values for KeyUsageType
const (
	KeyUsageTypeEncryptDecrypt KeyUsageType = "ENCRYPT_DECRYPT"
)

func (enum KeyUsageType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum KeyUsageType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type OriginType string

// Enum values for OriginType
const (
	OriginTypeAwsKms      OriginType = "AWS_KMS"
	OriginTypeExternal    OriginType = "EXTERNAL"
	OriginTypeAwsCloudhsm OriginType = "AWS_CLOUDHSM"
)

func (enum OriginType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OriginType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type WrappingKeySpec string

// Enum values for WrappingKeySpec
const (
	WrappingKeySpecRsa2048 WrappingKeySpec = "RSA_2048"
)

func (enum WrappingKeySpec) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum WrappingKeySpec) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
