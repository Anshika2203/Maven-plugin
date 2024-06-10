// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Plugin struct {
	ProxyHost      string `envconfig:"PLUGIN_PROXY_HOST"`
	ProxyPort      string `envconfig:"PLUGIN_PROXY_PORT"`
	ProxyUser      string `envconfig:"PLUGIN_PROXY_USER"`
	ProxyPassword  string `envconfig:"PROXY_PASSWORD"`
	NonProxyHosts  string `envconfig:"PROXY_NON_PROXY_HOSTS"`
	ServerUser     string `envconfig:"PLUGIN_SERVER_USER"`
	ServerPassword string `envconfig:"PLUGIN_SERVER_PASSWORD"`
	MavenMirrorURL string `envconfig:"PLUGIN_MAVEN_MIRROR_URL"`
	Goals          string `envconfig:"PLUGIN_GOALS"`
	MavenModules   string `envconfig:"PLUGIN_MAVEN_MODULES"`
	ContextDir     string `envconfig:"PLUGIN_CONTEXT_DIR"`
	LogLevel       string `envconfig:"PLUGIN_LOG_LEVEL"`
}

func (p *Plugin) Exec(ctx context.Context) error {
	if err := initMavenSettings(p); err != nil {
		return err
	}

	if err := runMavenCommand(p); err != nil {
		return err
	}

	return nil
}

func initMavenSettings(p *Plugin) error {
	MAVEN_CONFIG := os.Getenv("MAVEN_CONFIG")
	if MAVEN_CONFIG == "" {
		// Use the home directory of the current user
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get home directory: %w", err)
		}
		MAVEN_CONFIG = fmt.Sprintf("%s/.m2", homeDir)
	}

	settingsPath := MAVEN_CONFIG + "/settings.xml"

	if err := os.MkdirAll(MAVEN_CONFIG, 0755); err != nil {
		return fmt.Errorf("failed to create maven config directory: %w", err)
	}

	if _, err := os.Stat(settingsPath); os.IsNotExist(err) {
		f, err := os.Create(settingsPath)
		if err != nil {
			return err
		}
		defer f.Close()

		// Constructing proxy settings if provided
		var proxyXML string
		if p.ProxyHost != "" && p.ProxyPort != "" {
			proxyXML = fmt.Sprintf(`<proxy>
	<id>genproxy</id>
	<active>true</active>
	<protocol>http</protocol>
	<host>%s</host>
	<port>%s</port>`, p.ProxyHost, p.ProxyPort)
			if p.ProxyUser != "" && p.ProxyPassword != "" {
				proxyXML += fmt.Sprintf(`
	<username>%s</username>
	<password>%s</password>`, p.ProxyUser, p.ProxyPassword)
			}
			if p.NonProxyHosts != "" {
				proxyXML += fmt.Sprintf(`
	<nonProxyHosts>%s</nonProxyHosts>`, strings.ReplaceAll(p.NonProxyHosts, ",", "|"))
			}
			proxyXML += `
</proxy>`
		}

		// Constructing server settings if provided
		var serverXML string
		if p.ServerUser != "" && p.ServerPassword != "" {
			serverXML = fmt.Sprintf(`<server>
	<id>serverid</id>
	<username>%s</username>
	<password>%s</password>
</server>`, p.ServerUser, p.ServerPassword)
		}

		// Constructing mirror settings if provided
		var mirrorXML string
		if p.MavenMirrorURL != "" {
			mirrorXML = fmt.Sprintf(`<mirror>
	<id>mirror.default</id>
	<url>%s</url>
	<mirrorOf>*</mirrorOf>
</mirror>`, p.MavenMirrorURL)
		}

		xml := fmt.Sprintf(`<settings>
	<servers>
%s
	</servers>
	<mirrors>
%s
	</mirrors>
	<proxies>
%s
	</proxies>
</settings>`, serverXML, mirrorXML, proxyXML)

		if _, err := f.WriteString(xml); err != nil {
			return err
		}
	}

	return nil
}

func runMavenCommand(p *Plugin) error {
	mvnCommand := "mvn -B -s " + os.Getenv("MAVEN_CONFIG") + "/settings.xml"

	if p.Goals != "" {
		mvnCommand += " " + p.Goals
	} else {
		// Default goals if not provided
		mvnCommand += " -DskipTests clean install"
	}

	if p.MavenModules != "" {
		mvnCommand += " -pl " + p.MavenModules
	}

	if p.ContextDir != "" {
		mvnCommand += " -f " + p.ContextDir
	}

	cmd := exec.Command("bash", "-c", mvnCommand)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Running command: %s\n", mvnCommand)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute Maven command: %w", err)
	}

	return nil
}
