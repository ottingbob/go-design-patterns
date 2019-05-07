package prototype

import (
	"fmt"
	"errors"
)

// Similar to builder pattern where we want a default template
// for a given algorithm or build process
//
// Objects are cloned for the user at runtime
//
// You may build a cache-like solution storing information using
// a prototype

// Maintain a set of objects that will be cloned to create new instances
//
// Provde a default value of some type to start working on top of it
//
// Free CPU of complex object initialization to take more memory resources

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

type ShirtsCache struct {}
func (sc *ShirtsCache)GetClone(s int) (ItemInfoGetter, error) {
	switch s {
	case White:
		newItem := *WhitePrototype
		return &newItem, nil
	case Black:
		newItem := *BlackPrototype
		return &newItem, nil
	case Blue:
		newItem := *BluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n",
		s.SKU, s.Color, s.Price)
}

func GetShirtsCloner() ShirtCloner {
	return new(ShirtsCache)
}

var WhitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var BlackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

var BluePrototype *Shirt = &Shirt {
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

func (i *Shirt) GetPrice() float32 {
	return i.Price
}