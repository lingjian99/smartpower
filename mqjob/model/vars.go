package model

import (
	"xorm.io/xorm"
)

var ErrNotFound = xorm.ErrNotExist

var (
	TableFormalSysStructure   = "formal_sys_structure"
	TableGCRvcElectricalError = "gc_rvc_electrical_error"
)

