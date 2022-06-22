module tcp

go 1.18

require (
	github/Gavinyjb/Gopkg/proto v0.0.1
	github/Gavinyjb/Gopkg/server v0.0.1
	github/Gavinyjb/Gopkg/client v0.0.1
)
replace (
	github/Gavinyjb/Gopkg/proto => ./proto
	github/Gavinyjb/Gopkg/server => ./server
	github/Gavinyjb/Gopkg/client => ./client
)
