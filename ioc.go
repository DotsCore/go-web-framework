package foundation

import (
	"github.com/RobyFerro/dig"
	"github.com/RobyFerro/go-web-framework/kernel"
	"github.com/gorilla/sessions"
	"log"
)

// RetrieveSingletonContainer returns a IOC container that contains every IOC singleton services.
func RetrieveSingletonContainer() *dig.Container {
	return kernel.SingletonIOC
}

// RetrieveConfig provides a shortcut to retrieve the global configuration.
func RetrieveConfig() *kernel.ServerConf {
	var config *kernel.ServerConf
	if err := RetrieveSingletonContainer().Invoke(func(c *kernel.ServerConf) {
		config = c
	}); err != nil {
		log.Fatal(err)
	}

	return config
}

// RetrieveCookieStore provides a shortcut to retrieve the CookieStore object.
func RetrieveCookieStore() *sessions.CookieStore {
	var store *sessions.CookieStore
	if err := RetrieveSingletonContainer().Invoke(func(c *sessions.CookieStore) {
		store = c
	}); err != nil {
		log.Fatal(err)
	}

	return store
}
