package blockchain

import (
	"errors"
)

//Err Bad Block Cache Nil
var ErrBadBlocksCacheNil = errors.New("badBlockCache nil")

//Err Header Unit Nil
var ErrHeaderUnitNil = errors.New("header unit nil")

//Err Wrong Type In Set
var ErrWrongTypeInSet = errors.New(" wrong type in setter")

//Err Nil App Status Handler
var ErrNilAppStatusHanler = errors.New("nil AppStatusHandler")