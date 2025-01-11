/*
Copyright © 2025 qshuai <qishuai231@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cobra"
)

var (
	claims []string
	alg    string
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign JWT using alg, secret and claims",
	Long: `By providing secret data or a private key file, along with all claims information and 
a specified signing algorithm, a JWT is generated and output to the console.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c := parseClaims(claims)
		method := jwt.GetSigningMethod(alg)
		if method == nil {
			fatalf("invalid alg: %s", alg)
		}

		var key interface{}
		switch method.(type) {
		case *jwt.SigningMethodHMAC:
			key = []byte(secret)
		case *jwt.SigningMethodECDSA, *jwt.SigningMethodRSA:
			file, err := os.OpenFile(keyfile, os.O_RDONLY, 0750)
			if err != nil {
				fatalf("open private key file err: %v", err)
			}
			defer file.Close()

			data, err := io.ReadAll(file)
			if err != nil {
				fatalf("read private key file err: %v", err)
			}

			switch method.(type) {
			case *jwt.SigningMethodECDSA:
				key, err = jwt.ParseECPrivateKeyFromPEM(data)
			case *jwt.SigningMethodRSA:
				key, err = jwt.ParseRSAPrivateKeyFromPEM(data)
			}

			if err != nil {
				fatalf("parse private key err: %v", err)
			}
		}

		token := jwt.NewWithClaims(method, jwt.MapClaims(c))
		signedString, err := token.SignedString(key)
		if err != nil {
			return err
		}
		fmt.Println(signedString)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(signCmd)

	signCmd.Flags().StringVar(&alg, "alg", "HS256", "specify signature method")
	signCmd.Flags().StringArrayVarP(&claims, "claims", "c", nil, "provide claims will fill payload. claims name and value should be seperated by colon")
	signCmd.Flags().StringVarP(&secret, "secret", "s", "", "the secret string specified to sign or parse token")
	signCmd.Flags().StringVarP(&keyfile, "key-file", "f", "private_key.pem", "the private key file")
	signCmd.MarkFlagsOneRequired("secret", "key-file")
	signCmd.MarkFlagsMutuallyExclusive("secret", "key-file")
}

func parseClaims(claims []string) map[string]interface{} {
	res := make(map[string]interface{}, len(claims))
	for _, claim := range claims {
		items := strings.SplitN(claim, ":", 2)
		if len(items) != 2 {
			fatalf("malformed claim(usage: --claims foo:bar)")
		}

		number, err := strconv.Atoi(strings.TrimSpace(items[1]))
		if err != nil {
			res[strings.TrimSpace(items[0])] = strings.TrimSpace(items[1])
			continue
		}
		res[strings.TrimSpace(items[0])] = number
	}

	return res
}
