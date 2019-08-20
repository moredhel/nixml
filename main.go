package main

import (
	"errors"
	"fmt"
	"os"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"text/template"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type NixmlVersion string

type PackageSet struct {
	Name string
	Modules []string
}

type PackageImportSet struct {
	Lang string `yaml:"lang"`
	Version string `yaml:"version,omitempty"`
	Modules []string `yaml:"modules"`
}

type ApplicationConfig struct {
	Nixml NixmlVersion `yaml:"nixml"`
	Name string `yaml:"name"`
	Version string `yaml:"version"`
	Snapshot string `yaml:"snapshot"`
	Packages []PackageImportSet `yaml:"packages"`
}

type ImportSet struct {
	Name string
	Version string
	Url string
	Sha256 string
	PackageSets map[string]PackageSet
}

type ApplicationUrl struct {
	Url string
	Sha256 string
}

func generateUrl(cfg *ApplicationConfig) (ApplicationUrl, error) {
	// TODO: make this more smart!
	return ApplicationUrl{
		Url: fmt.Sprintf("https://github.com/nixos/nixpkgs/archive/%s.tar.gz", cfg.Snapshot),
		// HACK: this needs to be dynamically generated
		Sha256: "0iwn8lrhdldgdgz5rg7k8h5wavxw5y73j453ji7z84z805falfwi",
	}, nil
}

type LangHandler interface {
	Handle(*PackageImportSet) (PackageSet, error);
}

type PythonHandler struct {
	versionMap map[string]string
}
func (Self PythonHandler) Handle(pkg *PackageImportSet) (PackageSet, error) {
	version, exists := Self.versionMap[pkg.Version]
	if !exists {
		return PackageSet{}, errors.New(fmt.Sprintf("unknown version '%s'", pkg.Version))
	}

	prefix := fmt.Sprintf("python%sPackages", version)

	modules := []string{}

	for _, module := range pkg.Modules {
		newModule := fmt.Sprintf("%s.%s", prefix, module)
		modules = append(modules, newModule)
	}

	return PackageSet{
		Name: pkg.Lang,
		Modules: modules,
	}, nil
}

type GoHandler struct {}
func (GoHandler) Handle(pkg *PackageImportSet) (PackageSet, error) {
	return PackageSet{
		Name: pkg.Lang,
		Modules: pkg.Modules,
	}, nil
}

type DefaultHandler struct {}
func (DefaultHandler) Handle(pkg *PackageImportSet) (PackageSet, error) {
	return PackageSet{
		Name: pkg.Lang,
		Modules: pkg.Modules,
	}, nil
}

var handlers = map[string]LangHandler{
	"python": PythonHandler{
		versionMap: map[string]string{
			"2.7": "27",
			"3.6": "36",
			"3.7": "37",
		},
	},
	"go": GoHandler{},
	"default": DefaultHandler{},
}

func generatePackageSets(cfg *ApplicationConfig) (map[string]PackageSet, error) {
	pkgs := make(map[string]PackageSet)

	for _, pkg := range cfg.Packages {
		// lang handler!
		handler, bErr := handlers[pkg.Lang]
		if !bErr {
			handler = handlers["default"]
		}

		pkgOut, err := handler.Handle(&pkg)
		check(err)
		pkgs[pkg.Lang] = pkgOut
	}
	return pkgs, nil
}

func main() {
	// TODO pull this out into arg-parse
	dat, err := ioutil.ReadFile("./env.nml")
	check(err)

	applicationConfig := ApplicationConfig{}

	err = yaml.Unmarshal(dat, &applicationConfig)
	check(err)

	applicationUrl, err := generateUrl(&applicationConfig)
	check(err)

	packageSets, err := generatePackageSets(&applicationConfig)
		check(err)

	importSet := ImportSet{
		Name: applicationConfig.Name,
		Version: applicationConfig.Version,
		Url: applicationUrl.Url,
		Sha256: applicationUrl.Sha256,
		PackageSets: packageSets,
	}


	// ------------------------------ Template -------------------
	rawTemplate, err := ioutil.ReadFile("./templates/import.nix")

	// pin to nixpkgs version
	tmpl, err := template.
		New("import").
		Delims("||", "||").
		Parse(string(rawTemplate))
	check(err)

	err = tmpl.Execute(os.Stdout, importSet)
	check(err)
	
	// build dependencies package
}
