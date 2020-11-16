package idgener

//采用MongoDb类似的方法
//ID由24个字符组成，前八位是时间，接着6位是机器名的哈希，在接着4位是pid，最后6位是自增序号。这样保证在时间，主机，进程这单个维度上是唯一的，
//而且还保留了2的24次方的自增序号，这样即使在分布式环境下每一个id生成器产生的id都是唯一的。

import (
	"encoding/binary"
	"encoding/hex"
	"hash/crc32"
	"os"
	"time"
)

type IGener struct {
	machineCode [3]byte
	pidCode     [2]byte
	second      int64
	idChan      chan string
	inc         uint32
}

func NewIGener() <-chan string {
	ig := &IGener{
		second: time.Now().Unix(),
		inc:    0,
		idChan: make(chan string),
	}
	ig.machineCode, ig.pidCode = ig.machinePidEncode()

	go ig.generUniqueId()

	return ig.idChan
}

func (ig *IGener) machinePidEncode() ([3]byte, [2]byte) {
	hostname, _ := os.Hostname()

	buf := [4]byte{}
	binary.BigEndian.PutUint32(buf[:], crc32.ChecksumIEEE([]byte(hostname)))
	machineCode := [3]byte{buf[1], buf[2], buf[3]}

	pid := os.Getpid()
	pidCode := [2]byte{}
	binary.BigEndian.PutUint16(pidCode[:], uint16(pid))

	return machineCode, pidCode
}

func (ig *IGener) timeIncEncode() ([4]byte, [3]byte) {
	now := time.Now().Unix()
	timeEncode := [4]byte{}
	incEncode := [3]byte{}
	buf := [8]byte{}

	if ig.second != now {
		ig.second = now
		ig.inc = 0
	}

	binary.BigEndian.PutUint64(buf[:], uint64(ig.second))
	timeEncode = [4]byte{buf[4], buf[5], buf[6], buf[7]}

	binary.BigEndian.PutUint32(buf[:], uint32(ig.inc))
	incEncode = [3]byte{buf[1], buf[2], buf[3]}
	ig.inc++

	return timeEncode, incEncode
}

func (ig *IGener) generUniqueId() {
	var uniqueId [12]byte
	for {
		timeEncode, incEncode := ig.timeIncEncode()
		var i int = 0
		for _, b := range timeEncode {
			uniqueId[i] = b
			i++
		}

		for _, b := range ig.machineCode {
			uniqueId[i] = b
			i++
		}

		for _, b := range ig.pidCode {
			uniqueId[i] = b
			i++
		}

		for _, b := range incEncode {
			uniqueId[i] = b
			i++
		}

		ig.idChan <- hex.EncodeToString(uniqueId[:])

	}
}
