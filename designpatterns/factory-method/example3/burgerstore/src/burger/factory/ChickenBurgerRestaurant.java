package burger.factory;

import burger.product.Burger;
import burger.product.ChickenBurger;

// Restaurant but for ChickenBurgers

public class ChickenBurgerRestaurant extends Restaurant{

    @Override
    public Burger createBurger() {
        return new ChickenBurger();
    }
}
