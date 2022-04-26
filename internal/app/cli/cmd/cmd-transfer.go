package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
)

// transferCmd represents the mmp cmdTransfer
var transferCmd = &cobra.Command{
	Use:   "cp <nodeSpec>:(<dir>|<file>) <nodeSpec>:<dir>",
	Short: "Transfer file and directories between nodes",
	Long: appHeader(`Transfer files and directories between nodes.

An existing writable directory must be specified in <dst> when a file
is specified in <src>.

It supports recursive transfers when a directory is specified in <src>.
If an existing directory is specified in <src>, recursive mode is automatically
enabled and <dst> directory is created if not found on the target node.

Syntax:
  cp <src> <dst>

Parameters:
  <src>       <nodeSpec>:(<dir>|<file>)
  <dst>       <nodeSpec>:<dir>

  <nodeSpec>  ( localhost | <account>:<tenant>:<network>:<subnet>:<node> )
  <dir>       existing directory
  <file>      existing file

Examples:
  mm cp localhost:/tmp/file.txt account1:default:default:subnet-000:node1:/tmp/
  mm cp localhost:/tmp/dir/ account1:default:default:subnet-000:node1:/tmp/dir/
  mm cp account1:default:default:subnet-000:node1:/tmp/dir1/ localhost:/tmp/dir2/
  mm cp account1:default:default:subnet-000:node1:/tmp/f.txt localhost:/tmp/dir3/

	`),
	Args: cobra.ExactArgs(2),
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Command().Transfer(args)
	},
}

func init() {}
