package define

// Outlines the layout of a bakefile for HCL decoding
type BakeFile struct{
	Groups []GroupOptions `hcl:"group,block"`
	Targets []TargetOptions `hcl:"target,block"`	
	Variables []Variable `hcl:"variable,block"`
}

// This struct lists all the groups defined in the HCL bakefile and their members
// Groups used to execute multiple targets at once
type GroupOptions struct{
	GroupName string `hcl:",label"`
	GroupMembers []string `hcl:"targets,attr"`
}

// TargetOptions contains the information describe targets as defined in the docker bake reference
// Each target has the equivelent of a BuildOptions worth of information
type TargetOptions struct {
	TargetName string `hcl:",label"`
	// Used to define build arguments for target // --build-arg flag
	Args []string `hcl:"args,attr"`
	// KEY-VALUE Pair of annotations to images
	Annotations []string `hcl:"annotations,attr"`
	// Build Attestations to target
	Attest []string `hcl:"attest,attr"`
	// Specifies the frontend options to use
	Call string `hcl:"call,attr"`
	// Specifies the location of the build context to use
	Context string `hcl:"context,attr"`
	// Specifies the location of additional build contexts to use
	Contexts []string `hcl:"contexts,attr"`
	// Defines a human-readable description for hte target
	Description string `hcl:"description,attr"`
	// Specifies whether the dockerfile defined inline
	InlineDocker string `hcl:"dockerfile-inline,attr"`
	// Specifies path to dockerfile
	DockerFile string `hcl:"dockerfile,attr"`
	// Permissions that the build process requires to run
	// network.host: Allows build to use host network commands
	// security.insecure: allows the build to run commands in privileged containers
	Entitlements []string `hcl:"entitlements,attr"`
	// This target inherits attributes from another target
	Inherits []string `hcl:"inherits,attr"`
	// Assigns image labels to the build
	Labels []string `hcl:"labels,attr"`
	// MATRIX to be implemented later
	// Includes target.name and target.matrix
	// Specifies the network mode for the build request
	Network string `hcl:"network,attr"`
	// Dont use build cache for specified stages
	NoCacheFilter []string `hcl:"no-cache-filter,attr"`
	// Dont use cache while building images
	NoCache bool `hcl:"no-cache,attr"`
	// Config for exporting build output
	Output []string `hcl:"output,attr"`
	// Set target platforms for build target
	Platforms []string `hcl:"platforms,attr"`
	// Configures whether the builder should pull while building target
	PullPolicy bool `hcl:"pull,attr"`
	// Defines secrets to expose to the build target
	Secrets []string `hcl:"secret,attr"`
	// ShmSize is the "size" value to use when mounting an shmfs on the container's /dev/shm directory.
	ShmSize string `hcl:"shm-size,attr"`
	// SSHSources is the available ssh agent connections to forward in the build
	SSHSources []string `hcl:"ssh,attr"`
	// Tags to add to the image that we write, if we know of a
	// way to add them.
	Tags []string `hcl:"tags,attr"`
	// Target the targeted FROM in the Dockerfile to build.
	Target string `hcl:"target,attr"`
	// Ulimit specifies resource limit options, in the form type:softlimit[:hardlimit].
	// These types are recognized:
	// "core": maximum core dump size (ulimit -c)
	// "cpu": maximum CPU time (ulimit -t)
	// "data": maximum size of a process's data segment (ulimit -d)
	// "fsize": maximum size of new files (ulimit -f)
	// "locks": maximum number of file locks (ulimit -x)
	// "memlock": maximum amount of locked memory (ulimit -l)
	// "msgqueue": maximum amount of data in message queues (ulimit -q)
	// "nice": niceness adjustment (nice -n, ulimit -e)
	// "nofile": maximum number of open files (ulimit -n)
	// "nproc": maximum number of processes (ulimit -u)
	// "rss": maximum size of a process's (ulimit -m)
	// "rtprio": maximum real-time scheduling priority (ulimit -r)
	// "rttime": maximum amount of real-time execution between blocking syscalls
	// "sigpending": maximum number of pending signals (ulimit -i)
	// "stack": maximum stack size (ulimit -s)
	Ulimit []string `hcl:"ulimits,attr"`
}

// Defines an environmental variable used in a HCL bakefile
type Variable struct{
	Key string `hcl:",label"`
	Value string `hcl:"default,attr"`
}