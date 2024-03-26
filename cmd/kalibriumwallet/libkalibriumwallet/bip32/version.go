package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// KalibriumMainnetPrivate is the version that is used for
// Kalibrium mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var KalibriumMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// KalibriumMainnetPublic is the version that is used for
// Kalibrium mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var KalibriumMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// KalibriumTestnetPrivate is the version that is used for
// Kalibrium testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var KalibriumTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// KalibriumTestnetPublic is the version that is used for
// Kalibrium testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var KalibriumTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// KalibriumDevnetPrivate is the version that is used for
// Kalibrium devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var KalibriumDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// KalibriumDevnetPublic is the version that is used for
// Kalibrium devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var KalibriumDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// KalibriumSimnetPrivate is the version that is used for
// Kalibrium simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var KalibriumSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// KalibriumSimnetPublic is the version that is used for
// Kalibrium simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var KalibriumSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case KalibriumMainnetPrivate:
		return KalibriumMainnetPublic, nil
	case KalibriumTestnetPrivate:
		return KalibriumTestnetPublic, nil
	case KalibriumDevnetPrivate:
		return KalibriumDevnetPublic, nil
	case KalibriumSimnetPrivate:
		return KalibriumSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case KalibriumMainnetPrivate:
		return true
	case KalibriumTestnetPrivate:
		return true
	case KalibriumDevnetPrivate:
		return true
	case KalibriumSimnetPrivate:
		return true
	}

	return false
}
