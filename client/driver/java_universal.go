// +build !linux

package driver

import cstructs "github.com/ncodes/nomad/client/structs"

func (d *JavaDriver) FSIsolation() cstructs.FSIsolation {
	return cstructs.FSIsolationNone
}
