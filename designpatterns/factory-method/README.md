# Factory Method

Factory method is one of the mainly used creational design patterns.

It help us defining and creating new options of a common product.

As examples in this folder we have:

- Dialog button -> Java;
- Pizza Store (Simple Factory) -> Java;
- BurgerRestaurant -> Java;

## Simple Factory

Simple factory is a step earlier from the Factory Method design pattern.
It does not require the inheritance to work, so it can be implemented
in languages that does not fully support OOP like Go and Rust.

The difference between simple and factory method is that Simple factory
have a class specific to decide which factory to invoke based on
a value:

pseudocode:


```
function order(pizza) Pizza {
  PizzaFactory factory;

  if pizza == "cheese" {
    factory = new CheeseFactory();
    return factory.prepare()
  } else if pizza == "pepperoni" {
    factory = new PepperoniFactory()
    return factory.prepare()
  }
  // and so on...
}
```

While it is a working factory, it hurts the Open/Closed principle
since we have to modify the `order` function adding/removing
new flavors to properly direct the Pizza creation.

On a factory method in fully oop languages we can solve this
with inheritance like in the ./example3/ poc.

