package octopusdeploy

import (
	"fmt"

	"github.com/dghubble/sling"
	"gopkg.in/go-playground/validator.v9"
)

type CertificateService struct {
	sling *sling.Sling
}

func NewCertificateService(sling *sling.Sling) *CertificateService {
	return &CertificateService{
		sling: sling,
	}
}

type Certificates struct {
	Items []Certificate `json:"Items"`
	PagedResults
}

type Certificate struct {
	ID                              string                 `json:"Id,omitempty"`
	Name                            string                 `json:"Name,omitempty"`
	Notes                           string                 `json:"Notes,omitempty"`
	CertificateData                 SensitiveValue         `json:"CertificateData,omitempty"`
	Password                        SensitiveValue         `json:"Password,omitempty"`
	EnvironmentIds                  []string               `json:"EnvironmentIds,omitempty"`
	TenantedDeploymentParticipation TenantedDeploymentMode `json:"TenantedDeploymentParticipation,omitempty"`
	TenantIds                       []string               `json:"TenantIds,omitempty,omitempty"`
	TenantTags                      []string               `json:"TenantTags,omitempty,omitempty"`
	CertificateDataFormat           string                 `json:"CertificateDataFormat,omitempty"`
	Archived                        string                 `json:"Archived,omitempty"`
	ReplacedBy                      string                 `json:"ReplacedBy,omitempty"`
	SubjectDistinguishedName        string                 `json:"SubjectDistinguishedName,omitempty"`
	SubjectCommonName               string                 `json:"SubjectCommonName,omitempty"`
	SubjectOrganization             string                 `json:"SubjectOrganization,omitempty"`
	IssuerDistinguishedName         string                 `json:"IssuerDistinguishedName,omitempty"`
	IssuerCommonName                string                 `json:"IssuerCommonName,omitempty"`
	IssuerOrganization              string                 `json:"IssuerOrganization,omitempty"`
	SelfSigned                      bool                   `json:"SelfSigned,omitempty"`
	Thumbprint                      string                 `json:"Thumbprint,omitempty"`
	NotAfter                        string                 `json:"NotAfter,omitempty"`
	NotBefore                       string                 `json:"NotBefore,omitempty"`
	IsExpired                       bool                   `json:"IsExpired,omitempty"`
	HasPrivateKey                   bool                   `json:"HasPrivateKey,omitempty"`
	Version                         int                    `json:"Version,omitempty"`
	SerialNumber                    string                 `json:"SerialNumber,omitempty"`
	SignatureAlgorithmName          string                 `json:"SignatureAlgorithmName,omitempty"`
	SubjectAlternativeNames         []string               `json:"SubjectAlternativeNames,omitempty"`
}

func (t *Certificate) Validate() error {
	validate := validator.New()

	err := validate.Struct(t)

	if err != nil {
		return err
	}

	return nil
}

func NewCertificate(name string, certificateData SensitiveValue, password SensitiveValue) *Certificate {
	return &Certificate{
		Name:            name,
		CertificateData: certificateData,
		Password:        password,
	}
}

func (s *CertificateService) Get(certificateId string) (*Certificate, error) {
	path := fmt.Sprintf("certificates/%s", certificateId)
	resp, err := apiGet(s.sling, new(Certificate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Certificate), nil
}

func (s *CertificateService) GetAll() (*[]Certificate, error) {
	var p []Certificate

	path := "certificates"

	loadNextPage := true

	for loadNextPage {
		resp, err := apiGet(s.sling, new(Certificates), path)

		if err != nil {
			return nil, err
		}

		r := resp.(*Certificates)

		for _, item := range r.Items {
			p = append(p, item)
		}

		path, loadNextPage = LoadNextPage(r.PagedResults)
	}

	return &p, nil
}

func (s *CertificateService) GetByName(certificateName string) (*Certificate, error) {
	var foundCertificate Certificate
	certificates, err := s.GetAll()

	if err != nil {
		return nil, err
	}

	for _, certificate := range *certificates {
		if certificate.Name == certificateName {
			return &certificate, nil
		}
	}

	return &foundCertificate, fmt.Errorf("no certificate found with certificate name %s", certificateName)
}

func (s *CertificateService) Add(certificate *Certificate) (*Certificate, error) {
	resp, err := apiAdd(s.sling, certificate, new(Certificate), "certificates")

	if err != nil {
		return nil, err
	}

	return resp.(*Certificate), nil
}

func (s *CertificateService) Delete(certificateId string) error {
	path := fmt.Sprintf("certificates/%s", certificateId)
	err := apiDelete(s.sling, path)

	if err != nil {
		return err
	}

	return nil
}

func (s *CertificateService) Update(certificate *Certificate) (*Certificate, error) {
	path := fmt.Sprintf("certificates/%s", certificate.ID)
	resp, err := apiUpdate(s.sling, certificate, new(Certificate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Certificate), nil
}

func (s *CertificateService) Replace(certificate *Certificate) (*Certificate, error) {
	path := fmt.Sprintf("certificates/%s/replace", certificate.ID)
	resp, err := apiPost(s.sling, certificate, new(Certificate), path)

	if err != nil {
		return nil, err
	}

	return resp.(*Certificate), nil
}
