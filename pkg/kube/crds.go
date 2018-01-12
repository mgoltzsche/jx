package kube

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"github.com/jenkins-x/jx/pkg/apis/jx"
)

// RegisterEnvironmentCRD ensures that the CRD is registered for Environments
func RegisterEnvironmentCRD(apiClient *apiextensionsclientset.Clientset) error {
	name := "environments." + jx.GroupName
	names := &v1beta1.CustomResourceDefinitionNames{
		Kind:       "Environment",
		ListKind:   "EnvironmentList",
		Plural:     "environments",
		Singular:   "environment",
		ShortNames: []string{"env"},
	}

	return registerCRD(apiClient, name, names)
}

func registerCRD(apiClient  *apiextensionsclientset.Clientset, name string, names *v1beta1.CustomResourceDefinitionNames) error {
	_, err := apiClient.ApiextensionsV1beta1().CustomResourceDefinitions().Get(name, metav1.GetOptions{})
	if err == nil {
		return nil
	}

	crd := &v1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: v1beta1.CustomResourceDefinitionSpec{
			Group:   jx.GroupName,
			Version: jx.Version,
			Scope:   v1beta1.NamespaceScoped,
			Names:   *names,
		},
	}
	_, err = apiClient.ApiextensionsV1beta1().CustomResourceDefinitions().Create(crd)
	return err
}
