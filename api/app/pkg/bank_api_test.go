// TestBankApi_GetBankID is a unit test function that tests the GetBankID method of the BankApi struct.
// It verifies that the returned bank ID matches the expected bank ID.
//
// The test creates a new instance of BankApi and sets the expected bank ID to "1".
// It then calls the GetBankID method and compares the returned bank ID with the expected bank ID.
// If the bank IDs do not match, an error is reported using the testing.T.Errorf function.
package pkg

import (
	"testing"
)

func TestBankApi_GetBankID(t *testing.T) {
	api := NewBankApi()
	expectedBankID := "1"

	bankID := api.GetBankID()

	if bankID != expectedBankID {
		t.Errorf("Expected bank ID to be %s, got %s", expectedBankID, bankID)
	}
}
