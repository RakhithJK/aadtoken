# Azure Active Directory (Azure AD) Token

Get AAD token quickly.

## Docs

```txt
aadtoken [scope]

Get AAD token by providing a scope.
```

For more details, [it is based on Azure SDK](https://docs.microsoft.com/en-us/azure/developer/go/azure-sdk-authentication).

## Usage

### Azure Kubernetes Service (AKS)

Using [AAD Pod Identity](https://github.com/Azure/aad-pod-identity).

```bash
$ kubectl run aadtoken \
  --rm \
  --attach \
  --image=skibish/aadtoken
  --image-pull-policy=Always \
  --labels="aadpodidbinding=<your value>" \
  --restart=Never \
  "a6a0b2e1-d5d6-4b33-93c8-68125d41b9e7"

eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

### CLI

```bash
$ aadtoken "a6a0b2e1-d5d6-4b33-93c8-68125d41b9e7"
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
```

## Motivation

Working with Azure, Azure AD based services from time to time you need to generate the token for various reasons.
For example:

* verify connection to services
* check the token

This tool is a helper to do it faster.
