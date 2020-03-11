// SPDX-License-Identifier: Unlicense OR MIT

package js

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func BuildJS() error {
	if err := os.MkdirAll("html", 0700); err != nil {
		return err
	}
	cmd := exec.Command(
		"go",
		"build",
		//"-ldflags="+bi.ldflags,
		"-o", filepath.Join("html", "main.wasm"),
		//bi.pkg,
	)
	cmd.Env = append(
		os.Environ(),
		"GOOS=js",
		"GOARCH=wasm",
	)
	_, err := runCmd(cmd)
	if err != nil {
		return err
	}
	const indexhtml = `<!doctype html>
<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, user-scalable=no">
		<meta name="mobile-web-app-capable" content="yes">
		<script src="wasm_exec.js"></script>
		<script>
			if (!WebAssembly.instantiateStreaming) { // polyfill
				WebAssembly.instantiateStreaming = async (resp, importObject) => {
					const source = await (await resp).arrayBuffer();
					return await WebAssembly.instantiate(source, importObject);
				};
			}

			const go = new Go();
			WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
				go.run(result.instance);
			});
		</script>
		<style>
			body,pre { margin:0;padding:0; }
		</style>
	</head>
	<body>
	</body>
</html>`
	if err := ioutil.WriteFile(filepath.Join("html", "index.html"), []byte(indexhtml), 0600); err != nil {
		return err
	}
	goroot, err := runCmd(exec.Command("go", "env", "GOROOT"))
	if err != nil {
		return err
	}
	wasmjs := filepath.Join(goroot, "misc", "wasm", "wasm_exec.js")
	if _, err := os.Stat(wasmjs); err != nil {
		return fmt.Errorf("failed to find $GOROOT/misc/wasm/wasm_exec.js driver: %v", err)
	}
	return copyFile(filepath.Join("html", "wasm_exec.js"), wasmjs)
}

func runCmdRaw(cmd *exec.Cmd) ([]byte, error) {
	out, err := cmd.Output()
	if err == nil {
		return out, nil
	}
	if err, ok := err.(*exec.ExitError); ok {
		return nil, fmt.Errorf("%s failed: %s%s", strings.Join(cmd.Args, " "), out, err.Stderr)
	}
	return nil, err
}

func runCmd(cmd *exec.Cmd) (string, error) {
	out, err := runCmdRaw(cmd)
	return string(bytes.TrimSpace(out)), err
}
