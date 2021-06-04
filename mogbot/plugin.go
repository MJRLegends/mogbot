package mogbot

type Plugin interface {
	RegisterPlugin(Bot)
	PluginInfo() PluginInfo
}

type PluginInfo struct {
	Name string
}
