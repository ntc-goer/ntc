With Go's module support, go [build|run|test] automatically fetches the necessary dependencies when you add the import in your code:

import "github.com/ntc-goer/ntc"
Alternatively, use go get:

go get -u github.com/ntc-goer/ntc@latest
