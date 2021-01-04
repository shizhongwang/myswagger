package main

import (
	"fmt"
	"github.com/chilts/sid"
	"github.com/kjk/betterguid"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
	"log"
	"math/rand"
	"time"
)

func main() {
	guid := xid.New()
	fmt.Println(guid.String())  //随机字符串
	fmt.Println(guid.Machine()) //80 165 17
	fmt.Println(guid.Pid())     //随机五位数字
	fmt.Println(guid.Time())    //当前时间
	fmt.Println(guid.Counter()) //随机七位数字

	genXid()
	genKsuid()
	genBetterGUID()
	genUlid()
	genSonyflake()
	genSid()
	//genUUIDv4()
}

func genXid() {
	id := xid.New()
	fmt.Printf("github.com/rs/xid:           %s\n", id.String())
}

func genKsuid() {
	id := ksuid.New()
	fmt.Printf("github.com/segmentio/ksuid:  %s\n", id.String())
}

func genBetterGUID() {
	id := betterguid.New()
	fmt.Printf("github.com/kjk/betterguid:   %s\n", id)
}

func genUlid() {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	fmt.Printf("github.com/oklog/ulid:       %s\n", id.String())
}

func genSonyflake() {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	fmt.Printf("github.com/sony/sonyflake:   %x\n", id)
}

func genSid() {
	id := sid.Id()
	fmt.Printf("github.com/chilts/sid:       %s\n", id)
}

//func genUUIDv4() {
//	id, err := uuid.NewV4()
//	if err != nil {
//		fmt.Printf("get uuid error [%s]", err)
//	}
//	fmt.Printf("github.com/satori/go.uuid:   %s\n", id)
//}
