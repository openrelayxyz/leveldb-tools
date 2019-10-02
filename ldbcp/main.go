package main

import (
  "github.com/syndtr/goleveldb/leveldb"
  "log"
  "os"
  "strings"
)

func main() {
  if len(os.Args) != 3 {
    log.Fatalf("Usage: ldbcp /path/to/db#key /path/to/db2#key")
  }
  src := os.Args[1]
  dst := os.Args[2]
  srcSplit := strings.Split(src, "#")
  if len(srcSplit) != 2 {
    log.Fatalf("argument should be in form path#key, got %v", srcSplit)
  }
  srcPath, srcKey := srcSplit[0], []byte(srcSplit[1])
  dstSplit := strings.Split(dst, "#")
  if len(dstSplit) != 2 {
    log.Fatalf("argument should be in form path#key, got %v", srcSplit)
  }
  dstPath, dstKey := dstSplit[0], []byte(dstSplit[1])
  srcDB, err := leveldb.OpenFile(srcPath, nil)
  if err != nil {
    log.Fatalf("Error opening src ldb: %v", err.Error())
  }
  dstDB, err := leveldb.OpenFile(dstPath, nil)
  if err != nil {
    log.Fatalf("Error opening dst ldb: %v", err.Error())
  }
  data, err := srcDB.Get(srcKey, nil)
  if err != nil {
    log.Fatalf("Error copying src key: %v", err.Error())
  }
  dstDB.Put(data, dstKey, nil)
}
