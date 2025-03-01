package pizza;

public class PepperoniPizza implements Pizza{

    public PepperoniPizza(){

    }

    @Override
    public void prepare() {
        System.out.println("Preparing PepperoniPizza");
    }
}
