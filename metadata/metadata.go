package metadata

import (
	"embed"
	"fmt"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
)

//go:embed data/*
var dataFs embed.FS

func LoadMetadata(version string) (*types.Metadata, error) {
	h, err := dataFs.ReadFile("data/" + version + ".hex")
	if err != nil {
		return nil, fmt.Errorf("read metadata file for version %q: %w", version, err)
	}
	metadata := &types.Metadata{}
	err = codec.DecodeFromHex(string(h), metadata)
	if err != nil {
		return nil, fmt.Errorf("decoding metadata: %w", err)
	}
	return metadata, nil
}
