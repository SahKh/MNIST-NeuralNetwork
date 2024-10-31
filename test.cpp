#include <iostream>
#include <vector>
#include <cmath>  /
#include <string>
int main() {
    const int SIZE = 784;  
    const int ROWS = 28;   
    const int COLS = 28;   
    const float EPSILON = 1e-3; 
    
    std::vector<float> input(SIZE);
    std::vector<char> output(SIZE);
    std::cout << "784 floating-point numbers, Use input redirection to stream in data" << std::endl;
    for (int i = 0; i < SIZE; ++i) {
        std::cin >> input[i];
        
        if (input[i] > EPSILON) {
            output[i] = '#';
        } else {
            output[i] = '.';  
        }
    }
    

    std::cout << "28x28 Matrix:" << std::endl;
    for (int i = 0; i < ROWS; ++i) {
        for (int j = 0; j < COLS; ++j) {
            std::cout << output[i * COLS + j] << " ";
        }
        std::cout << std::endl;
    }
    
    return 0;
}
