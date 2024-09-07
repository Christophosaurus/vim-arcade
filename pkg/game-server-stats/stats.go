package gameserverstats

import (
	"context"
	"fmt"
)

type GameServerConfig struct {
	Id string

	Connections int

	// TODO possible?
	Load float32

    Host string

    Port int
}

func (g *GameServerConfig) Addr() string {
    return fmt.Sprintf("%s:%d", g.Host, g.Port)
}

// TODO I don't know what to call this thing...
type GSSRetriever interface {
    Iter() func(yield func(i int, s GameServerConfig) bool)
    Run(ctx context.Context)
    Update(stats GameServerConfig) error
}
