package service

import "github.com/itsLeonB/posyandu-api/module/home/model"

type HomeService interface {
	GetBidan(id int) (model.BidanHomeResponse, error)
	Get(id int) (model.HomeResponse, error)
}
