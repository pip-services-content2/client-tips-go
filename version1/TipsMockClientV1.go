package version1

import (
	"context"
	"strings"
	"time"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
	"github.com/pip-services3-gox/pip-services3-commons-gox/random"

	aclients "github.com/pip-services-content2/client-attachments-go/version1"
)

type TipsMockClientV1 struct {
	tips []*TipV1

	attachmentsConnector *AttachmentsConnector
	attachmentsClient    *aclients.AttachmentsMockClientV1
}

func NewTipsMockClientV1() *TipsMockClientV1 {
	c := &TipsMockClientV1{
		tips:              make([]*TipV1, 0),
		attachmentsClient: aclients.NewAttachmentsMockClientV1(),
	}

	c.attachmentsConnector = NewAttachmentsConnector(c.attachmentsClient)

	return c
}

func (c *TipsMockClientV1) contains(array1 []string, array2 []string) bool {
	if array1 == nil || array2 == nil {
		return false
	}

	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			if array1[i] == array2[i] {
				return true
			}
		}
	}

	return false
}

func (c *TipsMockClientV1) matchString(value string, search string) bool {
	if value == "" && search == "" {
		return true
	}
	if value == "" || search == "" {
		return false
	}
	return strings.Contains(strings.ToLower(value), strings.ToLower(search))
}

func (c *TipsMockClientV1) matchMultilanguageString(value map[string]string, search string) bool {
	for _, text := range value {
		if c.matchString(text, search) {
			return true
		}
	}

	return false
}

func (c *TipsMockClientV1) matchSearch(item *TipV1, search string) bool {
	search = strings.ToLower(search)
	if c.matchMultilanguageString(item.Title, search) {
		return true
	}
	if c.matchMultilanguageString(item.Content, search) {
		return true
	}
	if item.Creator != nil && c.matchString(item.Creator.Name, search) {
		return true
	}
	return false
}

func (c *TipsMockClientV1) composeFilter(filter *data.FilterParams) func(*TipV1) bool {
	if filter == nil {
		filter = data.NewEmptyFilterParams()
	}

	search, searchOk := filter.GetAsNullableString("search")
	id, idOk := filter.GetAsNullableString("id")
	status, statusOk := filter.GetAsNullableString("status")
	fromCreateTime, fromCreateTimeOk := filter.GetAsNullableDateTime("from_create_time")
	toCreateTime, toCreateTimeOk := filter.GetAsNullableDateTime("to_create_time")

	tagsString := filter.GetAsString("topics")
	tags := make([]string, 0)

	topicsString := filter.GetAsString("topics")
	topics := make([]string, 0)

	// Process tags filter
	if tagsString != "" {
		tags = data.TagsProcessor.CompressTags([]string{tagsString})
	}

	// Process topics
	if topicsString != "" {
		topics = strings.Split(topicsString, ",")
	}

	return func(item *TipV1) bool {
		if idOk && item.Id != id {
			return false
		}
		if len(topics) > 0 && !c.contains(item.Topics, topics) {
			return false
		}
		if statusOk && item.Status != status {
			return false
		}
		if len(tags) > 0 && !c.contains(item.AllTags, tags) {
			return false
		}
		if fromCreateTimeOk && item.CreateTime.Unix() >= fromCreateTime.Unix() {
			return false
		}
		if toCreateTimeOk && item.CreateTime.Unix() < toCreateTime.Unix() {
			return false
		}
		if searchOk && !c.matchSearch(item, search) {
			return false
		}
		return true
	}
}

func (c *TipsMockClientV1) GetTips(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*TipV1], error) {
	filterFunc := c.composeFilter(filter)

	items := make([]*TipV1, 0)
	for _, v := range c.tips {
		item := v
		if filterFunc(item) {
			items = append(items, item)
		}
	}
	return *data.NewDataPage(items, len(c.tips)), nil
}

func (c *TipsMockClientV1) GetRandomTip(ctx context.Context, correlationId string, filter *data.FilterParams) (*TipV1, error) {
	filterFunc := c.composeFilter(filter)

	items := make([]*TipV1, 0)
	for _, v := range c.tips {
		item := v
		if filterFunc(item) {
			items = append(items, item)
		}
	}

	buf := *items[random.Integer.Next(0, len(items))]
	return &buf, nil
}

func (c *TipsMockClientV1) GetTipById(ctx context.Context, correlationId string, tipId string) (result *TipV1, err error) {
	for _, v := range c.tips {
		if v.Id == tipId {
			buf := *v
			result = &buf
			break
		}
	}
	return result, nil
}

func (c *TipsMockClientV1) CreateTip(ctx context.Context, correlationId string, tip *TipV1) (*TipV1, error) {
	tip.CreateTime = time.Now()
	tip.AllTags = data.TagsProcessor.ExtractHashTags(
		"#title.en#title.sp#title.fr#title.de#title.ru#content.en#content.sp#content.fr#content.de#content.ru",
	)

	buf := *tip
	c.tips = append(c.tips, &buf)
	err := c.attachmentsConnector.AddAttachments(ctx, correlationId, &buf)
	return tip, err
}

func (c *TipsMockClientV1) UpdateTip(ctx context.Context, correlationId string, tip *TipV1) (*TipV1, error) {
	if tip == nil {
		return nil, nil
	}

	tip.CreateTime = time.Now()
	tip.AllTags = data.TagsProcessor.ExtractHashTags(
		"#title.en#title.sp#title.fr#title.de#title.ru#content.en#content.sp#content.fr#content.de#content.ru",
	)

	var index = -1
	for i, v := range c.tips {
		if v.Id == tip.Id {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, errors.NewNotFoundError(
			correlationId,
			"TIP_NOT_FOUND",
			"Tip "+tip.Id+" was not found",
		).WithDetails("tip_id", tip.Id)
	}

	oldTip := c.tips[index]

	buf := *tip
	c.tips[index] = &buf

	err := c.attachmentsConnector.UpdateAttachments(ctx, correlationId, oldTip, &buf)

	return tip, err
}

func (c *TipsMockClientV1) DeleteTipById(ctx context.Context, correlationId string, tipId string) (*TipV1, error) {
	var index = -1
	for i, v := range c.tips {
		if v.Id == tipId {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}

	var item = c.tips[index]
	if index < len(c.tips) {
		c.tips = append(c.tips[:index], c.tips[index+1:]...)
	} else {
		c.tips = c.tips[:index]
	}

	err := c.attachmentsConnector.RemoveAttachments(ctx, correlationId, item)
	return item, err
}
