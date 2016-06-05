package http_client

import (
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

// cd $GOPATH/src/github.com/konjoot/go-cookbook
// go test ./http_client/...   -bench=.  -cpu=1,2,4,8,16 -count=10
// testing: warning: no tests to run
// PASS
// BenchmarkDefaultHTTPClient   	   20000	     94395 ns/op
// BenchmarkDefaultHTTPClient   	   20000	     93106 ns/op
// BenchmarkDefaultHTTPClient   	   20000	     94232 ns/op
// BenchmarkDefaultHTTPClient   	   20000	     93875 ns/op
// BenchmarkDefaultHTTPClient   	   20000	     93952 ns/op
// BenchmarkDefaultHTTPClient   	   20000	     93747 ns/op
// BenchmarkDefaultHTTPClient   	   20000	     94513 ns/op
// BenchmarkDefaultHTTPClient   	   20000	     91912 ns/op
// BenchmarkDefaultHTTPClient   	   20000	     94323 ns/op
// BenchmarkDefaultHTTPClient   	   10000	    101954 ns/op
// BenchmarkDefaultHTTPClient-2 	   30000	     49202 ns/op
// BenchmarkDefaultHTTPClient-2 	   30000	     49978 ns/op
// BenchmarkDefaultHTTPClient-2 	   30000	     51742 ns/op
// BenchmarkDefaultHTTPClient-2 	   30000	     56796 ns/op
// BenchmarkDefaultHTTPClient-2 	   30000	     50957 ns/op
// BenchmarkDefaultHTTPClient-2 	   30000	     54633 ns/op
// BenchmarkDefaultHTTPClient-2 	   30000	     55838 ns/op
// BenchmarkDefaultHTTPClient-2 	   30000	     51552 ns/op
// BenchmarkDefaultHTTPClient-2 	   30000	     55721 ns/op
// BenchmarkDefaultHTTPClient-2 	   30000	     50069 ns/op
// BenchmarkDefaultHTTPClient-4 	    5000	    254246 ns/op
// BenchmarkDefaultHTTPClient-4 	    5000	    244609 ns/op
// BenchmarkDefaultHTTPClient-4 	    5000	    226438 ns/op
// BenchmarkDefaultHTTPClient-4 	    5000	    247377 ns/op
// BenchmarkDefaultHTTPClient-4 	    5000	    263824 ns/op
// BenchmarkDefaultHTTPClient-4 	    5000	    285493 ns/op
// BenchmarkDefaultHTTPClient-4 	   10000	    238567 ns/op
// BenchmarkDefaultHTTPClient-4 	   10000	    283335 ns/op
// BenchmarkDefaultHTTPClient-4 	    5000	    305496 ns/op
// BenchmarkDefaultHTTPClient-4 	    5000	    276740 ns/op
// BenchmarkDefaultHTTPClient-8 	    3000	    539342 ns/op
// BenchmarkDefaultHTTPClient-8 	    3000	    483940 ns/op
// BenchmarkDefaultHTTPClient-8 	    3000	    486770 ns/op
// BenchmarkDefaultHTTPClient-8 	    3000	    558159 ns/op
// BenchmarkDefaultHTTPClient-8 	    2000	    535347 ns/op
// BenchmarkDefaultHTTPClient-8 	    5000	    499869 ns/op
// BenchmarkDefaultHTTPClient-8 	    5000	    516867 ns/op
// BenchmarkDefaultHTTPClient-8 	    2000	    577805 ns/op
// BenchmarkDefaultHTTPClient-8 	    3000	    500532 ns/op
// BenchmarkDefaultHTTPClient-8 	    3000	    525765 ns/op
// BenchmarkDefaultHTTPClient-16	    1000	   1341974 ns/op
// BenchmarkDefaultHTTPClient-16	    1000	   1449688 ns/op
// BenchmarkDefaultHTTPClient-16	    1000	   1203171 ns/op
// BenchmarkDefaultHTTPClient-16	    2000	   1279075 ns/op
// BenchmarkDefaultHTTPClient-16	    1000	   1146673 ns/op
// BenchmarkDefaultHTTPClient-16	    1000	   1142640 ns/op
// BenchmarkDefaultHTTPClient-16	    1000	   1300106 ns/op
// BenchmarkDefaultHTTPClient-16	   10000	   1243376 ns/op
// BenchmarkDefaultHTTPClient-16	    1000	   1205065 ns/op
// BenchmarkDefaultHTTPClient-16	    1000	   1184279 ns/op
// BenchmarkPooledHTTPClient    	   20000	     95149 ns/op
// BenchmarkPooledHTTPClient    	   20000	     93536 ns/op
// BenchmarkPooledHTTPClient    	   20000	     93829 ns/op
// BenchmarkPooledHTTPClient    	   20000	     94768 ns/op
// BenchmarkPooledHTTPClient    	   10000	    102743 ns/op
// BenchmarkPooledHTTPClient    	   20000	    100482 ns/op
// BenchmarkPooledHTTPClient    	   10000	    104948 ns/op
// BenchmarkPooledHTTPClient    	   20000	     99934 ns/op
// BenchmarkPooledHTTPClient    	   10000	    104177 ns/op
// BenchmarkPooledHTTPClient    	   10000	    106969 ns/op
// BenchmarkPooledHTTPClient-2  	   30000	     59473 ns/op
// BenchmarkPooledHTTPClient-2  	   20000	     65686 ns/op
// BenchmarkPooledHTTPClient-2  	   30000	     60925 ns/op
// BenchmarkPooledHTTPClient-2  	   30000	     55351 ns/op
// BenchmarkPooledHTTPClient-2  	   30000	     60148 ns/op
// BenchmarkPooledHTTPClient-2  	   30000	     61460 ns/op
// BenchmarkPooledHTTPClient-2  	   30000	     57936 ns/op
// BenchmarkPooledHTTPClient-2  	   30000	     57296 ns/op
// BenchmarkPooledHTTPClient-2  	   30000	     61689 ns/op
// BenchmarkPooledHTTPClient-2  	   30000	     63440 ns/op
// BenchmarkPooledHTTPClient-4  	   30000	     40083 ns/op
// BenchmarkPooledHTTPClient-4  	   50000	     40438 ns/op
// BenchmarkPooledHTTPClient-4  	   30000	     41514 ns/op
// BenchmarkPooledHTTPClient-4  	   30000	     43262 ns/op
// BenchmarkPooledHTTPClient-4  	   30000	     43740 ns/op
// BenchmarkPooledHTTPClient-4  	   50000	     40010 ns/op
// BenchmarkPooledHTTPClient-4  	   30000	     40896 ns/op
// BenchmarkPooledHTTPClient-4  	   30000	     42291 ns/op
// BenchmarkPooledHTTPClient-4  	   30000	     39033 ns/op
// BenchmarkPooledHTTPClient-4  	   30000	     45751 ns/op
// BenchmarkPooledHTTPClient-8  	   50000	     36175 ns/op
// BenchmarkPooledHTTPClient-8  	   50000	     33608 ns/op
// BenchmarkPooledHTTPClient-8  	   50000	     34716 ns/op
// BenchmarkPooledHTTPClient-8  	   50000	     34970 ns/op
// BenchmarkPooledHTTPClient-8  	   50000	     36101 ns/op
// BenchmarkPooledHTTPClient-8  	   50000	     33794 ns/op
// BenchmarkPooledHTTPClient-8  	   50000	     35223 ns/op
// BenchmarkPooledHTTPClient-8  	   50000	     33354 ns/op
// BenchmarkPooledHTTPClient-8  	   30000	     37057 ns/op
// BenchmarkPooledHTTPClient-8  	   50000	     34559 ns/op
// BenchmarkPooledHTTPClient-16 	   30000	     37205 ns/op
// BenchmarkPooledHTTPClient-16 	   50000	     38523 ns/op
// BenchmarkPooledHTTPClient-16 	   50000	     36726 ns/op
// BenchmarkPooledHTTPClient-16 	   50000	     39020 ns/op
// BenchmarkPooledHTTPClient-16 	   50000	     39009 ns/op
// BenchmarkPooledHTTPClient-16 	   50000	     38578 ns/op
// BenchmarkPooledHTTPClient-16 	   50000	     37310 ns/op
// BenchmarkPooledHTTPClient-16 	   50000	     39430 ns/op
// BenchmarkPooledHTTPClient-16 	   50000	     39202 ns/op
// BenchmarkPooledHTTPClient-16 	   50000	     37971 ns/op
// ok  	github.com/konjoot/go-cookbook/http_client	221.322s

func init() {
	<-Srv()
}

func BenchmarkDefaultHTTPClient(b *testing.B) {
	url := "https://127.0.0.1:10443/hello"

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				b.Error(err)
			}

			resp, err := defaultClient.Do(req)
			if err == nil {
				io.Copy(ioutil.Discard, resp.Body)
				resp.Body.Close()
			} else {
				b.Error(err)
			}
		}
	})
}

func BenchmarkPooledHTTPClient(b *testing.B) {
	url := "https://127.0.0.1:10443/hello"

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				b.Error(err)
			}

			resp, err := pooledClient.Do(req)
			if err == nil {
				io.Copy(ioutil.Discard, resp.Body)
				resp.Body.Close()
			} else {
				b.Error(err)
			}
		}
	})
}
