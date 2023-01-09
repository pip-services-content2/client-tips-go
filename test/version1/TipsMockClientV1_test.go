package test_version1

import (
	"testing"

	"github.com/pip-services-content2/client-tips-go/version1"
)

type TipsMockClientV1Test struct {
	client  *version1.TipsMockClientV1
	fixture *TipsClientFixtureV1
}

func newTipsMockClientV1Test() *TipsMockClientV1Test {
	return &TipsMockClientV1Test{}
}

func (c *TipsMockClientV1Test) setup(t *testing.T) *TipsClientFixtureV1 {
	c.client = version1.NewTipsMockClientV1()

	c.fixture = NewTipsClientFixtureV1(c.client)

	return c.fixture
}

func (c *TipsMockClientV1Test) teardown(t *testing.T) {
	c.client = nil
}

func TestMockOperations(t *testing.T) {
	c := newTipsMockClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
