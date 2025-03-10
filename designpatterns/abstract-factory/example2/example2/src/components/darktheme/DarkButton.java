package components.darktheme;

import components.Button;

public class DarkButton implements Button {
    @Override
    public void renderButton() {
        System.out.println("<DarkButton></DarkButton>");
    }

}
