// This go.mod file is generated by ./hack/codegen.sh.
module go.pinniped.dev/generated/1.26/client

go 1.13

require (
	go.pinniped.dev/generated/1.26/apis v0.0.0
	k8s.io/apimachinery v0.26.5
	k8s.io/client-go v0.26.5
	k8s.io/kube-openapi v0.0.0-20221012153701-172d655c2280
)

replace go.pinniped.dev/generated/1.26/apis => ../apis