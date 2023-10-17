package main

import (
	"fmt"
	"log"
	"os"

	"github.com/open-component-model/ocm/pkg/contexts/ocm"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/compdesc"
	metav1 "github.com/open-component-model/ocm/pkg/contexts/ocm/compdesc/meta/v1"
	"github.com/open-component-model/ocm/pkg/contexts/ocm/repositories/ocireg"
)

const componentName = "github.com/skarlso/podinfo-helm-4"
const componentVersion = "v1.0.0"

const resourceName = "charts"

func main() {
	octx := ocm.DefaultContext()

	repoSpec := ocireg.NewRepositorySpec("ghcr.io/skarlso/helm-test")

	repo, err := octx.RepositoryForSpec(repoSpec)
	if err != nil {
		log.Fatal(err)
	}
	defer repo.Close()

	compvers, err := repo.LookupComponentVersion(componentName, componentVersion)
	if err != nil {
		log.Fatal(err)
	}
	defer compvers.Close()

	cd := compvers.GetDescriptor()
	data, err := compdesc.Encode(cd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("component descriptor:\n%s\n", string(data))

	res, err := compvers.GetResource(metav1.NewIdentity(resourceName))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("resource %s:\n  type: %s\n", resourceName, res.Meta().Type)

	meth, err := res.AccessMethod()
	if err != nil {
		log.Fatal(err)
	}
	defer meth.Close()

	fmt.Printf("  mime: %s\n", meth.MimeType())

	data, err = meth.Get()
	if err != nil {
		log.Fatal(err)
	}

	_ = os.WriteFile("charts.tgz", data, 0o755)
}
