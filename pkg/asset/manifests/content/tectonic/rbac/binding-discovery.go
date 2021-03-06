package rbac

const (
	// BindingDiscovery  is the variable/constant representing the contents of the respective file
	BindingDiscovery = `
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1beta1
metadata:
  name: discovery
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: system:discovery
subjects:
- kind: Group
  name: 'system:unauthenticated'
- kind: Group
  name: 'system:authenticated'
`
)
