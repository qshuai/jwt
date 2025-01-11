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
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"

	"github.com/spf13/cobra"
)

// rsaCmd represents the rsa command
var rsaCmd = &cobra.Command{
	Use:   "rsa",
	Short: "Generate RSA private key randomly and save the file specified",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.OpenFile(keyfile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0640)
		if err != nil {
			fatalf("open private file err: %v", err)
		}
		defer file.Close()

		privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
		if err != nil {
			fatalf("generate rsa private key err: %v", err)
		}

		payload := x509.MarshalPKCS1PrivateKey(privateKey)
		err = pem.Encode(file, &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: payload,
		})
		if err != nil {
			fatalf("write private key err: %v\n", err)
		}
	},
}

func init() {
	keygenCmd.AddCommand(rsaCmd)
}
