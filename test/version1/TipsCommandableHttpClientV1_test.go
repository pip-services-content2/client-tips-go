package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-content2/client-tips-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type tipsCommandableHttpClientV1Test struct {
	client  *version1.TipsCommandableHttpClientV1
	fixture *TipsClientFixtureV1
}

func newTipsCommandableHttpClientV1Test() *tipsCommandableHttpClientV1Test {
	return &tipsCommandableHttpClientV1Test{}
}

func (c *tipsCommandableHttpClientV1Test) setup(t *testing.T) *TipsClientFixtureV1 {
	var HTTP_HOST = os.Getenv("HTTP_HOST")
	if HTTP_HOST == "" {
		HTTP_HOST = "localhost"
	}
	var HTTP_PORT = os.Getenv("HTTP_PORT")
	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", HTTP_HOST,
		"connection.port", HTTP_PORT,
	)

	c.client = version1.NewTipsCommandableHttpClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewTipsClientFixtureV1(c.client)

	return c.fixture
}

func (c *tipsCommandableHttpClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableHttpCrudOperations(t *testing.T) {
	c := newTipsCommandableHttpClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
