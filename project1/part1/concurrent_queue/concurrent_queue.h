#ifndef CONCURRENT_QUEUE_H_
#define CONCURRENT_QUEUE_H_

#include <string>
#include <vector>
#include <atomic>
#include <thread>
#include <mutex>
#include <condition_variable>
#include <future>
#include <typeinfo>

namespace proj1 {

    class ConcurrentQueue {
    public:
        ConcurrentQueue(){}
        ~ConcurrentQueue(){delete []this->data; }
        void init(int capacity);
        int enqueue(std::string item);
        int dequeue(std::string& item);
        int size();
        int capacity();
        void close();

    private:
        int queue_size;
        int queue_capacity;
        int head, tail;
        std::string* data;
        bool closed;
        
        std::mutex queue_mtx;
        std::condition_variable en_cv;
        std::condition_variable de_cv;
    };
}

#endif