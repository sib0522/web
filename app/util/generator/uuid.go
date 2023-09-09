package generator

import "github.com/google/uuid"

// NewUuidV1 returns UUID based on current timestamp and MAC address.
func NewUuidV1() uuid.UUID {
	return uuid.Must(uuid.NewUUID())
}

// NewUuidV2 returns DCE Security UUID based on POSIX UID/GID.
func NewUuidV2(domain uuid.Domain, id uint32) uuid.UUID {
	return uuid.Must(uuid.NewDCESecurity(domain, id))
}

// NewUuidV3 returns UUID based on MD5 hash of namespace UUID and name.
func NewUuidV3(space uuid.UUID, name []byte) uuid.UUID {
	return uuid.NewMD5(space, name)
}

// NewUuidV4 returns random generated UUID.
func NewUuidV4() uuid.UUID {
	return uuid.New()
}

// NewUuidV4String returns random generated UUID string.
func NewUuidV4String() string {
	return uuid.NewString()
}

// NewUuidV5 returns UUID based on SHA-1 hash of namespace UUID and name.
func NewUuidV5(space uuid.UUID, name []byte) uuid.UUID {
	return uuid.NewSHA1(space, name)
}
