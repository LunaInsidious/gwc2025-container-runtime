# 3-2. UTS Namespace

## この説で使うsyscall

### [unshare](https://pkg.go.dev/golang.org/x/sys/unix#Unshare)

独立した**Namespaceを作り**、**今いるプロセス**をそのNamespaceに**所属させ**ます。

```go
func Unshare(flags int) (err error)
```

**どのNamespaceを分けるか**は`flags`で指定できます。  
使える`flags`は`unix`パッケージで`const`として定義されています。

```go
const (
  CLONE_NEWCGROUP = 0x2000000
  CLONE_NEWIPC    = 0x8000000
  CLONE_NEWNET    = 0x40000000
  CLONE_NEWNS     = 0x20000
  CLONE_NEWPID    = 0x20000000
  CLONE_NEWTIME   = 0x80
  CLONE_NEWUSER   = 0x10000000
  CLONE_NEWUTS    = 0x4000000
)
```
