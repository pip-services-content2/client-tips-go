package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-content2/client-tips-go/version1"
	"github.com/stretchr/testify/assert"
)

type TipsClientFixtureV1 struct {
	Client version1.ITipsClientV1
	TIP1   *version1.TipV1
	TIP2   *version1.TipV1
}

func NewTipsClientFixtureV1(client version1.ITipsClientV1) *TipsClientFixtureV1 {
	return &TipsClientFixtureV1{
		Client: client,
		TIP1: &version1.TipV1{
			Id:      "1",
			Topics:  []string{"maintenance"},
			Creator: version1.NewPartyReferenceV1("1", "Test User", ""),
			Title:   map[string]string{"en": "Tip 1"},
			Content: map[string]string{"en": "Sample Tip #1"},
		},

		TIP2: &version1.TipV1{
			Id:      "2",
			Tags:    []string{"TAG 1"},
			Topics:  []string{"maintenance"},
			Creator: version1.NewPartyReferenceV1("1", "Test User", ""),
			Title:   map[string]string{"en": "Tip 2"},
			Content: map[string]string{"en": "Sample Tip #2"},
		},
	}
}

func (c *TipsClientFixtureV1) clear() {
	page, _ := c.Client.GetTips(context.Background(), "", nil, nil)

	for _, t := range page.Data {
		c.Client.DeleteTipById(context.Background(), "", t.Id)
	}
}

func (c *TipsClientFixtureV1) TestCrudOperations(t *testing.T) {
	c.clear()
	defer c.clear()

	// Create one tip
	tip1, err := c.Client.CreateTip(context.Background(), "", c.TIP1)
	assert.Nil(t, err)

	assert.NotNil(t, tip1)
	assert.Equal(t, tip1.Topics, c.TIP1.Topics)
	assert.Equal(t, tip1.Content["en"], c.TIP1.Content["en"])

	// Create another tip
	tip2, err := c.Client.CreateTip(context.Background(), "", c.TIP2)
	assert.Nil(t, err)

	assert.NotNil(t, tip2)
	assert.Equal(t, tip2.Topics, c.TIP2.Topics)
	assert.Equal(t, tip2.Content["en"], c.TIP2.Content["en"])

	// Get all tips
	page, err1 := c.Client.GetTips(context.Background(), "", nil, nil)
	assert.Nil(t, err1)

	assert.NotNil(t, page)
	assert.Len(t, page.Data, 2)

	// Update the tip
	tip1.Content["en"] = "Updated Content 1"
	quote, err := c.Client.UpdateTip(context.Background(), "", tip1)
	assert.Nil(t, err)

	assert.NotNil(t, quote)
	assert.Equal(t, quote.Content["en"], "Updated Content 1")
	assert.Equal(t, quote.Topics, c.TIP1.Topics)

	// Delete quote
	_, err = c.Client.DeleteTipById(context.Background(), "", tip1.Id)
	assert.Nil(t, err)

	// Try to get deleted quote
	quote, err = c.Client.GetTipById(context.Background(), "", tip1.Id)
	assert.Nil(t, err)

	assert.Nil(t, quote)
}
