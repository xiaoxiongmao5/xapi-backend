package store

// 存储用户token信息
var TokenMemoryStore map[string]bool

// 存储所有接口id与客户端SDK包内对应的函数名的映射
var InterfaceFuncName map[int64]string
