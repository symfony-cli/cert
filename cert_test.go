package cert

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"testing"
	"time"
)

func TestCA_IsExpired(t *testing.T) {
	isMacOS = func() bool { return true }
	ca := CA{
		cert: &x509.Certificate{
			SerialNumber: big.NewInt(123),
			Subject: pkix.Name{
				Organization:       []string{"Symfony tests"},
				OrganizationalUnit: []string{"Symfony tests"},
			},

			NotBefore: time.Now(),
			NotAfter:  time.Now().AddDate(0, 0, 825),
		},
	}

	if got, want := ca.IsExpired(), false; got != want {
		t.Errorf("IsExpired() = %v, want %v", got, want)
	}
	if got, want := ca.MustBeRegenerated(), false; got != want {
		t.Errorf("MustBeRegenerated() = %v, want %v", got, want)
	}

	ca = CA{
		cert: &x509.Certificate{
			SerialNumber: big.NewInt(123),
			Subject: pkix.Name{
				Organization:       []string{"Symfony tests"},
				OrganizationalUnit: []string{"Symfony tests"},
			},

			NotBefore: time.Unix(0, 0),
			NotAfter:  time.Unix(0, 0).AddDate(0, 0, 3650),
		},
	}

	if got, want := ca.IsExpired(), true; got != want {
		t.Errorf("IsExpired() = %v, want %v", got, want)
	}
	if got, want := ca.MustBeRegenerated(), false; got != want {
		t.Errorf("MustBeRegenerated() = %v, want %v", got, want)
	}

	ca = CA{
		cert: &x509.Certificate{
			SerialNumber: big.NewInt(123),
			Subject: pkix.Name{
				Organization:       []string{"Symfony tests"},
				OrganizationalUnit: []string{"Symfony tests"},
			},

			NotBefore: time.Date(2019, 8, 1, 0, 0, 0, 0, time.UTC),
			NotAfter:  time.Date(2029, 8, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	if got, want := ca.MustBeRegenerated(), true; got != want {
		t.Errorf("MustBeRegenerated() [for macOS] = %v, want %v", got, want)
	}
	isMacOS = func() bool { return false }
	if got, want := ca.MustBeRegenerated(), false; got != want {
		t.Errorf("MustBeRegenerated() [for non macOS] = %v, want %v", got, want)
	}
}
