package burger.factory;

import burger.product.Burger;

// Restaurant is the abstract class with the agnostic order method which all extends implement

public abstract class Restaurant {

    public Burger order() {
        Burger burger = createBurger();
        burger.prepare();
        return burger;
    }

    // since restaurant is abstract, to create new options
    // we need to simple implement create new classes extending restaurant
    protected abstract Burger createBurger();

}
