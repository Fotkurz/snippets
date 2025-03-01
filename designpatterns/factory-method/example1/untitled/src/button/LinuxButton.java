package button;

public class LinuxButton implements Button {
    @Override
    public void render() {
        System.out.println("rendering LinuxBUtton");
    }

    @Override
    public void onClick() {
        System.out.println("on click linuxbutton");
    }
}
