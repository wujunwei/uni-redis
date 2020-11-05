package client

type Flag string

const (
	Write         Flag = "write"           // command may result in modifications
	Readonly      Flag = "readonly"        // command will never modify keys
	DenyOOM       Flag = "denyoom"         // reject command if currently OOM
	Admin         Flag = "admin"           // server admin command
	Pubsub        Flag = "pubsub"          // pubsub-related command
	Noscript      Flag = "noscript"        // deny this command from scripts
	Random        Flag = "random"          // command has random results, dangerous for scripts
	SortForScript Flag = "sort_for_script" // if called from script, sort output
	Loading       Flag = "loading"         // allow command while database is loading
	Stale         Flag = "stale"           // allow command while replica has stale data
	SkipMonitor   Flag = "skip_monitor"    // do not show this command in MONITOR
	Asking        Flag = "asking"          // cluster related - accept even if importing
	Fast          Flag = "fast"            // command operates in constant or log(N) time. Used for latency monitoring.
	MovableKeys   Flag = "movablekeys"     // keys have no pre-determined position. You must discover keys yourself.
)

type Command struct {
	Name                       string
	FixArity                   bool
	ArityCount                 int
	firstKey, lastKey, keyStep int
	Flags                      map[Flag]bool
}

func NewCommand(name string, arityCount, firstKey int, lastKey int, keyStep int, flags []string) *Command {
	fixArity := true
	if arityCount < 0 {
		fixArity = false
	}
	var temp = make(map[Flag]bool)
	for _, flag := range flags {
		temp[Flag(flag)] = true
	}
	return &Command{
		Name:       name,
		FixArity:   fixArity,
		ArityCount: arityCount,
		firstKey:   firstKey,
		lastKey:    lastKey,
		keyStep:    keyStep,
		Flags:      temp,
	}
}

func (c Command) String() string {
	return c.Name
}

func (c Command) KeysMovable() bool {
	return c.Flags[MovableKeys]
}

func (c Command) IsFast() bool {
	return c.Flags[Fast]
}

func (c Command) Writable() bool {
	return c.Flags[Write]
}

func (c Command) Readonly() bool {
	return c.Flags[Readonly]
}

func (c Command) IsRandom() bool {
	return c.Flags[Random]
}

func (c Command) DenyWhenOOM() bool {
	return c.Flags[DenyOOM]
}

func (c Command) CanSkipMonitor() bool {
	return c.Flags[SkipMonitor]
}
