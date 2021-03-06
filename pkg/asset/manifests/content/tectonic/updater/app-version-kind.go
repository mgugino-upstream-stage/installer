package updater

const (
	// AppVersionKind  is the variable/constant representing the contents of the respective file
	AppVersionKind = `
apiVersion: "apiextensions.k8s.io/v1beta1"
kind: "CustomResourceDefinition"
metadata:
  name: "appversions.tco.coreos.com"
spec:
  group: "tco.coreos.com"
  version: "v1"
  names:
    plural: "appversions"
    kind: "AppVersion"
    `
)
