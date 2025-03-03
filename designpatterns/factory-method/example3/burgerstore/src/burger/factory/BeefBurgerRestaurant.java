package burger.factory;

import burger.product.BeefBurger;
import burger.product.Burger;

// Restaurant but for BeefBurgers

public class BeefBurgerRestaurant extends Restaurant{
    // BeefBurgerRestaurant is an option in the menu
    // it is an extension of the Restaurant that simple
    // create new BeefBurgers and thats is
    // Restaurant now is closed for modifications but open for extension (open closed principle)
    @Override
    public Burger createBurger() {
        return new BeefBurger();
    }
}
