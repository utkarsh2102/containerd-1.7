/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package client

import (
	"fmt"
	"os/exec"
	"syscall"

	"github.com/containerd/cgroups/v3/cgroup1"
)

func getSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		Setpgid: true,
	}
}

func setCgroup(cgroupPath string, cmd *exec.Cmd) error {
	cg, err := cgroup1.Load(cgroup1.StaticPath(cgroupPath))
	if err != nil {
		return fmt.Errorf("failed to load cgroup %s: %w", cgroupPath, err)
	}
	if err := cg.AddProc(uint64(cmd.Process.Pid)); err != nil {
		return fmt.Errorf("failed to join cgroup %s: %w", cgroupPath, err)
	}
	return nil
}
