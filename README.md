# Platoon Go

This is a rewrite of the original Platoon library as a standalone application; rewritten entirely in Go.

## Why?

Platoon is currently developed in PHP and is very Laravel specific. The attempt to rewrite it serves numerous purposes:

- I'm using this project to learn more about Go
- Leveraging Go Routines will help to make Platoon even faster
- Rewriting in Go removed the dependency on Laravel and Ecnvoy
- There's no PHP and Composer requirement meaning Platoon could be used for any project using any technology.

## State

This is a very very early-stage project. I want to rebuild all the existing features into this version and I'm still learning alot about Go. This should not be used at all. Please wait for a proper release version.

## Build

To build, run:

```shell
go build -o bin/platoon main.go

./bin/platoon --version    # 0.0.0-0.0.1
```

## Config

The configuration for the original Platoon was written as a PHP array. For Platoon-go, the configuration is changing to a YAML file (`platoon.yml`) placed at the root of the project. An example would look like this:

```yml
repo: git@github.com:org/app.git

default: staging

targets:
  common:
    php:
      bin: /usr/bin/php
      composer: /usr/bin/composer
    keep: 2
  
  staging:
    host: staging.example.com
    port: 22
    username: deploy
    root: /var/www/myapp
    branch: main
    assets:
      - public/build:public/build
    scripts:
      pre-deploy-local:
        - npm run build
      post-deploy-remote:
        - @artisan config:cache
```

The scripts will change from the existing solution as the deployment process will no longer be Laravel or PHP specific.

### Deployment

To run a deployment

```shell
# to a specific target
platoon deploy staging

# to the default target
platoon deploy

# to multiple targets
platoon deploy staging,production
```

### Release management

Get a list of available releases:

```shell
# the default target
platoon release:list

# or a specific target
platoon release:list staging
```

Rollback to the previous release:

```shell
platoon release:rollback
```

Set a specific release as active:

```shell
platoon release:active 202512042144
```

## License

Platoon-Go, like the original, is licensed under MIT. You can find more in the [LICENCE.md]() file.

## Copyright

Copyright (c) 2026 Warrick Bayman