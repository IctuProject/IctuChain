package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ContractKeyPrefix is the prefix to retrieve all Contract
	ContractKeyPrefix = "Contract/value/"
)

// ContractKey returns the store key to retrieve a Contract from the index fields
func ContractKey(
	uid string,
	req string,
	prov string,
) []byte {
	var key []byte

	uidBytes := []byte(uid)
	key = append(key, uidBytes...)
	key = append(key, []byte("/")...)

	reqBytes := []byte(req)
	key = append(key, reqBytes...)
	key = append(key, []byte("/")...)

	provBytes := []byte(prov)
	key = append(key, provBytes...)
	key = append(key, []byte("/")...)

	return key
}
