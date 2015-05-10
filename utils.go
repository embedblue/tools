/// utils.go
/// Created by PanDaZhong on 2015/05/10.
/// Copyright (c) 2015年 PanDaZhong. All rights reserved.
///

package utils

import(
	"crypto/md5"
	"encoding/hex"
)

func GetMD5String(src string) string{
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil)))
}
