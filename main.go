package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	sprig "github.com/Masterminds/sprig/v3"
)

const (
	packageName  = "bench"
	templatesDir = "templates"
)

func main() {
	n := flag.Uint("n", 25, "Number of fields")
	flag.Parse()

	exitOnErr(recreatePackageDir())
	exitOnErr(renderTemplates(int(*n)))
}

func exitOnErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}

func recreatePackageDir() error {
	if err := os.RemoveAll(packageName); err != nil {
		return fmt.Errorf("failed to remove %s: %w", packageName, err)
	}

	if err := os.MkdirAll(packageName, 0744); err != nil {
		return fmt.Errorf("failed to create %s: %w", packageName, err)
	}

	return nil
}

func renderTemplates(n int) error {
	templates, err := filepath.Glob(filepath.Join(templatesDir, "*.tpl"))
	if err != nil {
		return fmt.Errorf("failed to glob templates from %s: %w", templatesDir, err)
	}

	for _, t := range templates {
		if err := render(t, n); err != nil {
			return fmt.Errorf("failed to render %s: %w", t, err)
		}
	}

	return nil
}

func render(fileName string, n int) error {
	t, err := template.New(filepath.Base(fileName)).Funcs(sprig.TxtFuncMap()).ParseFiles(fileName)
	if err != nil {
		return fmt.Errorf("failed to parse template %s: %w", fileName, err)
	}

	outFile := filepath.Join(packageName, filepath.Base(strings.TrimSuffix(fileName, ".tpl")))

	f, err := os.Create(outFile)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", outFile, err)
	}

	defer f.Close()

	if err := t.Execute(f, struct{ NumFields int }{NumFields: n}); err != nil {
		return fmt.Errorf("failed to render %s: %w", outFile, err)
	}

	return nil
}
