module github.com/ktye/i

go 1.12

require (
	github.com/eaburns/T v0.0.0-20190217122806-dbc7887ff15c
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/ktye/iv v0.0.0-20190506155233-ea1af1e66e83
	github.com/ktye/plot v0.0.0
	github.com/ktye/ui v1.0.0 // this is wrong: should be v0.0.0
	github.com/mattn/go-sixel v0.0.0-20190320171103-a8fac8fa7d81
	golang.org/x/exp v0.0.0-20190417140011-e40e924fdd3f
	golang.org/x/image v0.0.0-20190417020941-4e30a6eb7d9a
	golang.org/x/mobile v0.0.0-20190415191353-3e0bab5405d6
)

replace github.com/ktye/ui => ../ui

replace github.com/ktye/plot => ../plot
