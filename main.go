package main

import (
	"fmt"
	"github.com/wwqgtxx/mfp/combine_part"
	"github.com/wwqgtxx/mfp/find_pair"
	"github.com/wwqgtxx/mfp/format_file"
	"github.com/wwqgtxx/mfp/gb2utf8"
	"github.com/wwqgtxx/mfp/to_db"
	"os"

	"github.com/wwqgtxx/mfp/mr1"
	"github.com/wwqgtxx/mfp/mr2"
	"github.com/wwqgtxx/mfp/mr3"
	"github.com/wwqgtxx/mfp/sort_kv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("expected subcommands")
		os.Exit(1)
	}
	switch os.Args[1] {

	case "mapper1":
		mr1.Mapper()
	case "reducer1":
		mr1.Reducer()
	case "mapper2":
		mr2.Mapper()
	case "reducer2":
		mr2.Reducer()
	case "mapper3":
		mr3.Mapper()
	case "reducer3":
		mr3.Reducer()
	case "sort_kv":
		sort_kv.SortKV()
	case "combine_part":
		combine_part.CombinePart()
	case "gb2utf8":
		gb2utf8.GB2UTF8()
	case "format_file":
		format_file.FormatFile()
	case "to_db":
		to_db.ToDB()
	case "find_pair":
		find_pair.FindPair()
	default:
		fmt.Println("expected subcommands")
		os.Exit(1)
	}

}
