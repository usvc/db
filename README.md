# DB

[![latest release](https://badge.fury.io/gh/usvc%2Fdb.svg)](https://github.com/usvc/db/releases)
[![pipeline status](https://gitlab.com/usvc/utils/db/badges/master/pipeline.svg)](https://gitlab.com/usvc/utils/db/-/commits/master)
[![Build Status](https://travis-ci.org/usvc/db.svg?branch=master)](https://travis-ci.org/usvc/db)
<!-- // to add when ready
[![Test Coverage](https://api.codeclimate.com/v1/badges/f5311cc9a7bd253f3c4a/test_coverage)](https://codeclimate.com/github/usvc/db/test_coverage)
[![Maintainability](https://api.codeclimate.com/v1/badges/f5311cc9a7bd253f3c4a/maintainability)](https://codeclimate.com/github/usvc/db/maintainability)
-->

An easy-peasy database CLI tool.

| Find it on | URL |
| --- | --- |
| Github | [https://github.com/usvc/db](https://github.com/usvc/db) |
| Gitlab | [https://gitlab.com/usvc/utils/db](https://gitlab.com/usvc/utils/db) |

- - -

# Usage

## Quick Overview

```sh
# show help
db --help;

# check database credentials/connection
db check;

# migrate database to the next step only (using files at ./path/to/migrations)
db migrate ./path/to/migrations;

# migrate database upwards for 2 steps only (using files at ./path/to/migrations)
db migrate --steps 2 ./path/to/migrations;

# migrate database to the latest step (using files at ./path/to/migrations)
db migrate --latest ./path/to/migrations;
```

## Show Help

The `--help` flag shows all possible flags/sub-commands. The configurations can also be found below in alphabetical order:

| Flag | Default Value | Description |
| --- | --- | --- |
| `-d, --driver` | `"mysql"` | Defines the database driver to use (only MySQL is supported for now) |
| `-H, --host` | `"localhost"` | Defines the hostname at which the database server can be reached at |
| `-f, --log-format` | `"text"` | Defines the format of the logger according to [`github.com/usvc/go-logger`](https://github.com/usvc/go-logger). Other values: [`"json"`] |
| `-t, --log-type` | `"stdout"` | Defines the type of the logger according to [`github.com/usvc/go-logger`](github.com/usvc/go-logger). Other values: [`"levelled"`] |
| `-p, --password` | `"password"` | Defines the password of the user used to login to the database server |
| `-P, --port` | `3306` | Defines the port at which the database server can be reached |
| `-r, --retry-count` | `5` | Defines the number of times the application will re-attempt a failed connection to the database |
| `-R, --retry-interval-ms` | `3000` | Defines the number of milliseconds in between database connection retry attempts |
| `-u, --username` | `"user"` | Defines the username of the user used to login to the database server |

## Check Database Credentials/Connection

The `check` sub-command 

## Migrate Database



- - -

# Configuration

## Environment Variables

> `WIP`

| Key | Default Value | Description |
| --- | --- | --- |
| | |

- - -

# Development Runbook

## Getting Started

1. Clone this repository
2. Run `make deps` to pull in external dependencies
3. Write some awesome stuff
4. Run `make test` to ensure unit tests are passing
5. Push

## Continuous Integration (CI) Pipeline

### On Github

Github is used to deploy binaries/libraries because of it's ease of access by other developers.

#### Releasing

Releasing of the binaries can be done via Travis CI.

1. On Github, navigate to the [tokens settings page](https://github.com/settings/tokens) (by clicking on your profile picture, selecting **Settings**, selecting **Developer settings** on the left navigation menu, then **Personal Access Tokens** again on the left navigation menu)
2. Click on **Generate new token**, give the token an appropriate name and check the checkbox on **`public_repo`** within the **repo** header
3. Copy the generated token
4. Navigate to [travis-ci.org](https://travis-ci.org) and access the cooresponding repository there. Click on the **More options** button on the top right of the repository page and select **Settings**
5. Scroll down to the section on **Environment Variables** and enter in a new **NAME** with `RELEASE_TOKEN` and the **VALUE** field cooresponding to the generated personal access token, and hit **Add**

### On Gitlab

Gitlab is used to run tests and ensure that builds run correctly.

#### Version Bumping

1. Run `make .ssh`
2. Copy the contents of the file generated at `./.ssh/id_rsa.base64` into an environment variable named **`DEPLOY_KEY`** in **Settings > CI/CD > Variables**
3. Navigate to the **Deploy Keys** section of the **Settings > Repository > Deploy Keys** and paste in the contents of the file generated at `./.ssh/id_rsa.pub` with the **Write access allowed** checkbox enabled

- **`DEPLOY_KEY`**: generate this by running `make .ssh` and copying the contents of the file generated at `./.ssh/id_rsa.base64`

#### DockerHub Publishing

1. Login to [https://hub.docker.com](https://hub.docker.com), or if you're using your own private one, log into yours
2. Navigate to [your security settings at the `/settings/security` endpoint](https://hub.docker.com/settings/security)
3. Click on **Create Access Token**, type in a name for the new token, and click on **Create**
4. Copy the generated token that will be displayed on the screen
5. Enter the following varialbes into the CI/CD Variables page at **Settings > CI/CD > Variables** in your Gitlab repository:

- **`DOCKER_REGISTRY_URL`**: The hostname of the Docker registry (defaults to `docker.io` if not specified)
- **`DOCKER_REGISTRY_USERNAME`**: The username you used to login to the Docker registry
- **`DOCKER_REGISTRY_PASSWORD`**: The generated access token

- - -

# Licensing

Code here is licensed under the [MIT license](./LICENSE) by [@zephinzer](https://gitlab.com/zephinzer).
