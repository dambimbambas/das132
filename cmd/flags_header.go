package cmd

import (
	"encoding/hex"
	"fmt"

	"github.com/multiformats/go-multiaddr"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	"github.com/celestiaorg/celestia-node/node"
)

var (
	headersTrustedHashFlag = "headers.trusted-hash"
	headersTrustedPeerFlag = "headers.trusted-peer"
)

// HeadersFlags gives a set of hardcoded Header package flags.
func HeadersFlags() *flag.FlagSet {
	flags := &flag.FlagSet{}

	flags.String(
		headersTrustedHashFlag,
		"",
		"Hex encoded header hash. Used to subjectively initialize header synchronization",
	)

	flags.String(
		headersTrustedPeerFlag,
		"",
		"Multiaddr of a reliable peer to fetch headers from. (Format: multiformats.io/multiaddr)",
	)

	return flags
}

// ParseHeadersFlags parses Header package flags from the given cmd and applies values to Env.
func ParseHeadersFlags(cmd *cobra.Command, env *Env) error {
	hash := cmd.Flag(headersTrustedHashFlag).Value.String()
	if hash != "" {
		_, err := hex.DecodeString(hash)
		if err != nil {
			return fmt.Errorf("cmd: while parsing '%s': %w", headersTrustedHashFlag, err)
		}

		env.AddOptions(node.WithTrustedHash(hash))
	}

	tpeer := cmd.Flag(headersTrustedPeerFlag).Value.String()
	if tpeer != "" {
		_, err := multiaddr.NewMultiaddr(tpeer)
		if err != nil {
			return fmt.Errorf("cmd: while parsing '%s': %w", headersTrustedPeerFlag, err)
		}

		env.AddOptions(node.WithTrustedPeer(tpeer))
	}

	return nil
}
