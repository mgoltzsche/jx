package v1

import (
	v1 "github.com/jenkins-x/jx/pkg/apis/jx/v1"
	"github.com/jenkins-x/jx/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type JenkinsV1Interface interface {
	RESTClient() rest.Interface
	EnvironmentsGetter
}

// JenkinsV1Client is used to interact with features provided by the jenkins.io group.
type JenkinsV1Client struct {
	restClient rest.Interface
}

func (c *JenkinsV1Client) Environments(namespace string) EnvironmentInterface {
	return newEnvironments(c, namespace)
}

// NewForConfig creates a new JenkinsV1Client for the given config.
func NewForConfig(c *rest.Config) (*JenkinsV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &JenkinsV1Client{client}, nil
}

// NewForConfigOrDie creates a new JenkinsV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *JenkinsV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new JenkinsV1Client for the given RESTClient.
func New(c rest.Interface) *JenkinsV1Client {
	return &JenkinsV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *JenkinsV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
