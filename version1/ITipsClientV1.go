package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type ITipsClientV1 interface {
	GetTips(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*TipV1], error)

	GetRandomTip(ctx context.Context, correlationId string, filter *data.FilterParams) (*TipV1, error)

	GetTipById(ctx context.Context, correlationId string, tipId string) (*TipV1, error)

	CreateTip(ctx context.Context, correlationId string, tip *TipV1) (*TipV1, error)

	UpdateTip(ctx context.Context, correlationId string, tip *TipV1) (*TipV1, error)

	DeleteTipById(ctx context.Context, correlationId string, tipId string) (*TipV1, error)
}
