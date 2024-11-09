package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Gokert/gnss-radar/internal/pkg/model"
	"github.com/Gokert/gnss-radar/internal/store"
)

type ListRequest struct {
	X string
	Y string
	Z string
}

type RinexRequest struct{}

type IGnss interface {
	ListGnss(ctx context.Context, req ListRequest) ([]*store.ListResult, error)
	Rinexlist(ctx context.Context, req RinexRequest) ([]*model.RinexResults, error)
}

type GnssService struct {
	gnssStore store.IGnssStore
}

func NewGnssService(store store.IGnssStore) *GnssService {
	return &GnssService{gnssStore: store}
}

func (g *GnssService) ListGnss(ctx context.Context, req ListRequest) ([]*store.ListResult, error) {
	xf, err := strconv.ParseFloat(req.X, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat: %w", err)
	}
	yf, err := strconv.ParseFloat(req.Y, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat: %w", err)
	}
	zf, err := strconv.ParseFloat(req.Z, 64)
	if err != nil {
		return nil, fmt.Errorf("strconv.ParseFloat: %w", err)
	}

	gnss, err := g.gnssStore.ListGnssCoords(ctx, store.ListParams{X: xf, Y: yf, Z: zf})
	if err != nil {
		return nil, fmt.Errorf("gnssStore.List: %w", err)
	}

	return gnss, nil
}

func (g *GnssService) Rinexlist(ctx context.Context, req RinexRequest) ([]*model.RinexResults, error) {
	rinexlist, err := g.gnssStore.RinexList(ctx)
	if err != nil {
		return nil, fmt.Errorf("gnssStore.ListRinexlist: %w", err)
	}

	return rinexlist, nil
}
