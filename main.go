package main

import (
	"os"

	"github.com/nilic/kubectl-rsh/cmd"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	kcmdutil "k8s.io/kubectl/pkg/cmd/util"

	// Import to initialize client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	kubeConfigFlags := genericclioptions.NewConfigFlags(true).WithDiscoveryBurst(350).WithDiscoveryQPS(50.0)
	matchVersionKubeConfigFlags := kcmdutil.NewMatchVersionFlags(kubeConfigFlags)
	f := kcmdutil.NewFactory(matchVersionKubeConfigFlags)
	ioStreams := genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
	cmdRsh := cmd.NewCmdRsh(f, ioStreams)

	flags := cmdRsh.PersistentFlags()
	kubeConfigFlags.AddFlags(flags)
	matchVersionKubeConfigFlags.AddFlags(cmdRsh.PersistentFlags())

	cmdRsh.Execute()
}
