// cgo -godefs -- -Wall -Werror -static -I/tmp/include linux/types.go | go run mkpost.go
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build mips,linux

package unix

const (
	sizeofPtr      = 0x4
	sizeofShort    = 0x2
	sizeofInt      = 0x4
	sizeofLong     = 0x4
	sizeofLongLong = 0x8
	PathMax        = 0x1000
)

type (
	_C_short     int16
	_C_int       int32
	_C_long      int32
	_C_long_long int64
)

type Timespec struct {
	Sec  int32
	Nsec int32
}

type Timeval struct {
	Sec  int32
	Usec int32
}

type Timex struct {
	Modes     uint32
	Offset    int32
	Freq      int32
	Maxerror  int32
	Esterror  int32
	Status    int32
	Constant  int32
	Precision int32
	Tolerance int32
	Time      Timeval
	Tick      int32
	Ppsfreq   int32
	Jitter    int32
	Shift     int32
	Stabil    int32
	Jitcnt    int32
	Calcnt    int32
	Errcnt    int32
	Stbcnt    int32
	Tai       int32
	_         [44]byte
}

type Time_t int32

type Tms struct {
	Utime  int32
	Stime  int32
	Cutime int32
	Cstime int32
}

type Utimbuf struct {
	Actime  int32
	Modtime int32
}

type Rusage struct {
	Utime    Timeval
	Stime    Timeval
	Maxrss   int32
	Ixrss    int32
	Idrss    int32
	Isrss    int32
	Minflt   int32
	Majflt   int32
	Nswap    int32
	Inblock  int32
	Oublock  int32
	Msgsnd   int32
	Msgrcv   int32
	Nsignals int32
	Nvcsw    int32
	Nivcsw   int32
}

type Rlimit struct {
	Cur uint64
	Max uint64
}

type _Gid_t uint32

type Stat_t struct {
	Dev     uint32
	Pad1    [3]int32
	Ino     uint64
	Mode    uint32
	Nlink   uint32
	Uid     uint32
	Gid     uint32
	Rdev    uint32
	Pad2    [3]int32
	Size    int64
	Atim    Timespec
	Mtim    Timespec
	Ctim    Timespec
	Blksize int32
	Pad4    int32
	Blocks  int64
	Pad5    [14]int32
}

type Statfs_t struct {
	Type    int32
	Bsize   int32
	Frsize  int32
	_       [4]byte
	Blocks  uint64
	Bfree   uint64
	Files   uint64
	Ffree   uint64
	Bavail  uint64
	Fsid    Fsid
	Namelen int32
	Flags   int32
	Spare   [5]int32
	_       [4]byte
}

type StatxTimestamp struct {
	Sec  int64
	Nsec uint32
	_    int32
}

type Statx_t struct {
	Mask            uint32
	Blksize         uint32
	Attributes      uint64
	Nlink           uint32
	Uid             uint32
	Gid             uint32
	Mode            uint16
	_               [1]uint16
	Ino             uint64
	Size            uint64
	Blocks          uint64
	Attributes_mask uint64
	Atime           StatxTimestamp
	Btime           StatxTimestamp
	Ctime           StatxTimestamp
	Mtime           StatxTimestamp
	Rdev_major      uint32
	Rdev_minor      uint32
	Dev_major       uint32
	Dev_minor       uint32
	_               [14]uint64
}

type Dirent struct {
	Ino    uint64
	Off    int64
	Reclen uint16
	Type   uint8
	Name   [256]int8
	_      [5]byte
}

type Fsid struct {
	Val [2]int32
}

type Flock_t struct {
	Type   int16
	Whence int16
	_      [4]byte
	Start  int64
	Len    int64
	Pid    int32
	_      [4]byte
}

type FscryptPolicy struct {
	Version                   uint8
	Contents_encryption_mode  uint8
	Filenames_encryption_mode uint8
	Flags                     uint8
	Master_key_descriptor     [8]uint8
}

type FscryptKey struct {
	Mode uint32
	Raw  [64]uint8
	Size uint32
}

type KeyctlDHParams struct {
	Private int32
	Prime   int32
	Base    int32
}

const (
	FADV_NORMAL     = 0x0
	FADV_RANDOM     = 0x1
	FADV_SEQUENTIAL = 0x2
	FADV_WILLNEED   = 0x3
	FADV_DONTNEED   = 0x4
	FADV_NOREUSE    = 0x5
)

type RawSockaddrInet4 struct {
	Family uint16
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]uint8
}

type RawSockaddrInet6 struct {
	Family   uint16
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte /* in6_addr */
	Scope_id uint32
}

type RawSockaddrUnix struct {
	Family uint16
	Path   [108]int8
}

type RawSockaddrLinklayer struct {
	Family   uint16
	Protocol uint16
	Ifindex  int32
	Hatype   uint16
	Pkttype  uint8
	Halen    uint8
	Addr     [8]uint8
}

type RawSockaddrNetlink struct {
	Family uint16
	Pad    uint16
	Pid    uint32
	Groups uint32
}

type RawSockaddrHCI struct {
	Family  uint16
	Dev     uint16
	Channel uint16
}

type RawSockaddrL2 struct {
	Family      uint16
	Psm         uint16
	Bdaddr      [6]uint8
	Cid         uint16
	Bdaddr_type uint8
	_           [1]byte
}

type RawSockaddrCAN struct {
	Family  uint16
	_       [2]byte
	Ifindex int32
	Addr    [8]byte
}

type RawSockaddrALG struct {
	Family uint16
	Type   [14]uint8
	Feat   uint32
	Mask   uint32
	Name   [64]uint8
}

type RawSockaddrVM struct {
	Family    uint16
	Reserved1 uint16
	Port      uint32
	Cid       uint32
	Zero      [4]uint8
}

type RawSockaddr struct {
	Family uint16
	Data   [14]int8
}

type RawSockaddrAny struct {
	Addr RawSockaddr
	Pad  [96]int8
}

type _Socklen uint32

type Linger struct {
	Onoff  int32
	Linger int32
}

type Iovec struct {
	Base *byte
	Len  uint32
}

type IPMreq struct {
	Multiaddr [4]byte /* in_addr */
	Interface [4]byte /* in_addr */
}

type IPMreqn struct {
	Multiaddr [4]byte /* in_addr */
	Address   [4]byte /* in_addr */
	Ifindex   int32
}

type IPv6Mreq struct {
	Multiaddr [16]byte /* in6_addr */
	Interface uint32
}

type PacketMreq struct {
	Ifindex int32
	Type    uint16
	Alen    uint16
	Address [8]uint8
}

type Msghdr struct {
	Name       *byte
	Namelen    uint32
	Iov        *Iovec
	Iovlen     uint32
	Control    *byte
	Controllen uint32
	Flags      int32
}

type Cmsghdr struct {
	Len   uint32
	Level int32
	Type  int32
}

type Inet4Pktinfo struct {
	Ifindex  int32
	Spec_dst [4]byte /* in_addr */
	Addr     [4]byte /* in_addr */
}

type Inet6Pktinfo struct {
	Addr    [16]byte /* in6_addr */
	Ifindex uint32
}

type IPv6MTUInfo struct {
	Addr RawSockaddrInet6
	Mtu  uint32
}

type ICMPv6Filter struct {
	Data [8]uint32
}

type Ucred struct {
	Pid int32
	Uid uint32
	Gid uint32
}

type TCPInfo struct {
	State          uint8
	Ca_state       uint8
	Retransmits    uint8
	Probes         uint8
	Backoff        uint8
	Options        uint8
	_              [2]byte
	Rto            uint32
	Ato            uint32
	Snd_mss        uint32
	Rcv_mss        uint32
	Unacked        uint32
	Sacked         uint32
	Lost           uint32
	Retrans        uint32
	Fackets        uint32
	Last_data_sent uint32
	Last_ack_sent  uint32
	Last_data_recv uint32
	Last_ack_recv  uint32
	Pmtu           uint32
	Rcv_ssthresh   uint32
	Rtt            uint32
	Rttvar         uint32
	Snd_ssthresh   uint32
	Snd_cwnd       uint32
	Advmss         uint32
	Reordering     uint32
	Rcv_rtt        uint32
	Rcv_space      uint32
	Total_retrans  uint32
}

const (
	SizeofSockaddrInet4     = 0x10
	SizeofSockaddrInet6     = 0x1c
	SizeofSockaddrAny       = 0x70
	SizeofSockaddrUnix      = 0x6e
	SizeofSockaddrLinklayer = 0x14
	SizeofSockaddrNetlink   = 0xc
	SizeofSockaddrHCI       = 0x6
	SizeofSockaddrL2        = 0xe
	SizeofSockaddrCAN       = 0x10
	SizeofSockaddrALG       = 0x58
	SizeofSockaddrVM        = 0x10
	SizeofLinger            = 0x8
	SizeofIovec             = 0x8
	SizeofIPMreq            = 0x8
	SizeofIPMreqn           = 0xc
	SizeofIPv6Mreq          = 0x14
	SizeofPacketMreq        = 0x10
	SizeofMsghdr            = 0x1c
	SizeofCmsghdr           = 0xc
	SizeofInet4Pktinfo      = 0xc
	SizeofInet6Pktinfo      = 0x14
	SizeofIPv6MTUInfo       = 0x20
	SizeofICMPv6Filter      = 0x20
	SizeofUcred             = 0xc
	SizeofTCPInfo           = 0x68
)

const (
	IFA_UNSPEC           = 0x0
	IFA_ADDRESS          = 0x1
	IFA_LOCAL            = 0x2
	IFA_LABEL            = 0x3
	IFA_BROADCAST        = 0x4
	IFA_ANYCAST          = 0x5
	IFA_CACHEINFO        = 0x6
	IFA_MULTICAST        = 0x7
	IFLA_UNSPEC          = 0x0
	IFLA_ADDRESS         = 0x1
	IFLA_BROADCAST       = 0x2
	IFLA_IFNAME          = 0x3
	IFLA_MTU             = 0x4
	IFLA_LINK            = 0x5
	IFLA_QDISC           = 0x6
	IFLA_STATS           = 0x7
	IFLA_COST            = 0x8
	IFLA_PRIORITY        = 0x9
	IFLA_MASTER          = 0xa
	IFLA_WIRELESS        = 0xb
	IFLA_PROTINFO        = 0xc
	IFLA_TXQLEN          = 0xd
	IFLA_MAP             = 0xe
	IFLA_WEIGHT          = 0xf
	IFLA_OPERSTATE       = 0x10
	IFLA_LINKMODE        = 0x11
	IFLA_LINKINFO        = 0x12
	IFLA_NET_NS_PID      = 0x13
	IFLA_IFALIAS         = 0x14
	IFLA_NUM_VF          = 0x15
	IFLA_VFINFO_LIST     = 0x16
	IFLA_STATS64         = 0x17
	IFLA_VF_PORTS        = 0x18
	IFLA_PORT_SELF       = 0x19
	IFLA_AF_SPEC         = 0x1a
	IFLA_GROUP           = 0x1b
	IFLA_NET_NS_FD       = 0x1c
	IFLA_EXT_MASK        = 0x1d
	IFLA_PROMISCUITY     = 0x1e
	IFLA_NUM_TX_QUEUES   = 0x1f
	IFLA_NUM_RX_QUEUES   = 0x20
	IFLA_CARRIER         = 0x21
	IFLA_PHYS_PORT_ID    = 0x22
	IFLA_CARRIER_CHANGES = 0x23
	IFLA_PHYS_SWITCH_ID  = 0x24
	IFLA_LINK_NETNSID    = 0x25
	IFLA_PHYS_PORT_NAME  = 0x26
	IFLA_PROTO_DOWN      = 0x27
	IFLA_GSO_MAX_SEGS    = 0x28
	IFLA_GSO_MAX_SIZE    = 0x29
	IFLA_PAD             = 0x2a
	IFLA_XDP             = 0x2b
	IFLA_EVENT           = 0x2c
	IFLA_NEW_NETNSID     = 0x2d
	IFLA_IF_NETNSID      = 0x2e
	IFLA_MAX             = 0x31
	RT_SCOPE_UNIVERSE    = 0x0
	RT_SCOPE_SITE        = 0xc8
	RT_SCOPE_LINK        = 0xfd
	RT_SCOPE_HOST        = 0xfe
	RT_SCOPE_NOWHERE     = 0xff
	RT_TABLE_UNSPEC      = 0x0
	RT_TABLE_COMPAT      = 0xfc
	RT_TABLE_DEFAULT     = 0xfd
	RT_TABLE_MAIN        = 0xfe
	RT_TABLE_LOCAL       = 0xff
	RT_TABLE_MAX         = 0xffffffff
	RTA_UNSPEC           = 0x0
	RTA_DST              = 0x1
	RTA_SRC              = 0x2
	RTA_IIF              = 0x3
	RTA_OIF              = 0x4
	RTA_GATEWAY          = 0x5
	RTA_PRIORITY         = 0x6
	RTA_PREFSRC          = 0x7
	RTA_METRICS          = 0x8
	RTA_MULTIPATH        = 0x9
	RTA_FLOW             = 0xb
	RTA_CACHEINFO        = 0xc
	RTA_TABLE            = 0xf
	RTN_UNSPEC           = 0x0
	RTN_UNICAST          = 0x1
	RTN_LOCAL            = 0x2
	RTN_BROADCAST        = 0x3
	RTN_ANYCAST          = 0x4
	RTN_MULTICAST        = 0x5
	RTN_BLACKHOLE        = 0x6
	RTN_UNREACHABLE      = 0x7
	RTN_PROHIBIT         = 0x8
	RTN_THROW            = 0x9
	RTN_NAT              = 0xa
	RTN_XRESOLVE         = 0xb
	RTNLGRP_NONE         = 0x0
	RTNLGRP_LINK         = 0x1
	RTNLGRP_NOTIFY       = 0x2
	RTNLGRP_NEIGH        = 0x3
	RTNLGRP_TC           = 0x4
	RTNLGRP_IPV4_IFADDR  = 0x5
	RTNLGRP_IPV4_MROUTE  = 0x6
	RTNLGRP_IPV4_ROUTE   = 0x7
	RTNLGRP_IPV4_RULE    = 0x8
	RTNLGRP_IPV6_IFADDR  = 0x9
	RTNLGRP_IPV6_MROUTE  = 0xa
	RTNLGRP_IPV6_ROUTE   = 0xb
	RTNLGRP_IPV6_IFINFO  = 0xc
	RTNLGRP_IPV6_PREFIX  = 0x12
	RTNLGRP_IPV6_RULE    = 0x13
	RTNLGRP_ND_USEROPT   = 0x14
	SizeofNlMsghdr       = 0x10
	SizeofNlMsgerr       = 0x14
	SizeofRtGenmsg       = 0x1
	SizeofNlAttr         = 0x4
	SizeofRtAttr         = 0x4
	SizeofIfInfomsg      = 0x10
	SizeofIfAddrmsg      = 0x8
	SizeofRtMsg          = 0xc
	SizeofRtNexthop      = 0x8
)

type NlMsghdr struct {
	Len   uint32
	Type  uint16
	Flags uint16
	Seq   uint32
	Pid   uint32
}

type NlMsgerr struct {
	Error int32
	Msg   NlMsghdr
}

type RtGenmsg struct {
	Family uint8
}

type NlAttr struct {
	Len  uint16
	Type uint16
}

type RtAttr struct {
	Len  uint16
	Type uint16
}

type IfInfomsg struct {
	Family uint8
	_      uint8
	Type   uint16
	Index  int32
	Flags  uint32
	Change uint32
}

type IfAddrmsg struct {
	Family    uint8
	Prefixlen uint8
	Flags     uint8
	Scope     uint8
	Index     uint32
}

type RtMsg struct {
	Family   uint8
	Dst_len  uint8
	Src_len  uint8
	Tos      uint8
	Table    uint8
	Protocol uint8
	Scope    uint8
	Type     uint8
	Flags    uint32
}

type RtNexthop struct {
	Len     uint16
	Flags   uint8
	Hops    uint8
	Ifindex int32
}

const (
	SizeofSockFilter = 0x8
	SizeofSockFprog  = 0x8
)

type SockFilter struct {
	Code uint16
	Jt   uint8
	Jf   uint8
	K    uint32
}

type SockFprog struct {
	Len    uint16
	_      [2]byte
	Filter *SockFilter
}

type InotifyEvent struct {
	Wd     int32
	Mask   uint32
	Cookie uint32
	Len    uint32
}

const SizeofInotifyEvent = 0x10

type PtraceRegs struct {
	Regs     [32]uint64
	Lo       uint64
	Hi       uint64
	Epc      uint64
	Badvaddr uint64
	Status   uint64
	Cause    uint64
}

type FdSet struct {
	Bits [32]int32
}

type Sysinfo_t struct {
	Uptime    int32
	Loads     [3]uint32
	Totalram  uint32
	Freeram   uint32
	Sharedram uint32
	Bufferram uint32
	Totalswap uint32
	Freeswap  uint32
	Procs     uint16
	Pad       uint16
	Totalhigh uint32
	Freehigh  uint32
	Unit      uint32
	_         [8]int8
}

type Utsname struct {
	Sysname    [65]byte
	Nodename   [65]byte
	Release    [65]byte
	Version    [65]byte
	Machine    [65]byte
	Domainname [65]byte
}

type Ustat_t struct {
	Tfree  int32
	Tinode uint32
	Fname  [6]int8
	Fpack  [6]int8
}

type EpollEvent struct {
	Events uint32
	PadFd  int32
	Fd     int32
	Pad    int32
}

const (
	AT_EMPTY_PATH   = 0x1000
	AT_FDCWD        = -0x64
	AT_NO_AUTOMOUNT = 0x800
	AT_REMOVEDIR    = 0x200

	AT_STATX_SYNC_AS_STAT = 0x0
	AT_STATX_FORCE_SYNC   = 0x2000
	AT_STATX_DONT_SYNC    = 0x4000

	AT_SYMLINK_FOLLOW   = 0x400
	AT_SYMLINK_NOFOLLOW = 0x100
)

type PollFd struct {
	Fd      int32
	Events  int16
	Revents int16
}

const (
	POLLIN    = 0x1
	POLLPRI   = 0x2
	POLLOUT   = 0x4
	POLLRDHUP = 0x2000
	POLLERR   = 0x8
	POLLHUP   = 0x10
	POLLNVAL  = 0x20
)

type Sigset_t struct {
	Val [32]uint32
}

const RNDGETENTCNT = 0x40045200

const PERF_IOC_FLAG_GROUP = 0x1

type Termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Line   uint8
	Cc     [23]uint8
	Ispeed uint32
	Ospeed uint32
}

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

type Taskstats struct {
	Version                   uint16
	_                         [2]byte
	Ac_exitcode               uint32
	Ac_flag                   uint8
	Ac_nice                   uint8
	_                         [6]byte
	Cpu_count                 uint64
	Cpu_delay_total           uint64
	Blkio_count               uint64
	Blkio_delay_total         uint64
	Swapin_count              uint64
	Swapin_delay_total        uint64
	Cpu_run_real_total        uint64
	Cpu_run_virtual_total     uint64
	Ac_comm                   [32]int8
	Ac_sched                  uint8
	Ac_pad                    [3]uint8
	_                         [4]byte
	Ac_uid                    uint32
	Ac_gid                    uint32
	Ac_pid                    uint32
	Ac_ppid                   uint32
	Ac_btime                  uint32
	_                         [4]byte
	Ac_etime                  uint64
	Ac_utime                  uint64
	Ac_stime                  uint64
	Ac_minflt                 uint64
	Ac_majflt                 uint64
	Coremem                   uint64
	Virtmem                   uint64
	Hiwater_rss               uint64
	Hiwater_vm                uint64
	Read_char                 uint64
	Write_char                uint64
	Read_syscalls             uint64
	Write_syscalls            uint64
	Read_bytes                uint64
	Write_bytes               uint64
	Cancelled_write_bytes     uint64
	Nvcsw                     uint64
	Nivcsw                    uint64
	Ac_utimescaled            uint64
	Ac_stimescaled            uint64
	Cpu_scaled_run_real_total uint64
	Freepages_count           uint64
	Freepages_delay_total     uint64
}

const (
	TASKSTATS_CMD_UNSPEC                  = 0x0
	TASKSTATS_CMD_GET                     = 0x1
	TASKSTATS_CMD_NEW                     = 0x2
	TASKSTATS_TYPE_UNSPEC                 = 0x0
	TASKSTATS_TYPE_PID                    = 0x1
	TASKSTATS_TYPE_TGID                   = 0x2
	TASKSTATS_TYPE_STATS                  = 0x3
	TASKSTATS_TYPE_AGGR_PID               = 0x4
	TASKSTATS_TYPE_AGGR_TGID              = 0x5
	TASKSTATS_TYPE_NULL                   = 0x6
	TASKSTATS_CMD_ATTR_UNSPEC             = 0x0
	TASKSTATS_CMD_ATTR_PID                = 0x1
	TASKSTATS_CMD_ATTR_TGID               = 0x2
	TASKSTATS_CMD_ATTR_REGISTER_CPUMASK   = 0x3
	TASKSTATS_CMD_ATTR_DEREGISTER_CPUMASK = 0x4
)

type CGroupStats struct {
	Sleeping        uint64
	Running         uint64
	Stopped         uint64
	Uninterruptible uint64
	Io_wait         uint64
}

const (
	CGROUPSTATS_CMD_UNSPEC        = 0x3
	CGROUPSTATS_CMD_GET           = 0x4
	CGROUPSTATS_CMD_NEW           = 0x5
	CGROUPSTATS_TYPE_UNSPEC       = 0x0
	CGROUPSTATS_TYPE_CGROUP_STATS = 0x1
	CGROUPSTATS_CMD_ATTR_UNSPEC   = 0x0
	CGROUPSTATS_CMD_ATTR_FD       = 0x1
)

type Genlmsghdr struct {
	Cmd      uint8
	Version  uint8
	Reserved uint16
}

const (
	CTRL_CMD_UNSPEC            = 0x0
	CTRL_CMD_NEWFAMILY         = 0x1
	CTRL_CMD_DELFAMILY         = 0x2
	CTRL_CMD_GETFAMILY         = 0x3
	CTRL_CMD_NEWOPS            = 0x4
	CTRL_CMD_DELOPS            = 0x5
	CTRL_CMD_GETOPS            = 0x6
	CTRL_CMD_NEWMCAST_GRP      = 0x7
	CTRL_CMD_DELMCAST_GRP      = 0x8
	CTRL_CMD_GETMCAST_GRP      = 0x9
	CTRL_ATTR_UNSPEC           = 0x0
	CTRL_ATTR_FAMILY_ID        = 0x1
	CTRL_ATTR_FAMILY_NAME      = 0x2
	CTRL_ATTR_VERSION          = 0x3
	CTRL_ATTR_HDRSIZE          = 0x4
	CTRL_ATTR_MAXATTR          = 0x5
	CTRL_ATTR_OPS              = 0x6
	CTRL_ATTR_MCAST_GROUPS     = 0x7
	CTRL_ATTR_OP_UNSPEC        = 0x0
	CTRL_ATTR_OP_ID            = 0x1
	CTRL_ATTR_OP_FLAGS         = 0x2
	CTRL_ATTR_MCAST_GRP_UNSPEC = 0x0
	CTRL_ATTR_MCAST_GRP_NAME   = 0x1
	CTRL_ATTR_MCAST_GRP_ID     = 0x2
)

type cpuMask uint32

const (
	_CPU_SETSIZE = 0x400
	_NCPUBITS    = 0x20
)

const (
	BDADDR_BREDR     = 0x0
	BDADDR_LE_PUBLIC = 0x1
	BDADDR_LE_RANDOM = 0x2
)

type PerfEventAttr struct {
	Type               uint32
	Size               uint32
	Config             uint64
	Sample             uint64
	Sample_type        uint64
	Read_format        uint64
	Bits               uint64
	Wakeup             uint32
	Bp_type            uint32
	Ext1               uint64
	Ext2               uint64
	Branch_sample_type uint64
	Sample_regs_user   uint64
	Sample_stack_user  uint32
	Clockid            int32
	Sample_regs_intr   uint64
	Aux_watermark      uint32
	_                  uint32
}

type PerfEventMmapPage struct {
	Version        uint32
	Compat_version uint32
	Lock           uint32
	Index          uint32
	Offset         int64
	Time_enabled   uint64
	Time_running   uint64
	Capabilities   uint64
	Pmc_width      uint16
	Time_shift     uint16
	Time_mult      uint32
	Time_offset    uint64
	Time_zero      uint64
	Size           uint32
	_              [948]uint8
	Data_head      uint64
	Data_tail      uint64
	Data_offset    uint64
	Data_size      uint64
	Aux_head       uint64
	Aux_tail       uint64
	Aux_offset     uint64
	Aux_size       uint64
}

const (
	PerfBitDisabled               uint64 = CBitFieldMaskBit0
	PerfBitInherit                       = CBitFieldMaskBit1
	PerfBitPinned                        = CBitFieldMaskBit2
	PerfBitExclusive                     = CBitFieldMaskBit3
	PerfBitExcludeUser                   = CBitFieldMaskBit4
	PerfBitExcludeKernel                 = CBitFieldMaskBit5
	PerfBitExcludeHv                     = CBitFieldMaskBit6
	PerfBitExcludeIdle                   = CBitFieldMaskBit7
	PerfBitMmap                          = CBitFieldMaskBit8
	PerfBitComm                          = CBitFieldMaskBit9
	PerfBitFreq                          = CBitFieldMaskBit10
	PerfBitInheritStat                   = CBitFieldMaskBit11
	PerfBitEnableOnExec                  = CBitFieldMaskBit12
	PerfBitTask                          = CBitFieldMaskBit13
	PerfBitWatermark                     = CBitFieldMaskBit14
	PerfBitPreciseIPBit1                 = CBitFieldMaskBit15
	PerfBitPreciseIPBit2                 = CBitFieldMaskBit16
	PerfBitMmapData                      = CBitFieldMaskBit17
	PerfBitSampleIDAll                   = CBitFieldMaskBit18
	PerfBitExcludeHost                   = CBitFieldMaskBit19
	PerfBitExcludeGuest                  = CBitFieldMaskBit20
	PerfBitExcludeCallchainKernel        = CBitFieldMaskBit21
	PerfBitExcludeCallchainUser          = CBitFieldMaskBit22
	PerfBitMmap2                         = CBitFieldMaskBit23
	PerfBitCommExec                      = CBitFieldMaskBit24
	PerfBitUseClockID                    = CBitFieldMaskBit25
	PerfBitContextSwitch                 = CBitFieldMaskBit26
)

const (
	PERF_TYPE_HARDWARE   = 0x0
	PERF_TYPE_SOFTWARE   = 0x1
	PERF_TYPE_TRACEPOINT = 0x2
	PERF_TYPE_HW_CACHE   = 0x3
	PERF_TYPE_RAW        = 0x4
	PERF_TYPE_BREAKPOINT = 0x5

	PERF_COUNT_HW_CPU_CYCLES              = 0x0
	PERF_COUNT_HW_INSTRUCTIONS            = 0x1
	PERF_COUNT_HW_CACHE_REFERENCES        = 0x2
	PERF_COUNT_HW_CACHE_MISSES            = 0x3
	PERF_COUNT_HW_BRANCH_INSTRUCTIONS     = 0x4
	PERF_COUNT_HW_BRANCH_MISSES           = 0x5
	PERF_COUNT_HW_BUS_CYCLES              = 0x6
	PERF_COUNT_HW_STALLED_CYCLES_FRONTEND = 0x7
	PERF_COUNT_HW_STALLED_CYCLES_BACKEND  = 0x8
	PERF_COUNT_HW_REF_CPU_CYCLES          = 0x9

	PERF_COUNT_HW_CACHE_L1D  = 0x0
	PERF_COUNT_HW_CACHE_L1I  = 0x1
	PERF_COUNT_HW_CACHE_LL   = 0x2
	PERF_COUNT_HW_CACHE_DTLB = 0x3
	PERF_COUNT_HW_CACHE_ITLB = 0x4
	PERF_COUNT_HW_CACHE_BPU  = 0x5
	PERF_COUNT_HW_CACHE_NODE = 0x6

	PERF_COUNT_HW_CACHE_OP_READ     = 0x0
	PERF_COUNT_HW_CACHE_OP_WRITE    = 0x1
	PERF_COUNT_HW_CACHE_OP_PREFETCH = 0x2

	PERF_COUNT_HW_CACHE_RESULT_ACCESS = 0x0
	PERF_COUNT_HW_CACHE_RESULT_MISS   = 0x1

	PERF_COUNT_SW_CPU_CLOCK        = 0x0
	PERF_COUNT_SW_TASK_CLOCK       = 0x1
	PERF_COUNT_SW_PAGE_FAULTS      = 0x2
	PERF_COUNT_SW_CONTEXT_SWITCHES = 0x3
	PERF_COUNT_SW_CPU_MIGRATIONS   = 0x4
	PERF_COUNT_SW_PAGE_FAULTS_MIN  = 0x5
	PERF_COUNT_SW_PAGE_FAULTS_MAJ  = 0x6
	PERF_COUNT_SW_ALIGNMENT_FAULTS = 0x7
	PERF_COUNT_SW_EMULATION_FAULTS = 0x8
	PERF_COUNT_SW_DUMMY            = 0x9

	PERF_SAMPLE_IP           = 0x1
	PERF_SAMPLE_TID          = 0x2
	PERF_SAMPLE_TIME         = 0x4
	PERF_SAMPLE_ADDR         = 0x8
	PERF_SAMPLE_READ         = 0x10
	PERF_SAMPLE_CALLCHAIN    = 0x20
	PERF_SAMPLE_ID           = 0x40
	PERF_SAMPLE_CPU          = 0x80
	PERF_SAMPLE_PERIOD       = 0x100
	PERF_SAMPLE_STREAM_ID    = 0x200
	PERF_SAMPLE_RAW          = 0x400
	PERF_SAMPLE_BRANCH_STACK = 0x800

	PERF_SAMPLE_BRANCH_USER       = 0x1
	PERF_SAMPLE_BRANCH_KERNEL     = 0x2
	PERF_SAMPLE_BRANCH_HV         = 0x4
	PERF_SAMPLE_BRANCH_ANY        = 0x8
	PERF_SAMPLE_BRANCH_ANY_CALL   = 0x10
	PERF_SAMPLE_BRANCH_ANY_RETURN = 0x20
	PERF_SAMPLE_BRANCH_IND_CALL   = 0x40

	PERF_FORMAT_TOTAL_TIME_ENABLED = 0x1
	PERF_FORMAT_TOTAL_TIME_RUNNING = 0x2
	PERF_FORMAT_ID                 = 0x4
	PERF_FORMAT_GROUP              = 0x8

	PERF_RECORD_MMAP       = 0x1
	PERF_RECORD_LOST       = 0x2
	PERF_RECORD_COMM       = 0x3
	PERF_RECORD_EXIT       = 0x4
	PERF_RECORD_THROTTLE   = 0x5
	PERF_RECORD_UNTHROTTLE = 0x6
	PERF_RECORD_FORK       = 0x7
	PERF_RECORD_READ       = 0x8
	PERF_RECORD_SAMPLE     = 0x9

	PERF_CONTEXT_HV     = -0x20
	PERF_CONTEXT_KERNEL = -0x80
	PERF_CONTEXT_USER   = -0x200

	PERF_CONTEXT_GUEST        = -0x800
	PERF_CONTEXT_GUEST_KERNEL = -0x880
	PERF_CONTEXT_GUEST_USER   = -0xa00

	PERF_FLAG_FD_NO_GROUP = 0x1
	PERF_FLAG_FD_OUTPUT   = 0x2
	PERF_FLAG_PID_CGROUP  = 0x4
)

const (
	CBitFieldMaskBit0  = 0x8000000000000000
	CBitFieldMaskBit1  = 0x4000000000000000
	CBitFieldMaskBit2  = 0x2000000000000000
	CBitFieldMaskBit3  = 0x1000000000000000
	CBitFieldMaskBit4  = 0x800000000000000
	CBitFieldMaskBit5  = 0x400000000000000
	CBitFieldMaskBit6  = 0x200000000000000
	CBitFieldMaskBit7  = 0x100000000000000
	CBitFieldMaskBit8  = 0x80000000000000
	CBitFieldMaskBit9  = 0x40000000000000
	CBitFieldMaskBit10 = 0x20000000000000
	CBitFieldMaskBit11 = 0x10000000000000
	CBitFieldMaskBit12 = 0x8000000000000
	CBitFieldMaskBit13 = 0x4000000000000
	CBitFieldMaskBit14 = 0x2000000000000
	CBitFieldMaskBit15 = 0x1000000000000
	CBitFieldMaskBit16 = 0x800000000000
	CBitFieldMaskBit17 = 0x400000000000
	CBitFieldMaskBit18 = 0x200000000000
	CBitFieldMaskBit19 = 0x100000000000
	CBitFieldMaskBit20 = 0x80000000000
	CBitFieldMaskBit21 = 0x40000000000
	CBitFieldMaskBit22 = 0x20000000000
	CBitFieldMaskBit23 = 0x10000000000
	CBitFieldMaskBit24 = 0x8000000000
	CBitFieldMaskBit25 = 0x4000000000
	CBitFieldMaskBit26 = 0x2000000000
	CBitFieldMaskBit27 = 0x1000000000
	CBitFieldMaskBit28 = 0x800000000
	CBitFieldMaskBit29 = 0x400000000
	CBitFieldMaskBit30 = 0x200000000
	CBitFieldMaskBit31 = 0x100000000
	CBitFieldMaskBit32 = 0x80000000
	CBitFieldMaskBit33 = 0x40000000
	CBitFieldMaskBit34 = 0x20000000
	CBitFieldMaskBit35 = 0x10000000
	CBitFieldMaskBit36 = 0x8000000
	CBitFieldMaskBit37 = 0x4000000
	CBitFieldMaskBit38 = 0x2000000
	CBitFieldMaskBit39 = 0x1000000
	CBitFieldMaskBit40 = 0x800000
	CBitFieldMaskBit41 = 0x400000
	CBitFieldMaskBit42 = 0x200000
	CBitFieldMaskBit43 = 0x100000
	CBitFieldMaskBit44 = 0x80000
	CBitFieldMaskBit45 = 0x40000
	CBitFieldMaskBit46 = 0x20000
	CBitFieldMaskBit47 = 0x10000
	CBitFieldMaskBit48 = 0x8000
	CBitFieldMaskBit49 = 0x4000
	CBitFieldMaskBit50 = 0x2000
	CBitFieldMaskBit51 = 0x1000
	CBitFieldMaskBit52 = 0x800
	CBitFieldMaskBit53 = 0x400
	CBitFieldMaskBit54 = 0x200
	CBitFieldMaskBit55 = 0x100
	CBitFieldMaskBit56 = 0x80
	CBitFieldMaskBit57 = 0x40
	CBitFieldMaskBit58 = 0x20
	CBitFieldMaskBit59 = 0x10
	CBitFieldMaskBit60 = 0x8
	CBitFieldMaskBit61 = 0x4
	CBitFieldMaskBit62 = 0x2
	CBitFieldMaskBit63 = 0x1
)