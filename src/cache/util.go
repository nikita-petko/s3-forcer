package cache

import (
	"fmt"
	"hash/fnv"
	"strconv"

	"github.com/nikita-petko/s3-forcer/flags"
)

func computeCacheFileName() string {
	hash := fnv.New32a()
	hash.Write([]byte(fmt.Sprintf("%s_%s", *flags.ChannelPrefix, *flags.CharCombinations)))

	return ".cache." + strconv.Itoa(int(hash.Sum32()))
}
