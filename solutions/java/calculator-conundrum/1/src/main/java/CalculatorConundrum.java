class CalculatorConundrum {
    public String calculate(int operand1, int operand2, String operation) {
        if (operation=="/"&&operand2==0){
            throw new IllegalOperationException("Division by zero is not allowed",new ArithmeticException());
        }
        switch (operation){
            case "+":
                return operand1 + " + " + operand2 + " = " + (operand1+operand2);
            case "*":
                return operand1 + " * " + operand2 + " = " + (operand1*operand2);
            case "/":
                return operand1 + " / " + operand2 + " = " + (operand1/operand2);
            case null:
                throw new IllegalArgumentException("Operation cannot be null");
            case "":
                throw new IllegalArgumentException("Operation cannot be empty");
            default:
                throw new IllegalOperationException("Operation '"+operation+"' does not exist");
        }
    }
}
