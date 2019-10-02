ldbcp copies a key from one leveldb into another leveldb

Syntax:

```
ldbcp /path/to/db1#key /path/to/db2#key
```

db1 and db2 can be the same database. If db2 does not exist, it will be created
with the new key.
