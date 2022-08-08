## GoJSON
`GoJSON` is a fast and easy package to marshal/unmarshal struct to/from json. You can use `GoJSON` tool to generate marshal/unmarshal code, in benchmark tests, `GoJSON`'s generate code is almost 6~7 times faster than `encoding/json`.

## Example 
```golang
# install
  go get -u -v github.com/go-fish/gojson
  cd ${GOPATH}/src/github.com/go-fish/gojson/cmd/gojson
  go build -o gojson main.go


# usage
  gojson [flags] [file|directory]...

Flags:
      --escapehtml      escape html char when encode object to json
      --escapeunicode   escape unicode rune when decode object to json
  -h, --help            help for gojson
  -o, --output string   the filename of output file (default "gojson.generated.go")
      --unsafe          use input data directly without copy when decode json bytes (default true)
        
```

For expose structs, gojson will generate `MarshalJSON/UnmarshalJSON` methods for marshal/unmarshal json. You also can use `gojson.Marshal/gojson.Unmarshal` functions to marshal/unmarshal json.
If your field contains many escaped char, add `escape` to your field tag, it will improve preformance when unmarshal escape string.

## Benchmark
### Large Payload
#### Unmarshal
<table>
	<tr>
		<th>gojson</th>
		<th>91022 ns/op</th>
		<th>576.61 MB/s</th>
		<th>62393 B/op</th>
		<th>200 allocs/op</th>
	</tr>
	<tr>
		<th>jsonparser</th>
		<th>91804 ns/op</th>
		<th>571.70 MB/s</th>
		<th>57344 B/op</th>
		<th>1 allocs/op</th>
	</tr>
	<tr>
		<th>easyjson</th>
		<th>157800 ns/op</th>
		<th>332.60 MB/s</th>
		<th>65388 B/op</th>
		<th>286 allocs/op</th>
	</tr>
	<tr>
		<th>gojay</th>
		<th>84526 ns/op</th>
		<th>620.92 MB/s</th>
		<th>62393 B/op</th>
		<th>200 allocs/op</th>
	</tr>
	<tr>
		<th>jsoniter</th>
		<th>142165 ns/op</th>
		<th>369.18 MB/s</th>
		<th>81216 B/op</th>
		<th>2162 allocs/op</th>
	</tr>
	<tr>
		<th>encoding/json</th>
		<th>437117 ns/op</th>
		<th>120.07 MB/s</th>
		<th>62657 B/op</th>
		<th>205 allocs/op</th>
	</tr>
</table>

#### Marshal
<table>
	<tr>
		<th>gojson</th>
		<th>238.5 ns/op</th>
		<th>220076.86 MB/s</th>
		<th>1048 B/op</th>
		<th>2 allocs/op</th>
	</tr>
	<tr>
		<th>easyjson</th>
		<th>235 ns/op</th>
		<th>223039.53 MB/s</th>
		<th>272 B/op</th>
		<th>3 allocs/op</th>
	</tr>
	<tr>
		<th>encoding/json</th>
		<th>1232 ns/op</th>
		<th>42610.66 MB/s</th>
		<th>1369 B/op</th>
		<th>10 allocs/op</th>
	</tr>
</table>

### Medium Payload
#### Unmarshal

<table>
	<tr>
		<th>gojson</th>
		<th>3649 ns/op</th>
		<th>592.74 MB/s</th>
		<th>2457 B/op</th>
		<th>10 allocs/op</th>
	</tr>
	<tr>
		<th>jsonparser</th>
		<th>6861 ns/op</th>
		<th>315.25 MB/s</th>
		<th>2304 B/op</th>
		<th>1 allocs/op</th>
	</tr>
	<tr>
		<th>easyjson</th>
		<th>8951 ns/op</th>
		<th>241.64 MB/s</th>
		<th>2569 B/op</th>
		<th>11 allocs/op</th>
	</tr>
	<tr>
		<th>gojay</th>
		<th>4145 ns/op</th>
		<th>521.79 MB/s</th>
		<th>2472 B/op</th>
		<th>9 allocs/op</th>
	</tr>
	<tr>
		<th>jsoniter</th>
		<th>14565 ns/op</th>
		<th>148.51 MB/s</th>
		<th>5340 B/op</th>
		<th>86 allocs/op</th>
	</tr>
	<tr>
		<th>encoding/json</th>
		<th>28891 ns/op</th>
		<th>74.87 MB/s</th>
		<th>2721 B/op</th>
		<th>15 allocs/op</th>
	</tr>
</table>

#### Marshal
<table>
	<tr>
		<th>gojson</th>
		<th>282.5 ns/op</th>
		<th>7656.48 MB/s</th>
		<th>1048 B/op</th>
		<th>2 allocs/op</th>
	</tr>
	<tr>
		<th>easyjson</th>
		<th>155.3 ns/op</th>
		<th>13926.91 MB/s</th>
		<th>264 B/op</th>
		<th>3 allocs/op</th>
	</tr>
	<tr>
		<th>encoding/json</th>
		<th>1177 ns/op</th>
		<th>1837.14 MB/s</th>
		<th>1369 B/op</th>
		<th>10 allocs/op</th>
	</tr>
</table>

### Small Payload
#### Unmarshal

<table>
	<tr>
		<th>gojson</th>
		<th>598.3 ns/op</th>
		<th>239.01 MB/s</th>
		<th>296 B/op</th>
		<th>5 allocs/op</th>
	</tr>
	<tr>
		<th>jsonparser</th>
		<th>701.1 ns/op</th>
		<th>203.97 MB/s</th>
		<th>144 B/op</th>
		<th>1 allocs/op</th>
	</tr>
	<tr>
		<th>easyjson</th>
		<th>1058 ns/op</th>
		<th>135.19 MB/s</th>
		<th>272 B/op</th>
		<th>8 allocs/op</th>
	</tr>
	<tr>
		<th>gojay</th>
		<th>606.3 ns/op</th>
		<th>235.87 MB/s</th>
		<th>296 B/op</th>
		<th>5 allocs/op</th>
	</tr>
	<tr>
		<th>jsoniter</th>
		<th>1427 ns/op</th>
		<th>100.22 MB/s</th>
		<th>504 B/op</th>
		<th>16 allocs/op</th>
	</tr>
	<tr>
		<th>encoding/json</th>
		<th>1873 ns/op</th>
		<th>76.35 MB/s</th>
		<th>448 B/op</th>
		<th>7 allocs/op</th>
	</tr>
</table>

#### Marshal
<table>
	<tr>
		<th>gojson</th>
		<th>388.1 ns/op</th>
		<th>368.49 MB/s</th>
		<th>1048 B/op</th>
		<th>2 allocs/op</th>
	</tr>
	<tr>
		<th>easyjson</th>
		<th>288.2 ns/op</th>
		<th>496.20 MB/s</th>
		<th>264 B/op</th>
		<th>3 allocs/op</th>
	</tr>
	<tr>
		<th>encoding/json</th>
		<th>1009 ns/op</th>
		<th>141.75 MB/s</th>
		<th>1129 B/op</th>
		<th>4 allocs/op</th>
	</tr>
</table>

## Questions
Any questions or bugs can go though github issues.