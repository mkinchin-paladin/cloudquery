package shield

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/shield"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/plugin-sdk/v3/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildProtections(t *testing.T, ctrl *gomock.Controller) client.Services {
	m := mocks.NewMockShieldClient(ctrl)
	protection := shield.ListProtectionsOutput{}
	require.NoError(t, faker.FakeObject(&protection))
	protection.NextToken = nil
	m.EXPECT().ListProtections(gomock.Any(), gomock.Any(), gomock.Any()).Return(&protection, nil)

	tags := shield.ListTagsForResourceOutput{}
	require.NoError(t, faker.FakeObject(&tags))
	m.EXPECT().ListTagsForResource(gomock.Any(), gomock.Any(), gomock.Any()).Return(&tags, nil)
	return client.Services{
		Shield: m,
	}
}

func TestProtections(t *testing.T) {
	client.AwsMockTestHelper(t, Protections(), buildProtections, client.TestOptions{})
}
