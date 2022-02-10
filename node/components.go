package node

import (
	"context"

	nodecore "github.com/celestiaorg/celestia-node/node/core"
	"github.com/celestiaorg/celestia-node/node/fxutil"
	"github.com/celestiaorg/celestia-node/node/p2p"
	"github.com/celestiaorg/celestia-node/node/services"
	"github.com/celestiaorg/celestia-node/service/header"
)

// lightComponents keeps all the components as DI options required to built a Light Node.
func lightComponents(cfg *Config, store Store) fxutil.Option {
	return fxutil.Options(
		fxutil.Supply(Light),
		baseComponents(cfg, store),
		fxutil.Provide(services.DASer),
		fxutil.Provide(services.HeaderExchangeP2P(cfg.Services)),
		fxutil.Provide(services.LightAvailability),
	)
}

// bridgeComponents keeps all the components as DI options required to build a Bridge Node.
func bridgeComponents(cfg *Config, store Store) fxutil.Option {
	return fxutil.Options(
		fxutil.Supply(Bridge),
		baseComponents(cfg, store),
		nodecore.Components(cfg.Core, store.Core),
		fxutil.Provide(services.LightAvailability), // TODO(@Wondertan): Remove strict requirements to have Availability
	)
}

// fullComponents keeps all the components as DI options required to build a Full Node.
func fullComponents(cfg *Config, store Store) fxutil.Option {
	return fxutil.Options(
		fxutil.Supply(Full),
		baseComponents(cfg, store),
		fxutil.Provide(services.DASer),
		fxutil.Provide(services.HeaderExchangeP2P(cfg.Services)),
		fxutil.Provide(services.FullAvailability),
	)
}

// baseComponents keeps all the common components shared between different Node types.
func baseComponents(cfg *Config, store Store) fxutil.Option {
	return fxutil.Options(
		fxutil.Provide(context.Background),
		fxutil.Supply(cfg),
		fxutil.Supply(store.Config),
		fxutil.Provide(store.Datastore),
		fxutil.Provide(store.Keystore),
		fxutil.Provide(services.ShareService),
		fxutil.Provide(services.HeaderService),
		fxutil.Provide(services.HeaderStore),
		fxutil.Provide(services.HeaderSyncer(cfg.Services)),
		fxutil.ProvideAs(services.P2PSubscriber, new(header.Broadcaster), new(header.Subscriber)),
		fxutil.Provide(services.HeaderP2PExchangeServer),
		p2p.Components(cfg.P2P),
	)
}
