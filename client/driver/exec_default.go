//+build darwin dragonfly freebsd netbsd openbsd solaris windows

package driver

import (
	"github.com/ncodes/nomad/client/config"
	"github.com/ncodes/nomad/helper"
	"github.com/ncodes/nomad/nomad/structs"
)

func (d *ExecDriver) Fingerprint(cfg *config.Config, node *structs.Node) (bool, error) {
	d.fingerprintSuccess = helper.BoolToPtr(false)
	return false, nil
}
