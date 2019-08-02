module misc/go-test

go 1.12

require (
	github.com/golang/mock v1.3.1
	github.com/kr/pretty v0.1.0 // indirect
	github.com/stretchr/testify v1.3.0
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190418165655-df01cb2cc480
	golang.org/x/net => github.com/golang/net v0.0.0-20190420063019-afa5a82059c6
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190626221950-04f50cda93cb
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190731214159-1e85ed8060aa
)
