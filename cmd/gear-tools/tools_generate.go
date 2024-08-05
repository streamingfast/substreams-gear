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

var specVersions = []string{"1420", "1410", "1400", "1310", "1300", "1210", "1200", "1110", "1050", "1040", "1030", "1020", "1010", "1000", "350", "340", "330", "320", "310", "210", "140", "130", "120", "100"}

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
		for _, version := range specVersions {
			logger.Info("decoding metadata types for version", zap.String("version", version))

			md, err := metadata.LoadMetadata(version)
			if err != nil {
				return fmt.Errorf("loading metadata: %w", err)
			}
			metadataConverter := metadata.NewMetadataConverter(md, version, logger, tracer)
			messages, err := metadataConverter.Convert()
			logger.Info("done converting metadata", zap.String("version", version), zap.Int("messages", len(messages)))
			if err != nil {
				return fmt.Errorf("converting metadata: %w", err)
			}

			if len(messages) == 0 {
				logger.Info("no messages to generate")
				continue
			}

			// sort the messages to have them in the same order on each run
			sort.Slice(messages, func(i, j int) bool {
				return messages[i].FullTypeName() < messages[j].FullTypeName()
			})

			var sb strings.Builder
			sb.WriteString("syntax = \"proto3\";\n")
			sb.WriteString(fmt.Sprintf("package sf.substreams.vara.type.v%s;\n", version))
			sb.WriteString(fmt.Sprintf("option go_package = \"github.com/streamingfast/substreams-gear/pb/sf/substreams/vara/type/v%s;\";\n\n", version))
			for _, m := range messages {
				s := m.ToProto()
				sb.WriteString(s)
			}

			protoPath := fmt.Sprintf("proto/sf/substreams/vara/type/v%s/", version)

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

			dbg := generator.NewDecodedBlockGenerator("templates/extrinsic.proto.gotmpl", messages, version)
			content, err = dbg.Generate()
			if err != nil {
				return fmt.Errorf("generating extrinsic proto: %w", err)
			}

			err = os.WriteFile(filepath.Join(protoPath, "extrinsics.proto"), content, 0644)
			if err != nil {
				return fmt.Errorf("failed to write file: %w", err)
			}
		}

		return nil
	}
}
