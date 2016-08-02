ZBS
===

The **ZFS based Backup System** is a tiny backup system based on the functionality offered by [ZFS](https://www.google.es/url?sa=t&rct=j&q=&esrc=s&source=web&cd=7&cad=rja&uact=8&ved=0ahUKEwiZpdOLqqLOAhWBthQKHQhyCIUQFghNMAY&url=http%3A%2F%2Fzfsonlinux.org%2F&usg=AFQjCNHyOGcWrk6IrGr_XAE7LTg30Jgkxw).

The system is developed under Linux (Ubuntu) and OSX so these are the tested systems.

This service will launch remote sync backup copies and will maintain a registry in database. The service will provide the needed **ZFS** _buckets_ for backups storage.

It will manage backup, archive, deletion and restoring.

The system will use `rsync` to send the backups from the servers to this service.
