/// utils.go
/// Created by PanDaZhong on 2015/05/10.
/// Copyright (c) 2015年 PanDaZhong. All rights reserved.
///

package utils

import(
	"io"
	"time"
	"strconv"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/base64"
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

func GetZeroSecondsToday()int64{
	y, m, d := time.Now().Date()
	today := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	return today.Unix()
}

func NowSecondsString()string{
	return strconv.Itoa(int(time.Now().Unix()))
}
