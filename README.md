# golang-getfrompass

This is a small helper method to get a password from the pass passwordstore (https://www.passwordstore.org).

## Usage

This package contains one method (`GetFromPass(key string) (string, error)`) that returns the password from pass.
It calls the pass executable and extracts the first line (removing newline characters but not spaces).

Example:
```golang
import (
  "fmt"
  "log"
  "github.com/jduepmeier/golang-getfrompass"
)

func main() {
  pass, err := getfrompass.GetFromPass("test-password")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("Password: %s\n", pass)
}
```
