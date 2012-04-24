package kirkos

import "testing"
//import "strconv"
//import "github.com/jbussdieker/golibxml"

import "fmt"

func TestBasic(t *testing.T) {
	c := NewCache(2, 2)
	c.Set("test1", "value1")
	c.Set("test2", "value2")
	fmt.Println(c)
	c.Set("test3", "value3")
	fmt.Println(c)
}
/*
func TestXMLFree(t *testing.T) {
	c := NewCache(200, 2)
	for i := 0; true; i++ {
		c.Set("test" + strconv.Itoa(i), golibxml.NewDoc("1.0"))
	}
	c.Free()
}
*/
/*
func TestMixed(t *testing.T) {
	c := NewCache(20, 2)

	for i := 0; true; i++ {
		if (i % 3) == 0 {
			c.Set("test" + strconv.Itoa(i), golibxml.NewDoc("1.0"))
		} else {
			c.Set("test" + strconv.Itoa(i), "value" + strconv.Itoa(i))
		}
	}
	c.Free()
	//fmt.Println(c)
}
*/
