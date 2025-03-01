import factory.PepperoniFactory;
import factory.PizzaFactory;
import pizza.Pizza;

import java.text.Format;

public class PizzaStore {

    public Pizza prepare(String flavor) {
        switch (flavor) {
            case "pepperoni":
                PizzaFactory factory = new PepperoniFactory();
                return factory.prepare();
            default:
                throw new IllegalArgumentException("Invalid flavor: " + flavor);
        }
    }
}
