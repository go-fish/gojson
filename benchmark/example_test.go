package benchmark

import (
	json "encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/buger/jsonparser"
	"github.com/francoispqt/gojay"
	"github.com/go-fish/gojson"

	//"github.com/go-fish/gojson"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func nothing(_ ...interface{}) {}

var smallObject SmallPayload
var mediumObject MediumPayload
var largeObject LargePayload

func smartPrint(i interface{}) {
	var kv = make(map[string]interface{})
	vValue := reflect.ValueOf(i)
	vType := reflect.TypeOf(i)
	for i := 0; i < vValue.NumField(); i++ {
		kv[vType.Field(i).Name] = vValue.Field(i)
	}
	fmt.Println("获取到数据:")
	for k, v := range kv {
		fmt.Print(k)
		fmt.Print(":")
		fmt.Printf("%#v", v)
		fmt.Println()
	}
}

func genLargeTestData() []byte {
	return []byte(`{"@timestamp":"2019-12-14T02:18:17.554Z","@metadata":{"beat":"packetbeat","type":"_doc","version":"7.0.1","topic":"LB_HTTP"},"ecs":{"version":"1.0.0"},"host":{"name":"1.1.1.1"},"server":{"ip":"1.1.1.1","port":80,"domain":"somedomain","bytes":29533},"status":"OK","source":{"bytes":579,"ip":"1.1.1.1","port":14261},"method":"post","http":{"response":{"status_code":200,"body":{"content":"{\"cities\":[{\"id\":1,\"attLeaderIcon\":0,\"st\":\"IDLE\",\"avoidEndT\":0,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":2,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574386812,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":3,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574398041,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":4,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574387143,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":5,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574393656,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":6,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574485070,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":7,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574485651,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":8,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574395866,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":9,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574559403,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":10,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574385516,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":11,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"六角形\",\"flag\":\"陸,5,0,4\",\"st\":\"IDLE\",\"avoidEndT\":1574386927,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":12,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574386200,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":13,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574394923,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":14,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574400327,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":15,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574389086,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":17,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574773639,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":16,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574768089,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":19,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574487027,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":18,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574393715,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":21,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"天上人间\",\"flag\":\"天,1,1,2\",\"st\":\"IDLE\",\"avoidEndT\":1574388359,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":20,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574474595,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":23,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574555765,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":22,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574474702,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":25,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"天上人间\",\"flag\":\"天,1,1,2\",\"st\":\"IDLE\",\"avoidEndT\":1574475201,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":24,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574469380,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":27,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574393120,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":26,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574508259,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":29,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"席卷天下\",\"flag\":\"人,5,11,0\",\"st\":\"IDLE\",\"avoidEndT\":1574478708,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":28,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574472059,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":31,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574505053,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":30,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574384557,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":34,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574408868,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":35,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574393571,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":32,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574387777,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":33,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"兄弟一九一\",\"flag\":\"兄,0,1,2\",\"st\":\"IDLE\",\"avoidEndT\":1574471407,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":38,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574385401,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":39,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574657961,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":36,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574562433,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":37,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574558139,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":42,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574474157,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":43,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"天上人间\",\"flag\":\"天,1,1,2\",\"st\":\"IDLE\",\"avoidEndT\":1574683798,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":40,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574994449,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":41,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574390398,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":46,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574559854,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":47,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574385910,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":44,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574572162,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":45,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574563515,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":51,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574473330,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":50,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1575078427,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":49,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574471864,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":48,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574475369,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":55,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"六角形\",\"flag\":\"陸,5,0,4\",\"st\":\"IDLE\",\"avoidEndT\":1574593057,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":54,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574560460,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":53,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574396810,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":52,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"席卷天下\",\"flag\":\"人,5,11,0\",\"st\":\"IDLE\",\"avoidEndT\":1574564003,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":59,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"六角形\",\"flag\":\"陸,5,0,4\",\"st\":\"IDLE\",\"avoidEndT\":1574488231,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":58,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574390736,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":57,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574484328,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":56,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574568330,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":63,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574989334,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":62,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574488130,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":61,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574675842,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":60,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"六角形\",\"flag\":\"陸,5,0,4\",\"st\":\"IDLE\",\"avoidEndT\":1574474661,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":68,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574590902,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":69,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"兄弟一九一\",\"flag\":\"兄,0,1,2\",\"st\":\"IDLE\",\"avoidEndT\":1574484920,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":70,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574484982,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":71,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"娱乐厅\",\"flag\":\"王,9,1,2\",\"st\":\"IDLE\",\"avoidEndT\":1574488367,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":64,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574567890,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":65,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574574690,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":66,\"attServer\":\"服务器（1-470）\",\"attCorps\":\"迷糊三国\",\"attFlag\":\"迷,0,0,0\",\"attLeaderIcon\":600004,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"PREPARING\",\"avoidEndT\":1574474581,\"startT\":1576290606,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":67,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574489432,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":76,\"attLeaderIcon\":0,\"st\":\"IDLE\",\"avoidEndT\":0,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":77,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574646853,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":78,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"征战三国.S436\",\"flag\":\"征,2,4,5\",\"st\":\"IDLE\",\"avoidEndT\":1574568238,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":79,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574487772,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":72,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574412546,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":73,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574470835,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":74,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574387506,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":75,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574398219,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":85,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574745407,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":84,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574390607,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":87,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"天赐\",\"flag\":\"國,5,2,3\",\"st\":\"IDLE\",\"avoidEndT\":1575111267,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":86,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574561151,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":81,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574687692,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":80,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"六角形\",\"flag\":\"陸,5,0,4\",\"st\":\"IDLE\",\"avoidEndT\":1574484171,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":83,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574570518,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":82,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574410559,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":93,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574473668,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":92,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574486100,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":95,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"天上人间\",\"flag\":\"天,1,1,2\",\"st\":\"IDLE\",\"avoidEndT\":1574393782,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":94,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574485937,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":89,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574399362,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":88,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574385138,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":91,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574513454,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":90,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574991597,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":102,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574394205,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":103,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574394398,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":100,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574485965,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":101,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574557230,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":98,\"attLeaderIcon\":0,\"st\":\"IDLE\",\"avoidEndT\":0,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":99,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"天上人间\",\"flag\":\"天,1,1,2\",\"st\":\"IDLE\",\"avoidEndT\":1574399544,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":96,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574394451,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":97,\"attLeaderIcon\":0,\"st\":\"IDLE\",\"avoidEndT\":0,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":110,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574399830,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":111,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574399560,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":108,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574678195,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":109,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574770073,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":106,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1575075977,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":107,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"壹家人\",\"flag\":\"壹,7,0,9\",\"st\":\"IDLE\",\"avoidEndT\":1574769670,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":104,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574392694,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":105,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574418827,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":119,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574385328,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":118,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574680644,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":117,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574589063,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":116,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574493881,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":115,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574384452,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":114,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"兄弟一九一\",\"flag\":\"兄,0,1,2\",\"st\":\"IDLE\",\"avoidEndT\":1574470417,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":113,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574584034,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":112,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574398090,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":127,\"attLeaderIcon\":0,\"st\":\"IDLE\",\"avoidEndT\":0,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":126,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574397177,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":125,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574485171,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":124,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574564389,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":123,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"情义九洲\",\"flag\":\"秦,10,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574383119,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":122,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"自在就好\",\"flag\":\"自,0,0,3\",\"st\":\"IDLE\",\"avoidEndT\":1574383444,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":121,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574571736,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":120,\"attLeaderIcon\":0,\"ownerServer\":\"服务器（1-470）\",\"owner\":\"玉山书院\",\"flag\":\"玉,9,2,5\",\"st\":\"IDLE\",\"avoidEndT\":1574565614,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":129,\"attLeaderIcon\":0,\"st\":\"IDLE\",\"avoidEndT\":0,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false},{\"id\":128,\"attLeaderIcon\":0,\"st\":\"IDLE\",\"avoidEndT\":0,\"startT\":0,\"buff\":0,\"majorIcon\":0,\"countryId\":-1,\"destroyCountry\":false}],\"season\":\"WINTER\",\"year\":1803}","bytes":29283},"bytes":29533,"headers":{"content-length":29283,"transfer-encoding":"chunked","connection":"close","cache-control":"no-cache, no-cache, no-store, max-age=0, must-revalidate","pragma":"no-cache","server":"nginx/1.2.6","date":"Sat, 14 Dec 2019 02:18:17 GMT"},"status_phrase":"ok"},"version":"1.1","request":{"referrer":"http://somedomain/m.do","bytes":579,"headers":{"referer":"http://somedomain/m.do","x-requested-with":"ShockwaveFlash/21.0.0.0","yz_client_ip":"61.139.230.49","user-agent":"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.2; SV1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)","accept-language":"zh-CN","content-length":79,"yz-client-ip":"61.139.230.49","x-real-ip":"61.139.230.49","pragma":"no-cache","connection":"close","accept":"*/*","host":"somedomain","x-forwarded-for":"61.139.230.49","content-type":"application/x-www-form-urlencoded"},"method":"post","body":{"content":"{\"act\":\"World.worldSituation\",\"sid\":\"c71933d8401267e6f255dba5cdcfab6f27ee11a9\"}","bytes":79}}},"network":{"type":"ipv4","transport":"tcp","protocol":"http","community_id":"1:/eQtYOT0RhgMdRM/w4hjg7cogIo=","bytes":30112},"url":{"path":"/m.do","query":"%7B%22act%22%3A%22World.worldSituation%22%2C%22sid%22%3A%22c71933d8401267e6f255dba5cdcfab6f27ee11a9%22%7D=","full":"http://somedomain/m.do?%7B%22act%22%3A%22World.worldSituation%22%2C%22sid%22%3A%22c71933d8401267e6f255dba5cdcfab6f27ee11a9%22%7D=","scheme":"http","domain":"somedomain"},"client":{"bytes":579,"ip":"1.1.1.1","port":14261},"event":{"duration":3474000,"start":"2019-12-14T02:18:17.554Z","end":"2019-12-14T02:18:17.558Z","kind":"event","category":"network_traffic","dataset":"http"},"query":"POST /m.do","user_agent":{"original":"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.2; SV1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)"},"destination":{"ip":"1.1.1.1","port":80,"domain":"somedomain","bytes":29533},"type":"http","agent":{"hostname":"1.1.1.1","id":"d3ab5420-a50f-4730-bd6e-4ced53f0ef35","version":"7.0.1","type":"packetbeat","ephemeral_id":"ee1951d5-7f3e-41d1-898e-c93656ca75c5"}}`)
}

func genLargeFixture() []byte {
	return []byte(`
	{"users":[{"id":-1,"username":"system","avatar_template":"/user_avatar/discourse.metabase.com/system/{size}/6_1.png"},{"id":89,"username":"zergot","avatar_template":"https://avatars.discourse.org/v2/letter/z/0ea827/{size}.png"},{"id":1,"username":"sameer","avatar_template":"https://avatars.discourse.org/v2/letter/s/bbce88/{size}.png"},{"id":84,"username":"HenryMirror","avatar_template":"https://avatars.discourse.org/v2/letter/h/ecd19e/{size}.png"},{"id":73,"username":"fimp","avatar_template":"https://avatars.discourse.org/v2/letter/f/ee59a6/{size}.png"},{"id":14,"username":"agilliland","avatar_template":"/user_avatar/discourse.metabase.com/agilliland/{size}/26_1.png"},{"id":87,"username":"amir","avatar_template":"https://avatars.discourse.org/v2/letter/a/c37758/{size}.png"},{"id":82,"username":"waseem","avatar_template":"https://avatars.discourse.org/v2/letter/w/9dc877/{size}.png"},{"id":78,"username":"tovenaar","avatar_template":"https://avatars.discourse.org/v2/letter/t/9de0a6/{size}.png"},{"id":74,"username":"Ben","avatar_template":"https://avatars.discourse.org/v2/letter/b/df788c/{size}.png"},{"id":71,"username":"MarkLaFay","avatar_template":"https://avatars.discourse.org/v2/letter/m/3bc359/{size}.png"},{"id":72,"username":"camsaul","avatar_template":"/user_avatar/discourse.metabase.com/camsaul/{size}/70_1.png"},{"id":53,"username":"mhjb","avatar_template":"/user_avatar/discourse.metabase.com/mhjb/{size}/54_1.png"},{"id":58,"username":"jbwiv","avatar_template":"https://avatars.discourse.org/v2/letter/j/6bbea6/{size}.png"},{"id":70,"username":"Maggs","avatar_template":"https://avatars.discourse.org/v2/letter/m/bbce88/{size}.png"},{"id":69,"username":"andrefaria","avatar_template":"/user_avatar/discourse.metabase.com/andrefaria/{size}/65_1.png"},{"id":60,"username":"bencarter78","avatar_template":"/user_avatar/discourse.metabase.com/bencarter78/{size}/59_1.png"},{"id":55,"username":"vikram","avatar_template":"https://avatars.discourse.org/v2/letter/v/e47774/{size}.png"},{"id":68,"username":"edchan77","avatar_template":"/user_avatar/discourse.metabase.com/edchan77/{size}/66_1.png"},{"id":9,"username":"karthikd","avatar_template":"https://avatars.discourse.org/v2/letter/k/cab0a1/{size}.png"},{"id":23,"username":"arthurz","avatar_template":"/user_avatar/discourse.metabase.com/arthurz/{size}/32_1.png"},{"id":3,"username":"tom","avatar_template":"/user_avatar/discourse.metabase.com/tom/{size}/21_1.png"},{"id":50,"username":"LeoNogueira","avatar_template":"/user_avatar/discourse.metabase.com/leonogueira/{size}/52_1.png"},{"id":66,"username":"ss06vi","avatar_template":"https://avatars.discourse.org/v2/letter/s/3ab097/{size}.png"},{"id":34,"username":"mattcollins","avatar_template":"/user_avatar/discourse.metabase.com/mattcollins/{size}/41_1.png"},{"id":51,"username":"krmmalik","avatar_template":"/user_avatar/discourse.metabase.com/krmmalik/{size}/53_1.png"},{"id":46,"username":"odysseas","avatar_template":"https://avatars.discourse.org/v2/letter/o/5f8ce5/{size}.png"},{"id":5,"username":"jonthewayne","avatar_template":"/user_avatar/discourse.metabase.com/jonthewayne/{size}/18_1.png"},{"id":11,"username":"anandiyer","avatar_template":"/user_avatar/discourse.metabase.com/anandiyer/{size}/23_1.png"},{"id":25,"username":"alnorth","avatar_template":"/user_avatar/discourse.metabase.com/alnorth/{size}/34_1.png"},{"id":52,"username":"j_at_svg","avatar_template":"https://avatars.discourse.org/v2/letter/j/96bed5/{size}.png"},{"id":42,"username":"styts","avatar_template":"/user_avatar/discourse.metabase.com/styts/{size}/47_1.png"}],"topics":{"can_create_topic":false,"more_topics_url":"/c/uncategorized/l/latest?page=1","draft":null,"draft_key":"new_topic","draft_sequence":null,"per_page":30,"topics":[{"id":8,"title":"Welcome to Metabase's Discussion Forum","fancy_title":"Welcome to Metabase&rsquo;s Discussion Forum","slug":"welcome-to-metabases-discussion-forum","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":"/images/welcome/discourse-edit-post-animated.gif","created_at":"2015-10-17T00:14:49.526Z","last_posted_at":"2015-10-17T00:14:49.557Z","bumped":true,"bumped_at":"2015-10-21T02:32:22.486Z","unseen":false,"pinned":true,"unpinned":null,"excerpt":"Welcome to Metabase&#39;s discussion forum. This is a place to get help on installation, setting up as well as sharing tips and tricks.","visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":197,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"system","category_id":1,"pinned_globally":true,"posters":[{"extras":"latest single","description":"Original Poster, Most Recent Poster","user_id":-1}]},{"id":169,"title":"Formatting Dates","fancy_title":"Formatting Dates","slug":"formatting-dates","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":null,"created_at":"2016-01-14T06:30:45.311Z","last_posted_at":"2016-01-14T06:30:45.397Z","bumped":true,"bumped_at":"2016-01-14T06:30:45.397Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":11,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"zergot","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest single","description":"Original Poster, Most Recent Poster","user_id":89}]},{"id":168,"title":"Setting for google api key","fancy_title":"Setting for google api key","slug":"setting-for-google-api-key","posts_count":2,"reply_count":0,"highest_post_number":2,"image_url":null,"created_at":"2016-01-13T17:14:31.799Z","last_posted_at":"2016-01-14T06:24:03.421Z","bumped":true,"bumped_at":"2016-01-14T06:24:03.421Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":16,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"zergot","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest single","description":"Original Poster, Most Recent Poster","user_id":89}]},{"id":167,"title":"Cannot see non-US timezones on the admin","fancy_title":"Cannot see non-US timezones on the admin","slug":"cannot-see-non-us-timezones-on-the-admin","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":null,"created_at":"2016-01-13T17:07:36.764Z","last_posted_at":"2016-01-13T17:07:36.831Z","bumped":true,"bumped_at":"2016-01-13T17:07:36.831Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":11,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"zergot","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest single","description":"Original Poster, Most Recent Poster","user_id":89}]},{"id":164,"title":"External (Metabase level) linkages in data schema","fancy_title":"External (Metabase level) linkages in data schema","slug":"external-metabase-level-linkages-in-data-schema","posts_count":4,"reply_count":1,"highest_post_number":4,"image_url":null,"created_at":"2016-01-11T13:51:02.286Z","last_posted_at":"2016-01-12T11:06:37.259Z","bumped":true,"bumped_at":"2016-01-12T11:06:37.259Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":32,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"zergot","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"Original Poster, Most Recent Poster","user_id":89},{"extras":null,"description":"Frequent Poster","user_id":1}]},{"id":155,"title":"Query working on \"Questions\" but not in \"Pulses\"","fancy_title":"Query working on &ldquo;Questions&rdquo; but not in &ldquo;Pulses&rdquo;","slug":"query-working-on-questions-but-not-in-pulses","posts_count":3,"reply_count":0,"highest_post_number":3,"image_url":null,"created_at":"2016-01-01T14:06:10.083Z","last_posted_at":"2016-01-08T22:37:51.772Z","bumped":true,"bumped_at":"2016-01-08T22:37:51.772Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":72,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"agilliland","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":84},{"extras":null,"description":"Frequent Poster","user_id":73},{"extras":"latest","description":"Most Recent Poster","user_id":14}]},{"id":161,"title":"Pulses posted to Slack don't show question output","fancy_title":"Pulses posted to Slack don&rsquo;t show question output","slug":"pulses-posted-to-slack-dont-show-question-output","posts_count":2,"reply_count":0,"highest_post_number":2,"image_url":"/uploads/default/original/1X/9d2806517bf3598b10be135b2c58923b47ba23e7.png","created_at":"2016-01-08T22:09:58.205Z","last_posted_at":"2016-01-08T22:28:44.685Z","bumped":true,"bumped_at":"2016-01-08T22:28:44.685Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":34,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":87},{"extras":"latest","description":"Most Recent Poster","user_id":1}]},{"id":152,"title":"Should we build Kafka connecter or Kafka plugin","fancy_title":"Should we build Kafka connecter or Kafka plugin","slug":"should-we-build-kafka-connecter-or-kafka-plugin","posts_count":4,"reply_count":1,"highest_post_number":4,"image_url":null,"created_at":"2015-12-28T20:37:23.501Z","last_posted_at":"2015-12-31T18:16:45.477Z","bumped":true,"bumped_at":"2015-12-31T18:16:45.477Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":84,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":82},{"extras":"latest","description":"Most Recent Poster, Frequent Poster","user_id":1}]},{"id":147,"title":"Change X and Y on graph","fancy_title":"Change X and Y on graph","slug":"change-x-and-y-on-graph","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":null,"created_at":"2015-12-21T17:52:46.581Z","last_posted_at":"2015-12-21T17:52:46.684Z","bumped":true,"bumped_at":"2015-12-21T18:19:13.003Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":68,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"tovenaar","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest single","description":"Original Poster, Most Recent Poster","user_id":78}]},{"id":142,"title":"Issues sending mail via office365 relay","fancy_title":"Issues sending mail via office365 relay","slug":"issues-sending-mail-via-office365-relay","posts_count":5,"reply_count":2,"highest_post_number":5,"image_url":null,"created_at":"2015-12-16T10:38:47.315Z","last_posted_at":"2015-12-21T09:26:27.167Z","bumped":true,"bumped_at":"2015-12-21T09:26:27.167Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":122,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"Ben","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"Original Poster, Most Recent Poster","user_id":74},{"extras":null,"description":"Frequent Poster","user_id":1}]},{"id":137,"title":"I see triplicates of my mongoDB collections","fancy_title":"I see triplicates of my mongoDB collections","slug":"i-see-triplicates-of-my-mongodb-collections","posts_count":3,"reply_count":0,"highest_post_number":3,"image_url":null,"created_at":"2015-12-14T13:33:03.426Z","last_posted_at":"2015-12-17T18:40:05.487Z","bumped":true,"bumped_at":"2015-12-17T18:40:05.487Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":97,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"MarkLaFay","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"Original Poster, Most Recent Poster","user_id":71},{"extras":null,"description":"Frequent Poster","user_id":14}]},{"id":140,"title":"Google Analytics plugin","fancy_title":"Google Analytics plugin","slug":"google-analytics-plugin","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":null,"created_at":"2015-12-15T13:00:55.644Z","last_posted_at":"2015-12-15T13:00:55.705Z","bumped":true,"bumped_at":"2015-12-15T13:00:55.705Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":105,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"fimp","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest single","description":"Original Poster, Most Recent Poster","user_id":73}]},{"id":138,"title":"With-mongo-connection failed: bad connection details:","fancy_title":"With-mongo-connection failed: bad connection details:","slug":"with-mongo-connection-failed-bad-connection-details","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":null,"created_at":"2015-12-14T17:28:11.041Z","last_posted_at":"2015-12-14T17:28:11.111Z","bumped":true,"bumped_at":"2015-12-14T17:28:11.111Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":56,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"MarkLaFay","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest single","description":"Original Poster, Most Recent Poster","user_id":71}]},{"id":133,"title":"\"We couldn't understand your question.\" when I query mongoDB","fancy_title":"&ldquo;We couldn&rsquo;t understand your question.&rdquo; when I query mongoDB","slug":"we-couldnt-understand-your-question-when-i-query-mongodb","posts_count":3,"reply_count":0,"highest_post_number":3,"image_url":null,"created_at":"2015-12-11T17:38:30.576Z","last_posted_at":"2015-12-14T13:31:26.395Z","bumped":true,"bumped_at":"2015-12-14T13:31:26.395Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":107,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"MarkLaFay","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"Original Poster, Most Recent Poster","user_id":71},{"extras":null,"description":"Frequent Poster","user_id":72}]},{"id":129,"title":"My bar charts are all thin","fancy_title":"My bar charts are all thin","slug":"my-bar-charts-are-all-thin","posts_count":4,"reply_count":1,"highest_post_number":4,"image_url":"/uploads/default/original/1X/41bcf3b2a00dc7cfaff01cb3165d35d32a85bf1d.png","created_at":"2015-12-09T22:09:56.394Z","last_posted_at":"2015-12-11T19:00:45.289Z","bumped":true,"bumped_at":"2015-12-11T19:00:45.289Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":116,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"mhjb","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"Original Poster, Most Recent Poster","user_id":53},{"extras":null,"description":"Frequent Poster","user_id":1}]},{"id":106,"title":"What is the expected return order of columns for graphing results when using raw SQL?","fancy_title":"What is the expected return order of columns for graphing results when using raw SQL?","slug":"what-is-the-expected-return-order-of-columns-for-graphing-results-when-using-raw-sql","posts_count":3,"reply_count":0,"highest_post_number":3,"image_url":null,"created_at":"2015-11-24T19:07:14.561Z","last_posted_at":"2015-12-11T17:04:14.149Z","bumped":true,"bumped_at":"2015-12-11T17:04:14.149Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":153,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"jbwiv","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"Original Poster, Most Recent Poster","user_id":58},{"extras":null,"description":"Frequent Poster","user_id":14}]},{"id":131,"title":"Set site url from admin panel","fancy_title":"Set site url from admin panel","slug":"set-site-url-from-admin-panel","posts_count":2,"reply_count":0,"highest_post_number":2,"image_url":null,"created_at":"2015-12-10T06:22:46.042Z","last_posted_at":"2015-12-10T19:12:57.449Z","bumped":true,"bumped_at":"2015-12-10T19:12:57.449Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":77,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":70},{"extras":"latest","description":"Most Recent Poster","user_id":1}]},{"id":127,"title":"Internationalization (i18n)","fancy_title":"Internationalization (i18n)","slug":"internationalization-i18n","posts_count":2,"reply_count":0,"highest_post_number":2,"image_url":null,"created_at":"2015-12-08T16:55:37.397Z","last_posted_at":"2015-12-09T16:49:55.816Z","bumped":true,"bumped_at":"2015-12-09T16:49:55.816Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":85,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"agilliland","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":69},{"extras":"latest","description":"Most Recent Poster","user_id":14}]},{"id":109,"title":"Returning raw data with no filters always returns We couldn't understand your question","fancy_title":"Returning raw data with no filters always returns We couldn&rsquo;t understand your question","slug":"returning-raw-data-with-no-filters-always-returns-we-couldnt-understand-your-question","posts_count":3,"reply_count":1,"highest_post_number":3,"image_url":null,"created_at":"2015-11-25T21:35:01.315Z","last_posted_at":"2015-12-09T10:26:12.255Z","bumped":true,"bumped_at":"2015-12-09T10:26:12.255Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":133,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"bencarter78","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"Original Poster, Most Recent Poster","user_id":60},{"extras":null,"description":"Frequent Poster","user_id":14}]},{"id":103,"title":"Support for Cassandra?","fancy_title":"Support for Cassandra?","slug":"support-for-cassandra","posts_count":5,"reply_count":1,"highest_post_number":5,"image_url":null,"created_at":"2015-11-20T06:45:31.741Z","last_posted_at":"2015-12-09T03:18:51.274Z","bumped":true,"bumped_at":"2015-12-09T03:18:51.274Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":169,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"vikram","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"Original Poster, Most Recent Poster","user_id":55},{"extras":null,"description":"Frequent Poster","user_id":1}]},{"id":128,"title":"Mongo query with Date breaks [solved: Mongo 3.0 required]","fancy_title":"Mongo query with Date breaks [solved: Mongo 3.0 required]","slug":"mongo-query-with-date-breaks-solved-mongo-3-0-required","posts_count":5,"reply_count":0,"highest_post_number":5,"image_url":null,"created_at":"2015-12-08T18:30:56.562Z","last_posted_at":"2015-12-08T21:03:02.421Z","bumped":true,"bumped_at":"2015-12-08T21:03:02.421Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":102,"like_count":1,"has_summary":false,"archetype":"regular","last_poster_username":"edchan77","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"Original Poster, Most Recent Poster","user_id":68},{"extras":null,"description":"Frequent Poster","user_id":1}]},{"id":23,"title":"Can this connect to MS SQL Server?","fancy_title":"Can this connect to MS SQL Server?","slug":"can-this-connect-to-ms-sql-server","posts_count":7,"reply_count":1,"highest_post_number":7,"image_url":null,"created_at":"2015-10-21T18:52:37.987Z","last_posted_at":"2015-12-07T17:41:51.609Z","bumped":true,"bumped_at":"2015-12-07T17:41:51.609Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":367,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":9},{"extras":null,"description":"Frequent Poster","user_id":23},{"extras":null,"description":"Frequent Poster","user_id":3},{"extras":null,"description":"Frequent Poster","user_id":50},{"extras":"latest","description":"Most Recent Poster","user_id":1}]},{"id":121,"title":"Cannot restart metabase in docker","fancy_title":"Cannot restart metabase in docker","slug":"cannot-restart-metabase-in-docker","posts_count":5,"reply_count":1,"highest_post_number":5,"image_url":null,"created_at":"2015-12-04T21:28:58.137Z","last_posted_at":"2015-12-04T23:02:00.488Z","bumped":true,"bumped_at":"2015-12-04T23:02:00.488Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":96,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":66},{"extras":"latest","description":"Most Recent Poster, Frequent Poster","user_id":1}]},{"id":85,"title":"Edit Max Rows Count","fancy_title":"Edit Max Rows Count","slug":"edit-max-rows-count","posts_count":4,"reply_count":2,"highest_post_number":4,"image_url":null,"created_at":"2015-11-11T23:46:52.917Z","last_posted_at":"2015-11-24T01:01:14.569Z","bumped":true,"bumped_at":"2015-11-24T01:01:14.569Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":169,"like_count":1,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":34},{"extras":"latest","description":"Most Recent Poster, Frequent Poster","user_id":1}]},{"id":96,"title":"Creating charts by querying more than one table at a time","fancy_title":"Creating charts by querying more than one table at a time","slug":"creating-charts-by-querying-more-than-one-table-at-a-time","posts_count":6,"reply_count":4,"highest_post_number":6,"image_url":null,"created_at":"2015-11-17T11:20:18.442Z","last_posted_at":"2015-11-21T02:12:25.995Z","bumped":true,"bumped_at":"2015-11-21T02:12:25.995Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":217,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":51},{"extras":"latest","description":"Most Recent Poster, Frequent Poster","user_id":1}]},{"id":90,"title":"Trying to add RDS postgresql as the database fails silently","fancy_title":"Trying to add RDS postgresql as the database fails silently","slug":"trying-to-add-rds-postgresql-as-the-database-fails-silently","posts_count":4,"reply_count":2,"highest_post_number":4,"image_url":null,"created_at":"2015-11-14T23:45:02.967Z","last_posted_at":"2015-11-21T01:08:45.915Z","bumped":true,"bumped_at":"2015-11-21T01:08:45.915Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":162,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":46},{"extras":"latest","description":"Most Recent Poster, Frequent Poster","user_id":1}]},{"id":17,"title":"Deploy to Heroku isn't working","fancy_title":"Deploy to Heroku isn&rsquo;t working","slug":"deploy-to-heroku-isnt-working","posts_count":9,"reply_count":3,"highest_post_number":9,"image_url":null,"created_at":"2015-10-21T16:42:03.096Z","last_posted_at":"2015-11-20T18:34:14.044Z","bumped":true,"bumped_at":"2015-11-20T18:34:14.044Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":332,"like_count":2,"has_summary":false,"archetype":"regular","last_poster_username":"agilliland","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":5},{"extras":null,"description":"Frequent Poster","user_id":3},{"extras":null,"description":"Frequent Poster","user_id":11},{"extras":null,"description":"Frequent Poster","user_id":25},{"extras":"latest","description":"Most Recent Poster","user_id":14}]},{"id":100,"title":"Can I use DATEPART() in SQL queries?","fancy_title":"Can I use DATEPART() in SQL queries?","slug":"can-i-use-datepart-in-sql-queries","posts_count":2,"reply_count":0,"highest_post_number":2,"image_url":null,"created_at":"2015-11-17T23:15:58.033Z","last_posted_at":"2015-11-18T00:19:48.763Z","bumped":true,"bumped_at":"2015-11-18T00:19:48.763Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":112,"like_count":1,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":53},{"extras":"latest","description":"Most Recent Poster","user_id":1}]},{"id":98,"title":"Feature Request: LDAP Authentication","fancy_title":"Feature Request: LDAP Authentication","slug":"feature-request-ldap-authentication","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":null,"created_at":"2015-11-17T17:22:44.484Z","last_posted_at":"2015-11-17T17:22:44.577Z","bumped":true,"bumped_at":"2015-11-17T17:22:44.577Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":97,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"j_at_svg","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest single","description":"Original Poster, Most Recent Poster","user_id":52}]},{"id":87,"title":"Migrating from internal H2 to Postgres","fancy_title":"Migrating from internal H2 to Postgres","slug":"migrating-from-internal-h2-to-postgres","posts_count":2,"reply_count":0,"highest_post_number":2,"image_url":null,"created_at":"2015-11-12T14:36:06.745Z","last_posted_at":"2015-11-12T18:05:10.796Z","bumped":true,"bumped_at":"2015-11-12T18:05:10.796Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":111,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":42},{"extras":"latest","description":"Most Recent Poster","user_id":1}]},{"id":8,"title":"Welcome to Metabase's Discussion Forum","fancy_title":"Welcome to Metabase&rsquo;s Discussion Forum","slug":"welcome-to-metabases-discussion-forum","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":"/images/welcome/discourse-edit-post-animated.gif","created_at":"2015-10-17T00:14:49.526Z","last_posted_at":"2015-10-17T00:14:49.557Z","bumped":true,"bumped_at":"2015-10-21T02:32:22.486Z","unseen":false,"pinned":true,"unpinned":null,"excerpt":"Welcome to Metabase&#39;s discussion forum. This is a place to get help on installation, setting up as well as sharing tips and tricks.","visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":197,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"system","category_id":1,"pinned_globally":true,"posters":[{"extras":"latest single","description":"Original Poster, Most Recent Poster","user_id":-1}]},{"id":169,"title":"Formatting Dates","fancy_title":"Formatting Dates","slug":"formatting-dates","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":null,"created_at":"2016-01-14T06:30:45.311Z","last_posted_at":"2016-01-14T06:30:45.397Z","bumped":true,"bumped_at":"2016-01-14T06:30:45.397Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":11,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"zergot","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest single","description":"Original Poster, Most Recent Poster","user_id":89}]},{"id":168,"title":"Setting for google api key","fancy_title":"Setting for google api key","slug":"setting-for-google-api-key","posts_count":2,"reply_count":0,"highest_post_number":2,"image_url":null,"created_at":"2016-01-13T17:14:31.799Z","last_posted_at":"2016-01-14T06:24:03.421Z","bumped":true,"bumped_at":"2016-01-14T06:24:03.421Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":16,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"zergot","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest single","description":"Original Poster, Most Recent Poster","user_id":89}]},{"id":167,"title":"Cannot see non-US timezones on the admin","fancy_title":"Cannot see non-US timezones on the admin","slug":"cannot-see-non-us-timezones-on-the-admin","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":null,"created_at":"2016-01-13T17:07:36.764Z","last_posted_at":"2016-01-13T17:07:36.831Z","bumped":true,"bumped_at":"2016-01-13T17:07:36.831Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":11,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"zergot","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest single","description":"Original Poster, Most Recent Poster","user_id":89}]},{"id":164,"title":"External (Metabase level) linkages in data schema","fancy_title":"External (Metabase level) linkages in data schema","slug":"external-metabase-level-linkages-in-data-schema","posts_count":4,"reply_count":1,"highest_post_number":4,"image_url":null,"created_at":"2016-01-11T13:51:02.286Z","last_posted_at":"2016-01-12T11:06:37.259Z","bumped":true,"bumped_at":"2016-01-12T11:06:37.259Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":32,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"zergot","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"Original Poster, Most Recent Poster","user_id":89},{"extras":null,"description":"Frequent Poster","user_id":1}]},{"id":155,"title":"Query working on \"Questions\" but not in \"Pulses\"","fancy_title":"Query working on &ldquo;Questions&rdquo; but not in &ldquo;Pulses&rdquo;","slug":"query-working-on-questions-but-not-in-pulses","posts_count":3,"reply_count":0,"highest_post_number":3,"image_url":null,"created_at":"2016-01-01T14:06:10.083Z","last_posted_at":"2016-01-08T22:37:51.772Z","bumped":true,"bumped_at":"2016-01-08T22:37:51.772Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":72,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"agilliland","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":84},{"extras":null,"description":"Frequent Poster","user_id":73},{"extras":"latest","description":"Most Recent Poster","user_id":14}]},{"id":161,"title":"Pulses posted to Slack don't show question output","fancy_title":"Pulses posted to Slack don&rsquo;t show question output","slug":"pulses-posted-to-slack-dont-show-question-output","posts_count":2,"reply_count":0,"highest_post_number":2,"image_url":"/uploads/default/original/1X/9d2806517bf3598b10be135b2c58923b47ba23e7.png","created_at":"2016-01-08T22:09:58.205Z","last_posted_at":"2016-01-08T22:28:44.685Z","bumped":true,"bumped_at":"2016-01-08T22:28:44.685Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":34,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":87},{"extras":"latest","description":"Most Recent Poster","user_id":1}]},{"id":152,"title":"Should we build Kafka connecter or Kafka plugin","fancy_title":"Should we build Kafka connecter or Kafka plugin","slug":"should-we-build-kafka-connecter-or-kafka-plugin","posts_count":4,"reply_count":1,"highest_post_number":4,"image_url":null,"created_at":"2015-12-28T20:37:23.501Z","last_posted_at":"2015-12-31T18:16:45.477Z","bumped":true,"bumped_at":"2015-12-31T18:16:45.477Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":84,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":82},{"extras":"latest","description":"Most Recent Poster, Frequent Poster","user_id":1}]},{"id":147,"title":"Change X and Y on graph","fancy_title":"Change X and Y on graph","slug":"change-x-and-y-on-graph","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":null,"created_at":"2015-12-21T17:52:46.581Z","last_posted_at":"2015-12-21T17:52:46.684Z","bumped":true,"bumped_at":"2015-12-21T18:19:13.003Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":68,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"tovenaar","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest single","description":"Original Poster, Most Recent Poster","user_id":78}]},{"id":142,"title":"Issues sending mail via office365 relay","fancy_title":"Issues sending mail via office365 relay","slug":"issues-sending-mail-via-office365-relay","posts_count":5,"reply_count":2,"highest_post_number":5,"image_url":null,"created_at":"2015-12-16T10:38:47.315Z","last_posted_at":"2015-12-21T09:26:27.167Z","bumped":true,"bumped_at":"2015-12-21T09:26:27.167Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":122,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"Ben","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"Original Poster, Most Recent Poster","user_id":74},{"extras":null,"description":"Frequent Poster","user_id":1}]},{"id":137,"title":"I see triplicates of my mongoDB collections","fancy_title":"I see triplicates of my mongoDB collections","slug":"i-see-triplicates-of-my-mongodb-collections","posts_count":3,"reply_count":0,"highest_post_number":3,"image_url":null,"created_at":"2015-12-14T13:33:03.426Z","last_posted_at":"2015-12-17T18:40:05.487Z","bumped":true,"bumped_at":"2015-12-17T18:40:05.487Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":97,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"MarkLaFay","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"Original Poster, Most Recent Poster","user_id":71},{"extras":null,"description":"Frequent Poster","user_id":14}]},{"id":140,"title":"Google Analytics plugin","fancy_title":"Google Analytics plugin","slug":"google-analytics-plugin","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":null,"created_at":"2015-12-15T13:00:55.644Z","last_posted_at":"2015-12-15T13:00:55.705Z","bumped":true,"bumped_at":"2015-12-15T13:00:55.705Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":105,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"fimp","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest single","description":"Original Poster, Most Recent Poster","user_id":73}]},{"id":138,"title":"With-mongo-connection failed: bad connection details:","fancy_title":"With-mongo-connection failed: bad connection details:","slug":"with-mongo-connection-failed-bad-connection-details","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":null,"created_at":"2015-12-14T17:28:11.041Z","last_posted_at":"2015-12-14T17:28:11.111Z","bumped":true,"bumped_at":"2015-12-14T17:28:11.111Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":56,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"MarkLaFay","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest single","description":"Original Poster, Most Recent Poster","user_id":71}]},{"id":133,"title":"\"We couldn't understand your question.\" when I query mongoDB","fancy_title":"&ldquo;We couldn&rsquo;t understand your question.&rdquo; when I query mongoDB","slug":"we-couldnt-understand-your-question-when-i-query-mongodb","posts_count":3,"reply_count":0,"highest_post_number":3,"image_url":null,"created_at":"2015-12-11T17:38:30.576Z","last_posted_at":"2015-12-14T13:31:26.395Z","bumped":true,"bumped_at":"2015-12-14T13:31:26.395Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":107,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"MarkLaFay","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"Original Poster, Most Recent Poster","user_id":71},{"extras":null,"description":"Frequent Poster","user_id":72}]},{"id":129,"title":"My bar charts are all thin","fancy_title":"My bar charts are all thin","slug":"my-bar-charts-are-all-thin","posts_count":4,"reply_count":1,"highest_post_number":4,"image_url":"/uploads/default/original/1X/41bcf3b2a00dc7cfaff01cb3165d35d32a85bf1d.png","created_at":"2015-12-09T22:09:56.394Z","last_posted_at":"2015-12-11T19:00:45.289Z","bumped":true,"bumped_at":"2015-12-11T19:00:45.289Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":116,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"mhjb","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"Original Poster, Most Recent Poster","user_id":53},{"extras":null,"description":"Frequent Poster","user_id":1}]},{"id":106,"title":"What is the expected return order of columns for graphing results when using raw SQL?","fancy_title":"What is the expected return order of columns for graphing results when using raw SQL?","slug":"what-is-the-expected-return-order-of-columns-for-graphing-results-when-using-raw-sql","posts_count":3,"reply_count":0,"highest_post_number":3,"image_url":null,"created_at":"2015-11-24T19:07:14.561Z","last_posted_at":"2015-12-11T17:04:14.149Z","bumped":true,"bumped_at":"2015-12-11T17:04:14.149Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":153,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"jbwiv","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"Original Poster, Most Recent Poster","user_id":58},{"extras":null,"description":"Frequent Poster","user_id":14}]},{"id":131,"title":"Set site url from admin panel","fancy_title":"Set site url from admin panel","slug":"set-site-url-from-admin-panel","posts_count":2,"reply_count":0,"highest_post_number":2,"image_url":null,"created_at":"2015-12-10T06:22:46.042Z","last_posted_at":"2015-12-10T19:12:57.449Z","bumped":true,"bumped_at":"2015-12-10T19:12:57.449Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":77,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":70},{"extras":"latest","description":"Most Recent Poster","user_id":1}]},{"id":127,"title":"Internationalization (i18n)","fancy_title":"Internationalization (i18n)","slug":"internationalization-i18n","posts_count":2,"reply_count":0,"highest_post_number":2,"image_url":null,"created_at":"2015-12-08T16:55:37.397Z","last_posted_at":"2015-12-09T16:49:55.816Z","bumped":true,"bumped_at":"2015-12-09T16:49:55.816Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":85,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"agilliland","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":69},{"extras":"latest","description":"Most Recent Poster","user_id":14}]},{"id":109,"title":"Returning raw data with no filters always returns We couldn't understand your question","fancy_title":"Returning raw data with no filters always returns We couldn&rsquo;t understand your question","slug":"returning-raw-data-with-no-filters-always-returns-we-couldnt-understand-your-question","posts_count":3,"reply_count":1,"highest_post_number":3,"image_url":null,"created_at":"2015-11-25T21:35:01.315Z","last_posted_at":"2015-12-09T10:26:12.255Z","bumped":true,"bumped_at":"2015-12-09T10:26:12.255Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":133,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"bencarter78","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"Original Poster, Most Recent Poster","user_id":60},{"extras":null,"description":"Frequent Poster","user_id":14}]},{"id":103,"title":"Support for Cassandra?","fancy_title":"Support for Cassandra?","slug":"support-for-cassandra","posts_count":5,"reply_count":1,"highest_post_number":5,"image_url":null,"created_at":"2015-11-20T06:45:31.741Z","last_posted_at":"2015-12-09T03:18:51.274Z","bumped":true,"bumped_at":"2015-12-09T03:18:51.274Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":169,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"vikram","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"Original Poster, Most Recent Poster","user_id":55},{"extras":null,"description":"Frequent Poster","user_id":1}]},{"id":128,"title":"Mongo query with Date breaks [solved: Mongo 3.0 required]","fancy_title":"Mongo query with Date breaks [solved: Mongo 3.0 required]","slug":"mongo-query-with-date-breaks-solved-mongo-3-0-required","posts_count":5,"reply_count":0,"highest_post_number":5,"image_url":null,"created_at":"2015-12-08T18:30:56.562Z","last_posted_at":"2015-12-08T21:03:02.421Z","bumped":true,"bumped_at":"2015-12-08T21:03:02.421Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":102,"like_count":1,"has_summary":false,"archetype":"regular","last_poster_username":"edchan77","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest","description":"Original Poster, Most Recent Poster","user_id":68},{"extras":null,"description":"Frequent Poster","user_id":1}]},{"id":23,"title":"Can this connect to MS SQL Server?","fancy_title":"Can this connect to MS SQL Server?","slug":"can-this-connect-to-ms-sql-server","posts_count":7,"reply_count":1,"highest_post_number":7,"image_url":null,"created_at":"2015-10-21T18:52:37.987Z","last_posted_at":"2015-12-07T17:41:51.609Z","bumped":true,"bumped_at":"2015-12-07T17:41:51.609Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":367,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":9},{"extras":null,"description":"Frequent Poster","user_id":23},{"extras":null,"description":"Frequent Poster","user_id":3},{"extras":null,"description":"Frequent Poster","user_id":50},{"extras":"latest","description":"Most Recent Poster","user_id":1}]},{"id":121,"title":"Cannot restart metabase in docker","fancy_title":"Cannot restart metabase in docker","slug":"cannot-restart-metabase-in-docker","posts_count":5,"reply_count":1,"highest_post_number":5,"image_url":null,"created_at":"2015-12-04T21:28:58.137Z","last_posted_at":"2015-12-04T23:02:00.488Z","bumped":true,"bumped_at":"2015-12-04T23:02:00.488Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":96,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":66},{"extras":"latest","description":"Most Recent Poster, Frequent Poster","user_id":1}]},{"id":85,"title":"Edit Max Rows Count","fancy_title":"Edit Max Rows Count","slug":"edit-max-rows-count","posts_count":4,"reply_count":2,"highest_post_number":4,"image_url":null,"created_at":"2015-11-11T23:46:52.917Z","last_posted_at":"2015-11-24T01:01:14.569Z","bumped":true,"bumped_at":"2015-11-24T01:01:14.569Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":169,"like_count":1,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":34},{"extras":"latest","description":"Most Recent Poster, Frequent Poster","user_id":1}]},{"id":96,"title":"Creating charts by querying more than one table at a time","fancy_title":"Creating charts by querying more than one table at a time","slug":"creating-charts-by-querying-more-than-one-table-at-a-time","posts_count":6,"reply_count":4,"highest_post_number":6,"image_url":null,"created_at":"2015-11-17T11:20:18.442Z","last_posted_at":"2015-11-21T02:12:25.995Z","bumped":true,"bumped_at":"2015-11-21T02:12:25.995Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":217,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":51},{"extras":"latest","description":"Most Recent Poster, Frequent Poster","user_id":1}]},{"id":90,"title":"Trying to add RDS postgresql as the database fails silently","fancy_title":"Trying to add RDS postgresql as the database fails silently","slug":"trying-to-add-rds-postgresql-as-the-database-fails-silently","posts_count":4,"reply_count":2,"highest_post_number":4,"image_url":null,"created_at":"2015-11-14T23:45:02.967Z","last_posted_at":"2015-11-21T01:08:45.915Z","bumped":true,"bumped_at":"2015-11-21T01:08:45.915Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":162,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":46},{"extras":"latest","description":"Most Recent Poster, Frequent Poster","user_id":1}]},{"id":17,"title":"Deploy to Heroku isn't working","fancy_title":"Deploy to Heroku isn&rsquo;t working","slug":"deploy-to-heroku-isnt-working","posts_count":9,"reply_count":3,"highest_post_number":9,"image_url":null,"created_at":"2015-10-21T16:42:03.096Z","last_posted_at":"2015-11-20T18:34:14.044Z","bumped":true,"bumped_at":"2015-11-20T18:34:14.044Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":332,"like_count":2,"has_summary":false,"archetype":"regular","last_poster_username":"agilliland","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":5},{"extras":null,"description":"Frequent Poster","user_id":3},{"extras":null,"description":"Frequent Poster","user_id":11},{"extras":null,"description":"Frequent Poster","user_id":25},{"extras":"latest","description":"Most Recent Poster","user_id":14}]},{"id":100,"title":"Can I use DATEPART() in SQL queries?","fancy_title":"Can I use DATEPART() in SQL queries?","slug":"can-i-use-datepart-in-sql-queries","posts_count":2,"reply_count":0,"highest_post_number":2,"image_url":null,"created_at":"2015-11-17T23:15:58.033Z","last_posted_at":"2015-11-18T00:19:48.763Z","bumped":true,"bumped_at":"2015-11-18T00:19:48.763Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":112,"like_count":1,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":53},{"extras":"latest","description":"Most Recent Poster","user_id":1}]},{"id":98,"title":"Feature Request: LDAP Authentication","fancy_title":"Feature Request: LDAP Authentication","slug":"feature-request-ldap-authentication","posts_count":1,"reply_count":0,"highest_post_number":1,"image_url":null,"created_at":"2015-11-17T17:22:44.484Z","last_posted_at":"2015-11-17T17:22:44.577Z","bumped":true,"bumped_at":"2015-11-17T17:22:44.577Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":97,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"j_at_svg","category_id":1,"pinned_globally":false,"posters":[{"extras":"latest single","description":"Original Poster, Most Recent Poster","user_id":52}]},{"id":87,"title":"Migrating from internal H2 to Postgres","fancy_title":"Migrating from internal H2 to Postgres","slug":"migrating-from-internal-h2-to-postgres","posts_count":2,"reply_count":0,"highest_post_number":2,"image_url":null,"created_at":"2015-11-12T14:36:06.745Z","last_posted_at":"2015-11-12T18:05:10.796Z","bumped":true,"bumped_at":"2015-11-12T18:05:10.796Z","unseen":false,"pinned":false,"unpinned":null,"visible":true,"closed":false,"archived":false,"bookmarked":null,"liked":null,"views":111,"like_count":0,"has_summary":false,"archetype":"regular","last_poster_username":"sameer","category_id":1,"pinned_globally":false,"posters":[{"extras":null,"description":"Original Poster","user_id":42},{"extras":"latest","description":"Most Recent Poster","user_id":1}]}]}}
	`)
}

func genMediumFixture() []byte {
	return []byte(`{
		"person": {
		  "id": "d50887ca-a6ce-4e59-b89f-14f0b5d03b03",
		  "name": {
			"fullName": "Leonid Bugaev",
			"givenName": "Leonid",
			"familyName": "Bugaev"
		  },
		  "email": "leonsbox@gmail.com",
		  "gender": "male",
		  "location": "Saint Petersburg, Saint Petersburg, RU",
		  "geo": {
			"city": "Saint Petersburg",
			"state": "Saint Petersburg",
			"country": "Russia",
			"lat": 59.9342802,
			"lng": 30.3350986
		  },
		  "bio": "Senior engineer at Granify.com",
		  "site": "http://flickfaver.com",
		  "avatar": "https://d1ts43dypk8bqh.cloudfront.net/v1/avatars/d50887ca-a6ce-4e59-b89f-14f0b5d03b03",
		  "employment": {
			"name": "www.latera.ru",
			"title": "Software Engineer",
			"domain": "gmail.com"
		  },
		  "facebook": {
			"handle": "leonid.bugaev"
		  },
		  "github": {
			"handle": "buger",
			"id": 14009,
			"avatar": "https://avatars.githubusercontent.com/u/14009?v=3",
			"company": "Granify",
			"blog": "http://leonsbox.com",
			"followers": 95,
			"following": 10
		  },
		  "twitter": {
			"handle": "flickfaver",
			"id": 77004410,
			"bio": null,
			"followers": 2,
			"following": 1,
			"statuses": 5,
			"favorites": 0,
			"location": "",
			"site": "http://flickfaver.com",
			"avatar": null
		  },
		  "linkedin": {
			"handle": "in/leonidbugaev"
		  },
		  "googleplus": {
			"handle": null
		  },
		  "angellist": {
			"handle": "leonid-bugaev",
			"id": 61541,
			"bio": "Senior engineer at Granify.com",
			"blog": "http://buger.github.com",
			"site": "http://buger.github.com",
			"followers": 41,
			"avatar": "https://d1qb2nb5cznatu.cloudfront.net/users/61541-medium_jpg?1405474390"
		  },
		  "klout": {
			"handle": null,
			"score": null
		  },
		  "foursquare": {
			"handle": null
		  },
		  "aboutme": {
			"handle": "leonid.bugaev",
			"bio": null,
			"avatar": null
		  },
		  "gravatar": {
			"handle": "buger",
			"urls": [
			],
			"avatar": "http://1.gravatar.com/avatar/f7c8edd577d13b8930d5522f28123510",
			"avatars": [
			  {
				"url": "http://1.gravatar.com/avatar/f7c8edd577d13b8930d5522f28123510",
				"type": "thumbnail"
			  }
			]
		  },
		  "fuzzy": false
		},
		"company": null
	  }`)
}

func genSmallFixture() []byte {
	return []byte(`{"st": 1,"sid": 486,"tt": "active","gr": 0,"uuid": "de305d54-75b4-431b-adb2-eb6b9e546014","ip": "127.0.0.1","ua": "user_agent","tz": -6,"v": 1}`)
}

func genTestData() []byte {
	return []byte(`{"a":"b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\b\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcdb\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\bcd"}`)
}

func TestGOJSONUnmarshalLarge(t *testing.T) {
	var obj LargePayload
	largeFixture := genLargeFixture()
	err := gojson.Unmarshal(largeFixture, &obj)
	assert.Nil(t, err, "Err must be nil")

	var obj1 LargePayload
	largeFixture = genLargeFixture()
	err = easyjson.Unmarshal(largeFixture, &obj1)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, obj1, obj, "obj must be equal to the value expected")
}

func TestGOJSONUnmarshalMedium(t *testing.T) {
	var obj MediumPayload
	mediumFixture := genMediumFixture()
	err := gojson.Unmarshal(mediumFixture, &obj)
	assert.Nil(t, err, "Err must be nil")

	var obj1 MediumPayload
	mediumFixture = genMediumFixture()
	err = easyjson.Unmarshal(mediumFixture, &obj1)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, obj1, obj, "obj must be equal to the value expected")
}

func TestGOJSONUnmarshalSmall(t *testing.T) {
	var obj SmallPayload

	smallFixture := genSmallFixture()
	err := gojson.Unmarshal(smallFixture, &obj)
	assert.Nil(t, err, "Err must be nil")

	var obj1 SmallPayload
	smallFixture = genSmallFixture()
	err = easyjson.Unmarshal(smallFixture, &obj1)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, obj1, obj, "obj must be equal to the value expected")
}

func TestGOJSONMarshalLarge(t *testing.T) {
	var obj LargePayload

	err := gojson.Unmarshal(genLargeFixture(), &obj)
	assert.Nil(t, err, "Err must be nil")

	data, err := gojson.Marshal(&obj)
	assert.Nil(t, err, "Err must be nil")

	var obj1 LargePayload
	err = easyjson.Unmarshal(data, &obj1)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, obj1, obj, "obj must be equal to the value expected")
}

func TestUnmarshalSmall(t *testing.T) {
	var obj SmallPayload

	err := gojson.Unmarshal(genSmallFixture(), &obj)
	assert.Nil(t, err, "Err must be nil")

	data, err := gojson.Marshal(&obj)
	assert.Nil(t, err, "Err must be nil")

	var obj1 SmallPayload
	err = easyjson.Unmarshal(data, &obj1)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, obj1, obj, "obj must be equal to the value expected")
}

func TestUnmarshalTestStruct(t *testing.T) {
	var obj TestStruct

	err := gojson.Unmarshal(genTestData(), &obj)
	assert.Nil(t, err, "Err must be nil")

	data, err := gojson.Marshal(&obj)
	assert.Nil(t, err, "Err must be nil")

	var obj1 TestStruct
	err = easyjson.Unmarshal(data, &obj1)
	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, obj1, obj, "obj must be equal to the value expected")
}

func TestUnmarshalTestLargeStruct(t *testing.T) {
	var obj TestLargeStruct

	err := gojson.Unmarshal(genLargeTestData(), &obj)
	assert.Nil(t, err, "Err must be nil")

	data, err := gojson.Marshal(&obj)
	assert.Nil(t, err, "Err must be nil")

	var obj1 TestLargeStruct
	err = easyjson.Unmarshal(data, &obj1)

	assert.Nil(t, err, "Err must be nil")
	assert.Equal(t, obj1, obj, "obj must be equal to the value expected")
}

func BenchmarkGOJSONUnmarshalLarge(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genLargeFixture())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj LargePayload
		largeFixture := genLargeFixture()
		gojson.Unmarshal(largeFixture, &obj)

		for _, u := range obj.Users {
			nothing(u.Username)
		}

		for _, t := range obj.Topics.Topics {
			nothing(t.Id, t.Slug)
		}
	}
}

func BenchmarkJsonParserUnmarshalLarge(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genLargeFixture())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		largeFixture := genLargeFixture()

		jsonparser.ArrayEach(largeFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			jsonparser.Get(value, "username")
			nothing()
		}, "users")

		jsonparser.ArrayEach(largeFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			jsonparser.GetInt(value, "id")
			jsonparser.Get(value, "slug")
			nothing()
		}, "topics", "topics")
	}
}

func BenchmarkGoJayUnmarshalLarge(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genLargeFixture())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj LargePayload
		// gojay.UnmarshalJSONObject(largeFixture, &obj)
		largeFixture := genLargeFixture()
		gojay.Unsafe.UnmarshalJSONObject(largeFixture, &obj)

		for _, u := range obj.Users {
			nothing(u.Username)
		}

		for _, t := range obj.Topics.Topics {
			nothing(t.Id, t.Slug)
		}
	}
}

func BenchmarkEasyJsonUnmarshalLarge(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genLargeFixture())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		largeFixture := genLargeFixture()
		lexer := &jlexer.Lexer{Data: largeFixture}
		data := new(LargePayload)
		data.UnmarshalEasyJSON(lexer)

		for _, u := range data.Users {
			nothing(u.Username)
		}

		for _, t := range data.Topics.Topics {
			nothing(t.Id, t.Slug)
		}
	}
}

func BenchmarkJSONUnmarshalLarge(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genLargeFixture())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj LargePayload
		largeFixture := genLargeFixture()
		json.Unmarshal(largeFixture, &obj)

		for _, u := range obj.Users {
			nothing(u.Username)
		}

		for _, t := range obj.Topics.Topics {
			nothing(t.Id, t.Slug)
		}
	}
}

func BenchmarkJsonIterUnmarshalLarge(b *testing.B) {
	iter := jsoniter.ParseBytes(jsoniter.ConfigFastest, nil)

	b.ReportAllocs()
	b.SetBytes(int64(len(genLargeFixture())))
	b.ResetTimer()

	// var json = jsoniter.ConfigFastest
	for i := 0; i < b.N; i++ {
		largeFixture := genLargeFixture()
		iter.ResetBytes(largeFixture)
		count := 0
		for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
			if "topics" != field {
				iter.Skip()
				continue
			}
			for field := iter.ReadObject(); field != ""; field = iter.ReadObject() {
				if "topics" != field {
					iter.Skip()
					continue
				}
				for iter.ReadArray() {
					iter.Skip()
					count++
				}
				break
			}
			break
		}

	}
}

func BenchmarkGOJSONUnmarshalMedium(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genMediumFixture())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj MediumPayload
		mediumFixture := genMediumFixture()
		gojson.Unmarshal(mediumFixture, &obj)

		nothing(obj.Person.Name.FullName, obj.Person.Github.Followers, obj.Company)

		for _, el := range obj.Person.Gravatar.Avatars {
			nothing(el.Url)
		}
	}
}

func BenchmarkJsonParserUnmarshalMedium(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genMediumFixture())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mediumFixture := genMediumFixture()
		jsonparser.Get(mediumFixture, "person", "name", "fullName")
		jsonparser.GetInt(mediumFixture, "person", "github", "followers")
		jsonparser.Get(mediumFixture, "company")

		jsonparser.ArrayEach(mediumFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			jsonparser.Get(value, "url")
			nothing()
		}, "person", "gravatar", "avatars")
	}
}

func BenchmarkGoJayUnmarshalMedium(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genMediumFixture())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj MediumPayload
		mediumFixture := genMediumFixture()
		gojay.UnmarshalJSONObject(mediumFixture, &obj)

		nothing(obj.Person.Name.FullName, obj.Person.Github.Followers, obj.Company)

		for _, el := range obj.Person.Gravatar.Avatars {
			nothing(el.Url)
		}
	}
}

func BenchmarkEasyJsonUnmarshalMedium(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genMediumFixture())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		mediumFixture := genMediumFixture()
		lexer := &jlexer.Lexer{Data: mediumFixture}
		data := new(MediumPayload)
		data.UnmarshalEasyJSON(lexer)

		nothing(data.Person.Name.FullName, data.Person.Github.Followers, data.Company)

		for _, el := range data.Person.Gravatar.Avatars {
			nothing(el.Url)
		}
	}
}

func BenchmarkJSONUnmarshalMedium(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genMediumFixture())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj MediumPayload
		mediumFixture := genMediumFixture()
		json.Unmarshal(mediumFixture, &obj)

		nothing(obj.Person.Name.FullName, obj.Person.Github.Followers, obj.Company)

		for _, el := range obj.Person.Gravatar.Avatars {
			nothing(el.Url)
		}
	}
}

func BenchmarkJsonIterUnmarshalMedium(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genMediumFixture())))
	b.ResetTimer()

	var json = jsoniter.ConfigFastest
	for i := 0; i < b.N; i++ {
		var obj MediumPayload
		mediumFixture := genMediumFixture()
		json.Unmarshal(mediumFixture, &obj)

		nothing(obj.Person.Name.FullName, obj.Person.Github.Followers, obj.Company)

		for _, el := range obj.Person.Gravatar.Avatars {
			nothing(el.Url)
		}
	}
}

func BenchmarkGOJSONUnmarshalSmall(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genSmallFixture())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj SmallPayload
		smallFixture := genSmallFixture()
		gojson.Unmarshal(smallFixture, &obj)

		nothing(obj.Uuid, obj.Tz, obj.Ua, obj.St)
	}
}

func BenchmarkJsonParserUnmarshalSmall(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genSmallFixture())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		smallFixture := genSmallFixture()
		jsonparser.Get(smallFixture, "uuid")
		jsonparser.GetInt(smallFixture, "tz")
		jsonparser.Get(smallFixture, "ua")
		jsonparser.GetInt(smallFixture, "st")

		nothing()
	}
}

func BenchmarkGoJayUnmarshalSmall(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genSmallFixture())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj SmallPayload
		smallFixture := genSmallFixture()
		gojay.UnmarshalJSONObject(smallFixture, &obj)

		nothing(obj.Uuid, obj.Tz, obj.Ua, obj.St)
	}
}

func BenchmarkEasyJsonUnmarshalSmall(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genSmallFixture())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		smallFixture := genSmallFixture()
		lexer := &jlexer.Lexer{Data: smallFixture}
		data := new(SmallPayload)
		data.UnmarshalEasyJSON(lexer)

		nothing(data.Uuid, data.Tz, data.Ua, data.St)
	}
}

func BenchmarkJSONUnmarshalSmall(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genSmallFixture())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj SmallPayload
		smallFixture := genSmallFixture()
		json.Unmarshal(smallFixture, &obj)

		nothing(obj.Uuid, obj.Tz, obj.Ua, obj.St)
	}
}

func BenchmarkJsonIterUnmarshalSmall(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genSmallFixture())))
	b.ResetTimer()

	var json = jsoniter.ConfigFastest
	for i := 0; i < b.N; i++ {
		var obj SmallPayload
		smallFixture := genSmallFixture()
		json.Unmarshal(smallFixture, &obj)

		nothing(obj.Uuid, obj.Tz, obj.Ua, obj.St)
	}
}

func BenchmarkJSONUnmarshalTest(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genTestData())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj TestStruct
		testdata := genTestData()
		json.Unmarshal(testdata, &obj)

		nothing(obj.A)
	}
}

func BenchmarkJsonIterUnmarshalTest(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genTestData())))
	b.ResetTimer()

	var json = jsoniter.ConfigFastest
	for i := 0; i < b.N; i++ {
		var obj TestStruct
		testdata := genTestData()
		json.Unmarshal(testdata, &obj)

		nothing(obj.A)
	}
}

func BenchmarkGOJSONUnmarshalTest(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genTestData())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj TestStruct
		testdata := genTestData()
		gojson.Unmarshal(testdata, &obj)

		nothing(obj.A)
	}
}

func BenchmarkGJSONUnmarshalTest(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genTestData())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj TestStruct
		teststr := string(genTestData())
		gjson.Get(teststr, "a").String()

		nothing(obj.A)
	}
}

func BenchmarkJsonParserUnmarshalTest(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genTestData())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testdata := genTestData()
		jsonparser.Get(testdata, "a")

		nothing()
	}
}

func BenchmarkEasyJsonUnmarshalTest(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genTestData())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testdata := genTestData()
		lexer := &jlexer.Lexer{Data: testdata}
		data := new(TestStruct)
		data.UnmarshalEasyJSON(lexer)

		nothing(data.A)
	}
}

func BenchmarkGOJSONUnmarshalTestLarge(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genLargeTestData())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var obj TestLargeStruct
		testLargeData := genLargeTestData()
		gojson.Unmarshal(testLargeData, &obj)
	}
}

func BenchmarkEasyJsonUnmarshalTestLarge(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(len(genLargeTestData())))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		testLargeData := genLargeTestData()
		lexer := &jlexer.Lexer{Data: testLargeData}
		data := new(TestLargeStruct)
		data.UnmarshalEasyJSON(lexer)
	}
}

func BenchmarkGOJSONMarshal(b *testing.B) {
	b.Run("large", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(genLargeFixture())))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, err := gojson.Marshal(&largeObject)
			nothing(data, err)
		}
	})

	b.Run("medium", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(genMediumFixture())))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, err := gojson.Marshal(&mediumObject)
			nothing(data, err)
		}
	})

	b.Run("small", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(genSmallFixture())))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, err := gojson.Marshal(&smallObject)
			nothing(data, err)
		}
	})
}

func BenchmarkEasyJSONMarshal(b *testing.B) {
	b.Run("large", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(genLargeFixture())))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, err := easyjson.Marshal(&largeObject)
			nothing(data, err)
		}
	})

	b.Run("medium", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(genMediumFixture())))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, err := easyjson.Marshal(&mediumObject)
			nothing(data, err)
		}
	})

	b.Run("small", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(genSmallFixture())))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, err := easyjson.Marshal(&smallObject)
			nothing(data, err)
		}
	})
}

func BenchmarkJSONMarshal(b *testing.B) {
	b.Run("large", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(genLargeFixture())))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, err := json.Marshal(&largeObject)
			nothing(data, err)
		}
	})

	b.Run("medium", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(genMediumFixture())))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, err := json.Marshal(&mediumObject)
			nothing(data, err)
		}
	})

	b.Run("small", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(int64(len(genSmallFixture())))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			data, err := json.Marshal(&smallObject)
			nothing(data, err)
		}
	})
}
