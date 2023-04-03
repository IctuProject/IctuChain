package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BalanceKeyPrefix is the prefix to retrieve all Balance
	BalanceKeyPrefix = "Balance/value/"
)

// BalanceKey returns the store key to retrieve a Balance from the index fields
func BalanceKey(
	uid string,
	idContract string,
	requester string,
) []byte {
	var key []byte

	uidBytes := []byte(uid)
	key = append(key, uidBytes...)
	key = append(key, []byte("/")...)

	idContractBytes := []byte(idContract)
	key = append(key, idContractBytes...)
	key = append(key, []byte("/")...)

	requesterBytes := []byte(requester)
	key = append(key, requesterBytes...)
	key = append(key, []byte("/")...)

	return key
}
