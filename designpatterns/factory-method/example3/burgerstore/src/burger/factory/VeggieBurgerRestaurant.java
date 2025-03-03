package burger.factory;

import burger.product.Burger;
import burger.product.VeggieBurger;

// Restaurant but for VeggieBurgers

public class VeggieBurgerRestaurant extends Restaurant{
    @Override
    public Burger createBurger() {
        return new VeggieBurger();
    }
}
