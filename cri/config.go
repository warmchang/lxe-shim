package cri // import "github.com/automaticserver/lxe/cri"

// Config options that LXE will need to interface with LXD
type Config struct {
	// UnixSocket this LXE will be reachable under
	UnixSocket string
	// LXDSocket where LXD is reachable under
	LXDSocket string
	// LXDRemoteConfig file path where lxd remote settings are stored
	LXDRemoteConfig string
	// LXDImageRemote to use by default when ImageSpec doesn't provide an explicit remote
	LXDImageRemote string
	// LXDProfiles which all cri containers inherit
	LXDProfiles []string
	// LXEStreamingEndpoint contains the listen address for the streaming server
	LXEStreamingEndpoint string
	// LXEStreamingAddress is the base address for constructing streaming URLs
	LXEStreamingAddress string
	// LXEHostnetworkFile file path to use for lxc's raw.include
	LXEHostnetworkFile string
	// Which LXENetworkPlugin to use
	LXENetworkPlugin string
	// LXEBridgeName is the name of the bridge to create and use
	LXEBridgeName string
	// LXEBridgeDHCPRange to configure for lxebr0 if NetworkPlugin is default
	LXEBridgeDHCPRange string
	// CNIConfDir is the path where the cni configuration files are
	CNIConfDir string
	// CNIBinDir is the path where the cni plugins are
	CNIBinDir string
}
