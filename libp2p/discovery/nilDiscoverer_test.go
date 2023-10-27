package discovery_test

import (
	"testing"

	"github.com/Dharitri-org/sme-core-p2p-go/libp2p/discovery"
	"github.com/Dharitri-org/sme-core/core/check"
	"github.com/stretchr/testify/assert"
)

func TestNilDiscoverer(t *testing.T) {
	t.Parallel()

	nd := discovery.NewNilDiscoverer()

	assert.False(t, check.IfNil(nd))
	assert.Equal(t, discovery.NullName, nd.Name())
	assert.Nil(t, nd.Bootstrap())
}
