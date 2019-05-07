package decorator

import (
	"strings"
	"testing"

	d "decorator"
)

func TestPizzaDecorator_AddIngredient(t *testing.T) {
	pizza := &d.PizzaDecorator{}
	pizzaResult, err := pizza.AddIngredient()
	if err != nil {
		t.Error(err)
	}
	expectedText := "Pizza with the following ingredients:"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When calling the add ingredient of the pizza decorator it " +
			"must return the text %s, not '%s'", expectedText, pizzaResult)
	}
}

func TestOnion_AddIngredient(t *testing.T) {
	onion := &d.Onion{}
	onionResult, err := onion.AddIngredient()
	if err == nil {
		t.Errorf("When calling AddIngredient on the onion decorator without " +
			"an IngredientAdder on its Ingredient field it must return an error, " +
			"not a string with '%s'", onionResult)
	}
	t.Log(err)

	onion = &d.Onion{&d.PizzaDecorator{}}
	onionResult, err = onion.AddIngredient()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(onionResult, "onion") {
		t.Errorf("When calling the add ingredient of the onion decorator it "+
			"must return a text with the word 'onion', not '%s'", onionResult)
	}
}

func TestMeat_AddIngredient(t *testing.T) {
	meat := &d.Meat{}
	meatResult, err := meat.AddIngredient()
	if err == nil {
		t.Errorf("When calling AddIngredient on the meat decorator without " +
			"an IngredientAdder on its Ingredient field it must return an error, " +
			"not a string with '%s'", meatResult)
	}

	meat = &d.Meat{&d.PizzaDecorator{}}
	meatResult, err = meat.AddIngredient()
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(meatResult, "meat") {
		t.Errorf("When calling the add ingredient of the meat decorator it "+
			"must return a text with the word 'meat', not '%s'", meatResult)
	}
}

func TestPizzaDecorator_FullStack(t *testing.T) {

	pizza := &d.Onion{&d.Meat{&d.PizzaDecorator{}}}
	pizzaResult, err := pizza.AddIngredient()
	if err != nil {
		t.Error(err)
	}

	expectedText := "Pizza with the following ingredients: meat, onion"
	if !strings.Contains(pizzaResult, expectedText) {
		t.Errorf("When asking for a pizza with onion and meat the returned "+
			"string must contain the text '%s' but '%s' didn't have it", 
			expectedText, pizzaResult)
	}

	t.Log(pizzaResult)
}