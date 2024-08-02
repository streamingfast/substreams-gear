package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
	firecore "github.com/streamingfast/firehose-core"
	"github.com/streamingfast/logging"
	"github.com/streamingfast/substreams-gear/generator"
	"github.com/streamingfast/substreams-gear/metadata"
	"go.uber.org/zap"
)

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
	return func(cmd *cobra.Command, args []string) error {
		logger.Info("decoding metadata types")

		version := "1420"
		md, err := metadata.LoadMetadata(version)
		if err != nil {
			return fmt.Errorf("loading metadata: %w", err)
		}
		metadataConverter := metadata.NewMetadataConverter(md, logger, tracer)
		messages, err := metadataConverter.Convert()
		if err != nil {
			return fmt.Errorf("converting metadata: %w", err)
		}

		// sort the messages to have them in the same order on each run
		sort.Slice(messages, func(i, j int) bool {
			return messages[i].FullTypeName() < messages[j].FullTypeName()
		})

		var sb strings.Builder
		sb.WriteString("syntax = \"proto3\";\n")
		sb.WriteString("package sf.substreams.gear.type.v1;\n")
		sb.WriteString("option go_package = \"github.com/streamingfast/substreams-gear/pb/sf/substreams/gear/type/v1;\";\n\n")
		for _, m := range messages {
			s := m.ToProto()
			sb.WriteString(s)
		}

		protoPath := "proto/sf/substreams/gear/type/v1/"

		err = os.MkdirAll(protoPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", protoPath, err)
		}

		err = os.WriteFile(filepath.Join(protoPath, "vara.proto"), []byte(sb.String()), 0644)
		if err != nil {
			return fmt.Errorf("failed to write file: %w", err)
		}

		logger.Info("running types generator")
		gen := generator.NewGenerator("templates/gen_types.go.gotmpl", messages, md, version)
		content, err := gen.Generate()
		if err != nil {
			return fmt.Errorf("generating output proto: %w", err)
		}

		goPackage := "generated/convert" + version
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
