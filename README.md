# merkle-stream

[![build status][1]][2]

stream which constructs a merkle tree based on the incoming data.

adapted from https://github.com/datrs/merkle-tree-stream.

## Installation

``` sh
go get github.com/kiambogo/merkle-stream
```

## Usage

``` go
import github.com/kiambogo/merkle-stream

hasher := merkle.BLAKE2b512{}
stream := merkle.NewStream(hasher, nil, nil)
stream.Append([]byte("hello, world!"))
stream.Append([]byte("foo"))
stream.Append([]byte("bar"))
log.Println(stream.Nodes())
```

## License
[MIT](./LICENSE)

[1]: https://github.com/kiambogo/merkle-stream/actions/workflows/test.yml/badge.svg
[2]: https://github.com/kiambogo/merkle-stream/actions/workflows/test.yml
