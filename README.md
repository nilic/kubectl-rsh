# kubectl rsh 🐚
A [kubectl plugin](https://kubernetes.io/docs/tasks/extend-kubectl/kubectl-plugins/) for opening a remote shell session to a container running in a Kubernetes cluster.

This is a port of the `oc rsh` command from [OpenShift Client](https://github.com/openshift/oc) to a standalone binary, so all credits go to `oc` maintainers. Licensing and modification information is available in the [NOTICE](NOTICE) file.

## Installation

### Manual

Just download the binary for your OS and architecture from the [Releases](https://github.com/nilic/kubectl-rsh/releases) page and place it in your `PATH`.

## Usage

```
Usage:
  kubectl rsh [-c CONTAINER] [flags] (POD | TYPE/NAME) COMMAND [args...]

Flags:
      --as string                      Username to impersonate for the operation. User could be a regular user or a service account in a namespace.
      --as-group stringArray           Group to impersonate for the operation, this flag can be repeated to specify multiple groups.
      --as-uid string                  UID to impersonate for the operation.
      --cache-dir string               Default cache directory (default "$HOME/.kube/cache")
      --certificate-authority string   Path to a cert file for the certificate authority
      --client-certificate string      Path to a client certificate file for TLS
      --client-key string              Path to a client key file for TLS
      --cluster string                 The name of the kubeconfig cluster to use
  -c, --container string               Container name; defaults to first container
      --context string                 The name of the kubeconfig context to use
      --disable-compression            If true, opt-out of response compression for all requests to the server
  -f, --filename strings               to use to rsh into the resource
  -h, --help                           help for kubectl
      --insecure-skip-tls-verify       If true, the server's certificate will not be checked for validity. This will make your HTTPS connections insecure
      --kubeconfig string              Path to the kubeconfig file to use for CLI requests.
      --match-server-version           Require server version to match client version
  -n, --namespace string               If present, the namespace scope for this CLI request
  -T, --no-tty                         Disable pseudo-terminal allocation
      --pod-running-timeout duration   The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running (default 1m0s)
      --request-timeout string         The length of time to wait before giving up on a single server request. Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h). A value of zero means don't timeout requests. (default "0")
  -s, --server string                  The address and port of the Kubernetes API server
      --shell string                   Path to the shell command (default "/bin/sh")
      --tls-server-name string         Server name to use for server certificate validation. If it is not provided, the hostname used to contact the server is used
      --token string                   Bearer token for authentication to the API server
  -t, --tty                            Force a pseudo-terminal to be allocated
      --user string                    The name of the kubeconfig user to use
  -v, --version                        Print version information
```

This command will attempt to start a shell session in a pod for the specified resource. It works with pods, deployments, jobs, daemon sets, replication controllers and replica sets. Any of the aforementioned resources (apart from pods) will be resolved to a ready pod. It will default to the first container if none is specified, and will attempt to use '/bin/sh' as the default shell. You may pass any flags supported by this command before the resource name, and an optional command after the resource name, which will be executed instead of a login shell. 

A TTY will be automatically allocated if standard input is interactive - use `-t` and `-T` to override. A `TERM` variable is sent to the environment where the shell (or command) will be executed. By default its value is the same as the `TERM` variable from the local environment; if not set, `xterm` is used.

## Examples

```
  # Open a shell session on the first container in pod 'foo'
  kubectl rsh foo

  # Open a shell session on the first container in pod 'foo' and namespace 'bar'
  # (Note that kubectl client specific arguments must come before the resource name and its arguments)
  kubectl rsh -n bar foo

  # Run the command 'cat /etc/resolv.conf' inside pod 'foo'
  kubectl rsh foo cat /etc/resolv.conf

  # See the configuration of your internal registry
  kubectl rsh deployment/docker-registry cat config.yml

  # Open a shell session on the container named 'index' inside a pod of your job
  kubectl rsh -c index job/sheduled
```
