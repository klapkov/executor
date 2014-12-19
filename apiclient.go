package executor

import (
	"io"

	"github.com/pivotal-golang/lager"
)

//go:generate counterfeiter -o fakes/fake_client.go . Client

type Client interface {
	Ping() error
	AllocateContainer(request Container) (Container, error)
	GetContainer(guid string) (Container, error)
	RunContainer(guid string) error
	StopContainer(guid string) error
	DeleteContainer(guid string) error
	ListContainers(Tags) ([]Container, error)
	RemainingResources() (ExecutorResources, error)
	TotalResources() (ExecutorResources, error)
	GetFiles(guid string, path string) (io.ReadCloser, error)
	SubscribeToEvents() (<-chan Event, error)
}

type ClientProvider interface {
	WithLogger(logger lager.Logger) Client
}
