package models

import (
	proto "golang-learning/market_place/proto"
)

type Detail struct {
	Name            string  `json:"name" bson:"name"`
	Category        string  `json:"category" bson:"category"`
	UnitPrice       float32 `json:"unitPrice" bson:"unitPrice"`
	UnitPriceWeight string  `json:"unitPriceWeight" bson:"unitPriceWeight"`
	Count           float32 `json:"count" bson:"count"`
	CountUnit       string  `json:"countUnit" bson:"countUnit"`
}

func (m *Detail) ToPb() *proto.Detail {
	return &proto.Detail{
		Name:            m.Name,
		Category:        m.Category,
		UnitPrice:       m.UnitPrice,
		UnitPriceWeight: m.UnitPriceWeight,
		Count:           m.Count,
		CountUnit:       m.CountUnit,
	}
}

func DetailPbToStruct(pbDetail *proto.Detail) *Detail {
	if pbDetail == nil {
		return nil
	}

	return &Detail{
		Name:            pbDetail.GetName(),
		Category:        pbDetail.GetCategory(),
		UnitPrice:       pbDetail.GetUnitPrice(),
		UnitPriceWeight: pbDetail.GetUnitPriceWeight(),
		Count:           pbDetail.GetCount(),
		CountUnit:       pbDetail.GetCountUnit(),
	}
}
