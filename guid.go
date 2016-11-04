package guid

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

//guid comes from snowflake but made some changes
// 1、39bits timestamp, pump id for about 10yeas; 12bits sequences, generate 4096 sequence per Millisecond;12bits workerid, can deploy 4096 servers; highest 1bit reserve;
// 2、workerid in high position in order to make id global increment;
// 3、in order to make id be more hashed, random sequence from [0,10) where millisecond change,not set sequence to 0;
const (
	workerIDBits   = uint64(12)
	sequenceBits   = uint64(12)
	timestampShift = sequenceBits
	workerIDShift  = uint64(63) - workerIDBits
	sequenceMask   = int64(-1) ^ (int64(-1) << sequenceBits)

	twepoch = int64(1407875197154)
)

var TimeBackwardsErr = errors.New("time has gone backward")
var IDBackwardsErr = errors.New("ID went backward")

type IDFactory struct {
	workerID      int64
	sequence      int64
	lastTimestamp int64
	lastID        int64
	mutex         *sync.Mutex
}

func (i *IDFactory) Sequence() int64 {
	return i.sequence
}

func NewIDFactory(workerID int64) *IDFactory {
	factory := &IDFactory{
		workerID:      workerID,
		sequence:      0,
		lastTimestamp: 0,
		lastID:        0,
		mutex:         new(sync.Mutex),
	}
	return factory
}

func getTimestamp() int64 {
	return time.Now().UnixNano() >> 20
}

func nextTimestamp(lastTimestamp int64) int64 {
	timestamp := getTimestamp()
	for timestamp <= lastTimestamp {
		timestamp = getTimestamp()
	}

	return timestamp
}

func (i *IDFactory) IdPump() (int64, error) {
	i.mutex.Lock()

	timestamp := getTimestamp()
	if timestamp < i.lastTimestamp {
		i.mutex.Unlock()
		return 0, TimeBackwardsErr
	}

	if timestamp == i.lastTimestamp {
		i.sequence = (i.sequence + 1) & sequenceMask
		if i.sequence == 0 {
			timestamp = nextTimestamp(timestamp)
		}
	} else {
		rand.Seed(time.Now().UnixNano())
		i.sequence = rand.Int63n(10)
	}

	i.lastTimestamp = timestamp

	id := (i.workerID << workerIDShift) | ((timestamp - twepoch) << timestampShift) | i.sequence

	if id <= i.lastID {
		i.mutex.Unlock()
		return 0, IDBackwardsErr
	}

	i.lastID = id

	i.mutex.Unlock()
	return id, nil
}
