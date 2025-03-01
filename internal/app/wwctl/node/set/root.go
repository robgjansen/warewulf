package set

import "github.com/spf13/cobra"

var (
	baseCmd = &cobra.Command{
		Use:   "set [flags] [node pattern]...",
		Short: "Configure node properties",
		Long: "This command will allow you to set configuration properties for nodes.\n\n" +
			"Note: use the string 'UNSET' to remove a configuration",
		Args: cobra.MinimumNArgs(1),
		RunE: CobraRunE,
	}
	SetComment        string
	SetContainer      string
	SetKernel         string
	SetKernelArgs     string
	SetNetDev         string
	SetIpaddr         string
	SetNetmask        string
	SetGateway        string
	SetHwaddr         string
	SetType           string
	SetNetDevDel      bool
	SetNetDevDefault  bool
	SetClusterName    string
	SetIpxe           string
	SetRuntimeOverlay string
	SetSystemOverlay  string
	SetIpmiIpaddr     string
	SetIpmiNetmask    string
	SetIpmiPort       string
	SetIpmiGateway    string
	SetIpmiUsername   string
	SetIpmiPassword   string
	SetIpmiInterface  string
	SetNodeAll        bool
	SetYes            bool
	SetProfile        string
	SetAddProfile     []string
	SetDelProfile     []string
	SetForce          bool
	SetInit           string
	SetDiscoverable   bool
	SetUndiscoverable bool
	SetRoot           string
	SetKey            string
	SetValue          string
	SetKeyDel         bool
)

func init() {
	baseCmd.PersistentFlags().StringVar(&SetComment, "comment", "", "Set a comment for this node")
	baseCmd.PersistentFlags().StringVarP(&SetContainer, "container", "C", "", "Set the container (VNFS) for this node")
	baseCmd.PersistentFlags().StringVarP(&SetKernel, "kernel", "K", "", "Set Kernel version for nodes")
	baseCmd.PersistentFlags().StringVarP(&SetKernelArgs, "kernelargs", "A", "", "Set Kernel argument for nodes")
	baseCmd.PersistentFlags().StringVarP(&SetClusterName, "cluster", "c", "", "Set the node's cluster group")
	baseCmd.PersistentFlags().StringVar(&SetIpxe, "ipxe", "", "Set the node's iPXE template name")
	baseCmd.PersistentFlags().StringVarP(&SetInit, "init", "i", "", "Define the init process to boot the container")
	baseCmd.PersistentFlags().StringVar(&SetRoot, "root", "", "Define the rootfs")

	baseCmd.PersistentFlags().StringVarP(&SetRuntimeOverlay, "runtime", "R", "", "Set the node's runtime overlay")
	baseCmd.PersistentFlags().StringVarP(&SetSystemOverlay, "system", "S", "", "Set the node's system overlay")
	baseCmd.PersistentFlags().StringVar(&SetIpmiIpaddr, "ipmi", "", "Set the node's IPMI IP address")
	baseCmd.PersistentFlags().StringVar(&SetIpmiNetmask, "ipminetmask", "", "Set the node's IPMI netmask")
	baseCmd.PersistentFlags().StringVar(&SetIpmiPort, "ipmiport", "", "Set the node's IPMI port")
	baseCmd.PersistentFlags().StringVar(&SetIpmiGateway, "ipmigateway", "", "Set the node's IPMI gateway")
	baseCmd.PersistentFlags().StringVar(&SetIpmiUsername, "ipmiuser", "", "Set the node's IPMI username")
	baseCmd.PersistentFlags().StringVar(&SetIpmiPassword, "ipmipass", "", "Set the node's IPMI password")
	baseCmd.PersistentFlags().StringVar(&SetIpmiInterface, "ipmiinterface", "", "Set the node's IPMI interface (defaults to 'lan' if empty)")

	baseCmd.PersistentFlags().StringSliceVar(&SetAddProfile, "addprofile", []string{}, "Add Profile(s) to node")
	baseCmd.PersistentFlags().StringSliceVar(&SetDelProfile, "delprofile", []string{}, "Remove Profile(s) to node")
	baseCmd.PersistentFlags().StringVarP(&SetProfile, "profile", "P", "", "Set the node's profile members (comma separated)")

	baseCmd.PersistentFlags().StringVarP(&SetNetDev, "netdev", "N", "", "Define the network device to configure")
	baseCmd.PersistentFlags().StringVarP(&SetIpaddr, "ipaddr", "I", "", "Set the node's network device IP address")
	baseCmd.PersistentFlags().StringVarP(&SetNetmask, "netmask", "M", "", "Set the node's network device netmask")
	baseCmd.PersistentFlags().StringVarP(&SetGateway, "gateway", "G", "", "Set the node's network device gateway")
	baseCmd.PersistentFlags().StringVarP(&SetHwaddr, "hwaddr", "H", "", "Set the node's network device HW address")
	baseCmd.PersistentFlags().StringVarP(&SetType, "type", "T", "", "Set the node's network device type")
	baseCmd.PersistentFlags().BoolVar(&SetNetDevDel, "netdel", false, "Delete the node's network device")
	baseCmd.PersistentFlags().BoolVar(&SetNetDevDefault, "netdefault", false, "Set this network to be default")

	baseCmd.PersistentFlags().StringVarP(&SetKey, "key", "k", "", "Define custom key")
	baseCmd.PersistentFlags().BoolVar(&SetKeyDel, "keydel", false, "Delete custom key")

	baseCmd.PersistentFlags().StringVarP(&SetValue, "value", "", "", "Set value")

	baseCmd.PersistentFlags().BoolVarP(&SetNodeAll, "all", "a", false, "Set all nodes")

	baseCmd.PersistentFlags().BoolVarP(&SetYes, "yes", "y", false, "Set 'yes' to all questions asked")
	baseCmd.PersistentFlags().BoolVarP(&SetForce, "force", "f", false, "Force configuration (even on error)")
	baseCmd.PersistentFlags().BoolVar(&SetDiscoverable, "discoverable", false, "Make this node discoverable")
	baseCmd.PersistentFlags().BoolVar(&SetUndiscoverable, "undiscoverable", false, "Remove the discoverable flag")

}

// GetRootCommand returns the root cobra.Command for the application.
func GetCommand() *cobra.Command {
	return baseCmd
}
