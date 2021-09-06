// Copyright (c) Inlets Author(s) 2021. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package cmd

import (
	"fmt"
	"os"

	"github.com/morikuni/aec"
	"github.com/spf13/cobra"
)

var (
	Version   string
	GitCommit string
)

func init() {
	inletsCmd.AddCommand(versionCmd)
}

// inletsCmd represents the base command when called without any sub commands.
var inletsCmd = &cobra.Command{
	Use:   "cevka",
	Short: "Expose your local endpoints to the Internet.",
	Long: `
Expose your local endpoints to the Internet by creating a tunnel between 
your local machine and an exit-server with its own public IP address.

WARNING: When used over the Internet, you must add TLS to secure the
cevka control-plane. 

cevka is made available for free, so support the tools that you 
use through GitHub Sponsors: https://github.com/sponsors/inlets`,
	Run: runInlets,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the clients version information.",
	Run:   parseBaseCommand,
}

func getVersion() string {
	if len(Version) != 0 {
		return Version
	}
	return "dev"
}

func parseBaseCommand(_ *cobra.Command, _ []string) {
	printLogo()

	fmt.Println("Version:", getVersion())
	fmt.Println("Git Commit:", GitCommit)
	os.Exit(0)
}

// Execute adds all child commands to the root command(InletsCmd) and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the InletsCmd.
func Execute(version, gitCommit string) error {

	// Get Version and GitCommit values from main.go.
	Version = version
	GitCommit = gitCommit

	if err := inletsCmd.Execute(); err != nil {
		return err
	}
	return nil
}

func runInlets(cmd *cobra.Command, args []string) {
	printLogo()
	cmd.Help()
}

func printLogo() {
	inletsLogo := aec.WhiteF.Apply(inletsFigletStr)
	fmt.Println(inletsLogo)
}

const inletsFigletStr = `       _ _   _           _                            ___                     
  __ _(_) |_| |__  _   _| |__   ___ ___  _ __ ___    / / |__  _ __ ___   __ _ 
 / _' | | __| '_ \| | | | '_ \ / __/ _ \| '_ ' _ \  / /| '_ \| '_ ' _ \ / _' |
| (_| | | |_| | | | |_| | |_) | (_| (_) | | | | | |/ / | |_) | | | | | | (_| |
 \__, |_|\__|_| |_|\__,_|_.__(_)___\___/|_| |_| |_/_/  |_.__/|_| |_| |_|\__,_|
 |___/                                                                        
             _ _            _         __             _         
 _ __   ___ (_) | _____   _(_) ___   / /__ _____   _| | ____ _ 
| '_ \ / _ \| | |/ _ \ \ / / |/ __| / / __/ _ \ \ / / |/ / _' |
| | | | (_) | | | (_) \ V /| | (__ / / (_|  __/\ V /|   < (_| |
|_| |_|\___// |_|\___/ \_/ |_|\___/_/ \___\___| \_/ |_|\_\__,_|
          |__/                                                 
`
