package button;

public class WindowsButton implements Button{

    public WindowsButton() {
    }

    @Override
    public void render() {
        System.out.println("rendering windows button");
    }

    @Override
    public void onClick() {
        System.out.println("onclicking windows button");
    }
}
