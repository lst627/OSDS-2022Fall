#include <vector>
#include <tuple>

#include <string>   // string
#include <chrono>   // timer
#include <iostream> // cout, endl
#include <atomic>
#include <thread>
#include <mutex>
#include <condition_variable>
#include <future>

#include "concurrent_queue/concurrent_queue.h"
#include "utils/utils.h"

namespace proj1 {

void run_instructions(Instruction* inst, ConcurrentQueue* Q) {
    int return_value;

    for (int i=0; i < inst->num; ++i) {
        switch(inst->order[i]) {
            case ENQUEUE: {
                //printf("ENQUEUE\n");
                return_value = Q->enqueue(inst->items[i]);
                if (return_value) {printf("ENQUEUE ERROR!\n"); exit(1);}
                break;
            }
            case DEQUEUE: {
                //printf("DEQUEUE\n");
                return_value = Q->dequeue(inst->items[i]);
                if (return_value) {printf("DEQUEUE ERROR!\n"); exit(1);}
                break;
            }
            case ASK_SIZE: {
                //printf("ASK_SIZE\n");
                return_value = Q->size();
                break;
            }
            case ASK_CAPACITY: {
                //printf("ASK_CAPACITY\n");
                return_value = Q->capacity();
                break;
            }
        }
        inst->res.push_back(return_value);
    }

}
} // namespace proj1


int main(int argc, char *argv[]) {

    int load = 3;
    int capacity = 7;

    proj1::Instruction* I = new proj1::Instruction[load];
    for (int i=0; i< load; ++i) 
        I[i].GetInstruction();
    proj1::ConcurrentQueue* Q = new proj1::ConcurrentQueue();
    Q->init(capacity);
    {
    proj1::AutoTimer timer("proj1");  // using this to print out timing of the block

    // Run all the instructions
    std::vector<std::thread *> threadArr;
    for (int i=0; i< load; ++i) {
        std::thread *t = new std::thread(proj1::run_instructions, &I[i], Q);
        threadArr.push_back(t);
    }
    int len = threadArr.size();
    for (int j=0;j<len;++j) 
        threadArr[j]->join();
    }

    delete Q;

    // Output 
    for (int i=0; i< load; ++i) {
        printf("%d:\n", i);
        I[i].Output();
    }
    return 0;
}