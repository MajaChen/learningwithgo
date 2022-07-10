package leetcode

import (
	"fmt"
	"hash/crc32"
	"math"
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
type Codec struct {
	num2URL map[int]string
}

const (
	prefix = "https://tinyURL."
	bigNum = math.MaxInt32
)

func String(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	// v == MinInt
	return 0
}

func EncodeConstructor() Codec {
	return Codec{num2URL: make(map[int]string)}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	h := String(longUrl)
	for {
		if _, ok := this.num2URL[h]; ok {
			if this.num2URL[h] == longUrl {
				return fmt.Sprintf("%s%v", prefix, h)
			} else {
				h = (h + 1) % bigNum
			}
		} else {
			this.num2URL[h] = longUrl
			return fmt.Sprintf("%s%v", prefix, h)
		}
	}
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	shortUrl = shortUrl[strings.Index(shortUrl, ".")+1:]
	num, _ := strconv.Atoi(shortUrl)
	return this.num2URL[num]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := RangeRandConstructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
//leetcode submit region end(Prohibit modification and deletion)
