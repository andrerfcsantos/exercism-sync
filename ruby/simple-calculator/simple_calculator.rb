class SimpleCalculator
  ALLOWED_OPERATIONS = %w[+ / *].freeze

  class UnsupportedOperation < StandardError
  end

  def self.calculate(first_operand, second_operand, operation)
    unless ALLOWED_OPERATIONS.include?(operation)
      raise UnsupportedOperation, 'Only operations +, / and * are supported'
    end
    unless first_operand.is_a?(Numeric) && second_operand.is_a?(Numeric)
      raise ArgumentError, 'Arguments must be numbers'
    end

    result = begin
      case operation
      when '+'
        first_operand + second_operand
      when '/'
        first_operand / second_operand
      when '*'
        first_operand * second_operand
      end
    end

    "#{first_operand} #{operation} #{second_operand} = #{result}"
  rescue ZeroDivisionError
    'Division by zero is not allowed.'
  end
end
