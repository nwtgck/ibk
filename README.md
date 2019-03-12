# ibk
[![CircleCI](https://circleci.com/gh/nwtgck/ibk.svg?style=shield)](https://circleci.com/gh/nwtgck/ibk)

Incremental Backup CLI with `tar`

## Motivation

We don't want to lose data. That is the biggest motivation to **backup**. In addition, we don't want to waste our storage by backups. That is the main reason to use an **incremental** backup.

Although an amazing file system such as ZFS supports snapshot functionality, every OS doesn't support it. Thus, `ibk` uses snapshot functionality of `tar` since `tar` is widely installed, used and matured. `ibk` assist you to backup incrementally using `tar` inside. This means **you can backup/restore with the natural `tar` command**.

## Requirements

### Linux
* `tar`

### macOS
* `brew install gnu-tar`

## Usage

### Backup

```bash
ibk backup ./mydir
```
Then, you have `mydir.ibk` directory which contains .tar and .snar files as follows.

- `./mydir.ibk/mydir_20190302_0757_21_407103_UTC.tar`
- `./mydir.ibk/mydir.snar`

You can also use `backup -b mybackup` or `backup --backup-path mybackup` to specify backup directory path not default `ooo.ibk`.  
Specify `--local-time` to use local time not UTC.

### Incremental Backup
When you type the command again, the backup should be conducted incrementally.
```bash
ibk backup ./mydir
```

### Restore

```bash
./ibk restore ./mydir.ibk/
```

Then, you have `ibk_restored` directory which contains the following hierarchy.
- `ibk_restored/mydir`

You can also use `restore -r myrestore` or `restore --restored-path myrestore` to specify restored directory path not default `ibk_restored`.


## Backup with only `tar`

(NOTE: You should replace `tar` with `gtar` in macOS)

```bash
# First backup
mkdir mybackup
tar -g ./mybackup/mydir.snar -cf ./mybackup/mydir_first.tar ./mydir

## Second backup (incremental)
tar -g ./mybackup/mydir.snar -cf ./mybackup/mydir_second.tar ./mydir

## Third backup (incremental)
tar -g ./mybackup/mydir.snar -cf ./mybackup/mydir_third.tar ./mydir
```

## Restore with only `tar`

```bash
cd mybackup
tar -g mydir.snar -xf mydir_first.tar
tar -g mydir.snar -xf mydir_second.tar
tar -g mydir.snar -xf mydir_third.tar
```

Then, you have `./mybackup/mydir`.
