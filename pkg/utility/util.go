package utility

import (
	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	v1 "k8s.io/api/core/v1"
	"os"
	"strings"
)

var authOption remote.Option
var NewRegistry string

const (
	registryUser     = "REGISTRY_USER"
	registryPassword = "REGISTRY_PASSWORD"
	registry         = "REGISTRY"
)

func init() {
	rgs := os.Getenv(registry)
	username := os.Getenv(registryUser)
	password := os.Getenv(registryPassword)
	if rgs != "" && username != "" && password != "" {
		authOption = remote.WithAuth(&authn.Basic{
			Username: username,
			Password: password,
		})
		NewRegistry = rgs
	} else {
		authOption = remote.WithAuthFromKeychain(authn.DefaultKeychain)
	}
}


// copy repo & tag into registry
func CacheImage(src string, dst string) error {
	ref, err := name.ParseReference(src)
	if err != nil {
		panic(err)
	}

	//Fetching src image reference
	img, err := remote.Image(ref, remote.WithAuthFromKeychain(authn.DefaultKeychain))
	if err != nil {
		return err
	}
	ref, err = name.ParseReference(dst)
	if err != nil {
		return err
	}

	//Copying image to backup registry
	err = remote.Write(ref, img, authOption)
	if err != nil {
		return err
	}
	return nil
}

var CacheFunc = CacheImage

// update image specifications
func ModifyImage(containers *[]v1.Container) error {
	for i := range *containers {
		c := *containers
		temp := strings.Split(c[i].Image, "/")
		oldRegistry := temp[0]

		if oldRegistry != strings.Split(NewRegistry,"/")[0] {
			newImage := NewRegistry + "/" + temp[len(temp)-1]
			err := CacheFunc(c[i].Image, newImage)
			if err != nil {
				return err
			}
			c[i].Image = newImage

		}
	}
	return nil
}
