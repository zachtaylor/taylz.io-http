package websocket

import "taylz.io/types"

// Transport returns literal export data
func Transport(uri string, data types.Dict) types.Bytes {
	return types.BytesDict(types.Dict{
		"uri":  uri,
		"data": data,
	})
}
