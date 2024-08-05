package main

import (
	"fmt"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/streamingfast/substreams-gear/protobuf"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	firecore "github.com/streamingfast/firehose-core"
	"github.com/streamingfast/logging"
	"github.com/streamingfast/substreams-gear/generator"
	"github.com/streamingfast/substreams-gear/metadata"
	"go.uber.org/zap"
)

var specVersions = []string{"100", "120", "130", "140", "210", "310", "320", "330", "340", "350", "1000", "1010", "1020", "1030", "1040", "1050", "1110", "1200", "1210", "1300", "1310", "1400", "1410", "1420"}

func NewToolsGenerate(logger *zap.Logger, tracer logging.Tracer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate output proto, extrinsic proto and gen_types.go file for decoding all the types from the output proto.",
		RunE:  generateE(logger, tracer),
	}

	//cmd.Flags().String("output-path", "vara.proto", "Output decoded vara proto file location")
	//cmd.Flags().String("extrinsic-path", "extrinsics.proto", "Extrinsics proto file location")
	//cmd.Flags().String("gen-path", "gen_types.go", "Generated types files file location")

	return cmd
}

func generateE(logger *zap.Logger, tracer logging.Tracer) firecore.CommandExecutor {
	knownMessages := map[string]*protobuf.Message{}
	metadatas := map[string]*types.Metadata{}
	return func(cmd *cobra.Command, args []string) error {
		for _, version := range specVersions {
			logger.Info("decoding metadata types for version", zap.String("version", version))

			md, err := metadata.LoadMetadata(version)
			if err != nil {
				return fmt.Errorf("loading metadata: %w", err)
			}
			metadatas[version] = md

			metadataConverter := metadata.NewMetadataConverter(md, version, logger, tracer)
			err = metadataConverter.Convert(knownMessages)
			logger.Info("done converting metadata", zap.String("version", version))
			if err != nil {
				return fmt.Errorf("converting metadata: %w", err)
			}
		}

		messages := make([]*protobuf.Message, 0)
		for _, seenMsg := range knownMessages {
			if seenMsg != nil {

				messages = append(messages, seenMsg)
			}
		}
		// sort the messages to have them in the same order on each run
		sort.Slice(messages, func(i, j int) bool {
			iV, err := strconv.ParseInt(messages[i].SpecVersion, 10, 64)
			if err != nil {
				panic("parsing iV")
			}
			jV, err := strconv.ParseInt(messages[j].SpecVersion, 10, 64)
			if err != nil {
				panic("parsing jV")
			}

			is := fmt.Sprintf("%s%s%06d", messages[i].Pallet, messages[i].Name, iV)
			js := fmt.Sprintf("%s%s%06d", messages[j].Pallet, messages[j].Name, jV)

			return is < js
		})

		var sb strings.Builder
		sb.WriteString("syntax = \"proto3\";\n")
		sb.WriteString("package sf.substreams.vara.type.v1;\n")
		sb.WriteString("option go_package = \"github.com/streamingfast/substreams-gear/pb/sf/substreams/vara/type/v1;\";\n\n")
		for _, m := range messages {
			s := m.ToProto()
			sb.WriteString(s)
		}

		protoPath := "proto/sf/substreams/vara/type/v1/"

		err := os.MkdirAll(protoPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", protoPath, err)
		}

		err = os.WriteFile(filepath.Join(protoPath, "vara.proto"), []byte(sb.String()), 0644)
		if err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}

		logger.Info("running types generator")
		gen := generator.NewGenerator("templates/gen_types.go.gotmpl", messages, metadatas)
		content, err := gen.Generate()
		if err != nil {
			return fmt.Errorf("generating output proto: %w", err)
		}

		goPackage := "generated/convert"
		err = os.MkdirAll(goPackage, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory generated: %w", err)
		}

		err = os.WriteFile(filepath.Join(goPackage, "converter.go"), content, 0644)
		if err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}

		dbg := generator.NewDecodedBlockGenerator("templates/extrinsic.proto.gotmpl", messages)
		content, err = dbg.Generate()
		if err != nil {
			return fmt.Errorf("generating extrinsic proto: %w", err)
		}

		err = os.WriteFile(filepath.Join(protoPath, "extrinsics.proto"), content, 0644)
		if err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}

		return nil
	}
}
