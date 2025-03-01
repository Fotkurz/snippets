package factory;

import button.Button;

public abstract class Dialog {

    public abstract void render();

    public abstract Button createButton();
}
