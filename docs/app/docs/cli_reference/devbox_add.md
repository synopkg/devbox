# devbox add

Add a new package to your devbox

```bash
devbox add <pkg>... [flags]
```

## Examples

```bash
# Add the latest version of the `ripgrep` package
devbox add ripgrep

# Install glibcLocales only on x86_64-linux and aarch64-linux
devbox add glibcLocales --platform x86_64-linux,aarch64-linux

# Exclude busybox from installation on macOS
devbox add busybox --exclude-platform aarch64-darwin,x86_64-darwin
```

## Options

<!-- Markdown Table of Options -->
| Option | Description |
| --- | --- |
| `--allow-insecure` | allows Devbox to install a package that is marked insecure by Nix |
| `-c, --config string` | path to directory containing a devbox.json config file |
| `-e, --exclude-platform strings` | exclude packages from a specific platform. |
| `-h, --help` | help for add |
| `-q, --quiet` | quiet mode: Suppresses logs. |
| `-p`, `--platform strings` | install packages only on specific platforms. |

Valid Platforms include:

* `aarch64-darwin`
* `aarch64-linux`
* `x86_64-darwin`
* `x86_64-linux`

The platforms below are also supported, but will build packages from source

* `i686-linux`
* `armv7l-linux`


## SEE ALSO

* [devbox](./devbox.md)	 - Instant, easy, predictable shells and containers

