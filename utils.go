/// utils.go
/// Created by PanDaZhong on 2015/05/10.
/// Copyright (c) 2015年 PanDaZhong. All rights reserved.
///

package utils

import(
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/base64"
	"io"
)

func GetMD5String(src string) string{
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

func GetGUID()string{
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); nil != err {
		return ""
	}
	
	return GetMD5String(base64.URLEncoding.EncodeToString(b))
}
