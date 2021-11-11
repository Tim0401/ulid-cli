package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/oklog/ulid/v2"
)

func main() {
	var (
		sb = flag.Bool("sb", false, "string->byte")
		bs = flag.Bool("bs", false, "byte->string")
	)
	flag.Parse()
	if !(*sb || *bs) || (*sb && *bs) {
		log.Fatal("Specify either -sb or -bs")
	}
	arg := flag.Arg(0)

	// string->byte
	if *sb {
		id, err := ulid.ParseStrict(arg)
		if err != nil {
			log.Fatal("Input error")
		}
		fmt.Println(hex.EncodeToString(id[:]))
		return
	}

	// byte->string
	if *bs {
		arg = strings.TrimPrefix(arg, "0x")
		slice, err := hex.DecodeString(arg)
		if err != nil || len(slice) != 16 {
			log.Fatal("Input error")
		}
		array := [16]byte{}
		copy(array[:], slice)
		str := ulid.ULID(array).String()
		fmt.Println(str)
		return
	}
}
