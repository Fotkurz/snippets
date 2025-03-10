//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
public class Main {
    public static void main(String[] args) {
        DocumentBuilder builder = new DocumentBuilder();

        Document doc = builder.addHeader("Lorem Ipsum")
                .addBody("This is the document body written")
                .addFooter("This is the document footer")
                .save();

        doc.print();
    }
}