package play

import (
	"flag"
	"strings"
)

type Publisher struct {
	allCaps bool
	data []string
}

func NewPublisherFromArgs(args []string) (Publisher, error) {
	fset := flag.NewFlagSet("publish", flag.ContinueOnError)
	allCaps := fset.Bool("c", false, "Capitalise output")
	err := fset.Parse(args)
	if err != nil {
		return Publisher{}, err
	}
	p := Publisher{
		allCaps: *allCaps,
		data: fset.Args(),
	}
	return p, nil
}

func (p Publisher) Publish() string {
	result := strings.Join(p.data, " ")
	if p.allCaps {
		return strings.ToUpper(result)
	}
	return result
}