#include <fstream>
#include <iostream>
#include <sstream>
#include <cmath>
#include <atomic>
#include <thread>
#include <mutex>
#include <condition_variable>
#include <future>

#include "concurrent_queue.h"

namespace proj1 {

    void ConcurrentQueue::init(int capacity) {
        this->queue_capacity = capacity;
        this->data = new std::string[capacity];
        this->queue_size = 0;
        this->head = 0;
        this->tail = capacity - 1;
        this->closed = false;
    }

    int ConcurrentQueue::enqueue(std::string item) {
        if (this->closed) return 1; // error number
        {
            std::unique_lock<std::mutex> lock(this->queue_mtx);

            this->en_cv.wait(lock, [this]{return this->queue_size < this->queue_capacity;});
            if (this->closed) return -1;
            if (this->queue_size >= this->queue_capacity)
                return -1; // error number

            this->queue_size ++;
            this->tail ++;
            if (this->tail >= this->queue_capacity)
                this->tail = 0;
            this->data[this->tail] = item;
        }
        this->de_cv.notify_one();
        return 0;
    }

    int ConcurrentQueue::dequeue(std::string& item) {
        if (this->closed) return 1; // closed number
        {
            std::unique_lock<std::mutex> lock(this->queue_mtx);

            this->de_cv.wait(lock, [this]{return (this->queue_size > 0)||(this->closed);});
            if (this->closed) return 1;
            if (this->queue_size == 0)
                return -1; // error number

            this->queue_size --;
            item = this->data[this->head]; 
            this->head ++;
            if (this->head >= this->queue_capacity)
                this->head = 0;
        }
        this->en_cv.notify_one();
        return 0;
    }

    int ConcurrentQueue::size() {
        return this->queue_size;
    }

    int ConcurrentQueue::capacity() {
        return this->queue_capacity;
    }

    void ConcurrentQueue::close() {
        this->closed = true;
        this->de_cv.notify_all();
    }

}