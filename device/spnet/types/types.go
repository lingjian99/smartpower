package types

type NetConfig struct {
	Network    string
	ListenOn   string
	Multicore  bool
	PriKeys    map[string]string
	MaxConnect int `json:",optional"`
}
