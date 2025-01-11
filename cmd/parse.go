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
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse JWT and output formated information to console",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		token, err := jwt.Parse(args[0], func(token *jwt.Token) (interface{}, error) {
			switch token.Method.(type) {
			case *jwt.SigningMethodHMAC:
				return []byte(secret), nil
			case *jwt.SigningMethodECDSA, *jwt.SigningMethodRSA:
				file, err := os.OpenFile(keyfile, os.O_RDONLY, 0440)
				if err != nil {
					fatalf("read private file err: %v", err)
				}
				defer file.Close()

				data, err := io.ReadAll(file)
				if err != nil {
					fatalf("read private key file err: %v", err)
				}

				switch token.Method.(type) {
				case *jwt.SigningMethodECDSA:
					privkey, err := jwt.ParseECPrivateKeyFromPEM(data)
					if err != nil {
						fatalf("parse ecdsa private key err: %v", err)
					}
					return privkey.Public(), nil
				case *jwt.SigningMethodRSA:
					privkey, err := jwt.ParseRSAPrivateKeyFromPEM(data)
					if err != nil {
						fatalf("parse rsa private key err: %v", err)
					}
					return privkey.Public(), nil
				default:
					return nil, fmt.Errorf("unreached path")
				}
			default:
				return nil, fmt.Errorf("unimplementated sign method: %s", token.Method.Alg())
			}
		}, jwt.WithJSONNumber())
		if err != nil {
			fatal(err)
		}

		fmt.Printf("Token Validation: %t\n", token.Valid)

		fmt.Println("HEADER:")
		for key, value := range token.Header {
			fmt.Printf("\t%q: %v\n", key, value)
		}
		fmt.Println()

		fmt.Println("PAYLOAD:")
		switch v := token.Claims.(type) {
		case jwt.MapClaims:
			for key, value := range v {
				switch vv := value.(type) {
				case json.Number:
					fmt.Printf("\t%q: %s\n", key, vv)
				default:
					fmt.Printf("\t%q: %v\n", key, value)
				}
			}
		default:
			// 过期时间
			exp, err := v.GetExpirationTime()
			if err == nil {
				fmt.Printf("\t%q: %v\n", "exp", exp.UnixMilli())
			}

			// iat: issue at
			iat, err := v.GetIssuedAt()
			if err == nil {
				fmt.Printf("\t%q: %v\n", "iat", iat.UnixMilli())
			}

			// nbf: not before
			nbf, err := v.GetNotBefore()
			if err == nil {
				fmt.Printf("\t%q: %v\n", "nbf", nbf.UnixMilli())
			}

			// iss: issuer
			iss, err := v.GetIssuer()
			if err == nil {
				fmt.Printf("\t%q: %v\n", "iss", iss)
			}

			// sub: subject
			sub, err := v.GetSubject()
			if err == nil {
				fmt.Printf("\t%q: %v\n", "sub", sub)
			}

			// aud: audience
			aud, err := v.GetAudience()
			if err == nil {
				fmt.Printf("\t%q: %v\n", "aud", strings.Join(aud, ","))
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(parseCmd)

	parseCmd.Flags().StringVarP(&secret, "secret", "s", "", "the secret string specified to sign or parse token")
	parseCmd.Flags().StringVarP(&keyfile, "key-file", "f", "private_key.pem", "the private key file")
	parseCmd.MarkFlagsOneRequired("secret", "key-file")
	parseCmd.MarkFlagsMutuallyExclusive("secret", "key-file")
}
