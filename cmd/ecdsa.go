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
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/spf13/cobra"
)

// ecdsaCmd represents the ecdsa command
var ecdsaCmd = &cobra.Command{
	Use:   "ecdsa",
	Short: "Generate ECDSA private key randomly and save to the file specified",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.OpenFile(keyfile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0640)
		if err != nil {
			fatalf("open key file err: %v\n", err)
		}
		defer file.Close()

		var curve elliptic.Curve
		switch bitSize {
		case 224:
			curve = elliptic.P224()
		case 256:
			curve = elliptic.P256()
		case 384:
			curve = elliptic.P384()
		case 512:
			curve = elliptic.P521()
		default:
			fatalf("malformed bit size")
		}

		privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
		if err != nil {
			fatalf("generate ecdsa key err: %v\n", err)
		}

		payload, err := x509.MarshalECPrivateKey(privateKey)
		if err != nil {
			fatalf("marshal ecdsa key err: %v\n", err)
		}
		err = pem.Encode(file, &pem.Block{
			Type:  "EC PRIVATE KEY",
			Bytes: payload,
		})
		if err != nil {
			fatalf("write private key err: %v\n", err)
		}
	},
}

func init() {
	keygenCmd.AddCommand(ecdsaCmd)
}
