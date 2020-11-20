package types

// ProtoModule is a git submodule referencing protobuf repository
// such as the official protobuf/src, api-commons or validators.
// Examples of such repos are:
//  - github.com/googleapis/api-common-protos
//  - github.com/protocolbuffers/protobuf
type ProtoModule struct {
	// Repository is the path to the repository that should be downloaded
	// e.g. https://github.com/protocolbuffers/protobuf
	Repository string `json:"repository"`
	// Path is a path that the module is saved in. Commonly all modules are
	// saved into the 3rdparty/ folder
	Path string `json:"path"`
	// IncludePath is a relative path that should be included using the -I
	// protoc command. If not set, the repo root will be used instead.
	// This variable needs to be set for certain repos, such as the official
	// https://github.com/protocolbuffers/protobuf, where the protobufs reside
	// in a special folder called 'src'.
	//
	// All protobuf imports are then relative to this includePath
	// (Defaults to '')
	IncludePath string `json:"includePath,omitempty"`
}

// API is a reference to a gRPC definition that was created by the user.
type API struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Service is a definition of an implementation of one or more API objects
type Service struct {
	Name string              `json:"name"`
	APIs []APIImplementation `json:"apis"`
}

// APIImplementation defines an API that is implemented for a particular service
// with its options
type APIImplementation struct {
	Name    string `json:"name"`
	Version string `json:"version"`

	HTTPGateway bool `json:"httpGateway,omitempty"`
}

type Config struct {
	Project      string        `json:"project"`
	APIs         []API         `json:"apis"`
	Services     []Service     `json:"services"`
	ProtoModules []ProtoModule `json:"protomodules"`
}
