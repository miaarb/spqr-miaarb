package main

type CoordinatorInterface interface {
	initKeyRanges() (map[Shard][]KeyRange, error)
	isReloadRequired() (bool, error)

	lockKeyRange(rng KeyRange) error
	unlockKeyRange(rng KeyRange) error

	splitKeyRange(border *string) error
	mergeKeyRanges(border *string) error
	moveKeyRange(rng KeyRange, shardFrom, shardTo Shard) error
}