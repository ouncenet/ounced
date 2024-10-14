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

// OunceMainnetPrivate is the version that is used for
// Ounce mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var OunceMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// OunceMainnetPublic is the version that is used for
// Ounce mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var OunceMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// OunceTestnetPrivate is the version that is used for
// Ounce testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var OunceTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// OunceTestnetPublic is the version that is used for
// Ounce testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var OunceTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// OunceDevnetPrivate is the version that is used for
// Ounce devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var OunceDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// OunceDevnetPublic is the version that is used for
// Ounce devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var OunceDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// OunceSimnetPrivate is the version that is used for
// Ounce simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var OunceSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// OunceSimnetPublic is the version that is used for
// Ounce simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var OunceSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case OunceMainnetPrivate:
		return OunceMainnetPublic, nil
	case OunceTestnetPrivate:
		return OunceTestnetPublic, nil
	case OunceDevnetPrivate:
		return OunceDevnetPublic, nil
	case OunceSimnetPrivate:
		return OunceSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case OunceMainnetPrivate:
		return true
	case OunceTestnetPrivate:
		return true
	case OunceDevnetPrivate:
		return true
	case OunceSimnetPrivate:
		return true
	}

	return false
}
