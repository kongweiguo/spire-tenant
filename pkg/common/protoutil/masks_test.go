package protoutil_test

import (
	"testing"

	"github.com/kongweiguo/spire-api-sdk-tenant/proto/spire/api/types"
	"github.com/kongweiguo/spire-tenant/pkg/common/protoutil"
	"github.com/kongweiguo/spire-tenant/proto/spire/common"
	"github.com/kongweiguo/spire-tenant/test/spiretest"
)

func TestAllTrueMasks(t *testing.T) {
	spiretest.AssertProtoEqual(t, &types.AgentMask{
		AttestationType:      true,
		X509SvidSerialNumber: true,
		X509SvidExpiresAt:    true,
		Selectors:            true,
		Banned:               true,
		CanReattest:          true,
	}, protoutil.AllTrueAgentMask)

	spiretest.AssertProtoEqual(t, &types.BundleMask{
		X509Authorities: true,
		JwtAuthorities:  true,
		RefreshHint:     true,
		SequenceNumber:  true,
	}, protoutil.AllTrueBundleMask)

	spiretest.AssertProtoEqual(t, &types.EntryMask{
		SpiffeId:       true,
		ParentId:       true,
		Selectors:      true,
		X509SvidTtl:    true,
		JwtSvidTtl:     true,
		FederatesWith:  true,
		Admin:          true,
		CreatedAt:      true,
		Downstream:     true,
		ExpiresAt:      true,
		DnsNames:       true,
		RevisionNumber: true,
		StoreSvid:      true,
		Hint:           true,
	}, protoutil.AllTrueEntryMask)

	spiretest.AssertProtoEqual(t, &common.BundleMask{
		RootCas:         true,
		JwtSigningKeys:  true,
		RefreshHint:     true,
		SequenceNumber:  true,
		X509TaintedKeys: true,
	}, protoutil.AllTrueCommonBundleMask)

	spiretest.AssertProtoEqual(t, &common.AttestedNodeMask{
		AttestationDataType: true,
		CertSerialNumber:    true,
		CertNotAfter:        true,
		NewCertSerialNumber: true,
		NewCertNotAfter:     true,
		CanReattest:         true,
	}, protoutil.AllTrueCommonAgentMask)

	spiretest.AssertProtoEqual(t, &types.FederationRelationshipMask{
		BundleEndpointUrl:     true,
		BundleEndpointProfile: true,
		TrustDomainBundle:     true,
	}, protoutil.AllTrueFederationRelationshipMask)
}
