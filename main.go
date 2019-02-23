package main

//reference
//https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html

// bhoe1aaf1nh4rg0jt7bg
//  1HZErJjycI7c3GAtqc14tHpOSqf
// -LZNvOJAXcwCLAzJIWgQ
// 01D4CEV58CY0MKWBCDHMFJQS5N
// 34aa4461b004b01
// 1550901417231298000-7123594115641680708
// 5759cae5-f999-4f02-88d8-9a406de9a182

// bhoe1aaf1nh4rg0jt7bg
//  1HZErJjycI7c3GAtqc14tHpOSqf
// -LZNvOJAXcwCLAzJIWgQ
// 01D4CEV58CY0MKWBCDHMFJQS5N
// 34aa4461b004b01
// 1550901417231298000-7123594115641680708
// 5759cae5-f999-4f02-88d8-9a406de9a182

// github.com/segmentio/ksuid	0pPKHjWprnVxGH7dEsAoXX2YQvU	4個字節的時間（秒）+16個隨機字節
// github.com/rs/xid	b50vl5e54p1000fo3gh0	4字節時間（秒）+ 3字節機器ID + 2字節進程ID + 3字節隨機
// github.com/kjk/betterguid	-Kmdih_fs4ZZccpx2Hl1	8個字節的時間（毫秒）+ 9個隨機字節
// github.com/sony/sonyflake	20f8707d6000108	~6個字節的時間（10 ms）+ 1個字節的序列+ 2個字節的機器ID
// github.com/oklog/ulid	01BJMVNPBBZC3E36FJTGVF0C4S	6個字節的時間（毫秒）+ 8個字節隨機
// github.com/chilts/sid	1JADkqpWxPx-4qaWY47〜FQI	8個字節的時間（ns）+ 8個隨機字節
// github.com/satori/go.uuid	5b52d72c-82b3-4f8e-beb5-437a974842c	UUIDv4

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/chilts/sid"
	"github.com/kjk/betterguid"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
	uuid "github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
)

// To run:
// go run main.go

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
	// Note: this is base16, could shorten by encoding as base62 string
	fmt.Printf("github.com/sony/sonyflake:   %x\n", id)
}

func genSid() {
	id := sid.Id()
	fmt.Printf("github.com/chilts/sid:       %s\n", id)
}

func genUUIDv4() {
	id, _ := uuid.NewV4()
	fmt.Printf("github.com/satori/go.uuid:   %s\n", id)
}

func main() {
	genXid()
	genKsuid()
	genBetterGUID()
	genUlid()
	genSonyflake()
	genSid()
	genUUIDv4()
}
