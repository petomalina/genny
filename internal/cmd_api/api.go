package cmd_api

import (
	"fmt"
	"strings"
)

func ServiceDefinition(name, version string) string {
	pkgName := strings.Replace(name, "/", ".", -1) + "." + version
	// this is now in format such as ['documents', 'v1']
	// or ['documents', 'sheets', 'v1']
	svcSlice := strings.Split(name, "/")

	// create grpc service name
	svcName := ""
	for _, s := range svcSlice {
		svcName += strings.Title(s)
	}

	return fmt.Sprintf(`syntax = "proto3";

package %s;

import "google/protobuf/empty.proto";

service %sService {
  rpc GetHealth(google.protobuf.Empty) returns (ServiceHealth) {}
}

message ServiceHealth {
  string status = 1;
}
`, pkgName, svcName)
}
