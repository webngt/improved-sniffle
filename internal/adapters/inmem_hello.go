package adapters

import (
	"context"
	"github.com/webngt/improved-sniffle/internal/domain"
)

type saveHistoryRequest struct {
	user   string
	record *domain.HistoryRecord
}

type InmemHello struct {
	history        map[string][]*domain.HistoryRecord
	getHistoryReq  chan string
	getHistoryResp chan *domain.History
	saveHistoryReq chan
}

func (i *InmemHello) GetHello(ctx context.Context, name string) (*domain.Hello, error) {
	hello, err := domain.NewMessageForUser(name)

	if err != nil {
		return nil, err
	}

	return hello, nil
}

func (i *InmemHello) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case helloReq := <-i.getHelloReq:

		}
	}
}
