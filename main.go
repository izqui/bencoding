package bencoding

import "strconv"

func Decode(b []byte) (interface{}, []byte) {

	switch b[0] {
	case 'i':

		i := 1
		p := []byte{}
		for ; b[i] != 'e'; i++ {
			p = append(p, b[i])
		}

		n, _ := strconv.Atoi(string(p))

		return n, b[i+1:] //Don't return the 'e'

	case 'l':

		l := []interface{}{}
		r := b[1:]

		for r[0] != 'e' {

			var item interface{}
			item, r = Decode(r)
			l = append(l, item)
		}

		return l, r[1:]

	case 'd':

		d := map[string]interface{}{}

		r := b[1:]

		for r[0] != 'e' {

			var key, value interface{}
			key, r = Decode(r)
			value, r = Decode(r)

			d[key.(string)] = value
		}

		return d, r[1:]

	default:

		i := 0

		p := []byte{}
		for ; b[i] != ':'; i++ {
			p = append(p, b[i])
		}

		i += 1 //Ignore the colon

		n, _ := strconv.Atoi(string(p))
		n += i

		s := []byte{}
		for ; i < n; i++ {
			s = append(s, b[i])
		}

		return string(s), b[i:]

	}

	panic("Unknown")
}
