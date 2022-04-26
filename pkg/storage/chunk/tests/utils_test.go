package tests

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/mar4uk/loki/pkg/storage/chunk/client"
	"github.com/mar4uk/loki/pkg/storage/chunk/client/aws"
	"github.com/mar4uk/loki/pkg/storage/chunk/client/cassandra"
	"github.com/mar4uk/loki/pkg/storage/chunk/client/gcp"
	"github.com/mar4uk/loki/pkg/storage/chunk/client/local"
	"github.com/mar4uk/loki/pkg/storage/chunk/client/testutils"
	"github.com/mar4uk/loki/pkg/storage/stores/series/index"
)

const (
	userID    = "userID"
	tableName = "test"
)

type storageClientTest func(*testing.T, index.Client, client.Client)

func forAllFixtures(t *testing.T, storageClientTest storageClientTest) {
	var fixtures []testutils.Fixture
	fixtures = append(fixtures, aws.Fixtures...)
	fixtures = append(fixtures, gcp.Fixtures...)
	fixtures = append(fixtures, local.Fixtures...)
	fixtures = append(fixtures, cassandra.Fixtures()...)
	fixtures = append(fixtures, Fixtures...)

	for _, fixture := range fixtures {
		t.Run(fixture.Name(), func(t *testing.T) {
			indexClient, objectClient, closer, err := testutils.Setup(fixture, tableName)
			require.NoError(t, err)
			defer closer.Close()
			storageClientTest(t, indexClient, objectClient)
		})
	}
}
