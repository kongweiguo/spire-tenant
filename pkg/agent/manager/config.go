package manager

import (
	"crypto/x509"
	"sync"
	"time"

	"github.com/andres-erbsen/clock"
	"github.com/sirupsen/logrus"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/kongweiguo/spire-tenant/pkg/agent/catalog"
	managerCache "github.com/kongweiguo/spire-tenant/pkg/agent/manager/cache"
	"github.com/kongweiguo/spire-tenant/pkg/agent/manager/storecache"
	"github.com/kongweiguo/spire-tenant/pkg/agent/plugin/keymanager"
	"github.com/kongweiguo/spire-tenant/pkg/agent/plugin/nodeattestor"
	"github.com/kongweiguo/spire-tenant/pkg/agent/storage"
	"github.com/kongweiguo/spire-tenant/pkg/agent/svid"
	"github.com/kongweiguo/spire-tenant/pkg/agent/trustbundlesources"
	"github.com/kongweiguo/spire-tenant/pkg/agent/workloadkey"
	"github.com/kongweiguo/spire-tenant/pkg/common/rotationutil"
	"github.com/kongweiguo/spire-tenant/pkg/common/telemetry"
	"github.com/kongweiguo/spire-tenant/pkg/common/tlspolicy"
)

// Config holds a cache manager configuration
type Config struct {
	// Agent SVID and key resulting from successful attestation.
	SVID                     []*x509.Certificate
	SVIDKey                  keymanager.Key
	Bundle                   *managerCache.Bundle
	Reattestable             bool
	Catalog                  catalog.Catalog
	TrustDomain              spiffeid.TrustDomain
	Log                      logrus.FieldLogger
	Metrics                  telemetry.Metrics
	ServerAddr               string
	Storage                  storage.Storage
	TrustBundleSources       *trustbundlesources.Bundle
	RebootstrapMode          string
	RebootstrapDelay         time.Duration
	WorkloadKeyType          workloadkey.KeyType
	SyncInterval             time.Duration
	UseSyncAuthorizedEntries bool
	RotationInterval         time.Duration
	SVIDStoreCache           *storecache.Cache
	X509SVIDCacheMaxSize     int
	JWTSVIDCacheMaxSize      int
	DisableLRUCache          bool
	NodeAttestor             nodeattestor.NodeAttestor
	RotationStrategy         *rotationutil.RotationStrategy
	TLSPolicy                tlspolicy.Policy

	// Clk is the clock the manager will use to get time
	Clk clock.Clock
}

// New creates a cache manager based on c's configuration
func New(c *Config) Manager {
	return newManager(c)
}

func newManager(c *Config) *manager {
	if c.SyncInterval == 0 {
		c.SyncInterval = 5 * time.Second
	}

	if c.RotationInterval == 0 {
		c.RotationInterval = svid.DefaultRotatorInterval
	}

	if c.Clk == nil {
		c.Clk = clock.New()
	}

	cache := managerCache.NewLRUCache(c.Log.WithField(telemetry.SubsystemName, telemetry.CacheManager), c.TrustDomain, c.Bundle,
		c.Metrics, c.X509SVIDCacheMaxSize, c.JWTSVIDCacheMaxSize, c.Clk)

	rotCfg := &svid.RotatorConfig{
		SVIDKeyManager:   keymanager.ForSVID(c.Catalog.GetKeyManager()),
		Log:              c.Log,
		Metrics:          c.Metrics,
		SVID:             c.SVID,
		SVIDKey:          c.SVIDKey,
		BundleStream:     cache.SubscribeToBundleChanges(),
		ServerAddr:       c.ServerAddr,
		TrustDomain:      c.TrustDomain,
		Interval:         c.RotationInterval,
		Clk:              c.Clk,
		NodeAttestor:     c.NodeAttestor,
		Reattestable:     c.Reattestable,
		RotationStrategy: c.RotationStrategy,
		TLSPolicy:        c.TLSPolicy,
	}
	svidRotator, client := svid.NewRotator(rotCfg)

	m := &manager{
		cache:          cache,
		c:              c,
		mtx:            new(sync.RWMutex),
		svid:           svidRotator,
		storage:        c.Storage,
		client:         client,
		clk:            c.Clk,
		svidStoreCache: c.SVIDStoreCache,

		processedTaintedX509Authorities: make(map[string]struct{}),
		processedTaintedJWTAuthorities:  make(map[string]struct{}),
	}

	return m
}
