package networksharding_test

import (
	"testing"

	"github.com/Dharitri-org/sme-core-p2p-go/libp2p/networksharding"
	"github.com/Dharitri-org/sme-core/core/check"
	"github.com/stretchr/testify/assert"
)

func TestNilListSharderSharder(t *testing.T) {
	nls := networksharding.NewNilListSharder()

	assert.False(t, check.IfNil(nls))
	assert.Equal(t, 0, len(nls.ComputeEvictionList(nil)))
	assert.False(t, nls.Has("", nil))
	assert.Nil(t, nls.SetPeerShardResolver(nil))
}
