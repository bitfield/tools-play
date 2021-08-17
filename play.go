package play

import (
	"strings"
)

type Publisher struct {
	data []string
}

func NewPublisherFromArgs(args []string) (Publisher, error) {
	return Publisher{
		data: args,
	}, nil
}

func (p Publisher) Publish() string {
	return strings.Join(p.data, " ")
}