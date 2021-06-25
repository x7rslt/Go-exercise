package repository

import (
	"food/model"
)

type MarketRepo struct {
	DB model.DataBase
}

func (m *MarketRepo) GetMarketInfo(hotelid string) model.Market {
	var market model.Market
	m.DB.MyDB.Where("hotelid=?", hotelid).Find(&market)
	return market
}