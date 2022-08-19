package gen_xb_js

const XB_JS = `

function _0x237a87(_0x5c3d2a) {
    // console.log('_0x237a87-param', _0x5c3d2a)
    var _0x50ff23 =   [null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,0,1,2,3,4,5,6,7,8,9,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,null,10,11,12,13,14,15]
    for (var _0x1204d6 = _0x5c3d2a['length'] >> 0xa1e * -0x2 + 0x1937 + 0x4fa * -0x1, _0x700552 = _0x1204d6 << 0x68 + -0xa29 + 0x9c2, _0x1673dd = new Uint8Array(_0x1204d6), _0x19eb71 = 0xe9e + -0x167 + -0xd37, _0x249396 = 0x1 * 0x104a + 0xaa9 + 0x1af3 * -0x1; _0x249396 < _0x700552;)
        _0x1673dd[_0x19eb71++] = _0x50ff23[_0x5c3d2a['charCodeAt'](_0x249396++)] << -0x1938 + 0x10c4 * -0x1 + 0x2a00 | _0x50ff23[_0x5c3d2a['charCodeAt'](_0x249396++)];
    return _0x1673dd;
}

function _0x238632(_0x4cdef5, _0x268c9c) {
    let _0x2b4641, _0xbb44d8 = [], _0x138ea3 = 0x140 + -0x2038 + -0x7be * -0x4, _0xc9f8ff = '';
    for (let _0x332e7e = 0x17 * -0x8 + -0x910 + 0x2 * 0x4e4; _0x332e7e < 0x24e + -0x2533 + 0xbf7 * 0x3; _0x332e7e++)
        _0xbb44d8[_0x332e7e] = _0x332e7e;

    // console.log("FIRST", JSON.stringify(_0xbb44d8))
    // console.log("LEN OF", _0x4cdef5, _0x4cdef5.length)
    let intArr = [];
    for (let _0x369701 = 0x77a * -0x5 + -0x3 * 0x689 + -0x12ff * -0x3; _0x369701 < 0x1b3 * 0x12 + -0x1e3 + 0x3f5 * -0x7; _0x369701++)
        _0x138ea3 = (_0x138ea3 + _0xbb44d8[_0x369701] + _0x4cdef5['charCodeAt'](_0x369701 % _0x4cdef5['length'])) % (0x16d0 + 0x12 * -0xf4 + 0x254 * -0x2),
            intArr.push(_0x4cdef5['charCodeAt'](_0x369701 % _0x4cdef5['length'])),
            _0x2b4641 = _0xbb44d8[_0x369701],
            _0xbb44d8[_0x369701] = _0xbb44d8[_0x138ea3],
            _0xbb44d8[_0x138ea3] = _0x2b4641;
// console.log("intArr", JSON.stringify(intArr))
//     console.log('_0xbb44d8', JSON.stringify(_0xbb44d8)) //PASS
    let _0x1a0256 = 0x1ca3 + 0x1a34 + -0x36d7;
    _0x138ea3 = 0xdc * 0x28 + -0x15d * 0x1 + 0x3ab * -0x9;
    var aa = [];
    for (let _0x1b288d = 0x9 * 0x349 + 0x1e7f + -0x3c10 * 0x1; _0x1b288d < _0x268c9c['length']; _0x1b288d++)
        _0x1a0256 = (_0x1a0256 + (0x14ef * -0x1 + -0x1752 + -0x37 * -0xce)) % (-0x312 + 0x171d + -0x130b),
            _0x138ea3 = (_0x138ea3 + _0xbb44d8[_0x1a0256]) % (0x3 * 0x66d + -0x3 * -0x13d + -0x15fe),
            _0x2b4641 = _0xbb44d8[_0x1a0256],
            _0xbb44d8[_0x1a0256] = _0xbb44d8[_0x138ea3],
            _0xbb44d8[_0x138ea3] = _0x2b4641,
            aa.push(_0x268c9c['charCodeAt'](_0x1b288d) ^ _0xbb44d8[(_0xbb44d8[_0x1a0256] + _0xbb44d8[_0x138ea3]) % (-0x1db3 + 0x1 * -0x733 + -0x15 * -0x1ce)]),
            _0xc9f8ff += String['fromCharCode'](_0x268c9c['charCodeAt'](_0x1b288d) ^ _0xbb44d8[(_0xbb44d8[_0x1a0256] + _0xbb44d8[_0x138ea3]) % (-0x1db3 + 0x1 * -0x733 + -0x15 * -0x1ce)]);
    // console.log("aa", JSON.stringify(aa));
    return _0xc9f8ff;
}

function md5(str) {
    var str =  CryptoJS.enc.Utf8.parse(str);
    var str = CryptoJS.MD5(str).toString();
    return str
}

function hex_md5(str) {
    var str =  CryptoJS.enc.Hex.parse(str);
    var str = CryptoJS.MD5(str).toString();
    return str
}

function get_one_list(url_para) {
    var str_1 = md5(url_para);
    // console.log("get_one_list,url_para", url_para)
    // console.log("get_one_list", str_1)
    str_1 = hex_md5(str_1);
    return _0x237a87(str_1)
}

function get_two_list() {
    var str_2 = hex_md5("d41d8cd98f00b204e9800998ecf8427e");
    return _0x237a87(str_2)
}

function get_three_list(ua) {
    var str_3= _0x238632("\u0001\u0001\b", ua);
    //str_3 = Buffer.from(str_3, 'ASCII').toString('base64');
    str_3 = ethereumjs.Buffer.Buffer.from(str_3,'ascii').toString('base64');
    // console.log('base64', str_3)
    str_3 = md5(str_3);
    // console.log("md5", str_3)

    str_3 = _0x237a87(str_3);
    return str_3

}

function get_time_sign(time_now) {
    var list_time = [];
    for (var i of [24, 16, 8, 0]){
        var num_1 = time_now >> i;
        list_time = list_time.concat(num_1 & 255)
    }
    return list_time
}

function get_num_sign() {
    var num = 3963386674;
    var list_time = [];
    for (var i of [24, 16, 8, 0]){
        var num_1 = num >> i;
        list_time = list_time.concat(num_1 & 255)
    }
    return list_time
}

function get_last_sign(index_list) {
    var num = 0;
    for (var i of index_list){
        if (num == 0){
            num = i;
            continue
        }
        num = num^i;
    }
    return num
}

function get_index_str(url_para, ua, time_now) {
    var get_one_list_list = get_one_list(url_para);
    var get_two_list_list = get_two_list();
    var get_three_list_list = get_three_list(ua);
    var index_list_1 = [64,1.00390625,1,8,get_one_list_list[14], get_one_list_list[15],get_two_list_list[14], get_two_list_list[15],get_three_list_list[14], get_three_list_list[15]];
    var index_list_2 = get_time_sign(time_now);
    var index_list_3 = get_num_sign();
    // console.log("index_list_1:", index_list_1)
    // console.log("index_list_2:", index_list_2)
    // console.log("index_list_3:", index_list_3)

    var index_list = index_list_1.concat(index_list_2,index_list_3);
    var index_list_last = get_last_sign(index_list);
    index_list = index_list.concat(index_list_last);
    var last_str = ""
    for (var i of index_list){
        last_str += String.fromCharCode(i);
    }

    last_str = "\u0002ÿ" + _0x238632("ÿ", last_str);

    return last_str
}

function all_num(last_str){
    var num_list = [];
    for (var i=0;i<last_str.length;i++) {
        var num_list = num_list.concat(last_str.charCodeAt(i));
    }
    var result = [];
    for (var i=0;i<num_list.length;i+=3) {
        result.push(num_list.slice(i, i+3))
    }
    var result_list = [];
    for (var i of result){
        var num_1 = i[0]<<16;
        var num_2 = i[1]<<8;
        var num_3 = num_1^num_2;
        var num_4 = num_3^i[2];
        result_list.push(num_4)
    }
    return result_list

}

function get_xb(all_num_list) {
    var _str = "Dkdpgh4ZKsQB80/Mfvw36XI1R25-WUAlEi7NLboqYTOPuzmFjJnryx9HVGcaStCe=";
    var result_str = "";
    for (var i of all_num_list){
        for (var j of [[16515072,18],[258048,12],[4032,6],[63,0]]) {
            var num_1 = i & j[0];
            var num_2 = num_1>>j[1];
            result_str += _str[num_2];
        }
    }
    return result_str
}

function xb_main(url_para, ua) {
    var time_now = (new Date().getTime()/1000).toFixed(3);
    var last_str = get_index_str(url_para, ua, time_now);
    var all_num_list = all_num(last_str);
    var xb = get_xb(all_num_list);
    return xb
}

`
