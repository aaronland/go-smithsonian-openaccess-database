# go-smithsonian-openaccess-database

## Important

Work in progress.

## Tools

### oembed-populate

Populate an [oembed database](oembed/oembed.go) from the output of the [go-smithsonian-openaccess emit tool](https://github.com/aaronland/go-smithsonian-openaccess#emit) for all the [Smithsonian OpenAccess](https://github.com/Smithsonian/OpenAccess) records from the [National Air and Space Museum](https://airandspace.si.edu/.

```
$> sqlite3 nasm.db < schema/sqlite/oembed.sqlite

$> /usr/local/go-smithsonian-openaccess/bin/emit -bucket-uri file:///usr/local/OpenAccess \
   -oembed \
   metadata/objects/NASM \
   | bin/oembed-populate \
   -database-dsn sql://sqlite3/usr/local/go-smithsonian-openaccess-database/nasm.db

$> sqlite3 nasm.db 
sqlite> SELECT COUNT(url) FROM oembed;
2407
```

## See also

* https://github.com/aaronland/go-smithsonian-openaccess
* https://github.com/Smithsonian/OpenAccess