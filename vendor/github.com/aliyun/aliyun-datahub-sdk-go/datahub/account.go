package datahub

/*
only support aliyun account now
*/

import (
	"fmt"
)

type Account interface {
	fmt.Stringer
	GetAccountId() string
	GetAccountKey() string
	GetSecurityToken() string
}

type AliyunAccount struct {
	// Aliyun Access key ID
	AccessId string

	// Aliyun Secret Access Key
	AccessKey string
}

// create new instance
func NewAliyunAccount(accessId, accessKey string) *AliyunAccount {
	return &AliyunAccount{
		AccessId:  accessId,
		AccessKey: accessKey,
	}
}

func (a AliyunAccount) String() string {
	return fmt.Sprintf("accessId: %s, accessKey: %s", a.AccessId, a.AccessKey)
}

func (a AliyunAccount) GetAccountId() string {
	return a.AccessId
}

func (a AliyunAccount) GetAccountKey() string {
	return a.AccessKey
}

func (a AliyunAccount) GetSecurityToken() string {
	return ""
}

type StsCredential struct {
	// Access key ID
	AccessId string

	// Secret Access Key
	AccessKey string

	// Security Token
	SecurityToken string
}

// create new instance
func NewStsCredential(accessId, accessKey, securityToken string) *StsCredential {
	return &StsCredential{
		AccessId:      accessId,
		AccessKey:     accessKey,
		SecurityToken: securityToken,
	}
}

func (a StsCredential) String() string {
	return fmt.Sprintf("accessId: %s, accessKey: %s, securityToken: %s", a.AccessId, a.AccessKey, a.SecurityToken)
}

func (a StsCredential) GetAccountId() string {
	return a.AccessId
}

func (a StsCredential) GetAccountKey() string {
	return a.AccessKey
}

func (a StsCredential) GetSecurityToken() string {
	return a.SecurityToken
}
