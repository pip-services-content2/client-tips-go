package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type TipsCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewTipsCommandableHttpClientV1() *TipsCommandableHttpClientV1 {
	return &TipsCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/tips"),
	}
}

func (c *TipsCommandableHttpClientV1) GetTips(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*TipV1], error) {
	res, err := c.CallCommand(ctx, "get_tips", correlationId, data.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	))

	if err != nil {
		return *data.NewEmptyDataPage[*TipV1](), err
	}

	return clients.HandleHttpResponse[data.DataPage[*TipV1]](res, correlationId)
}

func (c *TipsCommandableHttpClientV1) GetRandomTip(ctx context.Context, correlationId string, filter *data.FilterParams) (*TipV1, error) {
	res, err := c.CallCommand(ctx, "get_random_tip", correlationId, data.NewAnyValueMapFromTuples(
		"filter", filter,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*TipV1](res, correlationId)
}

func (c *TipsCommandableHttpClientV1) GetTipById(ctx context.Context, correlationId string, tipId string) (*TipV1, error) {
	res, err := c.CallCommand(ctx, "get_tip_by_id", correlationId, data.NewAnyValueMapFromTuples(
		"tip_id", tipId,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*TipV1](res, correlationId)
}

func (c *TipsCommandableHttpClientV1) CreateTip(ctx context.Context, correlationId string, tip *TipV1) (*TipV1, error) {
	res, err := c.CallCommand(ctx, "create_tip", correlationId, data.NewAnyValueMapFromTuples(
		"tip", tip,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*TipV1](res, correlationId)
}

func (c *TipsCommandableHttpClientV1) UpdateTip(ctx context.Context, correlationId string, tip *TipV1) (*TipV1, error) {
	res, err := c.CallCommand(ctx, "update_tip", correlationId, data.NewAnyValueMapFromTuples(
		"tip", tip,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*TipV1](res, correlationId)
}

func (c *TipsCommandableHttpClientV1) DeleteTipById(ctx context.Context, correlationId string, tipId string) (*TipV1, error) {
	res, err := c.CallCommand(ctx, "delete_tip_by_id", correlationId, data.NewAnyValueMapFromTuples(
		"tip_id", tipId,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*TipV1](res, correlationId)
}
