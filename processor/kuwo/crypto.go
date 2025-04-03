package kuwo

import (
	"fmt"
	"math"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/foamzou/audio-get/logger"
)

type crypto struct {
	cookie string
	key    string
}

// var hash = parseInt(numericKey.charAt(r) + numericKey.charAt(r * 2) + numericKey.charAt(r * 3) + numericKey.charAt(r * 4) + numericKey.charAt(r * 5));
func (c *crypto) createKeyHash(key string) (float64, error) {
	if len(key) == 0 {
		return 0, fmt.Errorf("please enter a password with which to encrypt the message")
	}
	r := len(key) / 5
	digest := ""
	for i := 1; i <= 5 && r*i < len(key); i++ {
		digest += string(key[r*i])
	}

	hash, err := strconv.ParseFloat(digest, 64)
	if err != nil {
		return 0, err
	}

	if hash < 2 {
		return 0, fmt.Errorf("algorithm cannot find a suitable hash. Please choose a different password. Possible considerations are to choose a more complex or longer password")
	}
	return hash, nil
}

func (c *crypto) encrypt(text string, key string) string {
	if len(key) == 0 {
		logger.Error("Please enter a password with which to encrypt the message.")
		return ""
	}

	// Build a numerical string from the key by converting each character to its char code.
	numericKey := ""
	for _, ch := range key {
		numericKey += strconv.Itoa(int(ch))
	}

	// hash
	hash, err := c.createKeyHash(numericKey)
	if err != nil {
		logger.Error("Error creating key hash:", err)
		return ""
	}

	// addend
	addend := math.Ceil(float64(len(key)) / 2.0)

	// modulus = 2^31 - 1
	modulus := math.Pow(2, 31) - 1

	salt := c.createSalt()
	baseNumber := c.calculateBaseNumber(numericKey, salt)
	baseNumber = math.Mod(baseNumber*hash+addend, modulus)

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
		baseNumber = math.Mod(hash*baseNumber+addend, modulus)
	}

	// Convert salt to hex string and pad with zeros until it is 8 characters.
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
			baseNumber = strconv.FormatFloat(addedUp, 'g', -1, 64)
		} else {
			baseNumber = fmt.Sprintf("%.0f", addedUp)
		}
	}
	baseNumberFloat, err := strconv.ParseFloat(baseNumber, 64)
	if err != nil {
		logger.Error("Error parsing base number:", err)
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
			logger.Error("Error unescaping cookie value:", err)
			return ""
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
