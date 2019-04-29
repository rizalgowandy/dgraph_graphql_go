package apitest

import (
	"testing"

	"github.com/romshark/dgraph_graphql_go/apitest/setup"
	"github.com/stretchr/testify/require"
)

// TestCreateUserErr tests all possible user account creation errors
func TestCreateUserErr(t *testing.T) {
	// Test reserved email on creation
	t.Run("reservedEmail", func(t *testing.T) {
		ts := setup.New(t, tcx)
		defer ts.Teardown()

		ts.Help.OK.CreateUser("fooBarowich", "foo@bar.buz")
		res, err := ts.Help.CreateUser("bazBuzowich", "foo@bar.buz")
		require.Nil(t, res)
		verifyError(t, "InvalidInput", err)
	})

	// Test reserved displayName on creation
	t.Run("reservedDisplayName", func(t *testing.T) {
		ts := setup.New(t, tcx)
		defer ts.Teardown()

		ts.Help.OK.CreateUser("fooBarowich", "foo@bar.buz")
		res, err := ts.Help.CreateUser("fooBarowich", "baz@buzowich.buz")
		require.Nil(t, res)
		verifyError(t, "InvalidInput", err)
	})
}
