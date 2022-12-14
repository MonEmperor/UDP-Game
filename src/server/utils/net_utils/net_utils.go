package net_utils

import (
	"encoding/json"
	"fmt"
	rs "server/routes/route_structs"
)

func FormatResJson[T any](resData T, header rs.Header) []byte {
	/*
		marshals the sent data to json
		appends the prefix to the data
	*/
	res, _ := json.Marshal(resData)
	marshalled_head, _ := json.Marshal(header)
	stringified := fmt.Sprintf("%s|%s", string(marshalled_head), string(res))
	return []byte(stringified)
}

func FormatResBinary(body []byte, header rs.Header) []byte {

	marshalled_head, _ := json.Marshal(header)

	ret := marshalled_head
	ret = append(ret, byte('|'))
	ret = append(ret, body...)

	return ret
}
