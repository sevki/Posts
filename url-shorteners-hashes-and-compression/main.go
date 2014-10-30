package main

import (
	"bitbucket.org/binet/go-gnuplot/pkg/gnuplot"
	"bufio"
	"bytes"
	"compress/gzip"
	"compress/lzw"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/cespare/go-smaz"
	"github.com/howeyc/crc16"
	"hash/adler32"
	"hash/crc32"
	"hash/crc64"
	"os"
	"reflect"
	"runtime"

	"strconv"
)

func main() {
//	h := hashcrc32([]byte("http://fistfulofbytes.com/how-to-bypass-ssl-validation-for-exchange-webservices-managed-api"))

//	printHash(h)
//	printHash(h[:3])
//	printHash(h[:2])

	var crcs []hashFunc
	crcs = append(crcs, hashcrc16IBM, hashcrc32, hashcrc64)

	var cryptic []hashFunc
	cryptic = append(cryptic, hashmd5, hashsha256, hashsha512)

	var all []hashFunc
	all = append(all, hashcrc32_first_3, hashcrc64_first_3, hashmd5_first_3, hashsha256_first_3, hashsha512_first_3)

	var compressions []compressFunction
	compressions = append(compressions, compressSmaz)
//	testCompression(compressions, "smaz", 880, 500)
	//	sames()
	testHash(crcs, "crc", 250621, 5000, 880, 250)
//	testHash(crcs, "crcfine", 1000, 50, 880, 250)
//	testHash(cryptic, "cryptic", 250621, 5000, 880, 250)
//	testHash(cryptic, "crypticfine", 1000, 50, 880, 250)

	testHash(all, "all", 250621, 10000, 440, 250)
	testHash(all, "allfine", 10000, 10, 440, 250)

}
func printHash(s []byte) {
	b := make([]byte, 8)
	base64.StdEncoding.Encode(b, s)
	fmt.Println(s)
	fmt.Println(b)
	printBits(s)
	printBits(b)
}
func printInt(s []byte) {
	for _, x := range s {
		t := strconv.FormatInt(int64(x), 2)
		for i := len(t); i < 8; i++ {
			t += "0"
		}

		fmt.Printf("%s", t)
	}
	fmt.Println()
}
func printBits(s []byte) {
	w := ""
	for _, x := range s {
		t := strconv.FormatInt(int64(x), 2)
		padding := ""
		for i := len(t); i < 8; i++ {
			padding += "0"
		}
		t = padding + t
		w += "<td  class=\"bits\">" + t + "</td>"
	}
	p := len(w)
	for i := p; i%6 == 0; i++ {
		w += "0"
	}

	fmt.Println(w)
}
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func testCompression(comrpessions []compressFunction, name string, width, height int) {
	fname := ""
	persist := false
	debug := false

	p, err := gnuplot.NewPlotter(fname, persist, debug)

	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()
	for i := 0; i < len(comrpessions); i++ {
		plot(compress(comrpessions[i], p))

	}
	p.CheckedCmd(fmt.Sprintf("set terminal png medium size %d,%d", width, height))
	p.CheckedCmd("set key right top")
	p.SetXLabel("Compression percentage")
	p.SetYLabel("URL length")

	p.CheckedCmd(fmt.Sprintf("set output '%s.png'", name))
	p.CheckedCmd("replot")
	p.CheckedCmd("q")
}

func testHash(hashs []hashFunc, name string, hi, freq, width, height int) {
	fname := ""
	persist := false
	debug := false

	p, err := gnuplot.NewPlotter(fname, persist, debug)
	//	p.SetStyle("lines")
	if err != nil {
		err_string := fmt.Sprintf("** err: %v\n", err)
		panic(err_string)
	}
	defer p.Close()
	for i := 0; i < len(hashs); i++ {
		plot(hash(hashs[i], p, hi, freq))
	}
	p.CheckedCmd(fmt.Sprintf("set terminal png medium size %d,%d", width, height))
	p.CheckedCmd("set key right bottom")
	p.SetXLabel("Collisions")
	p.SetYLabel("Added Items")
	p.CheckedCmd(fmt.Sprintf("set output '%s.png'", name))
	p.CheckedCmd("replot")
	p.CheckedCmd("q")
}

type hashFunc func(s []byte) []byte
type compressFunction func(s []byte) []byte

func plot(v, k []float64, x string, p *gnuplot.Plotter) {
	p.PlotXY(v, k, fmt.Sprintf("%s", x))
}
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func hash(fn hashFunc, p *gnuplot.Plotter, hi, freq int) (v, k []float64, x string, y *gnuplot.Plotter) {
	y = p
	hs := make(map[string]bool)
	f, _ := os.Open("urls.txt")
	s := bufio.NewScanner(f)
	rt := 0 // running total
	for i := 1; i <= hi; i++ {
		if !s.Scan() {
//			fmt.Println(i)
			break
		}
		hb := fn(s.Bytes())
		l := base64.URLEncoding.EncodeToString(hb)
		if _, ok := hs[l]; ok {
			rt += 1
		} else {
			hs[l] = true
		}
		if i%freq == 0 {

//			v = append(v, (float64(rt)/float64(i))*100)
			v = append(v, float64(rt))
			k = append(k, float64(i))
//			fmt.Printf("%d\t%f\n", i, (float64(rt)/float64(i))*100)
		}
	}
	x = GetFunctionName(fn)[len("main.hash"):]
	x = fmt.Sprintf("%s:%d", x, rt)
	return
}
func sames() {
	hs := make(map[string]bool)
	f, _ := os.Open("urls.txt")
	s := bufio.NewScanner(f)
	e, _ := os.Open("nodups.txt")
	for i := 0; s.Scan(); i++ {
		l := s.Text()
		if _, ok := hs[l]; ok {
			//			fmt.Printf("%d\n", i)
		} else {
			hs[l] = true
			fmt.Println(s.Text())
		}

	}
	e.Close()
}

func compress(fn compressFunction, p *gnuplot.Plotter) (v, k []float64, x string, y *gnuplot.Plotter) {
	y = p

	f, _ := os.Open("urls.txt")
	s := bufio.NewScanner(f)
	for i := 1; s.Scan(); i++ {
		hb := fn(s.Bytes())
		uc := len(s.Bytes())
		//		co := len(base64.URLEncoding.EncodeToString(hb))
		co := len(hb)
		if i%1 == 0 {
			if uc < 400 {
				p := (float64(co) / float64(uc)) * float64(100)
				v = append(v, p)
				k = append(k, float64(uc))
			}
		}
		//		fmt.Printf("%d %d %s\n", uc, co,  base64.URLEncoding.EncodeToString(hb))
	}
	x = GetFunctionName(fn)[len("main.compress"):]
	return
}
func hashmd5(s []byte) (res []byte) {
	h := md5.New()
	h.Write(s)
	res = h.Sum(nil)
	return
}
func hashsha256(s []byte) (res []byte) {
	h := sha256.New()
	h.Write(s)
	res = h.Sum(nil)
	return
}
func hashsha512(s []byte) (res []byte) {
	h := sha512.New()
	h.Write(s)
	res = h.Sum(nil)
	return
}
func hashmd5_first_3(s []byte) (res []byte) {
	res = hashmd5(s)[:3]
	return
}
func hashsha256_first_3(s []byte) (res []byte) {
	res = hashsha256(s)[:3]
	return
}
func hashsha512_first_3(s []byte) (res []byte) {
	res = hashsha512(s)[:3]
	return
}
func hashcrc32_first_3(s []byte) []byte {
	h := crc32.NewIEEE()
	h.Write(s)
	return h.Sum(nil)[:3]
}
func hashcrc64_first_3(s []byte) []byte {
	var tab = crc64.MakeTable(crc64.ECMA)
	h := crc64.New(tab)
	h.Write(s)
	return h.Sum(nil)[:3]
}
func hashcrc16IBM(s []byte) []byte {
	h := crc16.NewIBM()
	h.Write(s)
	return h.Sum(nil)
}
func hashcrc16CCITT(s []byte) []byte {
	h := crc16.NewCCITT()
	h.Write(s)
	return h.Sum(nil)
}
func hashcrc16SCSI(s []byte) []byte {
	h := crc16.NewSCSI()
	h.Write(s)
	return h.Sum(nil)
}
func hashcrc32(s []byte) []byte {
	h := crc32.NewIEEE()
	h.Write(s)
	return h.Sum(nil)
}
func hashcrc64(s []byte) []byte {
	var tab = crc64.MakeTable(crc64.ECMA)
	h := crc64.New(tab)
	h.Write(s)

	return h.Sum(nil)
}

func hashadler32(s []byte) (res []byte) {
	h := adler32.New()
	h.Write(s)
	res = h.Sum(nil)
	return
}
func compressSmaz(s []byte) []byte {
	return smaz.Compress(s)
}
func compressGzip(s []byte) []byte {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write([]byte("hello, world\n"))
	w.Close()
	return b.Bytes()
}
func compressLZW(s []byte) []byte {
	var b bytes.Buffer
	w := lzw.NewWriter(&b, lzw.LSB, 8)
	w.Write([]byte("hello, world\n"))
	defer w.Close()
	return b.Bytes()
}
