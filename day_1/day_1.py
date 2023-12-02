def extract_claibration_values(file):
    total_sum: int = 0
    input = []
    with open(file) as fp:
        for file_line in fp:
            input.append(file_line)
    for line in input:
        first_digit = next((char for char in line if char.isdigit()), None)
        last_digit = next((char for char in reversed(line) if char.isdigit()), None)
        if first_digit is not None and last_digit is not None:
            calibration_value = int(first_digit + last_digit)
            total_sum += calibration_value
    return total_sum


print("The result is: ", extract_claibration_values("input.txt"))
