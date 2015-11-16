package stringCommons

/*
 * pystring
 *
 * Python-like string methods for Go
 * Alexander Rødseth <rodseth@gmail.com>
 *
 * July 2012
 * September 2013
 *
 * Qiukeren <qiuker521@163.com>
 *
 * Started to do this
 *
 * GPL2
 */


import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

const (
	/* string.ascii_letters */
	ASCII_letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	/* string.ascii_lowercase */
	ASCII_lowercase = "abcdefghijklmnopqrstuvwxyz"
	/* string.ascii_uppercase */
	ASCII_uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	/* string.digits */
	Digits = "0123456789"
	/* string.hexdigits */
	HexDigits = "0123456789abcdefABCDEF"
	/* string.octdigits */
	OctDigits = "01234567"
	/* string.punctuation */
	Punctuation = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	/* string.printable */
	Printable = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ \t\n\r\x0b\x0c"
	/* string.whitespace */
	Whitespace = " \t\n\r\x0b\x0c"
)

func Capitalize(p string) string {
    return strings.Title(p)
}

