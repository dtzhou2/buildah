package main

import (
	// "fmt"
	// "os"
	"github.com/containers/buildah/imagebuildah"
	buildahcli "github.com/containers/buildah/pkg/cli"
	// "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type bakeOptions struct {
	files string
	print bool
}

func init() {
	var (
		bakeDescription = `
  			Creates a repeteable build configuration using a single HCL/YAML bakefile`
		opts bakeOptions
	)

	bakeCommand := &cobra.Command{
		Use:     "bake [CONTEXT]",
		Aliases: []string{""},
		Short:   "Build an image using instructions in a Containerfile",
		Long:    bakeDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			return bakeCmd(cmd, args, opts)
		},
		Args: cobra.MaximumNArgs(1),
		Example: `buildah bake myapp
		buildah bake -f myapp.hcl --print `,
	}
	bakeCommand.SetUsageTemplate(UsageTemplate())

	flags := bakeCommand.Flags()
	flags.SetInterspersed(false)

	// build is a all common flags
	flags.StringVarP(&opts.files, "files", "f", "compose.yaml", "specifies filepath of bakefile")
	flags.BoolVar(&opts.print, "print", false, "pretty prints specified bakefile")

	rootCmd.AddCommand(bakeCommand)
}

func bakeCmd(c *cobra.Command, inputArgs []string, opts bakeOptions) error {
	var err error

	if err := buildahcli.VerifyFlagsArgsOrder(inputArgs); err != nil {
		return err
	}
	
	store, err := getStore(c)
	if err != nil {
		return err
	}

	imagebuildah.BakeFiles(getContext(), store, opts.files, opts.print, inputArgs...);

	return err
}
