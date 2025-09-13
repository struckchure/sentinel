package sentinel

type Config struct {
	Host     string    `json:"host" yaml:"host"`
	Port     int       `json:"port" yaml:"port"`
	Backends []Backend `json:"backends" yaml:"backends"`
}

type Backend struct {
	Patterns     []Pattern             `json:"patterns" yaml:"patterns"`
	Methods      []Method              `json:"methods" yaml:"methods"`
	LoadBalancer LoadBalancerAlgorithm `json:"load_balancer" yaml:"load_balancer"`
	Services     []Service             `json:"services" yaml:"services"`
	Middlewares  []Middleware          `json:"middlewares" yaml:"middlewares"`
}

type Pattern struct {
	From string `json:"from" yaml:"from"`
	To   string `json:"to" yaml:"to"`
}

type Method string

const (
	MethodGet     Method = "GET"
	MethodHead    Method = "HEAD"
	MethodPost    Method = "POST"
	MethodPut     Method = "PUT"
	MethodPatch   Method = "PATCH"
	MethodDelete  Method = "DELETE"
	MethodConnect Method = "CONNECT"
	MethodOptions Method = "OPTIONS"
	MethodTrace   Method = "TRACE"
)

type LoadBalancerAlgorithm string

const (
	LoadBalancerAlgorithmRoundRobin LoadBalancerAlgorithm = "round-robin"
	LoadBalancerAlgorithmRandom     LoadBalancerAlgorithm = "random"
	LoadBalancerAlgorithmLeastConn  LoadBalancerAlgorithm = "least-connections"
	LoadBalancerAlgorithmIPHash     LoadBalancerAlgorithm = "ip-hash"
	LoadBalancerAlgorithmWeighted   LoadBalancerAlgorithm = "weighted"
)

type Service struct {
	Url    string `json:"url" yaml:"url"`
	Weight int    `json:"weight" yaml:"weight"`
}

type Middleware struct {
	Name   string         `json:"name" yaml:"name"`
	Config map[string]any `json:"config" yaml:"config"`
}
