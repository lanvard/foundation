package foundation

import (
	"lanvard/config"
)

type Application struct {
	// The service container
	Container Container

	// Indicates if the application has been bootstrapped before.
	HasBeenBootstrapped bool
}

func (a Application) Make(abstract interface{}) interface{} {
	return a.Container.Make(abstract)
}

// Bind all of the application paths in the container.
func (a *Application) BindPathsInContainer() {
	a.Container.Instance("path.app", config.App.BasePath.AppPath())
	a.Container.Instance("path.base", config.App.BasePath.BasePath())
	a.Container.Instance("path.lang", config.App.BasePath.LangPath())
	a.Container.Instance("path.config", config.App.BasePath.ConfigPath())
	a.Container.Instance("path.public", config.App.BasePath.PublicPath())
	a.Container.Instance("path.storage", config.App.BasePath.StoragePath())
	a.Container.Instance("path.database", config.App.BasePath.DatabasePath())
	a.Container.Instance("path.resources", config.App.BasePath.ResourcePath())
	a.Container.Instance("path.bootstrap", config.App.BasePath.BootstrapPath())
}
