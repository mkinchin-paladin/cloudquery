package resiliencehub

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildAppComponentCompliances(t *testing.T, mock *mocks.MockResiliencehubClient) {
	var l resiliencehub.ListAppComponentCompliancesOutput
	require.NoError(t, faker.FakeObject(&l))

	l.NextToken = nil
	mock.EXPECT().ListAppComponentCompliances(
		gomock.Any(),
		gomock.Any(),
		gomock.Any(),
	).Return(&l, nil)
}
