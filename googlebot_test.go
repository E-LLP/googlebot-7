package googlebot

import (
	"net"
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type TSuite struct{}

var _ = Suite(&TSuite{})

func (s *TSuite) TestGoogleBot(c *C) {
	for n, addr := range testData {
		if count := strings.Count(addr, "."); count <= 2 {
			c.Fatalf("invalid test data (line:%d)", n+1)
		}
		c.Log(addr)
		yes, err := IsGoogleBotWithCache(addr)
		switch err.(type) {
		case *net.DNSError:
			c.Logf("dns error: %v", err)
			continue
		}

		c.Assert(err, IsNil)
		c.Assert(yes, Equals, true)
	}
}

var testData = strings.Split(`216.239.33.96
216.239.33.97
216.239.33.98
216.239.33.99
216.239.37.98
216.239.37.99
216.239.39.98
216.239.39.99
216.239.41.96
216.239.41.97
216.239.41.98
216.239.41.99
216.239.51.96
216.239.51.97
216.239.51.98
216.239.51.99
216.239.53.98
216.239.53.99
216.239.57.96
216.239.57.97
216.239.57.98
216.239.57.99
216.239.59.98
216.239.59.99
216.33.229.163
64.68.90.1
64.68.90.10
64.68.90.11
64.68.90.12
64.68.90.129
64.68.90.13
64.68.90.130
64.68.90.131
64.68.90.132
64.68.90.133
64.68.90.134
64.68.90.135
64.68.90.136
64.68.90.137
64.68.90.138
64.68.90.139
64.68.90.14
64.68.90.140
64.68.90.141
64.68.90.142
64.68.90.143
64.68.90.144
64.68.90.145
64.68.90.146
64.68.90.147
64.68.90.148
64.68.90.149
64.68.90.15
64.68.90.150
64.68.90.151
64.68.90.152
64.68.90.153
64.68.90.154
64.68.90.155
64.68.90.156
64.68.90.157
64.68.90.158
64.68.90.159
64.68.90.16
64.68.90.160
64.68.90.161
64.68.90.162
64.68.90.163
64.68.90.164
64.68.90.165
64.68.90.166
64.68.90.167
64.68.90.168
64.68.90.169
64.68.90.17
64.68.90.170
64.68.90.171
64.68.90.172
64.68.90.173
64.68.90.174
64.68.90.175
64.68.90.18
64.68.90.180
64.68.90.181
64.68.90.182
64.68.90.183
64.68.90.184
64.68.90.185
64.68.90.186
64.68.90.187
64.68.90.188
64.68.90.189
64.68.90.19
64.68.90.190
64.68.90.191
64.68.90.192
64.68.90.193
64.68.90.194
64.68.90.195
64.68.90.196
64.68.90.197
64.68.90.198
64.68.90.199
64.68.90.2
64.68.90.20
64.68.90.200
64.68.90.201
64.68.90.202
64.68.90.203
64.68.90.204
64.68.90.205
64.68.90.206
64.68.90.207
64.68.90.208
64.68.90.21
64.68.90.22
64.68.90.23
64.68.90.24
64.68.90.25
64.68.90.26
64.68.90.27
64.68.90.28
64.68.90.29
64.68.90.3
64.68.90.30
64.68.90.31
64.68.90.32
64.68.90.33
64.68.90.34
64.68.90.35
64.68.90.36
64.68.90.37
64.68.90.38
64.68.90.39
64.68.90.4
64.68.90.40
64.68.90.41
64.68.90.42
64.68.90.43
64.68.90.44
64.68.90.45
64.68.90.46
64.68.90.47
64.68.90.48
64.68.90.49
64.68.90.5
64.68.90.50
64.68.90.51
64.68.90.52
64.68.90.53
64.68.90.54
64.68.90.55
64.68.90.56
64.68.90.57
64.68.90.58
64.68.90.59
64.68.90.6
64.68.90.60
64.68.90.61
64.68.90.62
64.68.90.63
64.68.90.64
64.68.90.65
64.68.90.66
64.68.90.67
64.68.90.68
64.68.90.69
64.68.90.7
64.68.90.70
64.68.90.71
64.68.90.72
64.68.90.73
64.68.90.74
64.68.90.75
64.68.90.76
64.68.90.77
64.68.90.78
64.68.90.79
64.68.90.8
64.68.90.80
64.68.90.9
66.249.76.196`, "\n")
