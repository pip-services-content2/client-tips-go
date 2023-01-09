package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type TipsNullClientV1 struct {
}

func NewTipsNullClientV1() *TipsNullClientV1 {
	return &TipsNullClientV1{}
}

func (c *TipsNullClientV1) GetTips(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*TipV1], error) {
	return *data.NewEmptyDataPage[*TipV1](), nil
}

func (c *TipsNullClientV1) GetRandomTip(ctx context.Context, correlationId string, filter *data.FilterParams) (*TipV1, error) {
	return nil, nil
}

func (c *TipsNullClientV1) GetTipById(ctx context.Context, correlationId string, tipId string) (*TipV1, error) {
	return nil, nil
}

func (c *TipsNullClientV1) CreateTip(ctx context.Context, correlationId string, tip *TipV1) (*TipV1, error) {
	return tip, nil
}

func (c *TipsNullClientV1) UpdateTip(ctx context.Context, correlationId string, tip *TipV1) (*TipV1, error) {
	return tip, nil
}

func (c *TipsNullClientV1) DeleteTipById(ctx context.Context, correlationId string, tipId string) (*TipV1, error) {
	return nil, nil
}
