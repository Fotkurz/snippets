package burger.product;

// Concrete product

public class BeefBurger implements Burger{

    @Override
    public void prepare() {
        System.out.println("Preparing BeefBurger");
    }
}
