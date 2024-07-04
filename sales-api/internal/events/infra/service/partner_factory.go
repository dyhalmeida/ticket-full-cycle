package service

import "fmt"

type IPartnerFactory interface {
	CreatePartner(partnerID int) (IPartner, error)
}

type DefaultPartnerFactory struct {
	partnerBaseURLs map[int]string
}

func NewPartnerFactory(partnerBaseURLs map[int]string) IPartnerFactory {
	return &DefaultPartnerFactory{partnerBaseURLs: partnerBaseURLs}
}

// CreatePartner implements IPartnerFactory.
func (f *DefaultPartnerFactory) CreatePartner(partnerID int) (IPartner, error) {
	baseURL, ok := f.partnerBaseURLs[partnerID]
	if !ok {
		return nil, fmt.Errorf("unsupported partner ID: %d", partnerID)
	}
	switch partnerID {
	case 1:
		return &Partner1{BaseURL: baseURL}, nil
	case 2:
		return &Partner1{BaseURL: baseURL}, nil
	default:
		return nil, fmt.Errorf("unsupported partner ID: %d", partnerID)
	}
}
