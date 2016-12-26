ZebraPack: a data description language and serialization format. This one is all black and white. No gray areas.
==========

ZebraPack is a data definition language and serialization format. It removes gray areas from msgpack2 serialized data, and provides for declared schemas, sane data evolution, and more compact encoding. It does all this while maintaining the possibility of easy compatibility with all the dynamic languages that already have msgpack2 support.

The code here descends from the fantastic msgpack code generator https://github.com/tinylib/msgp by Philip Hofer.

At this time it is feature complete, and in beta while we get experience with it. See the section below on background and motivation ideas that we are implementing here.

### `msgp:",omitempty"` tags on struct fields

In the following example,
```
type Hedgehog struct {
   Furriness string `msgp:",omitempty"`
}
```
If Furriness is the empty string, the field will not be serialized, thus saving the space of the field name on the wire.

Generally, most zero values that have the omitempty tag applied will be skipped. Recursive struct elements are an exception; they are always included and are never impacted by omitempty tagging recursively. Naturally the member's own fields may have tags that will be in-force locally once the serializer reaches them. Note that member structs are different from member pointers-to-structs. Nil pointers that are tagged `omitempty` will have their field skipped.

Under tuple encoding, all fields are serialized and the omitempty tag is ignored.

# background and motivation

# ZebraPack: a data description language and serialization format. This one is all black and white. No gray areas.

ZebraPack is a data definition language and serialization format. It removes gray areas from msgpack2 serialized data, and provides for declared schemas, sane data evolution, and more compact encoding. It does all this while maintaining the possibility of easy compatibility with all the dynamic languages that already have msgpack2 support.

# the main idea

```
//given this definition, defined in Go:
type A struct {
  Name     string      `zid:"0"`
  Bday     time.Time   `zid:"1"`
  Phone    string      `zid:"2"`
  Sibs     int         `zid:"3"`
  GPA      float64     `zid:"4" deprecated:"true"` // a deprecated field.
  Friend   bool        `zid:"5"`
}

then to serialize the following instance `a`, we would
print the schema information at the front of the file,
in the form of a zebra.StructT structure. Then
the data follows as a map whose keys are now integers
instead of strings:
         
original(msgpack2) ->        header(msgpack2)      +    each instance(msgpack2)
--------                     --------------             -------------
a := A{                      zebra.StructT{             map{
  "Name":  "Atlanta",          0: {"Name", String},       0: "Atlanta",
  "Bday":  tm("1990-12-20"),   1: {"Bday", Timestamp},    1: "1990-12-20",
  "Phone": "650-555-1212",     2: {"Phone", String},      2: "650-555-1212",
  "Sibs":  3,                  3: {"Sibs", Int64},        3: 3,
  "GPA" :  3.95,               4: {"GPA", Float64},       4: 3.95,
  "Friend":true,               5: {"Friend", Bool},       5: true,
}                            }                          }

Each instance of the A map later in the file is
preceeded by a small zebra.StructInstance
identifier (type: msgpack2 extension number 6) which
tells the decoder which of the StructT
to use to interpret the integer indexes.
```


The central idea of ZebraPack: start with msgpack2, but when encoding a struct (in msgpack2 a struct is represented as a map), replace the key strings with small integers. In maps, typically for JSON compatibility the keys are all strings.

By adding a small schema description (essentially a lookup table with int->string mappings and a type identifier) at the front of the serialization stream/file, we get known schema types up-front, plus compression and the ability to evolve our data without crashes. If you've ever had your msgpack crash your server because you tried to change the type of a field but keep the same name, you know how fragile the msgpack can be.

The second easy idea: use the Go language struct definition syntax as our serialization schema. Since https://github.com/tinylib/msgp already easily parses Go files, we can use this to our advantage. While we are focused on a serialization format for Go, other language's that can read msgpack2 could be easily modified to read the schema and translation file, since they themselves will be stored in msgpack2 format at the front of the file (extension code number 9).

# background

Starting point: [msgpack2](http://msgpack.org) is great.
It is has an easy to read spec, it defines a compact
serialization format, and it has wide language support from
both dynamic and compiled languages.

Nonetheless, data update
conflicts still happen and can be hard to
resolve. Encoders could use the guidance of a
schema to avoid signed versus unsigned integer
encodings.

For instance, sadly the widely emulated C-encoder
for msgpack chooses to encode signed positive integers
as unsigned integers. This causes crashes in readers
who were expected a signed integer, which they may
have originated themselves in the original struct.

That's right, msgpack2 lets data types changes as
they are read and re-serialized. Simple copying of
a serialized struct can change the types of its
numbers. This is horrible. Now we have to guess
whether an unsigned integer was really intended because
of the integer's range, or if data will be silently
truncated or lost when coercing a 64-bit integer to
a 63-bit signed integer--assuming such coercing ever
makes logical sense, which it may not.

This kind of tragedy happens because of a lack of
shared communication across time and space between
readers and writers. It is easily addressed with
a shared schema.

While not always necessary, a schema provides
many benefits, both for coordinating between
people and for machine performance.

* Stronger typing: readers know what is expected, in
both type and size of the data delivered. Writers
know what they should be writing.

* Performance and compression: replacing struct/map
field names with numbers provides immediate space
savings and compression.

* Conflict resolution: the Cap'nProto numbering and
update conflict resolution method is used here.
This method originated in the ProtocolBuffers
scheme, and was enhanced after experience in
Cap'nProto. How it works: Additions are always
made by incrementing by one the largest number available
prior to the addition. No gaps in numbering are
allowed, and no numbers are ever deleted.
To get the effect of deletion, use the `deprecated` tag
(effectively a tombstone) instead.
This allows the tools to help detect
merge conflicts as soon as possible. If
two people try to merge schemas where the same
struct or field number is re-used, a
schema compiler can automatically detect
this update conflict, and flag the human
to resolve the conflict before proceeding.

* All fields optional. Just as in msgpack2,
all fields are optional.

Design:

* Schema language: the schema language for
defining structs is identical to the Go
language. Go is expressive and yet easily parsed
by the standard library packages included
with Go itself. There are already
high-performance msgpack2 libraries available
for go, https://github.com/tinylib/msgp and
https://github.com/ugorji/go which
make schema compilation easy.

* Schema serialization: schemas are serialized
using the msgpack2 encoding. It is turtles all the
way down.

* Requirement: zerbapack requires that the msgpack2 standard
be adhered to. Strings and raw binary byte arrays
are distinct, and must be marked distinctly; msgpack1 encoding is
not allowed.

# benchmarking

Based on the implementation now available in https://github.com/glycerine/zebrapack, we measure read and write speed with the `-fast -fast-strings` optimizations on. Benchmarks adapted from https://github.com/glycerine/go_serialization_benchmarks of this struct:

```
type A struct {
	Name     string
	BirthDay time.Time
	Phone    string
	Siblings int
	GPA      float64
	Friend   bool
}
```

## read performance

`zebrapack -fast -fast-strings` comes in 2nd, behind only Gencode. This is very competitive. Note we are on par, or faster than, the highly tuned Gogoprotobufs library. As is typical for binary formats, ZebraPack is about 20x faster than Go's JSON handling.

```
benchmark                                       iter           time/iter         bytes alloc       allocs
---------                                       ----           ---------         -----------       ------
  BenchmarkGencodeUnmarshal-4              	10000000	       207 ns/op	     112 B/op	       3 allocs/op
* BenchmarkZebraPackUnmarshal-4            	10000000	       222 ns/op	       0 B/op	       0 allocs/op
  BenchmarkGogoprotobufUnmarshal-4         	 5000000	       253 ns/op	     112 B/op	       3 allocs/op
  BenchmarkCapNProtoUnmarshal-4            	 5000000	       267 ns/op	       0 B/op	       0 allocs/op
  BenchmarkFlatBuffersUnmarshal-4          	 5000000	       277 ns/op	     112 B/op	       3 allocs/op
  BenchmarkMsgpUnmarshal-4                 	 5000000	       354 ns/op	     112 B/op	       3 allocs/op
  BenchmarkGoprotobufUnmarshal-4           	 2000000	       655 ns/op	     432 B/op	       9 allocs/op
  BenchmarkProtobufUnmarshal-4             	 2000000	       734 ns/op	     192 B/op	      10 allocs/op
  BenchmarkGobUnmarshal-4                  	 1000000	      1066 ns/op	     112 B/op	       3 allocs/op
  BenchmarkHproseUnmarshal-4               	 1000000	      1066 ns/op	     320 B/op	      10 allocs/op
  BenchmarkCapNProto2Unmarshal-4           	 1000000	      1399 ns/op	     608 B/op	      12 allocs/op
  BenchmarkBinaryUnmarshal-4               	 1000000	      1562 ns/op	     336 B/op	      22 allocs/op
  BenchmarkXdrUnmarshal-4                  	 1000000	      1571 ns/op	     239 B/op	      11 allocs/op
  BenchmarkVmihailencoMsgpackUnmarshal-4   	 1000000	      2084 ns/op	     384 B/op	      13 allocs/op
  BenchmarkUgorjiCodecMsgpackUnmarshal-4   	  500000	      2593 ns/op	    3008 B/op	       6 allocs/op
  BenchmarkUgorjiCodecBincUnmarshal-4      	  500000	      2769 ns/op	    3168 B/op	       9 allocs/op
  BenchmarkSerealUnmarshal-4               	  300000	      3426 ns/op	    1008 B/op	      34 allocs/op
  BenchmarkJsonUnmarshal-4                 	  300000	      4224 ns/op	     495 B/op	       8 allocs/op
```

## write performance

`zebrapack -fast -fast-strings` ties for first place in write speed with Gogoprotobufs. This is mostly due to the use of the highly tuned https://github.com/tinylib/msgp library (in 4th place here), which is then sped up further by using integer keys instead of strings.

```
benchmark                                       iter           time/iter          bytes alloc      allocs
---------                                       ----           ---------          -----------      ------
  BenchmarkGogoprotobufMarshal-4           	10000000	       151 ns/op	      64 B/op	       1 allocs/op
* BenchmarkZebraPackMarshal-4              	10000000	       151 ns/op	      80 B/op	       1 allocs/op
  BenchmarkGencodeMarshal-4                	10000000	       168 ns/op	      80 B/op	       2 allocs/op
  BenchmarkMsgpMarshal-4                   	10000000	       191 ns/op	     128 B/op	       1 allocs/op
  BenchmarkFlatbuffersMarshal-4            	 5000000	       352 ns/op	       0 B/op	       0 allocs/op
  BenchmarkCapNProtoMarshal-4              	 3000000	       510 ns/op	      56 B/op	       2 allocs/op
  BenchmarkGoprotobufMarshal-4             	 2000000	       535 ns/op	     312 B/op	       4 allocs/op
  BenchmarkGobMarshal-4                    	 2000000	       918 ns/op	      48 B/op	       2 allocs/op
  BenchmarkProtobufMarshal-4               	 2000000	       920 ns/op	     200 B/op	       7 allocs/op
  BenchmarkCapNProto2Marshal-4             	 1000000	      1008 ns/op	     436 B/op	       7 allocs/op
  BenchmarkHproseMarshal-4                 	 1000000	      1013 ns/op	     476 B/op	       8 allocs/op
  BenchmarkBinaryMarshal-4                 	 1000000	      1370 ns/op	     256 B/op	      16 allocs/op
  BenchmarkXdrMarshal-4                    	 1000000	      1810 ns/op	     456 B/op	      21 allocs/op
  BenchmarkVmihailencoMsgpackMarshal-4     	 1000000	      1997 ns/op	     368 B/op	       6 allocs/op
  BenchmarkJsonMarshal-4                   	 1000000	      2328 ns/op	     536 B/op	       6 allocs/op
  BenchmarkUgorjiCodecBincMarshal-4        	  500000	      2523 ns/op	    2785 B/op	       8 allocs/op
  BenchmarkUgorjiCodecMsgpackMarshal-4     	  500000	      2578 ns/op	    2752 B/op	       8 allocs/op
  BenchmarkSerealMarshal-4                 	  500000	      3262 ns/op	     912 B/op	      21 allocs/op
```

deprecating fields
------------------

to actually deprecate a field, you start by adding the deprecated:"true" tag:
```
type A struct {
  Name     string      `zid:"0"`
  Bday     time.Time   `zid:"1"`
  Phone    string      `zid:"2"`
  Sibs     int         `zid:"3"`
  GPA      float64     `zid:"4" deprecated:"true"` // a deprecated field.
  Friend   bool        `zid:"5"`
}
```
*In addition,* you'll want to change the type of the deprecated field, substituting `struct{}` for the old type. By converting the type of the deprecated field to struct{}, it will no longer takes up any space in the Go struct. This saves space. Even if a struct evolves heavily in time (rare), the changes will cause no extra overhead in terms of memory. It also allows the compiler to detect and reject any new writes to the field that are using the old type. 
```
// best practice for deprecation of fields, to save space + get compiler support for deprectation
type A struct {
  Name     string      `zid:"0"`
  Bday     time.Time   `zid:"1"`
  Phone    string      `zid:"2"`
  Sibs     int         `zid:"3"`
  GPA      struct{}    `zid:"4" deprecated:"true"` // a deprecated field should have its type changed to struct{}, as well as being marked deprecated:"true".
  Friend   bool        `zid:"5"`
}
```

Rules for safe data changes: To preserve forwards/backwards compatible changes, you must *never remove a field* from a struct, once that field has been defined and used. In the example above, the `zid:"4"` tag must stay in place, to prevent someone else from ever using 4 again. This allows sane data forward evolution, without tears, fears, or crashing of servers. The fact that `struct{}` fields take up no space also means that there is no need to worry about loss of performance when deprecating. We retain all fields ever used for their zebra ids, and the compiled Go code wastes no extra space for the deprecated fields.

NB: There is one exception to this `struct{}` consumes no space rule: if the newly deprecated `struct{}` field happens to be *the very last field* in a struct, it will take up one pointer worth of space. If you want to deprecate the last field in a struct, if possible you should move it up in the field order (e.g. make it the first field in the Go struct), so it doesn't still consume space; reference https://github.com/golang/go/issues/17450.

schema
------

what does a schema look like? See  https://github.com/glycerine/zebrapack/blob/master/testdata/my.go and  https://github.com/glycerine/zebrapack/blob/master/testdata/my.z.json for example:

First here is (a shortened version of) the go file that we parsed:
~~~
package main

import (
	"time"
)

const zebraSchemaId64 = 0x6eb25cc0f9a3e

func main() {}

type A struct {
	Name   string    `zid:"0"`
	Bday   time.Time `zid:"1"`
	Phone  string    `zid:"2"`
	Sibs   int       `zid:"3"`
	GPA    float64   `zid:"4"`
	Friend bool      `zid:"5"`
}

~~~

Second, here is the (json version) of the zebrapack schema (stored canonically in msgpack2) that corresponds:
~~~
{
    "SourcePath": "testdata/my.go",
    "SourcePackage": "main",
    "ZebraSchemaId": 1947397430155838,
    "Structs": [
        {
            "StructName": "A",
            "Fields": [
                {
                    "Zid": 0,
                    "FieldGoName": "Name",
                    "FieldTagName": "Name",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    }
                },
                {
                    "Zid": 1,
                    "FieldGoName": "Bday",
                    "FieldTagName": "Bday",
                    "FieldTypeStr": "time.Time",
                    "FieldCategory": 23,
                    "FieldPrimitive": 20,
                    "FieldFullType": {
                        "Kind": 20,
                        "Str": "Time"
                    }
                },
                {
                    "Zid": 2,
                    "FieldGoName": "Phone",
                    "FieldTagName": "Phone",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    }
                },
                {
                    "Zid": 3,
                    "FieldGoName": "Sibs",
                    "FieldTagName": "Sibs",
                    "FieldTypeStr": "int",
                    "FieldCategory": 23,
                    "FieldPrimitive": 13,
                    "FieldFullType": {
                        "Kind": 13,
                        "Str": "int"
                    }
                },
                {
                    "Zid": 4,
                    "FieldGoName": "GPA",
                    "FieldTagName": "GPA",
                    "FieldTypeStr": "float64",
                    "FieldCategory": 23,
                    "FieldPrimitive": 4,
                    "FieldFullType": {
                        "Kind": 4,
                        "Str": "float64"
                    }
                },
                {
                    "Zid": 5,
                    "FieldGoName": "Friend",
                    "FieldTagName": "Friend",
                    "FieldTypeStr": "bool",
                    "FieldCategory": 23,
                    "FieldPrimitive": 18,
                    "FieldFullType": {
                        "Kind": 18,
                        "Str": "bool"
                    }
                }
            ]
        }
    ]
}
~~~


notices
-------

Copyright (c) 2016, Jason E. Aten, Ph.D.

LICENSE: MIT. See https://github.com/glycerine/zebrapack/blob/master/LICENSE




# from the original https://github.com/tinylib/msgp README

MessagePack Code Generator [![Build Status](https://travis-ci.org/tinylib/msgp.svg?branch=master)](https://travis-ci.org/tinylib/msgp)
=======

This is a code generation tool and serialization library for [MessagePack](http://msgpack.org). You can read more about MessagePack [in the wiki](http://github.com/tinylib/msgp/wiki), or at [msgpack.org](http://msgpack.org).

### Why?

- Use Go as your schema language
- Performance
- [JSON interop](http://godoc.org/github.com/tinylib/msgp/msgp#CopyToJSON)
- [User-defined extensions](http://github.com/tinylib/msgp/wiki/Using-Extensions)
- Type safety
- Encoding flexibility

### Quickstart

In a source file, include the following directive:

```go
//go:generate zebrapack
```

The `zebrapack` command will generate serialization methods for all exported type declarations in the file.

You can [read more about the code generation options here](http://github.com/tinylib/msgp/wiki/Using-the-Code-Generator).

### Use

Field names can be set in much the same way as the `encoding/json` package. For example:

```go
type Person struct {
	Name       string `msg:"name"`
	Address    string `msg:"address"`
	Age        int    `msg:"age"`
	Hidden     string `msg:"-"` // this field is ignored
	unexported bool             // this field is also ignored
}
```

By default, the code generator will satisfy `msgp.Sizer`, `msgp.Encodable`, `msgp.Decodable`, 
`msgp.Marshaler`, and `msgp.Unmarshaler`. Carefully-designed applications can use these methods to do
marshalling/unmarshalling with zero heap allocations.

While `msgp.Marshaler` and `msgp.Unmarshaler` are quite similar to the standard library's
`json.Marshaler` and `json.Unmarshaler`, `msgp.Encodable` and `msgp.Decodable` are useful for 
stream serialization. (`*msgp.Writer` and `*msgp.Reader` are essentially protocol-aware versions
of `*bufio.Writer` and `*bufio.Reader`, respectively.)

### Features

 - Extremely fast generated code
 - Test and benchmark generation
 - JSON interoperability (see `msgp.CopyToJSON() and msgp.UnmarshalAsJSON()`)
 - Support for complex type declarations
 - Native support for Go's `time.Time`, `complex64`, and `complex128` types 
 - Generation of both `[]byte`-oriented and `io.Reader/io.Writer`-oriented methods
 - Support for arbitrary type system extensions
 - [Preprocessor directives](http://github.com/tinylib/msgp/wiki/Preprocessor-Directives)
 - File-based dependency model means fast codegen regardless of source tree size.

Consider the following:
```go
const Eight = 8
type MyInt int
type Data []byte

type Struct struct {
	Which  map[string]*MyInt `msg:"which"`
	Other  Data              `msg:"other"`
	Nums   [Eight]float64    `msg:"nums"`
}
```
As long as the declarations of `MyInt` and `Data` are in the same file as `Struct`, the parser will determine that the type information for `MyInt` and `Data` can be passed into the definition of `Struct` before its methods are generated.

#### Extensions

MessagePack supports defining your own types through "extensions," which are just a tuple of
the data "type" (`int8`) and the raw binary. You [can see a worked example in the wiki.](http://github.com/tinylib/msgp/wiki/Using-Extensions)

### Status

Mostly stable, in that no breaking changes have been made to the `/msgp` library in more than a year. Newer versions
of the code may generate different code than older versions for performance reasons. I (@philhofer) am aware of a
number of stability-critical commercial applications that use this code with good results. But, caveat emptor.

You can read more about how `msgp` maps MessagePack types onto Go types [in the wiki](http://github.com/tinylib/msgp/wiki).

Here some of the known limitations/restrictions:

- Identifiers from outside the processed source file are assumed (optimistically) to satisfy the generator's interfaces. If this isn't the case, your code will fail to compile.
- Like most serializers, `chan` and `func` fields are ignored, as well as non-exported fields.
- Encoding of `interface{}` is limited to built-ins or types that have explicit encoding methods.


If the output compiles, then there's a pretty good chance things are fine. (Plus, we generate tests for you.) *Please, please, please* file an issue if you think the generator is writing broken code.

### Performance

If you like benchmarks, see [here](http://bravenewgeek.com/so-you-wanna-go-fast/) and [here](https://github.com/glycerine/zebra); [here for the benchmark source code](https://github.com/glycerine/go_serialization_benchmarks).

As one might expect, the generated methods that deal with `[]byte` are faster for small objects, but the `io.Reader/Writer` methods are generally more memory-efficient (and, at some point, faster) for large (> 2KB) objects.
