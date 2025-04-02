package kuwo

import (
	"fmt"
	"math"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type crypto struct {
	cookie string
	key    string
}

func (c *crypto) encrypt(text string, key string) string {
	if len(key) == 0 {
		fmt.Println("Please enter a password with which to encrypt the message.")
		return ""
	}

	// Build a numerical string from the key by converting each character to its char code.
	numericKey := ""
	for _, ch := range key {
		numericKey += strconv.Itoa(int(ch))
	}

	// multiplier
	r := len(numericKey) / 5
	multiplierStr := string(numericKey[r]) + string(numericKey[r*2]) + string(numericKey[r*3]) + string(numericKey[r*4]) + string(numericKey[r*5])
	multiplier, err := strconv.ParseFloat(multiplierStr, 64)
	if err != nil {
		fmt.Println("Error parsing o:", err)
		return ""
	}

	if multiplier < 2 {
		fmt.Println("Algorithm cannot find a suitable hash. Please choose a different password." +
			"\nPossible considerations are to choose a more complex or longer password.")
		return ""
	}

	// addend
	addend := math.Ceil(float64(len(key)) / 2.0)

	// modulus = 2^31 - 1
	modulus := math.Pow(2, 31) - 1

	salt := c.createSalt()
	baseNumber := c.calculateBaseNumber(numericKey, salt)
	baseNumber = math.Mod(baseNumber*multiplier+addend, modulus)
	// return fmt.Sprintf("%d,%d,%d,%d,%d", d, N2, l, c, o)

	encryptedResult := ""
	// Process each character in the text.
	for i := 0; i < len(text); i++ {
		// Get the character code.
		charCode := float64(text[i])
		// Calculate the XOR value using floor(N / modulus * 255).
		xorVal := math.Floor(baseNumber / modulus * 255)
		// XOR the character code with the computed value.
		xorResult := int(charCode) ^ int(xorVal)
		// Convert the result to its hexadecimal representation.
		hexPart := fmt.Sprintf("%x", xorResult)
		if xorResult < 16 {
			hexPart = "0" + hexPart
		}
		encryptedResult += hexPart

		// Update the base number again for the next character.
		baseNumber = math.Mod(multiplier*baseNumber+addend, modulus)
	}

	// Convert d to hex string and pad with zeros until it is 8 characters.
	dHex := fmt.Sprintf("%x", int(salt))
	for len(dHex) < 8 {
		dHex = "0" + dHex
	}

	return encryptedResult + dHex
}

func (c *crypto) createSalt() float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return math.Mod(math.Round(r.Float64()*1000000000), 100000000)
}

// calculateBaseNumber computes a base number from the input numeric string by appending a random salt
// and repeatedly reducing the digit length until it is 10 or fewer digits.
// Tricks: Javascript Number.prototype.toString() will convert numbers to scientific notation if they are too large.
func (c *crypto) calculateBaseNumber(baseNumber string, salt float64) float64 {
	baseNumber = baseNumber + strconv.FormatFloat(salt, 'f', 0, 64)

	for len(baseNumber) > 10 {
		firstPart, _ := c.jsParseInt10(baseNumber[:10])
		secondPart, _ := c.jsParseInt10(baseNumber[10:])
		addedUp := firstPart + secondPart
		if addedUp >= 1e21 {
			baseNumber = fmt.Sprintf("%.14e", addedUp)
		} else {
			baseNumber = fmt.Sprintf("%.0f", addedUp)
		}
	}
	baseNumberFloat, err := strconv.ParseFloat(baseNumber, 64)
	if err != nil {
		fmt.Println("Error parsing base number:", err)
		return 0
	}
	return baseNumberFloat
}

// JSParseInt10 mimics JavaScript's parseInt for base 10 without handling signs.
// It trims leading whitespace and stops at the first non-digit.
func (c *crypto) jsParseInt10(s string) (float64, error) {
	// Trim whitespace.
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("NaN")
	}

	// Extract consecutive digits.
	digits := ""
	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			digits += string(ch)
		} else {
			// Stop at first non-digit.
			break
		}
	}

	if digits == "" {
		// No valid digits found.
		return 0, fmt.Errorf("NaN")
	}

	n, err := strconv.ParseFloat(digits, 64)
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (c *crypto) createRawHeader(key string, cookie string) string {
	search := key + "="
	idx := strings.Index(cookie, search)
	if idx != -1 {
		start := idx + len(search)
		end := strings.Index(cookie[start:], ";")
		var sub string
		if end == -1 {
			sub = cookie[start:]
		} else {
			sub = cookie[start : start+end]
		}
		// Unescape using URL unescaping.
		unescaped, err := url.QueryUnescape(sub)
		if err != nil {
			return sub
		}
		return unescaped
	}
	return ""
}

// genSecretHeader
// hmCookie format: Hm_Iuvt_cdb524f42f0cer9b268e4v7y735ewrq2324
func (c *crypto) GenSecretHeader() string {
	header := c.createRawHeader(c.key, c.cookie)
	return c.encrypt(header, c.key)
}
