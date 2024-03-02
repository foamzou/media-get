package kuwo

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/foamzou/audio-get/logger"
	"github.com/foamzou/audio-get/utils"
)

var codeJS = `
function h(t, e) {
    if (e == null || e.length <= 0) {
      console.log("Please enter a password with which to encrypt the message.");
      return null;
    }
    for (var n = "", i = 0; i < e.length; i++) {
      n += e.charCodeAt(i).toString();
    }
    var r = Math.floor(n.length / 5);
    var o = parseInt(n.charAt(r) + n.charAt(r * 2) + n.charAt(r * 3) + n.charAt(r * 4) + n.charAt(r * 5));
    var l = Math.ceil(e.length / 2);
    var c = Math.pow(2, 31) - 1;
    if (o < 2) {
      console.log("Algorithm cannot find a suitable hash. Please choose a different password. \nPossible considerations are to choose a more complex or longer password.");
      return null;
    }
    var d = Math.round(Math.random() * 1000000000) % 100000000;
    for (n += d; n.length > 10;) {
      n = (parseInt(n.substring(0, 10)) + parseInt(n.substring(10, n.length))).toString();
    }
    n = (o * n + l) % c;
    var h = "";
    var f = "";
    for (i = 0; i < t.length; i++) {
      f += (h = parseInt(t.charCodeAt(i) ^ Math.floor(n / c * 255))) < 16 ? "0" + h.toString(16) : h.toString(16);
      n = (o * n + l) % c;
    }
    for (d = d.toString(16); d.length < 8;) {
      d = "0" + d;
    }
    return f += d;
  }

  function v(t, cookie) {
    const e = cookie;
    let n = e.indexOf(t + "=");
    if (n !== -1) {
      n = n + t.length + 1;
      let r = e.indexOf(";", n);
      return r === -1 && (r = e.length), unescape(e.substring(n, r));
    }
    return null;
  }

  function main(cookie, hmCookie) {
      dd = v(hmCookie, cookie)
      return h(dd, hmCookie);
  }
  console.log(main(process.argv[2], process.argv[3]));
`

// genSecretHeader
// hmCookie format: Hm_Iuvt_cdb524f42f0cer9b268e4v7y735ewrq2324
func genSecretHeader(cookie string, hmCookie string) string {
	// 计算jsCode的MD5哈希值
	hasher := md5.New()
	hasher.Write([]byte(codeJS))
	// 将哈希值转换为十六进制并得到文件名
	md5Hash := hex.EncodeToString(hasher.Sum(nil))
	filename := filepath.Join(os.TempDir(), "media-get-"+md5Hash+".js")

	// 如果文件不存在则创建
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		err := ioutil.WriteFile(filename, []byte(codeJS), 0644)
		if err != nil {
			logger.Error("Failed to create file", filename, ":", err)
			return ""
		}
	}
	secret, err := utils.ExecCmd("node", filename, cookie, hmCookie)
	if err != nil {
		logger.Error("Failed to execute js code:", err)
		return ""
	}
	return strings.Trim(secret, "\n")
}

func nodeEnvAvailable() bool {
	_, err := utils.ExecCmd("node", "-v")
	return err == nil
}
