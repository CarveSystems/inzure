// Code generated by "stringer -type ContainerPermission"; DO NOT EDIT.

package inzure

import "fmt"

const _ContainerPermission_name = "ContainerAccessUnknownContainerAccessPrivateContainerAccessBlobContainerAccessContainer"

var _ContainerPermission_index = [...]uint8{0, 22, 44, 63, 87}

func (i ContainerPermission) String() string {
	if i >= ContainerPermission(len(_ContainerPermission_index)-1) {
		return fmt.Sprintf("ContainerPermission(%d)", i)
	}
	return _ContainerPermission_name[_ContainerPermission_index[i]:_ContainerPermission_index[i+1]]
}
