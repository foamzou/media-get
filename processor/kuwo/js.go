package kuwo

import (
	"fmt"

	"github.com/dop251/goja"
	"github.com/foamzou/audio-get/logger"
)

var codeJS = `
// the function is for fix 1.7118116959910114e+99 -> 1.71181169599101e+99
function normalizeParseIntToString(numberString) {
  const match = numberString.match(/^(\d+\.\d{14})(\d*)e([+-]?\d+)$/);

  if (match) {
    return match[1] + "e" + match[3];
  }

  return numberString;
}

function encryptMessage(input, password) {
    const debugList = [];
    const dl = (m) => {
        debugList.push(m);
    }
    const ed = () => {
        return debugList.join("\n");
    }

    if (password == null || password.length === 0) {
      console.log("Please enter a password with which to encrypt the message.");
      return null;
    }
  
    let numericPassword = "";
    for (let i = 0; i < password.length; i++) {
      numericPassword += password.charCodeAt(i).toString();
    }
    const r = Math.floor(numericPassword.length / 5);
    const o = parseInt(
      numericPassword.charAt(r) +
        numericPassword.charAt(2 * r) +
        numericPassword.charAt(3 * r) +
        numericPassword.charAt(4 * r) +
        numericPassword.charAt(5 * r)
    );
  
    const l = Math.ceil(password.length / 2);
    const c = Math.pow(2, 31) - 1;
  
    if (o < 2) {
      console.log(
        "Algorithm cannot find a suitable hash. Please choose a different password.\nPossible considerations are to choose a more complex or longer password."
      );
      return null;
    }
  
    let randomNum = Math.round(1e9 * Math.random()) % 1e8;
	randomNum = 91610361; // fixed this value
    numericPassword += randomNum.toString();

    let counter = 0;
    while (numericPassword.length > 10) {
        dl("sss|" + numericPassword)
        numericPassword = normalizeParseIntToString((
            parseInt(numericPassword.substring(0, 10)) +
            parseInt(numericPassword.substring(10))
            ).toString());
        dl("eee|" + numericPassword)
        if (counter++ > 20) {
			break;
    	}
	}
    dl(numericPassword.toString())
    
    numericPassword = (o * numericPassword + l) % c;
    dl(numericPassword.toString())

    let encryptedMessage = "";
    let hexValue = "";
    for (let i = 0; i < input.length; i++) {
      hexValue = parseInt(input.charCodeAt(i) ^ Math.floor((numericPassword / c) * 255));
      if (hexValue < 16) {
        encryptedMessage += "0" + hexValue.toString(16);
      } else {
        encryptedMessage += hexValue.toString(16);
      }
  
      numericPassword = (o * numericPassword + l) % c;
    }
    randomNum = randomNum.toString(16);
  
    while (randomNum.length < 8) {
      randomNum = "0" + randomNum;
    }
	dl(randomNum)
    //return ed();
    return encryptedMessage + randomNum;
  }

function findCookieValue(cookie, name) {
    const cookieString = cookie;
    const cookieName = name + '=';
    let startIndex = cookieString.indexOf(cookieName);
    
    if (startIndex !== -1) {
      startIndex = startIndex + cookieName.length;
      let endIndex = cookieString.indexOf(";", startIndex);
      
      if (endIndex === -1) {
        endIndex = cookieString.length;
      }
      
      return decodeURIComponent(cookieString.substring(startIndex, endIndex));
    }
    
    return null;
  }

function main(cookie, hmCookie) {
    return encryptMessage(findCookieValue(cookie, hmCookie), hmCookie);
}
`

// genSecretHeader
// hmCookie format: Hm_Iuvt_cdb524f42f0cer9b268e4v7y735ewrq2324
func genSecretHeader(cookie string, hmCookie string) string {
	vm := goja.New()

	_, err := vm.RunString(codeJS) // executes a script on the global context
	if err != nil {
		logger.Warn("genSecretHeader failed", err)
		return ""
	}
	val, err := vm.RunString(fmt.Sprintf("main('%s', '%s')", cookie, hmCookie))
	if err != nil {
		logger.Warn("genSecretHeader failed", err)
		return ""
	}

	//nolint
	return val.Export().(string)
}
