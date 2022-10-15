#ifndef UTILS_H_
#define UTILS_H_

#include <string>
#include <vector>
#include <iostream>
#include <chrono>  // for AutoTimer function

namespace proj1 {
    enum InstructionOrder {
        ENQUEUE = 0,
        DEQUEUE,
        ASK_SIZE,
        ASK_CAPACITY
    };

    class Instruction {
        public:
        Instruction(){}
        ~Instruction(){}
        void GetInstruction();
        void Output();
        std::vector<int> order;
        std::vector<std::string> items;
        std::vector<int> res;
        int num;
        private:
    };

    class AutoTimer {
    public:
        AutoTimer(std::string name);
        ~AutoTimer(); 
    private:
        std::string m_name;
        std::chrono::time_point<std::chrono::high_resolution_clock> m_beg;
    };
}

#endif