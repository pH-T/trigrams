# Golang Trigrams

Trigrams provides a way to compare strings (or bytes).

Given two strings/bytes A, B this lib generates trigrams for A and B: t1 und t2.
Then the unique and common trigrams are calculated. Similarity is given by common / unique.

## Example

```go
package main

import (
	"fmt"

	"github.com/pH-T/trigrams"
)

func main() {
	a := "test"
	b := "test123"

	t1 := trigrams.ToTrigrams([]byte(a))
	t2 := trigrams.ToTrigrams([]byte(b))

	fmt.Println(t1.JaccardCompare(t2, -1))

}
```

Result:
```bash
> go run main.go 
0.36363636363636365
```


## References & Sources

* https://medium.com/@appaloosastore/string-similarity-algorithms-compared-3f7b4d12f0ff
* https://metacpan.org/pod/distribution/String-Trigram/Trigram.pm
* https://sistemanalize.wordpress.com/2017/12/10/big-data-set-similarity-q-grams-overlap-measure-jaccard-index-jaccard-distance/
* https://hpi.de/fileadmin/user_upload/fachgebiete/naumann/folien/SS13/DPDC/DPDC_12_Similarity.pdf (24)
* http://cs.uef.fi/sipu/pub/TitleSimilarity-ICPR.pdf
- https://stackoverflow.com/a/34002499
