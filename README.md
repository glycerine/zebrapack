ZebraPack: a data description language and serialization format. This one is all black and white. No gray areas.
==========

ZebraPack is a data definition language and serialization format. It removes gray areas from msgpack2 serialized data, and provides for declared schemas, sane data evolution, and more compact encoding. It does all this while maintaining the possibility of easy compatibility with all the dynamic languages that already have msgpack2 support.

The code here descends from the fantastic msgpack code generator https://github.com/tinylib/msgp by Philip Hofer.

At this time it is feature complete, and in beta while we get experience with it. See https://github.com/glycerine/zebra for the spec/ideas that we are implementing here.

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


# from the tinylib/msgp README

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
//go:generate msgp
```

The `msgp` command will generate serialization methods for all exported type declarations in the file.

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
