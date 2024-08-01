package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/spf13/cobra"
	"github.com/streamingfast/cli/sflags"
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

	cmd.Flags().String("output-path", "vara.proto", "Output decoded vara proto file location")
	cmd.Flags().String("extrinsic-path", "extrinsics.proto", "Extrinsics proto file location")
	cmd.Flags().String("gen-path", "gen_types.go", "Generated types files file location")

	return cmd
}

func generateE(logger *zap.Logger, tracer logging.Tracer) firecore.CommandExecutor {
	return func(cmd *cobra.Command, args []string) error {

		outputPath := sflags.MustGetString(cmd, "output-path")
		extrinsicPath := sflags.MustGetString(cmd, "extrinsic-path")
		genPath := sflags.MustGetString(cmd, "gen-path")

		logger.Info("decoding metadata types", zap.String("output_path", outputPath))
		metadataConverter := metadata.NewMetadataConverter(logger, tracer)
		contentBytes, err := metadataConverter.Convert()
		if err != nil {
			return fmt.Errorf("converting metadata: %w", err)
		}

		directoryPath := "proto/sf/substreams/gear/type/v1/"

		err = os.MkdirAll(directoryPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", directoryPath, err)
		}

		err = os.WriteFile(filepath.Join(directoryPath, outputPath), contentBytes, 0644)
		if err != nil {
			return fmt.Errorf("failed to write file %s: %w", outputPath, err)
		}

		logger.Info("wrote metadata proto", zap.String("output_path", outputPath))
		messages := metadataConverter.FetchMessages()
		metadata := metadataConverter.FetchMetadata()

		// sort the messages to have them in the same order on each run
		sort.Slice(messages, func(i, j int) bool {
			return messages[i].FullTypeName() < messages[j].FullTypeName()
		})

		logger.Info("running types generator", zap.String("gen_path", genPath))
		gen := generator.NewGenerator("templates/gen_types.go.gotmpl", messages, metadata)
		content, err := gen.Generate()
		if err != nil {
			return fmt.Errorf("generating output proto: %w", err)
		}

		err = os.MkdirAll("generated", os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory generated: %w", err)
		}

		err = os.WriteFile(filepath.Join("generated", genPath), content, 0644)
		if err != nil {
			return fmt.Errorf("failed to write file %s: %w", genPath, err)
		}

		dbg := generator.NewDecodedBlockGenerator("templates/extrinsic.proto.gotmpl", messages)
		content, err = dbg.Generate()
		if err != nil {
			return fmt.Errorf("generating extrinsic proto: %w", err)
		}

		err = os.WriteFile(filepath.Join(directoryPath, extrinsicPath), content, 0644)
		if err != nil {
			return fmt.Errorf("failed to write file %s: %w", extrinsicPath, err)
		}

		return nil
	}
}
