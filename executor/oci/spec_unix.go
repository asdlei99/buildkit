// +build !windows

package oci

import (
	"github.com/containerd/containerd/contrib/seccomp"
	"github.com/containerd/containerd/oci"
	"github.com/moby/buildkit/solver/pb"
	"github.com/moby/buildkit/util/entitlements/security"
	"github.com/moby/buildkit/util/system"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

func generateSecurityOpts(mode pb.SecurityMode) ([]oci.SpecOpts, error) {
	if mode == pb.SecurityMode_INSECURE {
		return []oci.SpecOpts{security.WithInsecureSpec()}, nil
	} else if system.SeccompSupported() && mode == pb.SecurityMode_SANDBOX {
		return []oci.SpecOpts{seccomp.WithDefaultProfile()}, nil
	}
	return nil, nil
}

func generateProcessModeOpts(mode ProcessMode) ([]oci.SpecOpts, error) {
	if mode == NoProcessSandbox {
		// Mount for /proc is replaced in GetMounts() anyway
		return []oci.SpecOpts{oci.WithHostNamespace(specs.PIDNamespace)}, nil
		// TODO(AkihiroSuda): Configure seccomp to disable ptrace (and prctl?) explicitly
	}
	return nil, nil
}
