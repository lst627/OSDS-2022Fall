#include <cmath>
#include <string>
#include <iostream>
#include <chrono>
#include <thread>
#include "utils.h"

namespace proj1 {
    void Instruction::GetInstruction() {
        // Random Initialization
        this->num = 10;
        int ord[10] = {ENQUEUE, ENQUEUE, ASK_SIZE, ENQUEUE, DEQUEUE, ASK_SIZE, DEQUEUE, ASK_CAPACITY, DEQUEUE, ASK_SIZE};
        std::string ch[10] = {"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"};
        
        for (int i=0; i < this->num; ++i) {
            this->order.push_back(ord[i]);
            this->items.push_back(ch[i]);
        }
    }

    void Instruction::Output() {
        for (int i=0; i < this->num; ++i) {
            switch(this->order[i]) {
                case ENQUEUE: {
                    printf("ENQUEUE: %d\n", this->res[i]);
                    break;
                }
                case DEQUEUE: {
                    printf("DEQUEUE: %d\n", this->res[i]);
                    std::cout << this->items[i] << std::endl;
                    break;
                }
                case ASK_SIZE: {
                    printf("ASK_SIZE: %d\n", this->res[i]);
                    break;
                }
                case ASK_CAPACITY: {
                    printf("ASK_CAPACITY: %d\n", this->res[i]);
                    break;
                }
            }
        }
    }

    AutoTimer::AutoTimer(std::string name) : 
        m_name(std::move(name)),
        m_beg(std::chrono::high_resolution_clock::now()) { 
    }

    AutoTimer::~AutoTimer() {
        auto end = std::chrono::high_resolution_clock::now();
        auto dur = std::chrono::duration_cast<std::chrono::microseconds>(end - m_beg);
        std::cout << m_name << " : " << dur.count() << " usec\n";
    }
}