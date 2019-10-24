package main

import (
  "github.com/syndtr/goleveldb/leveldb"
  "log"
  "os"
  "strings"
)

func main() {
  if len(os.Args) != 2 {
    log.Fatalf("Usage: ldbrm /path/to/db#key")
  }
  src := os.Args[1]
  srcSplit := strings.Split(src, "#")
  if len(srcSplit) != 2 {
    log.Fatalf("argument should be in form path#key, got %v", srcSplit)
  }
  srcPath, srcKey := srcSplit[0], []byte(srcSplit[1])
  srcDB, err := leveldb.OpenFile(srcPath, nil)
  if err != nil {
    log.Fatalf("Error opening src ldb: %v", err.Error())
  }
  defer srcDB.Close()
  if err := srcDB.Delete(srcKey, nil); err != nil {
    log.Fatalf("Error copying src key: %v", err.Error())
  }
}
