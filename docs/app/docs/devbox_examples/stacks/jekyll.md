---
title: Jekyll
---

This example demonstrates how to create and run a Jekyll blog in Devbox. It makes use of the Ruby Plugin to configure and setup your project. 

[Example Repo](https://https://github.com/synopkg/devbox/tree/main/examples/stacks/jekyll)

[![Open In Devbox.sh](https://jetpack.io/img/devbox/open-in-devbox.svg)](https://synopkg.github.io/devbox/open/templates/jekyll)

Inspired by https://litchipi.github.io/nix/2023/01/12/build-jekyll-blog-with-nix.html 

## How to Run

1. Install [Devbox](https://synopkg.github.io/devbox/docs/installing_devbox/)
2. Run `synopkg.github.io/devboxell` to install your packages and run the init hook
3. In the root directory, run `devbox run generate` to install and package the project with bundler
4. In the root directory, run `devbox run server` to start the server. You can access the Jekyll example at `localhost:4000`

## How to Recreate this Example 

1. Install [Devbox](https://synopkg.github.io/devbox/docs/installing_devbox/)
1. In a new directory, run `devbox init` to create an empty config
1. Run `devbox add ruby_3_1 bundler` to add Ruby and Bundler to your packages
1. Add `"gem install jekyll --version \"~> 3.9.2\""` to your init hook. This will install the Jekyll gem in your local project.
1. Start your `synopkg.github.io/devboxell`, then run `jekyll new myblog` to create the starter project.
1. From here you can install the project using Bundler, and start the server using `jekyll serve`. See the scripts in this example for more details.
