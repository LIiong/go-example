package main

import (
    "crypto"
    "crypto/md5"
    "crypto/rand"
    "crypto/rsa"
    "fmt"
)

//rsa加密解密 签名验签
func rsaCrypto() {
    //生成私钥
    priv, e := rsa.GenerateKey(rand.Reader, 1024)
    if e != nil {
        fmt.Println(e)
    }

    //根据私钥产生公钥
    pub := &priv.PublicKey

    //明文
    plaintext := []byte("Hello world")

    //加密生成密文
    fmt.Printf("%q\n加密:\n", plaintext)
    ciphertext, e := rsa.EncryptOAEP(md5.New(), rand.Reader, pub, plaintext, nil)
    if e != nil {
        fmt.Println(e)
    }
    fmt.Printf("\t%x\n", ciphertext)

    //解密得到明文
    fmt.Printf("解密:\n")
    plaintext, e = rsa.DecryptOAEP(md5.New(), rand.Reader, priv, ciphertext, nil)
    if e != nil {
        fmt.Println(e)
    }
    fmt.Printf("\t%q\n", plaintext)

    //消息先进行Hash处理
    h := md5.New()
    h.Write(plaintext)
    hashed := h.Sum(nil)
    fmt.Printf("%q MD5 Hashed:\n\t%x\n", plaintext, hashed)

    //签名
    opts := &rsa.PSSOptions{SaltLength: rsa.PSSSaltLengthAuto, Hash: crypto.MD5}
    sig, e := rsa.SignPSS(rand.Reader, priv, crypto.MD5, hashed, opts)
    if e != nil {
        fmt.Println(e)
    }
    fmt.Printf("签名:\n\t%x\n", sig)

    //认证
    fmt.Printf("验证结果:")
    if e := rsa.VerifyPSS(pub, crypto.MD5, hashed, sig, opts); e != nil {
        fmt.Println("失败:", e)
    } else {
        fmt.Println("成功.")
    }
}