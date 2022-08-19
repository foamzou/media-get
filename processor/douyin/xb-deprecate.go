package douyin

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
	"time"

	"github.com/foamzou/audio-get/logger"
)

//const CryptoJS = require('crypto-js');
//var MD5 = require("crypto-js/md5");
// var Base64 = require("crypto-js/enc-base64.js");

func _0x237a87(_0x5c3d2a string) []byte {
	//fmt.Println("_0x237a87-param", _0x5c3d2a)

	_0x50ff23 := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 10, 11, 12, 13, 14, 15}
	//_0x50ff23 := []byte{nil,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,0,1,2,3,4,5,6,7,8,9,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,10,11,12,13,14,15}

	_0x1204d6 := len(_0x5c3d2a) / 2
	_0x700552 := len(_0x5c3d2a)
	_0x1673dd := make([]byte, _0x1204d6)
	_0x19eb71 := 0
	_0x249396 := 0

	for _0x249396 < _0x700552 {
		f := _0x5c3d2a[_0x249396]
		_0x249396 += 1
		o := _0x5c3d2a[_0x249396]
		_0x249396 += 1

		_0x1673dd[_0x19eb71] = _0x50ff23[f]*16 + _0x50ff23[o]
		_0x19eb71 += 1
	}
	return _0x1673dd
}

func _0x238632(_0x4cdef5 string, _0x268c9c string, twoByte bool) string {
	var _0x2b4641 uint8
	var _0xbb44d8 = make([]uint8, 256)
	var _0x138ea3 uint32 = 0
	var arr []byte

	for i := 0; i < 256; i++ {
		_0xbb44d8[i] = uint8(i)
	}
	//fmt.Println("FIRST", _0xbb44d8)
	lenOf_0x4cdef5 := len(_0x4cdef5)
	if _0x4cdef5 == "ÿ" {
		lenOf_0x4cdef5 = 1
	}
	//fmt.Println("LEN OF", _0x4cdef5, len(_0x4cdef5))

	for i := 0; i < 256; i++ {
		var charCode uint32
		if _0x4cdef5 == "ÿ" {
			charCode = 255
		} else {
			charCode = uint32(_0x4cdef5[i%lenOf_0x4cdef5])
		}

		_0x138ea3 = (_0x138ea3 + uint32(_0xbb44d8[i]) + charCode) % 256

		_0x2b4641 = _0xbb44d8[i]
		_0xbb44d8[i] = _0xbb44d8[_0x138ea3]
		_0xbb44d8[_0x138ea3] = _0x2b4641
	}
	//fmt.Println("intArr", intArr)

	//fmt.Println("_0xbb44d8", _0xbb44d8) //PASS

	_0x1a0256 := 0
	_0x138ea3 = 0
	//fmt.Println("_0x268c9c", len(_0x268c9c), []byte(_0x268c9c))
	for i := 0; i < len(_0x268c9c); i++ {
		_0x1a0256 = (_0x1a0256 + (0x14ef*-0x1 + -0x1752 + -0x37*-0xce)) % (-0x312 + 0x171d + -0x130b)
		_0x138ea3 = (_0x138ea3 + uint32(_0xbb44d8[_0x1a0256])) % 256
		_0x2b4641 = _0xbb44d8[_0x1a0256]
		_0xbb44d8[_0x1a0256] = _0xbb44d8[_0x138ea3]
		_0xbb44d8[_0x138ea3] = _0x2b4641
		arr = append(arr, uint8(uint32(_0x268c9c[i])^uint32(_0xbb44d8[(uint32(_0xbb44d8[_0x1a0256])+uint32(_0xbb44d8[_0x138ea3]))%256])))
	}
	//fmt.Println("aa", arr)
	if twoByte {
		return string(convertForCompatibleAsJSFromCharCode(arr))
	}
	return string(arr)
	//return string(arr)
	//fmt.Println("arr", convertForCompatibleAsJSFromCharCode(arr))
	//return string(convertForCompatibleAsJSFromCharCode(arr))
}

func md5Str(str string) string {
	data := []byte(str)
	hash := md5.Sum(data)
	md5str := fmt.Sprintf("%x", hash)

	return md5str
}

func hex_md5(str string) string {
	hx, err := hex.DecodeString(str)
	if err != nil {
		logger.Warn("hex.DecodeString error: ", err)
		return ""
	}
	hash := md5.Sum(hx)

	return fmt.Sprintf("%x", hash)
}

func get_one_list(urlParam string) []byte {
	var str1 = md5Str(urlParam)
	//fmt.Println("get_one_list,url_para", urlParam)
	//fmt.Println("get_one_list", str1)

	str1 = hex_md5(str1)
	return _0x237a87(str1)
}

func get_two_list() []byte {
	var str2 = hex_md5("d41d8cd98f00b204e9800998ecf8427e")
	return _0x237a87(str2)
}

func get_three_list(ua string) []byte {
	str_3 := _0x238632(string([]byte{1, 1, 8}), ua, false)
	str_3 = base64.StdEncoding.EncodeToString([]byte(str_3))
	//fmt.Println("base64", str_3)

	str_3 = md5Str(str_3)
	//fmt.Println("md5", str_3)
	return _0x237a87(str_3)
}

func get_time_sign(timeNow float64) []byte {
	var list_time []byte
	timeNowInt := int(math.Round(timeNow))
	for _, b := range []byte{24, 16, 8, 0} {
		num_1 := timeNowInt >> b
		list_time = append(list_time, byte(int(num_1)&255))
	}

	return list_time
}

func get_num_sign() []byte {
	return []byte{236, 60, 123, 50}
	//var num = 3963386674;
	//var list_time = [];
	//for (var i of [24, 16, 8, 0]){
	//var num_1 = num >> i;
	//list_time = list_time.concat(num_1 & 255)
	//}
	//return list_time
}

func get_last_sign(index_list []byte) byte {
	var num byte = 0
	for _, i := range index_list {
		if num == 0 {
			num = i
		} else {
			num = num ^ i
		}
	}
	return num
}

func convertForCompatibleAsJSFromCharCode(bytes []byte) []byte {
	// for compatible with JS String.fromCharCode
	// https://www.smashingmagazine.com/2012/06/all-about-unicode-utf8-character-sets/
	var index_list_new []byte
	for _, i := range bytes {
		var newI []byte
		if i <= 127 {
			newI = []byte{i}
		} else if i <= 191 {
			newI = []byte{194, i}
		} else {
			newI = []byte{195, (i - 192) + 64*2}
		}
		index_list_new = append(index_list_new, newI...)
	}
	return index_list_new
}

func get_index_str(url_para string, ua string, time_now float64) string {
	get_one_list_list := get_one_list(url_para)
	get_two_list_list := get_two_list()
	get_three_list_list := get_three_list(ua)
	index_list_1 := []byte{64, 1, 1, 8, get_one_list_list[14], get_one_list_list[15], get_two_list_list[14], get_two_list_list[15], get_three_list_list[14], get_three_list_list[15]}
	index_list_2 := get_time_sign(time_now)
	index_list_3 := get_num_sign()

	//fmt.Println("index_list_1:", index_list_1)
	//fmt.Println("index_list_2:", index_list_2)
	//fmt.Println("index_list_3:", index_list_3)

	index_list := append(index_list_1, index_list_2...)
	index_list = append(index_list, index_list_3...)

	indexListLast := get_last_sign(index_list)

	index_list = append(index_list, indexListLast)
	//fmt.Println("before encode:", index_list)

	//fmt.Println("last_str", convertForCompatibleAsJSFromCharCode(index_list))
	lastStr := "\u0002ÿ" + _0x238632("ÿ", string(convertForCompatibleAsJSFromCharCode(index_list)), true)
	//fmt.Println("lastStr", []byte(lastStr))
	return lastStr
}

func all_num(last_str string) []int {
	num_list_byte := []byte(last_str)
	var num_list []int
	for i := 0; i < len(num_list_byte); i++ {
		num_list = append(num_list, int(num_list_byte[i]))
	}
	var result [][]int
	// push to result per 3 items
	for i := 0; i < len(num_list); i += 3 {
		result = append(result, num_list[i:i+3])
	}

	var resultList []int
	for _, i := range result {
		num_1 := i[0] << 16
		num_2 := i[1] << 8
		num_3 := num_1 ^ num_2
		num_4 := num_3 ^ i[2]
		resultList = append(resultList, num_4)
	}
	return resultList
}

func get_xb(all_num_list []int) string {
	_str := "Dkdpgh4ZKsQB80/Mfvw36XI1R25-WUAlEi7NLboqYTOPuzmFjJnryx9HVGcaStCe="
	resultStr := ""
	for _, i := range all_num_list {
		for _, j := range [][]int{{16515072, 18}, {258048, 12}, {4032, 6}, {63, 0}} {
			num1 := i & j[0]
			num2 := num1 >> j[1]
			resultStr += string(_str[num2])
		}
	}
	return resultStr
}

func xb_main(url_para string, ua string) string {
	// get unix timestamp now
	time_now := float64(time.Now().Unix())
	var last_str = get_index_str(url_para, ua, time_now)
	var all_num_list = all_num(last_str)
	var xb = get_xb(all_num_list)
	return xb
}
