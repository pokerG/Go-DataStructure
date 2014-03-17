#include <stdio.h>
#include <stdlib.h>
#define N 10
typedef struct
{
    int data;
    int next;
} spacestr;
spacestr SPACE[N];
typedef int position,Cursor;  //int
Cursor L,avail;
void Initialize()
{
    int j;
    for(j=0; j<N-1; j++)
        SPACE[j].next=j+1;
    SPACE[j].next=-1;
    avail=0;
}
Cursor GetNode()
{
    Cursor p;
    if(SPACE[avail].next==-1)
    {
        p=-1;
    }
    else
    {
        p=SPACE[avail].next;
        SPACE[avail].next=SPACE[p].next;
    }
    return p;
}
void FreeNode(Cursor q)
{
    SPACE[q].next=SPACE[avail].next;
    SPACE[avail].next=q;
}
void Insert(int x,position p,spacestr *SPACE)
{
    position q;
    q=GetNode();
    SPACE[q].data=x;
    SPACE[q].next=SPACE[p].next;
    SPACE[p].next=q;
}
void Delete(position p,spacestr *SPACE)
{
    position q;
    if(SPACE[p].next!=-1)
    {
        q=SPACE[p].next;
        SPACE[p].next=SPACE[q].next;
        SPACE[q].data=0;
        FreeNode(q);
    }
}
int main()
{
    Initialize();
    L= GetNode();
    SPACE[L].next=-1;
    int i,choice=1;
    int x,y;
    printf("************************\n");
    printf("1.Please  Insert a data\n");
    printf("2.Please delete a data\n");
    printf("3.Print these data\n");
    printf("0.Exit\n");
    printf("************************\n");
    printf("Your choice");
    while(choice!=0)
    {
        scanf("%d",&choice);
        if(choice==1)
        {
            printf("The Element of add:");
            scanf("%d",&x);
            Insert(x,L,SPACE);
        }
        else if(choice==2)
        {
            printf("The Element of delete:");
            scanf("%d",&y);
            Delete(L,SPACE);
        }

        else
        {
            for(i=0; i<N; i++)
            {
                printf("%d--->%d ## %d\n",i,SPACE[i].data,SPACE[i].next);
            }
        }

    }
    return 0;
}
