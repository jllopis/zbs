ZBS
===

The **ZFS based Backup System** is a tiny backup system based on the functionality offered by [ZFS](https://www.google.es/url?sa=t&rct=j&q=&esrc=s&source=web&cd=7&cad=rja&uact=8&ved=0ahUKEwiZpdOLqqLOAhWBthQKHQhyCIUQFghNMAY&url=http%3A%2F%2Fzfsonlinux.org%2F&usg=AFQjCNHyOGcWrk6IrGr_XAE7LTg30Jgkxw).

The system is developed under Linux (Ubuntu) and OSX so these are the tested systems.

This service will launch remote sync backup copies and will maintain a registry in database. The service will provide the needed **ZFS** _buckets_ for backups storage.

It will manage backup, archive, deletion and restoring.

The system will use `rsync` to send the backups from the servers to this service.

Packages used and references:

* [zazab/runcmd](https://github.com/zazab/runcmd): golang package for run local/remote shell commands
* [jdkanani/commandcast](https://github.com/jdkanani/commandcast): Run command on multiple hosts over SSH (Golang)
* [golang/sync/errgroup](https://github.com/golang/sync/tree/master/errgroup): Package errgroup provides synchronization, error propagation, and Context cancelation for groups of goroutines working on subtasks of a common task. Doc @ https://godoc.org/golang.org/x/sync/errgroup
* [Go-lang: Mocking exec.Command Using Interfaces](https://joshrendek.com/2014/06/go-lang-mocking-exec-dot-command-using-interfaces/) by Josh Rendek
* [Run a subprocess and connect to it with golang](https://coderwall.com/p/ik5xxa/run-a-subprocess-and-connect-to-it-with-golang) by Murphy Randle
* [Cron like package](https://github.com/robfig/cron) by Rob Figueiredo

