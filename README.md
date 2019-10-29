# Usage

```go
// +build darwin

package main

import (
	macos "github.com/guildencrantz/logrus-macos-hook"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.AddHook(macos.New())
}
```

# License

Use of this source code is governed by a BSD-2-Clause license that can be found
in the LICENSE file.
