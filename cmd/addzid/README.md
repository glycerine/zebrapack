addzid
======

Summary: addzid automatically adds `zid:"0"`, `zid:"1"`, ... tags to your Go structs.

Given a set of golang (Go) source files, addzid will tag the public
struct fields with sequential zid tags. This prepares your source
so that it can be fed to the `zebrapack` codegen tool.


to install: *run make*. This lets us record version info.
--------

use
---------

~~~
use: addzid {-o outdir} myGoSourceFile.go myGoSourceFile2.go ...
     # addzid makes it easy to add ZebraPack serialization[1] to Go source files.
     # addzid reads .go files and adds `zid` tags to struct fields.
     #
     # options:
     #   -o="odir" specifies the directory to write to (created if need be).
     #   -X add tag to private fields of Go structs as well as public.
     #   -version   shows build version with git commit hash.
     #   -debug     print lots of debug info as we process.
     #   -OVERWRITE modify .go files in-place, adding zid tags (by default
     #       we write to the to -o dir).
     #
     # required: at least one .go source file for struct definitions.
     #  note: the .go source file to process must be listed last, after any options.
     #
     # [1] https://github.com/glycerine/zebrapack 
~~~


zid tags on go structs
--------------------------

When you run `addzid`, it will generate a modified copy of your go source files in the output directory.

These new versions include zid tags on all public fields of structs. You should inspect the copy of the source file in the output directory, and then replace your original source with the tagged version.  Of course you can always manually add zid tags to fields, if you prefer. However `addzid` simplifies this chore.

If you are feeling especially bold, `addzid -OVERWRITE my.go` will replace my.go with the zid tagged version. For safety, only do this on backed-up and version controlled source files.

By default only public fields (with a Capital first letter in their name) are tagged. The -X flag ignores the public/private distinction, and tags all fields.

The zid tags allow the ZebraPack schema evolution to function properly as you add new fields to structs.

example of zid annotion use
~~~
type Job struct { 
   C int `zid:"2"`  // we added C later, thus it is numbered higher.
   A int `zid:"0"`
   B int `zid:"1"` 
}
~~~

other tags
----------

Also available tags: `zid:"skip"` or `zid:"-1"` (any negative number): this field will be skipped and not serialized or written to the schema.


windows build script
---------------------------
see `build.cmd`. Thanks to Klaus Post (http://klauspost.com) for contributing this.

-----
-----

Copyright (c) 2016, Jason E. Aten, Ph.D.

