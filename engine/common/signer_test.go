package common

import (
	"fmt"
	"testing"

	internal "github.com/zodiac163/bitget_api/engine"
)

func TestSigner_Sign(t *testing.T) {
	signer := new(Signer)
	result := signer.Sign("GET", "www.bitget.com", "aaaa", internal.TimesStamp())
	fmt.Print(result)
}
