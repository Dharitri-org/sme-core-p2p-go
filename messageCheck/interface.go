package messagecheck

import "github.com/Dharitri-org/sme-core/core"

type p2pSigner interface {
	Verify(payload []byte, pid core.PeerID, signature []byte) error
}
