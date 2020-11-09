# Scylla Manager

Welcome to Scylla Manager repository!
Scylla Manager user docs can be found in [Scylla docs portal](https://docs.scylladb.com/operating-scylla/manager/). 

Scylla Manager consists of tree components:

1. a server that manages Scylla clusters
1. an agent running on each node
1. `sctool` - a CLI interface to the server

## Installing and updating Go

The desired Go version is specified in `.go-version` file.
You can install or update Go version by running `make install` in `go` directory.

**Procedure**:

1. Define `GOROOT` environment variable as `/PATH/TO/GO/SDKS/latest`
1. Run `make -C go install`

This would download required go version and install it at `GOROOT`.

## Installing other packages needed for development

1. Install Docker and `docker-compose`
1. Run installation script `./install_deps.sh`

## Configuring imports formatting in GoLand

If using GoLand update import grouping policy:

1. Open File -> Settings -> Editor -> Code Style -> Go
1. Go to Imports Pane
1. Set "Sorting type" to goimports
1. Check every checkbox but "Group current project imports"
1. Press OK

## Running a development environment 

Let's start the development environment.

```bash
make start-dev-env
```

This command will:
1. Build custom Scylla Docker image (testing/scylla)
1. Compile server, agent and sctool binaries
1. Start Scylla cluster with 2 DCs 3 nodes each (6 containers)
1. Start MinIO and Prometheus containers
1. Start dedicated Scylla container for Scylla Manager datastore
1. Start Scylla Manager server on localhost

Docker compose environment for test cluster is located in the `testing` directory.

Once `scylla-manager` is bootstrapped use `./sctool` to add information about the cluster to the manager:

```bash
./sctool cluster add --host=192.168.100.11 --auth-token token --name my-cluster 
```

You can ask Scylla Manager to give the status of the cluster:

```bash
./sctool status
```

For other commands consult [user manual](https://docs.scylladb.com/operating-scylla/manager/).

### Helpful Makefile targets

Make is self-documenting run `make` or `make help` to see available targets with descriptions. 

More Makefile targets are available in `testing` directory.

## Running tests

If test environment is running correctly you can run tests with:

```bash
make test
```

This runs both unit and integration tests. You can run them separately with:

```bash
make unit-test
make integration-test
make pkg-integration-test PKG=pkg/service/foo
```

Project testing heavily depends on integration tests, which are slow.
Tests should run for couple of minutes.
All tests should succeed.

## Extending HTTP clients

There are two HTTP APIs that are specified with Swagger format, the API of the Scylla node and the API of the `scylla-manager` itself.

Client implementations are separated into packages:

- `scyllaclient` which contains client for the Scylla API along with Swagger specification file `scylla-api.json` and,
- `managerclient` which contains client for the `scylla-manager` API along with Swagger specification file `scylla-manager-api.json`.

Both clients are generated automatically by shell scripts.
To refresh generated packages from Swagger specification run:

```bash
make generate

# Or for generating specific client package ex.
go generate ./managerclient
go generate ./scyllaclient
```

## Sending patches

Develop on dedicated branch rooted at master.
Pull master regularly and rebase your work upon master whenever your dev branch is behind.
You are allowed and required to force push to your branches to keep the history neat and clean.

Document your work in commit messages.
Explain why you made the changes and mention IDs of the affected issues on Github.

Before pushing please run `make check` and `make test` to make sure all tests pass.

When satisfied create a pull request onto master.
All pull requests have to go trough review process before they are merged.
