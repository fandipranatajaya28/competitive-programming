#include<bits/stdc++.h>
using namespace std;

class LRUCache {
private:
    int capacity_;
    list<pair<int, int>> list_;
    unordered_map<int,list<pair<int, int>>::iterator> map_;
public:
    LRUCache(int capacity) : capacity_(capacity) {}
    
    int get(int key) {

        if (map_.find(key) != map_.end()) { // key in map_
            list_.splice(list_.begin(), list_, map_[key]);
            return map_[key]->second;
        }

        return -1;
    }
    
    void put(int key, int value) {

        if (map_.find(key) != map_.end()) { // key in map_
            map_[key]->second = value;
            list_.splice(list_.begin(), list_, map_[key]);
            return;
        }

        list_.push_front(make_pair(key, value));

        map_[key] = list_.begin();

        if (list_.size() > capacity_) { // Evict
            auto& back = list_.back();
            map_.erase(back.first);
            list_.pop_back();
        }

    }
};

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache* obj = new LRUCache(capacity);
 * int param_1 = obj->get(key);
 * obj->put(key,value);
 */