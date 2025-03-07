package main

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type MkDocsConfig struct {
	SiteName string   `yaml:"site_name"`
	SiteURL  string   `yaml:"site_url"`
	Plugins  []string `yaml:"plugins"`
	Theme    struct {
		Name      string `yaml:"name"`
		CustomDir string `yaml:"custom_dir"`
		Palette   struct {
			Scheme string `yaml:"scheme"`
		} `yaml:"palette"`
		Font struct {
			Text string `yaml:"text"`
			Code string `yaml:"code"`
		} `yaml:"font"`
		Logo    string `yaml:"logo"`
		Favicon string `yaml:"favicon"`
	} `yaml:"theme"`
	ExtraCSS           []string      `yaml:"extra_css"`
	RepoName           string        `yaml:"repo_name"`
	RepoURL            string        `yaml:"repo_url"`
	MarkdownExtensions []interface{} `yaml:"markdown_extensions"`
	Nav                []struct {
		Home               string        `yaml:"Home,omitempty"`
		Quickstart         string        `yaml:"Quickstart,omitempty"`
		Features           []interface{} `yaml:"Features,omitempty"`
		Examples           []string      `yaml:"Examples,omitempty"`
		Modules            []string      `yaml:"Modules,omitempty"`
		SystemRequirements []interface{} `yaml:"System Requirements,omitempty"`
		Contributing       []string      `yaml:"Contributing,omitempty"`
		GettingHelp        string        `yaml:"Getting help,omitempty"`
	} `yaml:"nav"`
	EditURI string `yaml:"edit_uri"`
	Extra   struct {
		LatestVersion string `yaml:"latest_version"`
	} `yaml:"extra"`
}

func getMkdocsConfigFile(rootDir string) string {
	return filepath.Join(rootDir, "mkdocs.yml")
}

func getExamples() ([]os.DirEntry, error) {
	return getModulesOrExamples(false)
}

func getExamplesDocs() ([]os.DirEntry, error) {
	parent, err := getRootDir()
	if err != nil {
		return nil, err
	}

	dir := filepath.Join(parent, "docs", "examples")

	return os.ReadDir(dir)
}

func getModules() ([]os.DirEntry, error) {
	return getModulesOrExamples(true)
}

func getRootDir() (string, error) {
	current, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Dir(current), nil
}

func readMkdocsConfig(rootDir string) (*MkDocsConfig, error) {
	configFile := getMkdocsConfigFile(rootDir)

	file, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	config := &MkDocsConfig{}

	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func writeMkdocsConfig(rootDir string, config *MkDocsConfig) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	file := getMkdocsConfigFile(rootDir)

	return os.WriteFile(file, data, 0777)
}
