mk_sys(
	name = "syscallheader",
	mode = "sys.h",
	arch = "amd64",
	sysconf = "//sys/src/sysconf.json",
	out = "//sys/src/libc/9syscall/sys.h"
)

mk_sys(
	name = "syscallfiles",
	mode = "syscallfiles",
	arch = "amd64",
	sysconf = "//sys/src/sysconf.json",
	out = "//sys/src/libc/9syscall",
	deps = [
	     "syscallheader"
	     ]
)