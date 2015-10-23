package wallet

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/FactomProject/btcutil/base58"
	"github.com/FactomProject/go-bip32"
	"github.com/FactomProject/go-bip39"

	"github.com/FactomProject/factomd/common/factoid"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

func MnemonicStringToPrivateKey(mnemonic string) ([]byte, error) {
	mnemonic = strings.ToLower(strings.TrimSpace(mnemonic))
	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, "")
	if err != nil {
		return nil, err
	}

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}

	child, err := masterKey.NewChildKey(bip32.FirstHardenedChild + 7)
	if err != nil {
		return nil, err
	}

	return child.Key, nil
}

func MnemonicStringToPrivateKeyString(mnemonic string) (string, error) {
	key, err := MnemonicStringToPrivateKey(mnemonic)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", key), nil
}

/******************************************************************************/
/********************Human-readible private keys*******************************/
/******************************************************************************/

func HumanReadableFactoidPrivateKeyToPrivateKey(human string) ([]byte, error) {
	human = strings.TrimSpace(human)
	base, v1, v2, err := base58.CheckDecode(human)
	if err != nil {
		return nil, err
	}

	if v1 != 0x64 || v2 != 0x78 {
		return nil, fmt.Errorf("Invalid prefix")
	}

	return base, nil
}

func HumanReadableFactoidPrivateKeyToPrivateKeyString(human string) (string, error) {
	key, err := HumanReadableFactoidPrivateKeyToPrivateKey(human)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", key), nil
}

func HumanReadableECPrivateKeyToPrivateKey(human string) ([]byte, error) {
	human = strings.TrimSpace(human)
	base, v1, v2, err := base58.CheckDecode(human)
	if err != nil {
		return nil, err
	}

	if v1 != 0x5d || v2 != 0xb6 {
		return nil, fmt.Errorf("Invalid prefix")
	}

	return base, nil
}

func HumanReadableECPrivateKeyToPrivateKeyString(human string) (string, error) {
	key, err := HumanReadableECPrivateKeyToPrivateKey(human)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", key), nil
}

func PrivateKeyStringToHumanReadableFactoidPrivateKey(priv string) (string, error) {
	return PrivateKeyStringToHumanReadablePrivateKey(priv, 0x64, 0x78)
}

func PrivateKeyStringToHumanReadableECPrivateKey(priv string) (string, error) {
	return PrivateKeyStringToHumanReadablePrivateKey(priv, 0x5d, 0xb6)
}

func PrivateKeyStringToHumanReadablePrivateKey(priv string, b1, b2 byte) (string, error) {
	priv = strings.TrimSpace(priv)

	h, err := hex.DecodeString(priv)
	if err != nil {
		return "", err
	}
	
	return string(h),fmt.Errorf("Issue with base58")
//	return base58.CheckEncodeWithVersionBytes(h, b1, b2), nil
}

/******************************************************************************/
/***************************To public key**************************************/
/******************************************************************************/

func PrivateKeyStringToPublicKeyString(private string) (string, error) {
	pub, err := PrivateKeyStringToPublicKey(private)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", pub), nil
}

func PrivateKeyStringToPublicKey(private string) ([]byte, error) {
	privHex, err := hex.DecodeString(private)
	if err != nil {
		return nil, err
	}
	return PrivateKeyToPublicKey(privHex)
}

func PrivateKeyToPublicKey(private []byte) ([]byte, error) {
	pub, _, err := GenerateKeyFromPrivateKey(private)
	if err != nil {
		return nil, err
	}
	return pub, nil
}

/******************************************************************************/
/****************************To addresses**************************************/
/******************************************************************************/

func PublicKeyStringToFactoidAddressString(public string) (string, error) {
	pubHex, err := hex.DecodeString(public)
	if err != nil {
		return "", err
	}
	add, err := PublicKeyToFactoidAddress(pubHex)
	if err != nil {
		return "", err
	}

	return primitives.ConvertFctAddressToUserStr(add), nil
}

func PublicKeyToFactoidAddress(public []byte) (interfaces.IAddress, error) {
	rcd := factoid.NewRCD_1(public)
	add, err := rcd.GetAddress()
	if err != nil {
		return nil, err
	}
	return add, nil
}

/******************************************************************************/
/******************************Combined****************************************/
/******************************************************************************/

func HumanReadiblePrivateKeyStringToEverythingString(private string) (string, string, string, error) {
	priv, err := HumanReadableFactoidPrivateKeyToPrivateKeyString(private)
	if err != nil {
		return "", "", "", err
	}
	return PrivateKeyStringToEverythingString(priv)
}

func PrivateKeyStringToEverythingString(private string) (string, string, string, error) {
	pub, err := PrivateKeyStringToPublicKeyString(private)
	if err != nil {
		return "", "", "", err
	}
	add, err := PublicKeyStringToFactoidAddressString(pub)
	if err != nil {
		return "", "", "", err
	}
	return private, pub, add, nil
}
