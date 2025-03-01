package factory;

import pizza.PepperoniPizza;
import pizza.Pizza;

public class PepperoniFactory extends PizzaFactory {
    @Override
    public Pizza prepare() {
        return new PepperoniPizza();
    }
}
