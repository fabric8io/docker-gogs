/*
 * Copyright 2015 Red Hat, Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "os/exec"
  "path/filepath"
  "regexp"
  "runtime"
  "strings"
  "syscall"
)

const (
  gogsBinary      = "/opt/gogs/gogs"
  configFile      = "/opt/gogs/custom/conf/app.ini"
  defaultUser     = "git"
  defaultCertFile = "/opt/gogs/custom/https/cert.pem"
  defaultKeyFile  = "/opt/gogs/custom/https/key.pem"
)

func init() {
  // make sure we only have one process and that it runs on the main thread (so that ideally, when we Exec, we keep our user switches and stuff)
  runtime.GOMAXPROCS(1)
  runtime.LockOSThread()
}

func main() {
  runAsUser := defaultUser
  if len(os.Getenv("GOGS_USER")) > 0 {
    runAsUser = os.Getenv("GOGS_USER")
    os.Setenv("GOGS_USER", "")
  }

  if os.Getenv("GOGS_SERVER__PROTOCOL") == "https" {
    createCerts()
  }

  if _, err := os.Stat(configFile); os.IsNotExist(err) {
    writeConfigFromEnvVars()
  }

  user, err := getUser(runAsUser, nil)
  if err != nil {
    log.Fatal(err)
  }
  if user == nil {
    log.Fatalf("Unknown user %s", runAsUser)
  }
  os.Chown(configFile, user.Uid, user.Gid)

  // clear HOME so that SetupUser will set it
  os.Setenv("HOME", "")

  err = SetupUser(runAsUser)
  if err != nil {
    log.Fatalf("error: failed switching to %q: %v", runAsUser, err)
  }

  os.Chdir(filepath.Dir(gogsBinary))

  name, err := exec.LookPath(gogsBinary)
  if err != nil {
    log.Fatalf("error: %v", err)
  }

  err = syscall.Exec(name, []string{name, "web"}, os.Environ())
  if err != nil {
    log.Fatalf("error: exec failed: %v", err)
  }

}

func createCerts() {
  basedir := filepath.Dir(defaultCertFile)
  if _, err := os.Stat(basedir); os.IsNotExist(err) {
    if err := os.MkdirAll(basedir, 0755); err != nil {
      log.Fatal(err)
    }
  }
  if _, err := os.Stat(defaultCertFile); os.IsExist(err) {
    return
  }
  os.Chdir(basedir)
  certHosts := []string{"localhost", "127.0.0.1"}
  if hostname, err := os.Hostname(); err == nil {
    certHosts = append(certHosts, hostname)
  }
  if len(os.Getenv("GOGS_SERVER__ROOT_URL")) > 0 {
    certHosts = append(certHosts, os.Getenv("GOGS_SERVER__ROOT_URL"))
  }
  cmd := exec.Command("/opt/gogs/gogs", "cert", fmt.Sprintf("%s=%s", "-host", strings.Join(certHosts, ",")))
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  err := cmd.Run()
  if err != nil {
    log.Fatal(err)
  }
  os.Chmod("cert.pem", 0644)
  os.Chmod("key.pem", 0644)
}

func writeConfigFromEnvVars() {
  basedir := filepath.Dir(configFile)
  if _, err := os.Stat(basedir); os.IsNotExist(err) {
    if err := os.MkdirAll(basedir, 0755); err != nil {
      log.Fatal(err)
    }
  }

  f, err := os.Create(configFile)
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()

  w := bufio.NewWriter(f)

  gogsConfig := make(map[string][]string)

  validConfigRE := regexp.MustCompile(`^GOGS_(?:([^=]+)__)?([^=]+)=(.+)$`)

  env := os.Environ()
  for _, envVar := range env {
    matches := validConfigRE.FindStringSubmatch(envVar)
    if matches != nil {
      section := strings.Replace(strings.ToLower(matches[1]), "_", ".", -1)
      key := strings.ToUpper(matches[2])
      value := matches[3]
      gogsConfig[section] =
        append(gogsConfig[section], fmt.Sprintf("%s=%s", key, value))
    }
  }

  if _, ok := gogsConfig[""]; ok {
    for _, value := range gogsConfig[""] {
      writeLine(w, value)
    }

    writeEmptyLine(w)
    delete(gogsConfig, "")
  }

  for section, values := range gogsConfig {
    writeLine(w, fmt.Sprintf("[%s]", section))
    for _, value := range values {
      writeLine(w, value)
    }
    writeEmptyLine(w)
  }

  w.Flush()
}

func writeLine(w *bufio.Writer, value string) {
  if _, err := w.WriteString(fmt.Sprintf("%s\n", value)); err != nil {
    log.Fatal(err)
  }
}

func writeEmptyLine(w *bufio.Writer) {
  writeLine(w, "")
}
