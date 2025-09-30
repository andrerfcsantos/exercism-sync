public class Twofer {
    public String twofer(String name) {
        String other = name == null ? "you": name;
        return String.format("One for %s, one for me.", other);
    }
}
