package pkg

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"os"
	"sort"
	"strings"
)

// GetPrivateKeyStr 获取私钥明文
func GetPrivateKeyStr(privateKey *rsa.PrivateKey) (string, error) {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privateKeyBytes,
		},
	)
	//return base64.StdEncoding.EncodeToString(privateKeyPEM), nil
	return string(privateKeyPEM), nil
}

// GetPrivateKey 获取私钥对象
func GetPrivateKey(password, path string) (*rsa.PrivateKey, error) {
	p12Data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	blocks, err := pkcs12.ToPEM(p12Data, password)
	if err != nil {
		return nil, err
	}

	for _, b := range blocks {
		if b.Type == "PRIVATE KEY" || b.Type == "RSA PRIVATE KEY" {
			var key *rsa.PrivateKey
			if key, err = x509.ParsePKCS1PrivateKey(b.Bytes); err != nil {
				return nil, err
			}
			return key, nil
		}
	}

	return nil, errors.New("private key not found")
}

// GetPublicKey 获取公钥对象
func GetPublicKey(path string) (*rsa.PublicKey, error) {
	certData, err := os.ReadFile(path)
	if err != nil {

		return nil, err
	}
	// PEM解码
	block, _ := pem.Decode(certData)
	if block == nil {
		fmt.Println("Failed to decode PEM block")
	}
	certBody, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 获取公钥
	publicKey, ok := certBody.PublicKey.(*rsa.PublicKey)
	if !ok {
		fmt.Println("Not an RSA public key")
	}

	return publicKey, nil
}

// GetPublicKeyStr 获取公钥明文
func GetPublicKeyStr(publicKey *rsa.PublicKey) string {
	// 提取公钥
	publicKeyDer, _ := x509.MarshalPKIXPublicKey(publicKey)
	publicKeyBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyDer,
	}
	publicKeyPem := pem.EncodeToMemory(&publicKeyBlock)
	return string(publicKeyPem)
}

// 将json的value自然排序后拼接返回字符串
func AssemblyData(dataJson []byte) string {
	// Unmarshal the JSON to a map.
	var demoMap map[string]interface{}
	err := json.Unmarshal(dataJson, &demoMap)
	if err != nil {
		panic(err)
	}

	//  按照key自然排序 map
	keys := MapKeysSort(demoMap)

	//fmt.Println("keys====", keys)

	// 按照排序后的顺序遍历 map
	plainTextBuilder := ""
	for _, key := range keys {
		plainTextBuilder += fmt.Sprintf("%v", demoMap[key])
	}
	//fmt.Println("plainTextBuilder====", plainTextBuilder)
	return plainTextBuilder
}

func EncryptData(publicKeyPath, str string) (string, error) {
	// 从证书解析公钥
	publicKey, err := GetPublicKey(publicKeyPath)
	if err != nil {
		fmt.Println(err)
	}
	encryptedStr, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(str))
	if err != nil {
		fmt.Printf("Failed to encrypt password: %v\n", err)
	}
	return base64.StdEncoding.EncodeToString(encryptedStr), nil
	//return string(encryptedStr), nil
}

func DecryptData(ciphertextStr, keyPath, privateKeyPath string) (string, error) {
	privateKey, err := GetPrivateKey(keyPath, privateKeyPath)
	if err != nil {
		panic(err)
	} //str, err := GetPrivateKeyStr(privateKey)
	//if err != nil {
	//	return nil, err
	//}
	//fmt.Println("privateKey:" + str)

	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextStr)
	if err != nil {
		return "", err
	}
	// 解密数据
	decryptedPassword, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return "", err
	}
	return string(decryptedPassword), nil
}

func MapKeysSort(value map[string]interface{}) []string {
	// 创建一个切片保存 map 的 keys
	keys := make([]string, 0, len(value))
	for k := range value {
		keys = append(keys, k)
	}

	// 使用 sort.Slice 函数定义我们自己的排序规则
	sort.Slice(keys, func(i, j int) bool {
		a, b := keys[i], keys[j]
		for k := 0; k < len(a) && k < len(b); k++ {
			if a[k] != b[k] {
				if a[k] >= 'a' && b[k] >= 'a' || a[k] < 'a' && b[k] < 'a' {
					return a[k] < b[k]
				}
				return a[k] < 'a'
			}
		}
		return len(a) < len(b)
	})
	return keys
}

// GetSortedValues 拼接带签名字符串
func GetSortedValues(data map[string]interface{}, excludedKey map[string]bool) string {
	var keys []string
	for key := range data {
		if excludedKey != nil && excludedKey[key] {
			continue
		}
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var plainTxtSb strings.Builder
	for _, key := range keys {
		value := data[key]
		switch v := value.(type) {
		case string:
			plainTxtSb.WriteString(v)
		case map[string]interface{}:
			subMap := make(map[string]interface{})
			for k, val := range v {
				subMap[k] = val
			}
			subPlain := GetSortedValues(subMap, excludedKey)
			plainTxtSb.WriteString(subPlain)
		case []interface{}:
			for _, item := range v {
				subMap := make(map[string]interface{})
				for k, val := range item.(map[string]interface{}) {
					subMap[k] = val
				}
				subPlain := GetSortedValues(subMap, excludedKey)
				plainTxtSb.WriteString(subPlain)
			}
		}
	}

	return plainTxtSb.String()
}

// Verify 验签
func Verify(data string, sign string, publicKey string) bool {
	// 解密base64编码的公钥
	keyBytes, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		fmt.Println("Failed to decode public key err =", err)
		return false
	}

	// 解析PKIX编码的公钥
	keyInterface, err := x509.ParsePKIXPublicKey(keyBytes)
	if err != nil {
		fmt.Println("Failed to parse public key err = ", err)
		return false
	}

	// 断言公钥类型
	pubKey, ok := keyInterface.(*rsa.PublicKey)
	if !ok {
		fmt.Println("Failed to get public key err = ", err)
		return false
	}

	// 解密签名
	signBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		fmt.Println("Failed to decode sign err = ", err)
		return false
	}

	// 创建哈希
	hashed := sha256.Sum256([]byte(data))

	// 验证签名
	err = rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hashed[:], signBytes)
	if err != nil {
		fmt.Println("Failed to verify sign err = ", err)
		return false
	}

	return true
}
