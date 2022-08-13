package iterators

import (
	"github.com/hellogo/pkg/generic/functions"
)

type Iterable[E any] interface {
	Iterator() Iterator[E]
	ForEach(action functions.Consumer[E])
}
