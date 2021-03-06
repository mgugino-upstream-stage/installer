package updater

const (
	// MigrationStatusKind  is the variable/constant representing the contents of the respective file
	MigrationStatusKind = `
apiVersion: "apiextensions.k8s.io/v1beta1"
kind: "CustomResourceDefinition"
metadata:
  name: "migrationstatuses.kvo.coreos.com"
spec:
  group: "kvo.coreos.com"
  version: "v1"
  names:
    plural: "migrationstatuses"
    kind: "MigrationStatus"
    `
)
