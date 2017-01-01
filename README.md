ZebraPack: a data description language and serialization format. Like Gobs version 2.0.
==========

ZebraPack is a data definition language and serialization format. It removes gray areas from msgpack2 serialized data, and provides for declared schemas, sane data evolution, and more compact encoding.

It does all this while maintaining the possibility of easy compatibility with all the dynamic languages that already have msgpack2 support. If your favorite language (after Go, of course) has a library that reads msgpack2, then it would be only a day's work to adapt the library to read zebrapack: the schema are in msgpack2, and then one simply keeps a hashmap to translate between small integer <-> field names/type.

Why start with [msgpack2](http://msgpack.org)?  Quite simple: msgpack2 is simple, fast, and extremely portable. It has an implementation in every language you've heard of, and some you haven't (some 50 libraries are available). It has a well defined and short spec. The mspack1 vs msgpack2 terminology is a distinction we make here: msgpack1 spec poorly distringuished between strings and raw binary bytes, but that was remedied in msgpack2. Importantly, msgpack2 is dynamic-language friendly because it is largely self-describing.

We find only two problems with msgpack2: weak support for data evolution, and insufficiently strong typing.

The ZebraPack format addresses these problems. Moreover, ZebraPack is actually still binary compatible with msgpack2 spec. It just adopts a new convention about how to encode the field names of structs. Structs are encoded in msgpack2 using maps, as usual. Hence all data is still encoded precisely in the msgpack2 format. The only difference in ZebraPack is this convention: maps that represent structs are now keyed by integers. Rather than have string keys -- the convention for most msgpack2 language bindings -- in ZebraPack we use integers as keys for those maps that are representing structs. These integers are associated with a field name and type in a (separable) schema. The schema is also defined and encoded in msgpack2.

The resulting binary encoding is very similar in style to protobufs/Thrift/Capn'Proto. However it is much more friendly to other (dynamic) languages. Also it is screaming fast (see benchmarks below).

Once we have a schema, we can be very strongly typed, and be very efficient. We borrow the idea of field deprecation from FlatBuffers. For conflicting update detection, we use CapnProto's field numbering discipline. We add support for the `omitempty` tag. In fact, in ZebraPack, all fields are `omitempty`. If they are empty they won't be serialized on the wire. Like FlatBuffers and Protobufs, this enables one to define a very large schema of possibilities, and then only transmit a very small (efficient) portion that is currently relevant over the wire.

Full credit: the code here descends from the fantastic msgpack2 code generator https://github.com/tinylib/msgp by Philip Hofer.

By default we generate ZebraPack format encoders and decoders when the `zebrapack` tool is run. Note that we continue to offer straight msgpack2 serialization and deserialization with the `-msgp` flag to `zebrapack`.

At this time it is feature complete and production ready.


# background and motivation

# ZebraPack serialization. This one is all black and white. No gray areas.

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
print the schema information at the front of the file --
or detached completely and kept in a separate file --
in the form of a zebra.Schema (see https://github.com/glycerine/zebrapack/blob/master/zebra/zebra.go for the spec) structure. Then
the data follows as a map whose keys are now integers
instead of strings. A simple example:
         
original(msgpack2) ->        schema(msgpack2)      +    each instance(msgpack2)
--------                     --------------             -------------
a := A{                      zebra.StructT{             map{
  "Name":  "Atlanta",          0: {"Name", String},       0: "Atlanta",
  "Bday":  tm("1990-12-20"),   1: {"Bday", Timestamp},    1: "1990-12-20",
  "Phone": "650-555-1212",     2: {"Phone", String},      2: "650-555-1212",
  "Sibs":  3,                  3: {"Sibs", Int64},        3: 3,
  "GPA" :  3.95,               4: {"GPA", Float64},       4: 3.95,
  "Friend":true,               5: {"Friend", Bool},       5: true,
}                            }                          }

```


The central idea of ZebraPack: start with msgpack2, but when encoding a struct (in msgpack2 a struct is represented as a map), replace the key strings with small integers.

By having a small schema description (essentially a lookup table with int->string mappings and a type identifier) either separate or serialized at the front of the serialization stream/file, we get known schema types up-front, plus compression and the ability to evolve our data without crashes. If you've ever had your msgpack crash your server because you tried to change the type of a field but keep the same name, then you know how fragile msgpack can be.

By default, today, we serialize the schema to a separate file so that the wire encoding is as fast as possible. However it is trivial to add/pre-pend the encoded schema to any file when you need to. The `zebrapack` generated Go code incorporates knowledge of the schema, so if you are only working in Go there is no need to `zebrapack -write-schema` to generate an external schema desription file. In summary, by default we behave like protobufs/thrift/capnproto, but dynamic languages and runtime type discovery can be supported in full fidelity.

The second easy idea: use the Go language struct definition syntax as our serialization schema. Why invent another format? Serialization for Go developers should be almost trivially easy. While we are focused on a serialization format for Go, because other language can read msgpack2, they can also readily parse the schema. The schema is stored in msgpack2 struct convention (and optionally json), rather than the ZebraPack struct convention, for bootstrapping.

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

Astonishing, but true: the existing practice for msgpack2
language bindings allows the data types to change as
they are read and re-serialized. Simple copying of
a serialized struct can change the types of data
from signed to unsigned. This is horrible. Now we have to guess
whether an unsigned integer was really intended because
of the integer's range, or if data will be silently
truncated or lost when coercing a 64-bit integer to
a 63-bit signed integer--assuming such coercing ever
makes logical sense, which it may not.

This kind of tragedy happens because of a lack of
shared communication across time and space between
readers and writers. It is easily addressed with
a shared schema. ZebraPack, in its essense, is the
agreement to follow that schema when binding
msgpack2 to a new language.

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
using the msgpack2 encoding.

* Requirement: zerbapack requires that the msgpack2 standard
be adhered to. Strings and raw binary byte arrays
are distinct, and must be marked distinctly; msgpack1 encoding is
not allowed.

* All language bindings must respect the declared type in
the ZebraPack schema when writing data. For example,
this means that signed and unsigned declarations
must be respected.



# benchmarking

Based on the implementation now available in https://github.com/glycerine/zebrapack, we measure read and write speed with the `-fast-strings` optimizations on. Benchmarks from https://github.com/glycerine/go_serialization_benchmarks of this struct:

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

`zebrapack -fast-strings -no-rtti` jockeys for the top position with go-capnproto-version-1, Gencode, FlatBuffers, and gogoprotobuf. In the sampling below it comes out fastest, but this varies occassionally run-by-run. Nonetheless, we see a very strong showing amongst strong company. Moreover, our zero allocation profile and serialization directly to and from Go structs are advantages. As is typical for binary formats, ZebraPack is about 20x faster than Go's JSON handling.

```
benchmark                                       iter           time/iter         bytes alloc       allocs
---------                                       ----           ---------         -----------       ------
BenchmarkZebraPackUnmarshal-4            	10000000	       227 ns/op	       0 B/op	       0 allocs/op
BenchmarkGencodeUnmarshal-4              	10000000	       229 ns/op	     112 B/op	       3 allocs/op
BenchmarkFlatBuffersUnmarshal-4          	10000000	       232 ns/op	      32 B/op	       2 allocs/op
BenchmarkGogoprotobufUnmarshal-4         	10000000	       232 ns/op	      96 B/op	       3 allocs/op
BenchmarkCapNProtoUnmarshal-4            	10000000	       258 ns/op	       0 B/op	       0 allocs/op
BenchmarkMsgpUnmarshal-4                 	 5000000	       296 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoprotobufUnmarshal-4           	 2000000	       688 ns/op	     432 B/op	       9 allocs/op
BenchmarkProtobufUnmarshal-4             	 2000000	       707 ns/op	     192 B/op	      10 allocs/op
BenchmarkGobUnmarshal-4                  	 2000000	       886 ns/op	     112 B/op	       3 allocs/op
BenchmarkHproseUnmarshal-4               	 1000000	      1045 ns/op	     320 B/op	      10 allocs/op
BenchmarkCapNProto2Unmarshal-4           	 1000000	      1359 ns/op	     608 B/op	      12 allocs/op
BenchmarkXdrUnmarshal-4                  	 1000000	      1659 ns/op	     239 B/op	      11 allocs/op
BenchmarkBinaryUnmarshal-4               	 1000000	      1907 ns/op	     336 B/op	      22 allocs/op
BenchmarkVmihailencoMsgpackUnmarshal-4   	 1000000	      2085 ns/op	     384 B/op	      13 allocs/op
BenchmarkUgorjiCodecMsgpackUnmarshal-4   	  500000	      2620 ns/op	    3008 B/op	       6 allocs/op
BenchmarkUgorjiCodecBincUnmarshal-4      	  500000	      2795 ns/op	    3168 B/op	       9 allocs/op
BenchmarkSerealUnmarshal-4               	  500000	      3271 ns/op	    1008 B/op	      34 allocs/op
BenchmarkJsonUnmarshal-4                 	  200000	      5576 ns/op	     495 B/op	       8 allocs/op
  ```

## write performance

`zebrapack -fast-strings -no-rtti` consistently dominates the field. This is mostly due to the use of the highly tuned https://github.com/tinylib/msgp library (in 3rd place here), which is then sped up further by using integer keys instead of strings.

```
benchmark                                       iter           time/iter          bytes alloc      allocs
---------                                       ----           ---------          -----------      ------
BenchmarkZebraPackMarshal-4              	10000000	       115 ns/op	       0 B/op	       0 allocs/op
BenchmarkGogoprotobufMarshal-4           	10000000	       148 ns/op	      64 B/op	       1 allocs/op
BenchmarkMsgpMarshal-4                   	10000000	       161 ns/op	     128 B/op	       1 allocs/op
BenchmarkGencodeMarshal-4                	10000000	       176 ns/op	      80 B/op	       2 allocs/op
BenchmarkFlatBufferMarshal-4             	 5000000	       347 ns/op	       0 B/op	       0 allocs/op
BenchmarkCapNProtoMarshal-4              	 3000000	       506 ns/op	      56 B/op	       2 allocs/op
BenchmarkGoprotobufMarshal-4             	 3000000	       617 ns/op	     312 B/op	       4 allocs/op
BenchmarkGobMarshal-4                    	 2000000	       887 ns/op	      48 B/op	       2 allocs/op
BenchmarkProtobufMarshal-4               	 2000000	       912 ns/op	     200 B/op	       7 allocs/op
BenchmarkHproseMarshal-4                 	 1000000	      1052 ns/op	     473 B/op	       8 allocs/op
BenchmarkCapNProto2Marshal-4             	 1000000	      1214 ns/op	     436 B/op	       7 allocs/op
BenchmarkBinaryMarshal-4                 	 1000000	      1427 ns/op	     256 B/op	      16 allocs/op
BenchmarkVmihailencoMsgpackMarshal-4     	 1000000	      1772 ns/op	     368 B/op	       6 allocs/op
BenchmarkXdrMarshal-4                    	 1000000	      1802 ns/op	     455 B/op	      20 allocs/op
BenchmarkJsonMarshal-4                   	 1000000	      2500 ns/op	     536 B/op	       6 allocs/op
BenchmarkUgorjiCodecBincMarshal-4        	  500000	      2514 ns/op	    2784 B/op	       8 allocs/op
BenchmarkSerealMarshal-4                 	  500000	      2729 ns/op	     912 B/op	      21 allocs/op
BenchmarkUgorjiCodecMsgpackMarshal-4     	  500000	      3274 ns/op	    2752 B/op	       8 allocs/op
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
// best practice for deprecation of fields, to save space + get compiler support for deprecation
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

First here is (a shortened version of) the go file that we parsed. The zebraSchemaId64 is a random number generated with a quick command line call to `zebrapack -genid`. Assigning a `zebraSchemaId64` in your Go source/schema can avoid format ambiguity.

~~~
package main

import (
	"time"
)

const zebraSchemaId64 = 0x6eb25cc0f9a3e

func main() {}

type A struct {
	Name   string    `zid:"0" msg:"name"` 
	Bday   time.Time `zid:"1"`
	Phone  string    `zid:"2" msg:"phone,omitempty"`
	Sibs   int       `zid:"3" msg:",omitempty"`
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
                    "FieldTagName": "name",
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
                    "FieldTagName": "phone",
                    "FieldTypeStr": "string",
                    "FieldCategory": 23,
                    "FieldPrimitive": 2,
                    "FieldFullType": {
                        "Kind": 2,
                        "Str": "string"
                    },
                    "OmitEmpty": true
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
                    },
                    "OmitEmpty": true
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

The official, machine-readable definition of the zebrapack format is given in this file: https://github.com/glycerine/zebrapack/blob/master/zebra/zebra.go

command line flags
------------------

~~~
  $ zebrapack -h

  Usage of zebrapack:

  -msgp
    	generate msgpack2 serializers instead of ZebraPack;
        for backward compatiblity or serializing the zebra
        schema itself.

  -fast-strings
    	for speed when reading a string in a message that won't be
     reused, this flag means we'll use unsafe to cast the string
     header and avoid allocation.

  -file go generate
    	input file (or directory); default is $GOFILE, which
     is set by the go generate command.

  -genid
    	generate a fresh random zebraSchemaId64 value to
     include in your Go source schema

  -io
    	create Encode and Decode methods (default true)

  -marshal
    	create Marshal and Unmarshal methods (default true)

  -no-embedded-schema
    	don't embed the schema in the generated files
        
  -o string
    	output file (default is {input_file}_gen.go

  -schema-to-go string
    	(standalone functionality) path to schema in msgpack2
     format; we will convert it to Go, write the Go on stdout,
     and exit immediately

  -tests
    	create tests and benchmarks (default true)

  -unexported
    	also process unexported types

  -write-schema string
		write schema to this file; - for stdout

~~~

## `zebrapack -msgp` as a msgpack2 code-generator

### `msg:",omitempty"` tags on struct fields

If you're using `zebrapack -msgp` to generate msgpack2 serialization code, then you can use the `omitempty` tag on your struct fields.

In the following example,
```
type Hedgehog struct {
   Furriness string `msg:",omitempty"`
}
```
If Furriness is the empty string, the field will not be serialized, thus saving the space of the field name on the wire.

It is safe to re-use structs even with `omitempty`. For reference:

from https://github.com/tinylib/msgp/issues/154:
> The only special feature of UnmarshalMsg and DecodeMsg (from a zero-alloc standpoint) is that they will use pre-existing fields in an object rather than allocating new ones. So, if you decode into the same object repeatedly, things like slices and maps won't be re-allocated on each decode; instead, they will be re-sized appropriately. In other words, mutable fields are simply mutated in-place.

This continues to hold true, and missing fields on the wire will zero the field in any re-used struct.

NB: Under tuple encoding (https://github.com/tinylib/msgp/wiki/Preprocessor-Directives), for example `//msgp:tuple Hedgehog`, then all fields are always serialized and the omitempty tag is ignored.

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

The `zebrapack` command will generate serialization methods for all exported type declarations in the file. If you add the flag `-msgp`, it will generate msgpack2 rather than ZebraPack format.

For other language's use, schemas can can be written to a separate file using `zebrapack -file my.go -write-schema` at the shell. (By default schemas are not written to the wire, just as in protobufs/CapnProto/Thrift.)

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

If you like benchmarks, see [here](http://bravenewgeek.com/so-you-wanna-go-fast/) and above in the ZebraPack benchmarks section; [see here for the benchmark source code](https://github.com/glycerine/go_serialization_benchmarks).

As one might expect, the generated methods that deal with `[]byte` are faster for small objects, but the `io.Reader/Writer` methods are generally more memory-efficient (and, at some point, faster) for large (> 2KB) objects.
