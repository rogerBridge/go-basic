package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
)

var METALIST = []byte{
	0xd1, 0x46, 0xc3, 0x51, 0x56, 0x98, 0xaf, 0x62, 0xfc, 0x2b, 0x45, 0x5a, 0x72, 0xb3, 0x79, 0xfa,
}

func decrypt(encryptString string) []byte {
	encryptBytes, err := hex.DecodeString(encryptString)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("接收到的加密字符串是:", encryptString, encryptBytes)
	R := encryptBytes[2]
	fmt.Println("随机数R:", R)
	// encryptContent 是加密字符串中的有效内容
	encryptContentBytes := encryptBytes[3 : len(encryptBytes)-2]
	fmt.Println("密文是: ", hex.EncodeToString(encryptContentBytes))
	decryptContentByteList := make([]byte, 0, 1000) // 制作一个cap为100的[]byte
	for i := 0; i < len(encryptContentBytes); i++ {
		q := i % 16          // 如果i大于metaList的长度的话, 就减去16
		s := METALIST[q] + R // 异或中的一个因子
		decryptContentByteList = append(decryptContentByteList, s^encryptContentBytes[i])
	}
	fmt.Println("解密后的数据: ", hex.EncodeToString(decryptContentByteList))
	return decryptContentByteList
}

// 全部操作建立在[]byte的类型上
func encrypt(originalData string) []byte {
	R := byte(0x7b)
	log.Println("生成的随机数R是:", R)
	// 首先, 将content转化为[]byte类型
	contentByte, _ := hex.DecodeString(originalData)
	log.Println("contentByte is:", contentByte, len(contentByte))
	encryptByte := make([]byte, 0, 1000)
	for i := 0; i < len(contentByte); i++ {
		q := i % 16
		encryptByte = append(encryptByte, contentByte[i]^(METALIST[q]+R))
		//fmt.Println(encryptByte, contentByte[i]^METALIST[q]+R)
	}
	log.Println("encrypt data is:", encryptByte)
	dataLen := byte(len(encryptByte) + 5)
	result := append([]byte{0xfa, dataLen, R}, encryptByte...) // 问题点, 校验值的byte数组如何加到数组末尾
	result = append(result, checkValue(result)...)
	fmt.Println("byte数组是:", result)
	//showStr := "["
	//for i:=0; i<len(result); i++ {
	//	showStr += string(result[i])+","
	//}
	//showStr += "]"
	//fmt.Println(showStr)
	//fmt.Println("加密后的数据:", hex.EncodeToString(result))
	return result
}

func checkValue(checkThing []byte) []byte {
	// 校验值是一个16位的整数
	var check uint16
	for i := 0; i < len(checkThing); i++ {
		check += uint16(checkThing[i])
	}
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, check)
	// 反转一下b, 使之符合高低位
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	log.Println("生成的校验值是:", b, check)
	return b
}
