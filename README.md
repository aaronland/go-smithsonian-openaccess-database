# go-smithsonian-openaccess-database

## Important

This package has been superseded by the [go-wunderkammer](https://github.com/aaronland/go-wunderkammer) package. You should use that instead. For example:

```
$> sqlite3 /usr/local/go-wunderkammer/npg.db \
	< /usr/local/go-wunderkammer/schema/sqlite/oembed.db
	
$> /usr/local/go-smithsonian/bin/emit \
	-oembed \
	-bucket-uri file:///usr/local/OpenAccess \
	metadata/objects/NPG \

   | /usr/local/go-wunderkammer/bin/wunderkammer-db \
	-database-dsn 'sql://sqlite3/usr/local/go-wunderkammer/npg.db'
	
$> sqlite3 /usr/local/go-wunderkammer/schema/sqlite/oembed.db
SQLite version 3.32.1 2020-05-25 16:19:56
Enter ".help" for usage hints.
sqlite> SELECT COUNT(url) FROM oembed;
10542
```

## See also

* https://github.com/aaronland/go-smithsonian-openaccess
* https://github.com/aaronland/go-wunderkammer
* https://github.com/Smithsonian/OpenAccess