package goutils

import "time"

const SysTimeform = "2006-01-02 15:04:05"

const SysTimeformShort = "2006-01-02"
//中国时区
var SysTimeLocation,_=time.LoadLocation("Asia/Chongqing")