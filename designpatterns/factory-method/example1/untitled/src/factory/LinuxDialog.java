package factory;

import button.Button;
import button.LinuxButton;

public class LinuxDialog extends Dialog{

    @Override
    public void render() {

    }

    @Override
    public Button createButton() {
        return new LinuxButton();
    }
}
