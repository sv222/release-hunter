# Release Hunter

Release Hunter is a Go-based command-line tool that helps you discover latest GitHub releases for any software. You can use it to search for GitHub repositories by keyword and access direct download URLs for release assets.

![Alt text](/extras/rh_logo.png)

## Command-Line Flags and Aliases

- `-help, -h`: Show usage and examples.

- `-token, -t <token>`: GitHub personal access token.

- `-repo, -r <repository>`: GitHub repository in the format `user/repo`.

- `-find, -f <keyword>`: Search GitHub repositories by keyword to find correct `user/repo`.

- `-keyword, -k <keyword>`: Filter links by an optional keyword, can be used with the -f or -r flags.

## Prerequisites

Before using Release Hunter, you need to set up a GitHub personal access token. You can create the token and set it either as an environment variable or use it as a flag when running the tool.

### Creating a GitHub Personal Access Token

1. Go to [GitHub Settings → New Token](https://github.com/settings/tokens/new) in your GitHub account.
2. Set "name", "expiration date" and grant permissions: "repo - public_repo" and "user - read:user".
3. Use the token:
- as an environment variable:
```sh
export GITHUB_TOKEN=<your-token>
```
- or as a flag:
```sh
rh -t <your-token>
```

## Installation

All available versions for different operating systems can be found in releases [here](https://github.com/sv222/release-hunter/releases)

To install Release Hunter, use the following commands (Linux/amd64 example):

```sh
curl -L -o rh https://github.com/sv222/release-hunter/releases/download/v0.1.0/rh-amd64 && chmod +x rh && sudo mv rh /usr/local/bin/rh
```

## Usage

### Search GitHub Repositories

Search for GitHub repositories using a keyword:

```sh
rh -find <keyword>
```

For example, to search for repositories related to Jenkins:

```sh
rh -find jenkins
```

Filter a search result by the keyword "java":

```sh
rh -find jenkins -k java
```

### Get the Latest Release for a Repository

Retrieve the latest release for a specific GitHub repository:

```sh
rh -repo <user/repo>
```

For example, to get the latest release for "helm":

```sh
rh -repo helm/helm
```

Filter result links by keyword "arm":

```sh
rh -repo helm/helm -k arm
```

## Contributing

We welcome contributions from the community. If you have ideas, bug reports, or feature requests, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
