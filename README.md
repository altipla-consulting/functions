
# functions

Utilities and helpers for apps running in Google Cloud Functions.


## Install

```shell
go get github.com/altipla-consulting/functions
```


## Usage

Receive a Firestore Event:

```go
package foo

import (
  "context"
  "fmt"

  "github.com/altipla-consulting/functions"
)

func Entrypoint(ctx context.Context, event *functions.FirestoreEvent) error {
  fmt.Printf("%#v\n", event.OldValue)
  fmt.Printf("%#v\n", event.Value)

  return nil
}
```


## Contributing

You can make pull requests or create issues in GitHub. Any code you send should be formatted using `make gofmt`.


## Running tests

Run the tests:

```shell
make test
```


## License

[MIT License](LICENSE)
