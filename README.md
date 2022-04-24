# NekosGo

incredibly barebones and simple wrapper around the Nekos.life v3 api

## Usage

```go
import (
	"fmt"

	"github.com/RinLovesYou/NekosGo"
)

func main() {
	res, err := NekosGo.Image(NekosGo.AnalGif)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
```
