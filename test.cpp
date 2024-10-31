#include <iostream>
#include <vector>
#include <cmath>  // for std::fabs
#include <string>
int main() {
    const int SIZE = 784;  // Input size
    const int ROWS = 28;   // Number of rows in the matrix
    const int COLS = 28;   // Number of columns in the matrix
    const float EPSILON = 1e-3; // Tolerance for floating point comparison
    
    // Create a vector to store 784 input values
    std::vector<float> input(SIZE);
    std::vector<char> output(SIZE);
    // Taking input from the user
    std::cout << "Enter 784 floating-point numbers (with up to 2-3 decimal places):" << std::endl;
    for (int i = 0; i < SIZE; ++i) {
        std::cin >> input[i];
        
        // If the input number is greater than epsilon (effectively greater than 0), replace it with 1
        if (input[i] > EPSILON) {
            output[i] = '#';
        } else {
            output[i] = '.';  // Explicitly setting it to 0 if it's not greater than EPSILON
        }
    }
    
    // Printing the 28x28 matrix
    std::cout << "28x28 Matrix:" << std::endl;
    for (int i = 0; i < ROWS; ++i) {
        for (int j = 0; j < COLS; ++j) {
            std::cout << output[i * COLS + j] << " ";
        }
        std::cout << std::endl;
    }
    
    return 0;
}
