package cachectl

import (
	"fmt"
	"runtime"
)

func PrintVersion(name string) {
	fmt.Printf(`%s %s
Compiler: %s %s
Copyright (C) 2014-2015 Tatsuhiko Kubo <cubicdaiya@gmail.com>
`,
		name,
		Version,
		runtime.Compiler,
		runtime.Version())
}
