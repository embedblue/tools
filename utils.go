/// utils.go
/// Created by PanDaZhong on 2015/05/10.
/// Copyright (c) 2015年 PanDaZhong. All rights reserved.
///

package utils

import(
	"io"
	"fmt"
	"time"
	"strconv"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/base64"
	"encoding/binary"
	"bytes"
	"math"
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

func ToString(raw interface{})string{
	return fmt.Sprintf("%v", raw)
}

func StringToUint32(raw string)uint32{
	res, _ := strconv.ParseUint(raw, 10, 32)
	return uint32(res)
}

func StringToInt64(raw string)int64{
	res, _ := strconv.ParseInt(raw, 10, 64)
	return res
}

func WriteString(w *bytes.Buffer, str string) {
	binary.Write(w, binary.LittleEndian, uint16(len(str)))
	binary.Write(w, binary.LittleEndian, []byte(str))
}

func ReadString(r *bytes.Buffer) string {
	var n uint16
	ReadLE(r, &n)
	return string(r.Next(int(n)))
}

func StrToUint32(str string)uint32{
	i, _ := strconv.Atoi(str)
	return uint32(i)
}

func WriteUint32LE(w *bytes.Buffer, i uint32){
	binary.Write(w, binary.LittleEndian, i)
}

func ReadLE(r io.Reader, n interface{}) error {
	return binary.Read(r, binary.LittleEndian, n)
}

func NowStringSeconds()string{
	return strconv.Itoa(int(time.Now().Unix() / 1000))
}

func SliceContains(s []string, e string)bool{
	if nil == s {
		return false
	}
	for _, v := range s {
		if v == e {
			return true
		}
	}

	return false
}

func Math_round(f float64, n int)float64{
	pow10_n := math.Pow10(n)
	return math.Trunc((f + 0.5/pow10_n) * pow10_n)/pow10_n;
}

/// 根据某天的某个时刻的时间戳，获取当天0:0:0的时间戳
func GetZeroSecondsFromTimestamp(ts int64)int64{
	today_0_ts := GetZeroSecondsToday()
	past_time := today_0_ts - ts
	if past_time > 0 {
		return today_0_ts - (int64(past_time/86400) + 1) * 86400
	}else {
		return today_0_ts + (int64(past_time * -1 / 86400)) * 86400
	}
}

func Pkt(msgType uint8, msgHead *bytes.Buffer, msgBody *bytes.Buffer)[]byte{
	binary.Write(msgHead, binary.LittleEndian, msgType)
	WriteUint32LE(msgHead, uint32(msgBody.Len()))
	binary.Write(msgHead, binary.LittleEndian, msgBody.Bytes())
	return msgHead.Bytes()
}
