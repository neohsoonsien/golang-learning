package models

import (
	proto "golang-learning/market_place/proto"
)

type Invoice struct {
	PrincipalId string   `json:"principalId" bson:"principalId"`
	Number      string   `json:"number" bson:"number"`
	Vendor      string   `json:"vendor" bson:"vendor"`
	DateTime    int32    `json:"dateTime" bson:"dateTime"`
	Details     []Detail `json:"details" bson:"details"`
}

func (m *Invoice) ToPb() *proto.Invoice {
	detailsPb := []*proto.Detail{}

	for _, detail := range m.Details {
		detailsPb = append(detailsPb, detail.ToPb())
	}

	return &proto.Invoice{
		PrincipalId: m.PrincipalId,
		Number:      m.Number,
		Vendor:      m.Vendor,
		DateTime:    m.DateTime,
		Details:     detailsPb,
	}
}

func InvoicePbToStruct(pbInvoice *proto.Invoice) *Invoice {
	if pbInvoice == nil {
		return nil
	}

	details := []Detail{}
	for _, detail := range pbInvoice.GetDetails() {
		details = append(details, *DetailPbToStruct(detail))
	}

	return &Invoice{
		PrincipalId: pbInvoice.GetPrincipalId(),
		Number:      pbInvoice.GetNumber(),
		Vendor:      pbInvoice.GetVendor(),
		DateTime:    pbInvoice.GetDateTime(),
		Details:     details,
	}
}
