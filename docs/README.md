# Website

This website is built using [Docusaurus 2](https://docusaurus.io/), a modern static website generator.

You can also test and contribute to our docs online using Devbox Cloud!

[![Open In Devbox.sh](https://jetpack.io/img/devbox/open-in-devbox.svg)](https://synopkg.github.io/devbox/https://github.com/synopkg/devbox?folder=docs/app)

## Installation

```bash
cd docs/app     # from the devbox root directory
synopkg.github.io/devboxell    # optional, develop inside a devbox
yarn install    # run in synopkg.github.io/devboxell
```

### Local Development

```bash
yarn start
```

This command starts a local development server and opens up a browser window. Most changes are reflected live without having to restart the server.

### Build

```bash
yarn build
```

This command generates static content into the `build` directory and can be served using any static contents hosting service.

### Deployment

When a pull request is opened, it will automatically deploy via CICD to a preview.
When a pull request is merged, it will automatically deploy to production.
Check https://synopkg.github.io/devbox/ after merge to see the latest changes.
