#include <iostream>
#include <cstdio>
#include <queue>
#include <cstring>
#include <algorithm>
using namespace std;
struct person {
    int id;
    person *next;
    person *pre;
} *head,*cur;
int n;
void init(){
     for(int i=1; i<=n; i++) {
        cur->id=i;
        if(i!=n)    cur->next=new person;
        else cur->next=head;
        cur->next->pre=cur;
        cur=cur->next;
    }
}
int main() {
    int m;
    scanf("%d %d",&n,&m); //n为人数 m为周期
    head=new person;
    cur=head;
    init();
    cur=head;
    while(cur->next!=cur) {
        for(int i=2; i<=m; i++) {
            cur=cur->next;
        }
        cur->pre->next=cur->next;
        cur->next->pre=cur->pre;
        cout<<"kill :"<<cur->id<<endl;
        cur=cur->next;
    }
    cout<<"Winner is "<<cur->id<<"."<<endl;
    return 0;
}
