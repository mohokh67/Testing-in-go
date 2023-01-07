package unique_id

import (
	"github.com/google/uuid"
	"github.com/segmentio/ksuid"
)

func UUIDGenerator() string {
	id := uuid.New()
	return id.String()
}

func KSUIDGenerator() string {
	ksuidExample := ksuid.New()
	return ksuidExample.String()
}
