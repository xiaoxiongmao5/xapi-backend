package loadconfig

import gstore "xj/xapi-backend/g_store"

func LoadInterfaceFuncNameMap() {
	gstore.InterfaceFuncName = map[int64]string{
		1: "GetNameByGet",
		2: "GetNameByGet",
		3: "GetNameByPost",
	}
}
