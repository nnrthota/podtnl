module github.com/narendranathreddythota/podtnl

go 1.14

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190819141724-e14f31a72a77

require (
	github.com/go-log/log v0.2.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/kataras/golog v0.0.15
	github.com/stretchr/testify v1.4.0
	golang.org/x/crypto v0.0.0-20200510223506-06a226fb4e37 // indirect
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/time v0.0.0-20200416051211-89c76fbcd5d1 // indirect
	k8s.io/api v0.17.2
	k8s.io/apiextensions-apiserver v0.17.2
	k8s.io/apimachinery v0.17.2
	k8s.io/client-go v11.0.0+incompatible
)
