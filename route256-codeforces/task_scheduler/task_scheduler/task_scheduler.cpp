#include <iostream>
#include <queue>
#include <fstream>
#include <sstream>

struct RunningProc {
    uint64_t proc;
    uint64_t freeTime;
};

struct Task {
    uint64_t start;
    uint64_t duration;
};

int main()
{
    //std::ofstream myfile;
    //myfile.open("06", std::fstream::in | std::fstream::out | std::fstream::app);

    //std::stringstream ss;
    //ss << myfile.rdbuf();
    long taskNum = 0;
    long procNum = 0;

    auto processorPriority = [](const RunningProc& left, const RunningProc& right) {
        return right.freeTime < left.freeTime;
    };

    std::priority_queue<uint64_t, std::vector<uint64_t>, std::greater<uint64_t>> freeProcs;
    std::priority_queue<RunningProc, std::vector<RunningProc>, decltype(processorPriority)> runningProcs(processorPriority);
    //ss >> procNum >> taskNum;
    std::cin >> procNum >> taskNum;
    std::vector<Task> tasks(taskNum);

    uint64_t total = 0;

    for (long i = 0; i < procNum; i++) {
        uint64_t proc = 0;
        //ss >> proc;
        std::cin >> proc;
        freeProcs.push(proc);
    }
    for (long i = 0; i < taskNum; i++) {
        Task task;
        //ss >> task.start;
        //ss >> task.duration;
        std::cin >> task.start;
        std::cin >> task.duration;
        tasks[i] = task;
    }
    for (long i = 0; i < taskNum; i++) {
        Task task = tasks[i];

        while (true) {
            if (runningProcs.empty()) {
                break;
            }
            RunningProc currentRunningTask = runningProcs.top();
            if (currentRunningTask.freeTime <= task.start) {
                freeProcs.push(currentRunningTask.proc);
                runningProcs.pop();
            }
            else {
                break;
            }
        }
        if (freeProcs.empty()) {
            continue;
        }
        uint64_t freeProc = freeProcs.top();
        freeProcs.pop();
        RunningProc running;
        running.freeTime = task.start + task.duration;
        running.proc = freeProc;
        runningProcs.push(running);
        uint64_t totalUpdate = freeProc * task.duration;
        total += totalUpdate;
    }
    std::cout << total << std::endl;
}
//
//if (!freeProcs.empty() && !runningProcs.empty()) {
//    uint64_t freeProc = freeProcs.top();
//    RunningProc running = runningProcs.top();
//
//    if (running.freeTime <= task.start && running.proc <= freeProc) {
//        runningProcs.pop();
//        running.freeTime = task.start + task.duration;
//        runningProcs.push(running);
//        uint64_t totalUpdate = running.proc * task.duration;
//        total += totalUpdate;
//    }
//    else {
//        freeProcs.pop();
//        RunningProc running;
//        running.freeTime = task.start + task.duration;
//        running.proc = freeProc;
//        runningProcs.push(running);
//        uint64_t totalUpdate = freeProc * task.duration;
//        total += totalUpdate;
//
//    }
//}
//else if (!freeProcs.empty()) {
//    uint64_t freeProc = freeProcs.top();
//    freeProcs.pop();
//    RunningProc running;
//    running.freeTime = task.start + task.duration;
//    running.proc = freeProc;
//    runningProcs.push(running);
//    uint64_t totalUpdate = freeProc * task.duration;
//    total += totalUpdate;
//
//}
//else {
//    RunningProc running = runningProcs.top();
//
//    if (running.freeTime <= task.start) {
//        runningProcs.pop();
//        running.freeTime = task.start + task.duration;
//        runningProcs.push(running);
//        uint64_t totalUpdate = running.proc * task.duration;
//        total += totalUpdate;
//
//    }
//}