#include <vector>
#include <tuple>
#include <opencv2/opencv.hpp>
#include <opencv2/core/core.hpp>
#include <opencv2/highgui/highgui.hpp>
#include <opencv2/imgproc/imgproc.hpp>
#include <iostream>
#include <string>   
#include <chrono>   
#include <atomic>
#include <thread>
#include <mutex>
#include <condition_variable>
#include <future>
#include <dirent.h>
#include <unistd.h>

#include "../part1/concurrent_queue/concurrent_queue.h"
#include "../part1/utils/utils.h"

namespace proj1 {

void MakeFolder(std::string path) {
    if (access(path.c_str(), 0) == -1)
        system(("mkdir "+ path).c_str());
    else {
        system(("rm -rf "+ path).c_str());
        system(("mkdir "+ path).c_str());
    }
}

void resize_images(ConcurrentQueue* Q, std::string path_from, std::string path_to) {
    int return_value;
    std::string str;
    return_value = Q->dequeue(str);
    while (return_value == 0) {
        // std::cout<< str <<std::endl;
        if (str == "Finish") {Q->close(); break;}
        cv::String img_path = path_from + str; //When using relative path, be relative to the EXECUTABLE file
        cv::Mat img = cv::imread(img_path);
        cv::Size r_size = cv::Size(128, 128);
        cv::Mat resized_img; 
        if (!img.data) {
            printf("Image reading error !\n");
            exit(1);
        }
        cv::resize(img, resized_img, r_size);
        cv::imwrite(path_to + str, resized_img);
        return_value = Q->dequeue(str);
    }
    if (return_value == -1 ) {
        printf("Dequeue error !\n");
        exit(1);
    }
}

void add_images(ConcurrentQueue* Q, std::string path_from) {
    struct dirent *dirp;
    DIR *dir;
    dir = opendir(path_from.c_str());
    if (dir==NULL) {
        printf("Open folder failed\n");
        Q->enqueue("Finish");
        exit(1);
    }
    dirp = readdir(dir);
    while(dirp != NULL) {
        std::string image_name = std::string(dirp->d_name);
        if (image_name.length() > 5) 
            if (image_name.substr(image_name.length()-5, image_name.length()) == ".JPEG") {
                Q->enqueue(image_name);
                // std::cout<< image_name << std::endl;
            }
        dirp = readdir(dir);
    }
    Q->enqueue("Finish");
    closedir(dir);
}

} // namespace proj1


int main(int argc, char *argv[]) {

    int num_workers = 50;
    int load = 10000;
    int capacity = 100;
    std::string path_from =  "./tiny-imagenet-200/test/images/"; // "../test_image"
	std::string path_to = "./tiny-imagenet-200/test/images_resize/"; //"../test_image_resize" 
	proj1::MakeFolder(path_to);

    proj1::ConcurrentQueue* Q = new proj1::ConcurrentQueue();
    Q->init(capacity);
    {
    proj1::AutoTimer timer("proj1");  // using this to print out timing of the block

    // Run all the instructions
    std::vector<std::thread *> threadArr;
    std::thread *p = new std::thread(proj1::add_images, Q, path_from);
    threadArr.push_back(p);
    for (int i=0; i< num_workers; ++i) {
        std::thread *t = new std::thread(proj1::resize_images, Q, path_from, path_to);
        threadArr.push_back(t);
    }
    int len = threadArr.size();
    for (int j=0;j<len;++j) 
        threadArr[j]->join();
    }

    delete Q;

    return 0;
}