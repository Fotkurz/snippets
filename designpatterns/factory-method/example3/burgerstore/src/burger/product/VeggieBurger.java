package burger.product;

// Concrete product

public class VeggieBurger implements Burger{
    @Override
    public void prepare() {
        System.out.println("Preparing VeggieBurger");
    }
}
