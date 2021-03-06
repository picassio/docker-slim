package fsutil

import (
	"syscall"
)

func SysStatInfo(raw *syscall.Stat_t) SysStat {
	return SysStat{
		Uid:   raw.Uid,
		Gid:   raw.Gid,
		Atime: raw.Atimespec,
		Mtime: raw.Mtimespec,
		Ctime: raw.Ctimespec,
	}
}
