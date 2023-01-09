package build

import (
	clients1 "github.com/pip-services-content2/client-tips-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type TipsServiceFactory struct {
	*cbuild.Factory
}

func NewTipsServiceFactory() *TipsServiceFactory {
	c := &TipsServiceFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-tips", "client", "null", "*", "1.0")
	mockClientDescriptor := cref.NewDescriptor("service-tips", "client", "mock", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-tips", "client", "commandable-http", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewTipsNullClientV1)
	c.RegisterType(mockClientDescriptor, clients1.NewTipsMockClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewTipsCommandableHttpClientV1)

	return c
}
